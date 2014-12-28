package gengo

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var (
	SandboxUrl    = "http://sandbox.api.gengo.com/v2/"
	ProductionUrl = "http://api.gengo.com/v2/"
)

type Gengo struct {
	Publickey  string
	Privatekey string
	Sandbox    bool
}

type GengoResponse struct {
	Err      *ErrorResponse
	Opstat   string                 `json:"opstat"`
	Response map[string]interface{} `json:"response"`
}

type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (g *Gengo) BaseURL() string {
	s := SandboxUrl

	if g.Sandbox == false {
		s = ProductionUrl
	}

	return s
}

// Create signature based on Private Key and Timestamp
func ComputeHmacSha1Hex(privatekey string, timestamp string) string {
	h := hmac.New(sha1.New, []byte(privatekey))
	h.Write([]byte(timestamp))
	return hex.EncodeToString(h.Sum(nil))
}

// Calculate timestamp and signature
func signatureAndTimestamp(privatekey string) (timestamp string, signature string) {
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	sig := ComputeHmacSha1Hex(privatekey, ts)
	return ts, sig
}

func getRequest(subpath string) (r *GengoResponse) {

	client := &http.Client{}

	ts, signature := signatureAndTimestamp(gengo.Privatekey)

	v := url.Values{}
	v.Set("api_key", gengo.Publickey)
	v.Set("api_sig", signature)
	v.Set("ts", ts)

	urlList := []string{gengo.BaseURL(), subpath, "?", v.Encode()}
	u := strings.Join(urlList, "")

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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(subpath)

	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(r.Opstat)

	return
}

func (g *Gengo) getAccountStats() (r *GengoResponse) {
	return getRequest("account/stats")
}

//    Get the current Unix epoch time as an integer
//    Insert the time as the value to a ‘ts’ key in your argument list
//    Calculate the SHA1 hash of the timestamp against your private key
//    Append the value of this hash to the argument list as a parameter named api_sig
