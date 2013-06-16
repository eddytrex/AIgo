package Fourier

import (
	"Matrix"
	"testing"
)

func TestFastFourierTrasnform(t *testing.T) {
	xi, _ := Matrix.FromFile("test/tDFT.txt")
	Xi, _ := Matrix.FromFile("test/tFFT.txt")

	Xip := Fourier.FFT(xi, 8)

	if !Matrix.Equal(Xi, Xip) {
		t.Errorf("The DFT has to be %v", Xi)
	}

}

func TestInverseFastFourierTrasnform(t *testing.T) {
	xi, _ := Matrix.FromFile("test/tDFT.txt")

	Xi, _ := Matrix.FromFile("test/tFFT.txt")

	xip := Fourier.IFFT(Xi, 8)

	if !Matrix.Equal(xi, xip) {
		t.Errorf("The DFT has to be %v", xi)
	}
}
