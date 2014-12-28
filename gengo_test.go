package gengo

import (
	"testing"
)

var (
	PubKey  = ""
	PrivKey = ""
	gengo   = Gengo{PubKey, PrivKey, false}
)

func TestGetAccountStats(t *testing.T) {
	r := gengo.getAccountStats()
	t.Log(r.Response)
}
