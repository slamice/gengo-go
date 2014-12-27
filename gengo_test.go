package gengo

import (
	"fmt"
	"testing"
)

var (
	pubKey  = ""
	privKey = ""
	gengo   = Gengo{pubKey, privKey, false}
)

func TestGetAccountStats(t *testing.T) {
	body := gengo.getAccountStats()

	body2 := string(body)
	fmt.Printf(body2)
}
