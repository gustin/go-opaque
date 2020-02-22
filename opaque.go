package opaque

import (
	"fmt"
	"github.com/bwesterb/go-ristretto"
)

func RegistrationInit(username string, alpha *ristretto.Point) {
	fmt.Println("~-> Initial Registration ~-")
	fmt.Println(alpha)

	var secretKey ristretto.Scalar
	var publicKey ristretto.Point

	secretKey.Rand()
	publicKey.ScalarMultBase(&secretKey)

	// let k = Scalar::random(&mut cspring); // salt, private
	// let v: RistrettoPoint = RISTRETTO_BASEPOINT_POINT * k; // salt 2, public
	// let beta = alpha * k;

	var k ristretto.Scalar
	k.Rand()

	var v ristretto.Point // salt, private
	v.ScalarMultBase(&k)  // salt, public
	var beta ristretto.Point
	beta.ScalarMult(&v, &k)
}

func Opaque() string {
	fmt.Println("~-- OPAQUE --~")

	return ""
}
