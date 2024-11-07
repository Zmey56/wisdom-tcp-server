package pow

import "testing"

func TestVerifyPoW(t *testing.T) {
	seed, _ := (PoWImpl{}).GenerateChallenge()
	proof := "example_proof"
	if (PoWImpl{}).VerifyPoW(seed, proof) {
		t.Error("Expected false for invalid proof")
	}
}
