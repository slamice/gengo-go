package gengo

import (
	"os"
	"testing"
)

var (
	pubKey  = os.Getenv("")
	privKey = os.Getenv("")
	gengo   = gengo{pubKey, privKey, true}
)

func TestGetAccountStats(t *testing.T) {
	stuff, err := gengo.getAccountStats()
	t.Log("Hello %d\n", err)
	t.Log("Hello %d\n", stuff)
}
