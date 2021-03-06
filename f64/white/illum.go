// Copyright 2012 The Chroma Authors. All rights reserved. See the LICENSE file.

// Illumination functions

package white

//Values for Daylight illuminant: s0 s1 s2
func daylight_spect(wavelength, m1, m2, temp float64) float64 {
	// should be const, not var

	//s0
	var s0 = [97]float64{61.80, 61.65, 61.50, 65.15, 68.80, 66.10, 63.40, 64.60, 65.80, 80.30, 94.80, 99.80, 104.80, 105.35, 105.90, 101.35, 96.80, 105.35, 113.90, 119.75, 125.60, 125.55, 125.50, 123.40, 121.30, 121.30, 121.30, 117.40, 113.50, 113.30,
		113.10, 111.95, 110.80, 108.65, 106.50, 107.65, 108.80, 107.05, 105.30, 104.85, 104.40, 102.20, 100.00, 98.00, 96.00, 95.55, 95.10, 92.10, 89.10, 89.80, 90.50, 90.40, 90.30, 89.35, 88.40, 86.20, 84.00, 84.55, 85.10,
		83.50, 81.90, 82.25, 82.60, 83.75, 84.90, 83.10, 81.30, 76.60, 71.90, 73.10, 74.30, 75.35, 76.40, 69.85, 63.30, 67.50, 71.70, 74.35, 77.00, 71.10, 65.20, 56.45, 47.70, 58.15, 68.60, 66.80, 65.00, 65.50, 66.00, 63.50, 61.00, 57.15,
		53.30, 56.10, 58.90, 60.40, 61.90}
	//s1
	var s1 = [97]float64{41.60, 39.80, 38.00, 40.70, 43.40, 40.95, 38.50, 36.75, 35.00, 39.20, 43.40, 44.85, 46.30, 45.10, 43.90, 40.50, 37.10, 36.90, 36.70, 36.30, 35.90, 34.25, 32.60, 30.25, 27.90, 26.10, 24.30, 22.20, 20.10, 18.15, 16.20, 14.70,
		13.20, 10.90, 8.60, 7.35, 6.10, 5.15, 4.20, 3.05, 1.90, 0.95, 0.00, -0.80, -1.60, -2.55, -3.50, -3.50, -3.50, -4.65, -5.80, -6.50, -7.20, -7.90, -8.60, -9.05, -9.50, -10.20, -10.90, -10.80, -10.70, -11.35, -12.00, -13.00, -14.00,
		-13.80, -13.60, -12.80, -12.00, -12.65, -13.30, -13.10, -12.90, -11.75, -10.60, -11.10, -11.60, -11.90, -12.20, -11.20, -10.20, -9.00, -7.80, -9.50, -11.20, -10.80, -10.50, -10.60, -10.15, -9.70, -9.00, -8.30,
		-8.80, -9.30, -9.55, -9.80}
	//s2
	var s2 = [97]float64{6.70, 6.00, 5.30, 5.70, 6.10, 4.55, 3.00, 2.10, 1.20, 0.05, -1.10, -0.80, -0.50, -0.60, -0.70, -0.95, -1.20, -1.90, -2.60, -2.75, -2.90, -2.85, -2.80, -2.70, -2.60, -2.60, -2.60, -2.20, -1.80, -1.65, -1.50, -1.40, -1.30,
		-1.25, -1.20, -1.10, -1.00, -0.75, -0.50, -0.40, -0.30, -0.15, 0.00, 0.10, 0.20, 0.35, 0.50, 1.30, 2.10, 2.65, 3.65, 4.10, 4.40, 4.70, 4.90, 5.10, 5.90, 6.70, 7.00, 7.30, 7.95, 8.60, 9.20, 9.80, 10.00, 10.20, 9.25, 8.30, 8.95,
		9.60, 9.05, 8.50, 7.75, 7.00, 7.30, 7.60, 7.80, 8.00, 7.35, 6.70, 5.95, 5.20, 6.30, 7.40, 7.10, 6.80, 6.90, 7.00, 6.70, 6.40, 5.95, 5.50, 5.80, 6.10, 6.30, 6.50}

	wlm := int((wavelength - 350.) / 5.)
	return s0[wlm] + m1*s1[wlm] + m2*s2[wlm]
}

// spectral data for Daylight direct Sun: I have chosen 5300K beacuse Nikon=5200K, Olympus=5300K, Panasonic=5500K, Leica=5400K, Minolta=5100K
func daylight5300_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{24.82, 26.27, 27.72, 28.97, 30.22, 29.71, 29.19, 31.95, 34.71, 45.49, 56.26, 59.97, 63.68, 65.30, 66.92, 65.39, 63.86, 72.59, 81.32, 87.53, 93.73, 95.15, 96.56, 96.55, 96.54, 98.13, 99.73, 97.70, 95.66, 97.19, 98.72,
		98.90, 99.08, 98.98, 98.87, 101.13, 103.39, 102.48, 101.57, 102.14, 102.71, 101.36, 100.00, 98.71, 97.42, 97.81, 98.21, 95.20, 92.20, 93.92, 95.63, 96.15, 96.67, 96.34, 96.01, 94.21, 92.41, 93.58, 94.74, 93.05, 91.36, 92.29,
		93.21, 95.25, 97.28, 95.30, 93.32, 87.92, 82.51, 84.29, 86.06, 86.94, 87.81, 80.24, 72.68, 77.32, 81.96, 84.88, 87.79, 81.01, 74.22, 64.41, 54.60, 66.55, 78.51, 76.35, 74.20, 74.79, 75.38, 72.48, 69.58, 65.11, 60.64,
		63.88, 67.13, 68.85, 70.57}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for Daylight Cloudy: I have chosen 6200K beacuse Nikon=6000K, Olympus=6000K, Panasonic=6200K, Leica=6400K, Minolta=6500K
func cloudy6200_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{39.50, 40.57, 41.63, 43.85, 46.08, 45.38, 44.69, 47.20, 49.71, 63.06, 76.41, 80.59, 84.77, 85.91, 87.05, 84.14, 81.23, 90.29, 99.35, 105.47, 111.58, 112.23, 112.87, 111.74, 110.62, 111.41, 112.20, 108.98, 105.76, 106.32,
		106.89, 106.34, 105.79, 104.62, 103.45, 105.09, 106.72, 105.24, 103.76, 103.75, 103.75, 101.87, 100.00, 98.29, 96.58, 96.46, 96.34, 92.85, 89.37, 90.25, 91.12, 91.06, 90.99, 90.17, 89.35, 87.22, 85.10, 85.48, 85.85,
		84.03, 82.20, 82.45, 82.69, 83.92, 85.15, 83.14, 81.13, 76.65, 72.17, 73.27, 74.36, 75.65, 76.95, 70.34, 63.74, 67.98, 72.22, 74.88, 77.54, 71.59, 65.65, 56.82, 47.99, 58.53, 69.06, 67.27, 65.47, 65.96, 66.44, 63.92, 61.41, 57.52,
		53.63, 56.47, 59.31, 60.80, 62.29}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for Daylight Shade: I have choose 7600K beacuse Nikon=8000K, Olympus=7500K, Panasonic=7500K, Leica=7500K, Minolta=7500K
func shade7600_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{64.42, 64.46, 64.51, 68.35, 72.20, 70.22, 68.24, 69.79, 71.35, 87.49, 103.64, 108.68, 113.72, 114.12, 114.53, 109.54, 104.55, 113.59, 122.63, 128.52, 134.41, 134.02, 133.63, 131.02, 128.41, 128.08, 127.75, 123.16,
		118.57, 117.89, 117.22, 115.72, 114.22, 111.60, 108.99, 109.84, 110.68, 108.57, 106.45, 105.71, 104.98, 102.49, 100.00, 97.78, 95.55, 94.82, 94.08, 90.47, 86.87, 86.94, 87.01, 86.45, 85.88, 84.57, 83.27, 80.83, 78.40, 78.21,
		78.03, 76.22, 74.42, 74.15, 73.89, 74.41, 74.92, 73.01, 71.09, 67.26, 63.42, 64.01, 64.60, 66.10, 67.60, 61.83, 56.06, 59.94, 63.82, 66.27, 68.71, 63.49, 58.26, 50.30, 42.34, 51.64, 60.95, 59.45, 57.95, 58.35, 58.76, 56.57,
		54.38, 51.00, 47.62, 50.10, 52.58, 53.88, 55.19}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for tungsten - incandescent 2856K
func tungsten2856_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{4.75, 5.42, 6.15, 6.95, 7.83, 8.78, 9.80, 10.91, 12.09, 13.36, 14.72, 16.16, 17.69, 19.30, 21.01, 22.80, 24.68, 26.65, 28.71, 30.86, 33.10, 35.42, 37.82, 40.31, 42.88, 45.53, 48.25, 51.05, 53.92, 56.87, 59.87, 62.94, 66.07, 69.26, 72.50,
		75.80, 79.14, 82.53, 85.95, 89.42, 92.91, 96.44, 100.00, 103.58, 107.18, 110.80, 114.43, 118.07, 121.72, 125.38, 129.03, 132.68, 136.33, 139.97, 143.60, 147.21, 150.81, 154.39, 157.95, 161.48, 164.99, 168.47, 171.92, 175.34,
		178.72, 182.07, 185.38, 188.65, 191.88, 195.06, 198.20, 201.30, 204.34, 207.34, 210.29, 213.19, 216.04, 218.84, 221.58, 224.28, 226.91, 229.49, 232.02, 234.49, 236.91, 239.27, 241.57, 243.82, 246.01, 248.14, 250.21, 252.23, 254.19,
		256.10, 257.95, 259.74, 261.47}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for fluo  F1 Daylight 6430K
func fluoF1_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.87, 2.36, 2.94, 3.47, 5.17, 19.49, 6.13, 6.24, 7.01, 7.79, 8.56, 43.67, 16.94, 10.72, 11.35, 11.89, 12.37, 12.75, 13.00, 13.15, 13.23, 13.17, 13.13, 12.85, 12.52,
		12.20, 11.83, 11.50, 11.22, 11.05, 11.03, 11.18, 11.53, 27.74, 17.05, 13.55, 14.33, 15.01, 15.52, 18.29, 19.55, 15.48, 14.91, 14.15, 13.22, 12.19, 11.12, 10.03, 8.95, 7.96, 7.02, 6.20, 5.42, 4.73, 4.15, 3.64, 3.20, 2.81,
		2.47, 2.18, 1.93, 1.72, 1.67, 1.43, 1.29, 1.19, 1.08, 0.96, 0.88, 0.81, 0.77, 0.75, 0.73, 0.68, 0.69, 0.64, 0.68, 0.69, 0.61, 0.52, 0.43, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for fluo  F2 Cool white 4230K
func fluoF2_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.18, 1.48, 1.84, 2.15, 3.44, 15.69, 3.85, 3.74, 4.19, 4.62, 5.06, 34.98, 11.81, 6.27, 6.63, 6.93, 7.19, 7.40, 7.54, 7.62, 7.65, 7.62, 7.62, 7.45, 7.28, 7.15, 7.05, 7.04, 7.16, 7.47, 8.04, 8.88, 10.01, 24.88, 16.64, 14.59, 16.16, 17.60, 18.62, 21.47, 22.79, 19.29, 18.66, 17.73, 16.54, 15.21, 13.80, 12.36, 10.95, 9.65, 8.40, 7.32, 6.31, 5.43, 4.68, 4.02, 3.45,
		2.96, 2.55, 2.19, 1.89, 1.64, 1.53, 1.27, 1.10, 0.99, 0.88, 0.76, 0.68, 0.61, 0.56, 0.54, 0.51, 0.47, 0.47, 0.43, 0.46, 0.47, 0.40, 0.33, 0.27, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for fluo  F3 White 3450K
func fluoF3_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.82, 1.02, 1.26, 1.44, 2.57, 14.36, 2.70, 2.45, 2.73, 3.00, 3.28, 31.85, 9.47, 4.02, 4.25, 4.44, 4.59, 4.72, 4.80, 4.86, 4.87, 4.85, 4.88, 4.77, 4.67, 4.62, 4.62, 4.73, 4.99, 5.48, 6.25,
		7.34, 8.78, 23.82, 16.14, 14.59, 16.63, 18.49, 19.95, 23.11, 24.69, 21.41, 20.85, 19.93, 18.67, 17.22, 15.65, 14.04, 12.45, 10.95, 9.51, 8.27, 7.11, 6.09, 5.22, 4.45, 3.80, 3.23, 2.75, 2.33, 1.99, 1.70, 1.55,
		1.27, 1.09, 0.96, 0.83, 0.71, 0.62, 0.54, 0.49, 0.46, 0.43, 0.39, 0.39, 0.35, 0.38, 0.39, 0.33, 0.28, 0.21, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for fluo F4 Warm white 2940K
func fluoF4_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.57, 0.70, 0.87, 0.98, 2.01, 13.75, 1.95, 1.59, 1.76, 1.93, 2.10, 30.28, 8.03, 2.55, 2.70, 2.82, 2.91, 2.99, 3.04, 3.08, 3.09, 3.09, 3.14, 3.06, 3.00, 2.98, 3.01,
		3.14, 3.41, 3.90, 4.69, 5.81, 7.32, 22.59, 15.11, 13.88, 16.33, 18.68, 20.64, 24.28, 26.26, 23.28, 22.94, 22.14, 20.91, 19.43, 17.74, 16.00, 14.42, 12.56, 10.93, 9.52, 8.18, 7.01, 6.00, 5.11, 4.36, 3.69, 3.13, 2.64,
		2.24, 1.91, 1.70, 1.39, 1.18, 1.03, 0.88, 0.74, 0.64, 0.54, 0.49, 0.46, 0.42, 0.37, 0.37, 0.33, 0.35, 0.36, 0.31, 0.26, 0.19, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for fluo F5 Daylight 6350K
func fluoF5_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.87, 2.35, 2.92, 3.45, 5.10, 18.91, 6.00, 6.11, 6.85, 7.58, 8.31, 40.76, 16.06, 10.32, 10.91, 11.40, 11.83, 12.17, 12.40, 12.54, 12.58, 12.52, 12.47, 12.20, 11.89,
		11.61, 11.33, 11.10, 10.96, 10.97, 11.16, 11.54, 12.12, 27.78, 17.73, 14.47, 15.20, 15.77, 16.10, 18.54, 19.50, 15.39, 14.64, 13.72, 12.69, 11.57, 10.45, 9.35, 8.29, 7.32, 6.41, 5.63, 4.90, 4.26,
		3.72, 3.25, 2.83, 2.49, 2.19, 1.93, 1.71, 1.52, 1.43, 1.26, 1.13, 1.05, 0.96, 0.85, 0.78, 0.72, 0.68, 0.67, 0.65, 0.61, 0.62, 0.59, 0.62, 0.64, 0.55, 0.47, 0.40, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for fluo F6 Lite white 4150K
func fluoF6_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.05, 1.31, 1.63, 1.90, 3.11, 14.8, 3.43, 3.30, 3.68, 4.07, 4.45, 32.61, 10.74, 5.48, 5.78, 6.03, 6.25, 6.41, 6.52, 6.58, 6.59, 6.56, 6.56, 6.42, 6.28, 6.20, 6.19, 6.30, 6.60, 7.12, 7.94, 9.07, 10.49, 25.22, 17.46, 15.63, 17.22, 18.53,
		19.43, 21.97, 23.01, 19.41, 18.56, 17.42, 16.09, 14.64, 13.15, 11.68, 10.25, 8.96, 7.74, 6.69, 5.71, 4.87, 4.16, 3.55, 3.02, 2.57, 2.20, 1.87, 1.60, 1.37, 1.29, 1.05, 0.91, 0.81, 0.71, 0.61, 0.54, 0.48, 0.44,
		0.43, 0.40, 0.37, 0.38, 0.35, 0.39, 0.41, 0.33, 0.26, 0.21, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for fluo F7 D65 Daylight simulator 6500K
func fluoF7_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 2.56, 3.18, 3.84, 4.53, 6.15, 19.37, 7.37, 7.05, 7.71, 8.41, 9.15, 44.14, 17.52, 11.35, 12.00, 12.58, 13.08, 13.45, 13.71, 13.88, 13.95, 13.93, 13.82, 13.64, 13.43, 13.25, 13.08, 12.93, 12.78, 12.60,
		12.44, 12.33, 12.26, 29.52, 17.05, 12.44, 12.58, 12.72, 12.83, 15.46, 16.75, 12.83, 12.67, 12.43, 12.19, 11.89, 11.60, 11.35, 11.12, 10.95, 10.76, 10.42, 10.11, 10.04, 10.02, 10.11, 9.87, 8.65, 7.27, 6.44, 5.83, 5.41,
		5.04, 4.57, 4.12, 3.77, 3.46, 3.08, 2.73, 2.47, 2.25, 2.06, 1.90, 1.75, 1.62, 1.54, 1.45, 1.32, 1.17, 0.99, 0.81, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for fluo F8 D50 simulator Sylvania F40 Design 5000K
func fluoF8_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.21, 1.5, 1.81, 2.13, 3.17, 13.08, 3.83, 3.45, 3.86, 4.42, 5.09, 34.1, 12.42, 7.68, 8.60, 9.46, 10.24, 10.84, 11.33, 11.71, 11.98, 12.17, 12.28, 12.32, 12.35, 12.44, 12.55, 12.68, 12.77, 12.72,
		12.60, 12.43, 12.22, 28.96, 16.51, 11.79, 11.76, 11.77, 11.84, 14.61, 16.11, 12.34, 13.61, 13.87, 14.07, 14.20, 14.16, 14.13, 14.34, 14.50, 14.46, 14.00, 12.58, 10.99, 9.98, 9.22, 8.62, 8.07, 7.39, 6.71, 6.16, 5.63, 5.03, 4.46, 4.02, 3.66,
		3.36, 3.09, 2.85, 2.65, 2.51, 2.37, 2.15, 1.89, 1.61, 1.32, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for fluo F9 Cool white deluxe 4150K
func fluoF9_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.9, 1.12, 1.36, 1.60, 2.59, 12.8, 3.05, 2.56, 2.86, 3.30, 3.82, 32.62, 10.77, 5.84, 6.57, 7.25, 7.86, 8.35, 8.75, 9.06, 9.31, 9.48, 9.61, 9.68, 10.04, 10.26, 10.48, 10.63, 10.76, 10.96,
		11.18, 27.71, 16.29, 12.28, 12.74, 13.21, 13.65, 16.57, 18.14, 14.55, 14.65, 14.66, 14.61, 14.50, 14.39, 14.40, 14.47, 14.62, 14.72, 14.55, 14.4, 14.58, 14.88, 15.51, 15.47, 13.20, 10.57, 9.18, 8.25, 7.57, 7.03,
		6.35, 5.72, 5.25, 4.80, 4.29, 3.80, 3.43, 3.12, 2.86, 2.64, 2.43, 2.26, 2.14, 2.02, 1.83, 1.61, 1.38, 1.12, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for fluo F10 Philips TL85 - 5000K
func fluoF10_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.11, 0.63, 0.62, 0.57, 1.48, 12.16, 2.12, 2.70, 3.74, 5.14, 6.75, 34.39, 14.86, 10.4, 10.76, 10.11, 9.27, 8.29, 7.29, 7.91, 16.64, 16.73, 10.44, 5.94, 3.34, 2.35, 1.88, 1.59, 1.47,
		1.80, 5.71, 40.98, 73.69, 33.61, 8.24, 3.38, 2.47, 4.86, 11.45, 14.79, 12.16, 8.97, 6.52, 8.81, 44.12, 34.55, 12.09, 12.15, 10.52, 4.43, 1.95, 2.19, 3.19, 2.77, 2.29, 2.00, 1.52, 1.35, 1.47, 1.79, 1.74, 1.02, 1.14,
		3.32, 4.49, 2.05, 0.49, 0.24, 0.21, 0.21, 0.24, 0.24, 0.21, 0.17, 0.21, 0.22, 0.17, 0.12, 0.09, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for fluo F11 Philips TL84 4150K
func fluoF11_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.91, 0.63, 0.46, 0.37, 1.29, 12.68, 1.59, 1.79, 2.46, 3.38, 4.49, 33.94, 12.13, 6.95, 7.19, 7.12, 6.72, 6.13, 5.46, 4.79, 5.66, 14.29, 14.96, 8.97, 4.72, 2.33, 1.47, 1.10, 0.89, 0.83, 1.18, 4.90, 39.59,
		72.84, 32.61, 7.52, 2.83, 1.96, 1.67, 4.43, 11.28, 14.76, 12.73, 9.74, 7.33, 9.72, 55.27, 42.58, 13.18, 13.16, 12.26, 5.11, 2.07, 2.34, 3.58, 3.01, 2.48, 2.14, 1.54, 1.33, 1.46, 1.94, 2.00, 1.20, 1.35, 4.10, 5.58,
		2.51, 0.57, 0.27, 0.23, 0.21, 0.24, 0.24, 0.20, 0.24, 0.32, 0.26, 0.16, 0.12, 0.09, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for fluo F12 Philips TL83 3000K
func fluoF12_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.96, 0.64, 0.45, 0.33, 1.19, 12.48, 1.12, 0.94, 1.08, 1.37, 1.78, 29.05, 7.90, 2.65, 2.71, 2.65, 2.49, 2.33, 2.10, 1.91, 3.01, 10.83, 11.88, 6.88, 3.43, 1.49, 0.92, 0.71, 0.60, 0.63, 1.10, 4.56, 34.4, 65.40, 29.48,
		7.16, 3.08, 2.47, 2.27, 5.09, 11.96, 15.32, 14.27, 11.86, 9.28, 12.31, 68.53, 53.02, 14.67, 14.38, 14.71, 6.46, 2.57, 2.75, 4.18, 3.44, 2.81, 2.42, 1.64, 1.36, 1.49, 2.14, 2.34, 1.42, 1.61, 5.04, 6.98, 3.19, 0.71, 0.30, 0.26, 0.23, 0.28, 0.28, 0.21,
		0.17, 0.21, 0.19, 0.15, 0.10, 0.05, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for HMI lamp studio "Osram" 4800K (for film, spectacle, studio...)
func hmi_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 9.66, 11.45, 13.24, 14.93, 16.63, 17.90, 19.20, 20.12, 21.03, 23.84, 26.65, 26.38, 26.12, 26.52, 27.92, 31.15, 34.37, 34.98, 35.61, 35.71, 35.81, 34.90, 34.02, 34.06, 34.08, 34.68, 35.28, 34.72, 34.20, 33.63,
		33.05, 34.70, 36.35, 38.01, 39.48, 37.29, 35.10, 36.22, 37.28, 38.76, 40.24, 39.56, 38.90, 39.35, 39.79, 38.90, 38.01, 38.05, 38.10, 37.45, 36.64, 35.82, 35.00, 35.06, 33.13, 33.85, 34.55, 35.26, 35.77, 34.92,
		34.09, 33.40, 32.72, 32.08, 31.45, 26.83, 22.23, 21.50, 20.79, 21.41, 22.03, 11.01, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for GTI lamp : Graphiclite & ColorMatch for Photography 5000K
func gti_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 3.26, 4.71, 6.17, 12.71, 19.27, 20.53, 21.80, 19.15, 16.53, 28.25, 39.97, 48.52, 57.06, 43.66, 30.27, 30.22, 30.16, 31.48, 32.98, 34.01, 35.04, 35.83, 36.62, 37.12, 37.62, 37.99, 38.19, 38.29, 38.48,
		38.82, 39.16, 45.40, 51.63, 51.83, 62.04, 52.41, 42.80, 42.95, 43.09, 45.64, 48.20, 46.23, 44.27, 43.74, 43.22, 43.30, 43.41, 43.10, 42.78, 42.03, 41.29, 40.29, 39.29, 37.89, 36.58, 34.92, 33.27, 31.47, 29.68, 27.90,
		26.13, 24.55, 22.98, 21.42, 19.86, 18.40, 16.92, 14.46, 13.99, 12.36, 11.73, 5.86, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for JudgeIII Lamp D50
func judgeIII_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 4.08, 4.25, 4.43, 6.90, 9.40, 9.75, 10.11, 9.30, 8.54, 14.90, 21.16, 26.01, 30.83, 24.90, 19.00, 19.00, 19.00, 19.56, 20.13, 20.28, 20.44, 20.64, 20.85, 21.05, 21.24, 21.65, 22.11, 22.85, 23.58, 24.00, 24.43,
		27.75, 31.27, 33.90, 36.59, 30.90, 25.32, 25.05, 24.76, 26.03, 27.31, 25.90, 24.48, 23.85, 23.29, 23.10, 22.94, 23.24, 23.53, 24.02, 24.52, 23.80, 23.13, 24.51, 25.76, 27.90, 29.15, 24.70, 20.25, 16.60, 12.98, 11.63, 10.27, 9.30, 8.34,
		7.60, 6.91, 6.25, 5.67, 5.15, 4.68, 2.34, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for Solux lamp : 3500K
func solux3500_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{0.5268, 0.93, 1.3278, 1.51, 1.6987, 2.65, 3.6100, 3.80, 3.9927, 6.08, 8.1680, 11.02, 13.863, 15.66, 17.4600, 18.78, 20.130, 21.43, 22.749, 24.02, 25.290, 27.40, 29.504, 31.77, 34.031, 36.35, 38.672, 40.55, 42.426, 44.15, 45.865, 47.37, 48.879,
		49.71, 50.531, 51.2, 51.872, 51.9, 51.928, 52.97, 54.015, 55.93, 57.846, 60.25, 62.650, 64.36, 66.065, 66.72, 67.369, 68.81, 70.260, 71.37, 72.487, 72.53, 72.578, 72.51, 72.447, 72.46, 72.471,
		72.76, 73.047, 74.25, 75.449, 76.5, 77.543, 78.79, 80.040, 80.72, 81.394, 82.12, 82.840, 83.23, 83.614, 83.36, 83.100, 82.36, 81.615, 80.11, 78.606, 75.91, 73.221, 69.61, 66.006, 62.43, 58.844, 56.07, 53.292,
		51.07, 48.839, 46.93, 45.013, 43.54, 42.070, 40.61, 39.150, 37.79, 36.425}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for Solux lamp : 4100K
func solux4100_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{0.5717, 0.70, 0.8286, 1.16, 1.522, 2.01, 2.384, 3.45, 4.57, 6.46, 8.4548, 11.31, 14.205, 16.10, 17.949, 19.51, 21.068, 22.60, 24.197, 25.37, 26.566, 28.15, 29.742, 30.90, 32.060, 33.26, 34.481, 34.80, 35.130, 35.42, 35.697, 36.20, 36.763,
		37.90, 39.004, 40.26, 41.494, 43.10, 44.690, 45.80, 46.900, 47.45, 47.885, 47.75, 47.635, 47.00, 46.410, 46.22, 46.058, 46.70, 47.344, 48.65, 50.005, 51.02, 52.045, 53.55, 55.075, 55.98, 56.823,
		56.85, 56.884, 56.15, 55.523, 54.60, 53.732, 52.55, 51.425, 50.30, 49.1830, 48.76, 48.273, 48.22, 48.169, 49.92, 49.915, 51.90, 53.099, 54.95, 56.852, 58.45, 60.090, 61.67, 63.2530, 63.55, 63.834, 63.55, 63.468,
		62.40, 61.373, 59.75, 58.1810, 56.25, 54.395, 51.90, 49.496, 47.05, 44.620}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for Solux lamp : near Daylight (for example  "musée d'Orsay..") - 4700K
func solux4700_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{0.4590, 0.83, 1.2011, 1.53, 1.8647, 2.15, 2.5338, 3.06, 3.5809, 3.99, 4.4137, 4.82, 5.2228, 5.63, 6.0387, 6.53, 6.9944, 7.55, 8.0266, 8.475, 8.9276, 8.90, 9.7840, 10.20, 10.6390, 11.00, 11.3600, 11.75, 12.1340, 12.36, 12.5880, 12.74, 12.8790,
		13.07, 13.2560, 13.38, 13.5220, 13.41, 13.3070, 13.35, 13.3990, 13.37, 13.3420, 13.39, 13.4220, 13.65, 13.2710, 13.25, 13.2330, 13.12, 13.0110, 12.93, 12.8470, 12.805, 12.7630, 12.66, 12.5760, 12.563, 12.5490,
		12.59, 12.6330, 12.617, 12.6010, 12.616, 12.6310, 12.6275, 12.6240, 12.70, 12.7710, 12.776, 12.7810, 12.786, 12.7950, 12.74, 12.6850, 12.64, 12.5950, 12.55, 12.5420, 12.43, 12.3180, 12.07, 11.8340, 11.72, 11.6190, 11.55, 11.5020,
		11.32, 11.1510, 11.05, 10.9530, 10.80, 10.6550, 10.495, 10.4390, 10.31, 10.1790}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for Solux lamp : near Daylight 4400K - test National Gallery
func ng_Solux4700_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 139, 152, 170, 185, 202, 223, 240, 257, 270, 282, 293, 305, 317, 329, 342, 355, 367, 378, 387, 395, 401, 405, 408, 411, 414, 415, 416, 415, 414, 414, 416, 419, 423, 427, 432, 437, 442, 447, 452,
		456, 461, 467, 475, 483, 488, 490, 491, 490, 487, 485, 481, 477, 472, 466, 461, 455, 449, 442, 434, 427, 419, 411, 403, 395, 386, 377, 367, 359, 351, 343, 335, 327, 322, 316, 312, 306, 305, 301, 299, 299, 298,
		0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.0}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for LED LSI Lumelex 2040 - test National Gallery
func ng_LEDLSI2040_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.5, 1.2, 0.5, 0.7, 0.6, 1.6, 1.7, 7.0, 16.6, 35.5, 64, 106, 162.5, 230.5, 272.2, 249, 213.4, 214, 227.6, 231.9, 233, 235.2, 241.4, 253.7, 270.3, 288.5, 306.2, 322.3, 337.6, 352.5, 367.2, 381.7, 395.9, 409.6, 416.2, 423.2, 429.7, 435.8, 442.8,
		451.7, 464.2, 480.3, 501, 526.3, 555.9, 587.5, 625.4, 655.1, 681.7, 705.3, 721.5, 728.5, 729, 719.8, 702.5, 676.7, 646.2, 611.5, 571.7, 530.3, 488.3, 445.9, 404, 365.2, 326.8, 290.8, 257.6, 226.9, 199.8, 175.2, 154.2, 133.8, 116.4, 101.5, 88.5, 76.6, 67.3, 57.9, 50.7, 44.2, 38.2,
		0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.0}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for LED CRS SP12 WWMR16 - test National Gallery
func ng_CRSSP12WWMR16_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0., 0., 0., 0., 0., 0.14, 0.33, 1.31, 3.34, 7.9, 17.4, 36.5, 72.6, 145.4, 260.5, 359.2, 365.3, 303.1, 256.1, 221.7, 193.6, 185.8, 191.4, 207.3, 232.8, 257.5, 285.1, 310.5, 333.4, 351.5, 368.8, 383.7, 398.8, 411.6, 424.7, 435.6, 447.9, 459.7, 471.7,
		484.6, 497.9, 512.3, 531.1, 548.9, 567.9, 587.5, 608.6, 625.3, 640.1, 648.6, 654.0, 654.3, 647.2, 633.9, 616.1, 590.5, 561.5, 526.5, 494.8, 457.9, 420.8, 382.4, 347.3, 309.9, 280.5, 249.2, 220.0, 194.9, 170.8, 149.1, 130.0, 112.3, 97.5, 84.9, 73.2, 63.1, 54.1, 45.6, 39.0, 32.6, 27.4,
		0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.0}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for Flash daylight 5500K (Leica...)
func flash5500_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{27.77, 29.17, 30.58, 32.02, 33.47, 33.00, 32.53, 35.28, 38.04, 49.46, 60.88, 64.68, 68.48, 69.99, 71.51, 69.68, 67.85, 76.70, 85.54, 91.74, 97.93, 99.17, 100.41, 100.14, 99.86, 101.28, 102.70, 100.37, 98.04, 99.35, 100.65,
		100.66, 100.67, 100.32, 99.97, 102.08, 104.20, 103.15, 102.09, 102.53, 102.96, 101.48, 100.00, 98.61, 97.22, 97.49, 97.76, 94.60, 91.44, 92.94, 94.44, 94.80, 95.16, 94.70, 94.25, 92.36, 90.48, 91.42, 92.36, 90.63,
		88.89, 89.62, 90.36, 92.18, 94.00, 92.00, 90.00, 84.86, 79.72, 81.30, 82.88, 83.88, 84.89, 77.58, 70.27, 74.80, 82.19, 85.03, 78.47, 71.91, 62.37, 52.82, 64.39, 75.96, 73.91, 71.85, 72.41, 72.97, 70.17, 67.38, 63.07,
		58.75, 61.89, 65.02, 66.68, 68.34}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for Flash daylight 6000K (Canon, Pentax, Olympus,...Standard)
func flash6000_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{36.00, 37.18, 38.36, 40.36, 42.35, 41.77, 41.19, 43.79, 46.40, 59.24, 72.09, 76.15, 80.22, 81.47, 82.71, 80.12, 77.52, 86.54, 95.56, 101.70, 107.85, 108.66, 109.46, 108.57, 107.68, 108.65, 109.61, 106.63, 103.65, 104.42,
		105.19, 104.79, 104.39, 103.45, 102.51, 104.28, 106.05, 104.68, 103.31, 103.42, 103.54, 101.77, 100.00, 98.38, 96.75, 96.74, 96.72, 93.30, 89.89, 90.92, 91.95, 91.99, 92.03, 91.30, 90.57, 88.51, 86.45, 86.96,
		87.47, 85.66, 83.85, 84.21, 84.57, 85.95, 87.32, 85.31, 83.30, 78.66, 74.03, 75.24, 76.45, 77.68, 78.91, 72.13, 65.35, 69.67, 73.98, 76.68, 79.39, 73.29, 67.19, 58.19, 49.19, 59.98, 70.77, 68.91, 67.05, 67.55, 68.05,
		65.47, 62.88, 58.89, 54.90, 57.81, 60.72, 62.25, 63.78}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

//spectral data for Flash daylight 6500K (Nikon, Minolta, Panasonic, Sony...)
func flash6500_spect(wavelength, m1, m2, temp float64) float64 {
	var val = [97]float64{44.86, 45.72, 46.59, 49.16, 51.74, 50.83, 49.92, 52.26, 54.60, 68.65, 82.69, 87.06, 91.42, 92.39, 93.37, 90.00, 86.63, 95.72, 104.81, 110.88, 116.96, 117.36, 117.76, 116.29, 114.82, 115.35, 115.89, 112.33, 108.78, 109.06,
		109.33, 108.56, 107.78, 106.28, 104.78, 106.23, 107.68, 106.04, 104.40, 104.22, 104.04, 102.02, 100.00, 98.17, 96.34, 96.06, 95.79, 92.24, 88.69, 89.35, 90.02, 89.81, 89.61, 88.66, 87.71, 85.51, 83.30, 83.51, 83.72,
		81.88, 80.05, 80.14, 80.24, 81.27, 82.30, 80.31, 78.31, 74.03, 69.74, 70.69, 71.63, 73.00, 74.37, 68.00, 61.62, 65.76, 69.91, 72.51, 75.11, 69.36, 63.61, 55.02, 46.43, 56.63, 66.83, 65.11, 63.40, 63.86, 64.32, 61.90, 59.47,
		55.72, 51.97, 54.72, 57.46, 58.89, 60.33}

	wlm := int((wavelength - 350.) / 5.)
	return val[wlm]
}

// Bradford transform between illuminants
var d65_d50 = [3][3]float64{{0.9555766, -0.0230393, 0.0631636},
	{-0.0282895, 1.0099416, 0.0210077},
	{0.0122982, -0.0204830, 1.3299098}}

var d50_d65 = [3][3]float64{{1.0478112, 0.0228866, -0.0501270},
	{0.0295424, 0.9904844, -0.0170491},
	{-0.0092345, 0.0150436, 0.7521316}}
