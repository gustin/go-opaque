package opaque

import (
	"fmt"
	"github.com/bwesterb/go-ristretto"
	"lukechampine.com/blake3"
	"testing"
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

	beta, v, publicKey := RegistrationInit("jerry-g", &alpha)
	fmt.Println(beta, v, publicKey)

	// U: upon receiving values beta and v, set the PRF output to
	// H(x, v, beta*v^{-r})

	// simplified:
	//  set the function output to H(x,v,beta^{1/r})
	fmt.Println("-) {{1/r}}:")
	fmt.Println("-) beta^{{1/r}}:")

	sub_beta := beta.ScalarMult(&beta, r.Inverse(&r))
	fmt.Println(sub_beta)

	hasher := blake3.New(64, nil)
	hasher.Write([]byte(password))
	hasher.Write(v.Bytes())
	hasher.Write(sub_beta.Bytes())
	var rwdU [64]byte
	hasher.Sum(rwdU[:0])

	fmt.Println("RwdU:", rwdU)
}
