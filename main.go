package main

import (
	"crypto"
	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var hasher hash.Hash
	hash, encoded, wordlistFile := strings.ToLower(os.Args[1]), os.Args[2], os.Args[3]

	if _, err := os.Stat(wordlistFile); errors.Is(err, os.ErrNotExist) {
		panic("You must provide a valid wordlist\nUsage: hacksher <hashType> <hash> <wordlist>\nhacksher md5 2f290dd624e33112d3c578b348b5d137")
	}

	wordlist, err := ioutil.ReadFile(wordlistFile)
	if err != nil {
		panic(err)
	}

	for _, word := range strings.Split(string(wordlist), "\n") {
		if hash == "md5" {
			hasher = crypto.MD5.New()
		} else if hash == "sha1" {
			hasher = crypto.SHA1.New()
		} else if hash == "sha256" {
			hasher = crypto.SHA256.New()
		} else if hash == "sha512" {
			hasher = crypto.SHA512.New()
		} else {
			panic(fmt.Sprintf("I don't know '%s'", strings.ToUpper(hash)))
		}

		if _, err := hasher.Write([]byte(word)); err != nil {
			panic(err)
		}

		hashed := hex.EncodeToString(hasher.Sum([]byte{}))
		if hashed == encoded {
			fmt.Printf("the decoded hash is: %s\n", word)
			os.Exit(0)
		}
	}
}
