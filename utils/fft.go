package utils

import "math"

func FFTDouble(data []float64) []float64 {
	n := len(data)
	complexData := make([]complex128, n)
	for i := 0; i < n; i++ {
		complexData[i] = complex(data[i], 0)
	}
	FFT(complexData)
	result := make([]float64, n)
	for i := 0; i < n; i++ {
		result[i] = real(complexData[i])
	}
	return result
}

func FFT(x []complex128) {
	n := len(x)
	if n <= 1 {
		return
	}

	bitReverse(x)

	for length := 2; length <= n; length <<= 1 {
		angle := -2 * math.Pi / float64(length)
		wlen := complex(math.Cos(angle), math.Sin(angle))

		for i := 0; i < n; i += length {
			w := complex(1.0, 0.0)
			half := length >> 1
			for j := 0; j < half; j++ {
				u := x[i+j]
				v := w * x[i+j+half]
				x[i+j] = u + v
				x[i+j+half] = u - v
				w *= wlen
			}
		}
	}
}

func IFFT(x []complex128) {
	n := len(x)
	if n <= 1 {
		return
	}

	for i := range x {
		x[i] = complex(real(x[i]), -imag(x[i]))
	}

	FFT(x)

	invN := 1.0 / float64(n)
	for i := range x {
		x[i] = complex(real(x[i]), -imag(x[i])) * complex(invN, 0)
	}
}

func bitReverse(x []complex128) {
	n := len(x)
	j := 0
	for i := 1; i < n; i++ {
		bit := n >> 1
		for ; j&bit != 0; bit >>= 1 {
			j &= ^bit
		}
		j |= bit

		if i < j {
			x[i], x[j] = x[j], x[i]
		}
	}
}
