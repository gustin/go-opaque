package opaque

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"testing"
)

func TestOpaque(t *testing.T) {
	Opaque()
}

func TestRegistrationInit(t *testing.T) {
	password := "YouCantSeeMe"

	var publicKey ristretto.Point

	secretKey.Rand()
	publicKey.ScalarMultBase(&secretKey)

	var hashPrime ristretto.Point
	hashPrime.DeriveDalek([]byte(password))

	RegistrationInit("jerry-g")

}
