package gengo

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Gengo struct {
	publickey  string
	privatekey string
	sandbox    bool
}

type GengoError struct {
	err  string
	code int
	body string
}

func (g *Gengo) BaseURL() string {
	if g.sandbox == false {
		return "https://api.gengo.com/v2/"
	}
	return "https://sandbox.api.gengo.com/v2/"
}

func ComputeHmacSha1Hex(privatekey string, timestamp string) string {
	h := hmac.New(sha1.New, []byte(privatekey))
	h.Write([]byte(timestamp))
	return hex.EncodeToString(h.Sum(nil))
}

func signatureAndTimestamp(privatekey string) (timestamp string, signature string) {
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	sig := ComputeHmacSha1Hex(privatekey, ts)
	return sig, ts
}

func handleAuthentication() {

}

func getRequest(url string) (body []byte) {

	client := &http.Client{}

	var ts, signature = signatureAndTimestamp(gengo.privatekey)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func (g *Gengo) getAccountStats() (body []byte) {
	return getRequest(g.BaseURL() + "account/stats")
}

//    Get the current Unix epoch time as an integer
//    Insert the time as the value to a ‘ts’ key in your argument list
//    Calculate the SHA1 hash of the timestamp against your private key
//    Append the value of this hash to the argument list as a parameter named api_sig
