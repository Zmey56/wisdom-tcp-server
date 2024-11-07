package pow

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
)

const difficulty = 2 // Количество ведущих нулей в хэше, указывающее на сложность задачи

type PoW interface {
	GenerateChallenge() (string, string)
	VerifyPoW(seed, proof string) bool
}

type PoWImpl struct{}

func (p PoWImpl) GenerateChallenge() (string, string) {
	seed := fmt.Sprintf("%d", rand.Int63())
	return seed, fmt.Sprintf("%0*d", difficulty, 0)
}

func (p PoWImpl) VerifyPoW(seed, proof string) bool {
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(seed+proof)))
	expectedPrefix := fmt.Sprintf("%0*d", difficulty, 0)
	fmt.Printf("Verifying proof: %s | Expected prefix: %s | Hash: %s\n", proof, expectedPrefix, hash)
	return hash[:difficulty] == expectedPrefix
}
