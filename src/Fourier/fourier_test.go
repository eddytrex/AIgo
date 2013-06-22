package Fourier

import (
	"Matrix"
	"testing"
)

func TestFastFourierTrasnform(t *testing.T) {
	xi, _ := Matrix.FromFile("test/tDFT.txt")
	Xi, _ := Matrix.FromFile("test/tFFT.txt")

	Xip, _ := FFT(xi, 8)

	if !Matrix.AlmostEqual(Xi, Xip) {
		t.Errorf("The DFT has to be \n", Xi.ToString(), Xip.ToString())
	}

}

func TestInverseFastFourierTrasnform(t *testing.T) {
	xi, _ := Matrix.FromFile("test/tDFT.txt")

	Xi, _ := Matrix.FromFile("test/tFFT.txt")

	xip, _ := IFFT(Xi, 8)

	if !Matrix.AlmostEqual(xi, xip) {
		t.Errorf("The DFT has to be %v", xi.ToString(), xip.ToString())
	}
}

func BenchmarkFFT(b *testing.B) {
	xi, _ := Matrix.FromFile("test/BMFFT.txt")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		FFT(xi, 1024)
	}

}
