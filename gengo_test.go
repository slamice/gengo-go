package gengo

import (
	"os"
	"testing"
)

var (
	pubKey  = os.Getenv("")
	privKey = os.Getenv("")
	gengo   = Gengo{pubKey, privKey, true}
)
