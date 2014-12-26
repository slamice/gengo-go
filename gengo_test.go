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
	t.Log("Hello %d\n", body2)
	fmt.Printf("%d", body2)
}
