package gosom

import "testing"

func TestGaussian(t *testing.T) {
	ris := Gaussian(0.1, 0.2)
	if ris != 0.8824969025845955 {
		t.Fatalf("Expected %v but having %v\n", 0.8824969025845955, ris)
	}
}

func TestBubble(t *testing.T) {
	ris := Bubble(0.1, 0.2)
	if ris != 1 {
		t.Fatalf("Expected %v but having %v\n", 1, ris)
	}
	ris = Bubble(0.2, 0.1)
	if ris != 0 {
		t.Fatalf("Expected %v but having %v\n", 0, ris)
	}

}
