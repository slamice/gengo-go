/*
 * All code provided from the http://gengo.com site, such as API example code
 * and libraries, is provided under the New BSD license unless otherwise
 * noted. Details are below.
 *
 * New BSD License
 * Copyright (c) 2009-2012, Gengo, Inc.
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 * Redistributions of source code must retain the above copyright notice,
 * this list of conditions and the following disclaimer.
 * Redistributions in binary form must reproduce the above copyright notice,
 * this list of conditions and the following disclaimer in the documentation
 * and/or other materials provided with the distribution.
 * Neither the name of Gengo, Inc. nor the names of its contributors may
 * be used to endorse or promote products derived from this software
 * without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS
 * IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO,
 * THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
 * PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR
 * CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
 * PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
 * LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
 * NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

type Gengo struct {
	publickey  string
	privatekey string
	sandbox    bool
}

func ComputeHmac256(privatekey string, timestamp string) string {
	h := hmac.New(sha1.New, []byte(privatekey))
	h.Write([]byte(timestamp))
	return hex.EncodeToString(h.Sum(nil))
}

func signatureAndTimestamp(privatekey string, gengo Gengo) (signature string, ts string) {
	ts := time.Now().Unix()
	signature := hmacSha1Hex(gengo.privatekey, timestamp)
	return
}

func main() {
	signatureAndTimestamp("ZZl6Us6e70XmBhEIgHhfkqiXr}O@ja$GJKL653wz$[RybM=lv}J7IufEaCG4YjAh")
}

//    Get the current Unix epoch time as an integer
//    Insert the time as the value to a ‘ts’ key in your argument list
//    Calculate the SHA1 hash of the timestamp against your private key
//    Append the value of this hash to the argument list as a parameter named api_sig
