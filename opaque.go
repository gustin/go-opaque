package opaque

import (
	"fmt"
	"github.com/bwesterb/go-ristretto"
)

func RegistrationInit(username string, alpha []uint32) {
	fmt.Println("~-> Initial Registration ~-")
}

func Opaque() string {
	fmt.Println("~-- OPAQUE --~")
	var secretKey ristretto.Scalar
	var publicKey ristretto.Point

	secretKey.Rand()
	publicKey.ScalarMultBase(&secretKey)

	return ""
}
