// Command loofah prints 2FA codes for supplied 2FA secret keys.
package main

// Large portions of code in this file are copied or adapted from
// github.com/rsc/2fa, whose license is reproduced below.

/*
Copyright (c) 2009 The Go Authors. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/atotto/clipboard"
)

var (
	fClip = flag.Bool("c", false, "also copy code to clipboard")
	f7    = flag.Bool("7", false, "generate 7-digit code")
	f8    = flag.Bool("8", false, "generate 8-digit code")
)

func printUsage() {
	fmt.Fprintf(os.Stderr, "usage: 2fa [-7] [-8] [-c]\n")
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("2fa: ")

	flag.Usage = printUsage
	flag.Parse()

	if flag.NArg() != 0 {
		printUsage()
		os.Exit(2)
	}

	size := 6
	if *f7 {
		size = 7
		if *f8 {
			log.Fatalf("cannot use -7 and -8 together")
		}
	} else if *f8 {
		size = 8
	}

	fmt.Fprintf(os.Stderr, "enter 2fa key: ")
	key, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatalf("error reading key: %v", err)
	}
	key = strings.Map(noSpace, key)
	key += strings.Repeat("=", -len(key)&7) // pad to 8 bytes

	raw, err := decodeKey(key)
	if err != nil {
		log.Fatalf("invalid key: %v", err)
	}

	code := fmt.Sprintf("%0*d", size, totp(raw, time.Now(), size))
	if *fClip {
		if err := clipboard.WriteAll(code); err != nil {
			log.Printf("error copying to clipboard: %s", err)
		}
	}
	fmt.Println(code)
}

func noSpace(r rune) rune {
	if unicode.IsSpace(r) {
		return -1
	}
	return r
}

func decodeKey(key string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(strings.ToUpper(key))
}

func hotp(key []byte, counter uint64, digits int) int {
	h := hmac.New(sha1.New, key)
	binary.Write(h, binary.BigEndian, counter)
	sum := h.Sum(nil)
	v := binary.BigEndian.Uint32(sum[sum[len(sum)-1]&0x0F:]) & 0x7FFFFFFF
	d := uint32(1)
	for i := 0; i < digits && i < 8; i++ {
		d *= 10
	}
	return int(v % d)
}

func totp(key []byte, t time.Time, digits int) int {
	return hotp(key, uint64(t.UnixNano())/30e9, digits)
}
