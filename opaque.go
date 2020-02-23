package opaque

import (
	"fmt"
	"github.com/bwesterb/go-ristretto"
)

func RegistrationInit(username string, alpha *ristretto.Point) (ristretto.Point, ristretto.Point, ristretto.Point) {
	fmt.Println("~-> Initial Registration ~-")
	fmt.Println(alpha)

	var secretKey ristretto.Scalar
	var publicKey ristretto.Point

	secretKey.Rand()
	publicKey.ScalarMultBase(&secretKey)

	var k ristretto.Scalar
	k.Rand()

	var v ristretto.Point // salt, private
	v.ScalarMultBase(&k)  // salt, public
	var beta ristretto.Point
	beta.ScalarMult(&v, &k)

	return beta, v, publicKey
}

func Opaque() string {
	fmt.Println("~-- OPAQUE --~")

	return ""
}
