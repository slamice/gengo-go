package gengo

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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
	s := "http://sandbox.api.gengo.com/v2/"

	if g.sandbox == false {
		s = "http://api.gengo.com/v2/"
	}
	return s
}

func ComputeHmacSha1Hex(privatekey string, timestamp string) string {
	h := hmac.New(sha1.New, []byte(privatekey))
	h.Write([]byte(timestamp))
	return hex.EncodeToString(h.Sum(nil))
}

func signatureAndTimestamp(privatekey string) (timestamp string, signature string) {
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println(ts)
	sig := ComputeHmacSha1Hex(privatekey, ts)
	return ts, sig
}

func getRequest(subpath string) (body []byte) {

	client := &http.Client{}

	ts, signature := signatureAndTimestamp(gengo.privatekey)

	// Set URL values
	v := url.Values{}
	v.Set("api_key", gengo.publickey)
	v.Set("api_sig", signature)
	v.Set("ts", ts)

	u := gengo.BaseURL() + subpath + "?" + v.Encode()

	fmt.Println(u)

	req, err := http.NewRequest("GET", u, nil)

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
	return getRequest("account/stats")
}

//    Get the current Unix epoch time as an integer
//    Insert the time as the value to a ‘ts’ key in your argument list
//    Calculate the SHA1 hash of the timestamp against your private key
//    Append the value of this hash to the argument list as a parameter named api_sig
