package opaque

import (
	"testing"

	"github.com/bwesterb/go-ristretto"
)

func TestOpaque(t *testing.T) {
	Opaque()
}

func TestRegistrationInit(t *testing.T) {
	password := "YouCantSeeMe"

	var hashPrime ristretto.Point
	hashPrime.DeriveDalek([]byte(password))

	var r ristretto.Scalar
	r.Rand()

	var alpha ristretto.Point
	alpha.ScalarMult(&hashPrime, &r)

	var privU ristretto.Scalar
	var pubU ristretto.Point

	privU.Rand()
	pubU.ScalarMultBase(&privU)

	RegistrationInit("jerry-g", &alpha)
}
