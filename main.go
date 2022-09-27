package main

import (
	"crypto"
	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"encoding/hex"
	"hash"
	"fmt"
	"strings"
	"os"
	"io/ioutil"
	"flag"
	"errors"
)

var (
	hasher hash.Hash

	hashType string
	wordlistFile string
	encoded string
	encodedFile string

	clearcmd string
	help bool
	asciiart = `
  _   _    __    ___  _  _  ___  _   _  ____  ____
 ( )_( )  /__\  / __)( )/ )/ __)( )_( )( ___)(  _ \
  ) _ (  /(__)\( (__  )  ( \__ \ ) _ (  )__)  )   /
 (_) (_)(__)(__)\___)(_)\_)(___/(_) (_)(____)(_)\_)

	`
	helptext = `
    Usage:
        |  --hash  : The hash type (MD5, SHA1, etc.)
        |  --wlist : The hash type to bruteforce
        |  -f      : The hash file
        |  -t      : The hash text
	`
)

func main() {
	//hash, encoded, wordlistFile := strings.ToLower(os.Args[1]), os.Args[2], os.Args[3]

	flagset := flag.NewFlagSet("Hacksher", flag.ExitOnError)
	flagset.StringVar(&hashType, "hash", "", "The hash type to bruteforce")
	flagset.StringVar(&wordlistFile, "wlist", "", "The wordlist to bforce hash")
	flagset.StringVar(&encodedFile, "f", "", "The hash file")
	flagset.StringVar(&encoded, "t", "", "The hash text")
	flagset.BoolVar(&help, "h", false, "Hacksher help")

	flagset.Parse(os.Args[1:])

	if help {
		fmt.Println(asciiart + helptext)
	}

	if !(strings.Contains(strings.Join(os.Args[1:], " "), "f")) && !(strings.Contains(strings.Join(os.Args[1:], " "), "t")) {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if _, err := os.Stat(wordlistFile); errors.Is(err, os.ErrNotExist) {
		panic("You must provide a valid wordlist")
	}

	wordlist, err := ioutil.ReadFile(wordlistFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("[*] Hang on, we're breaking your hash...")
	fmt.Println("[-] Trying to break", encoded, "with", len(strings.Split(string(wordlist), "\n")), "Words...")

	for _, word := range strings.Split(string(wordlist), "\n") {
		if hashType == "md5" {
			hasher = crypto.MD5.New()
		} else if hashType == "sha1" {
			hasher = crypto.SHA1.New()
		} else if hashType == "sha256" {
			hasher = crypto.SHA256.New()
		} else if hashType == "sha512" {
			hasher = crypto.SHA512.New()
		} else {
			panic(fmt.Sprintf("I don't know '%s'", strings.ToUpper(hashType)))
		}

		if _, err := hasher.Write([]byte(word)); err != nil {
			panic(err)
		}

		hashed := hex.EncodeToString(hasher.Sum([]byte{}))
		if hashed == encoded {
			fmt.Printf("[!] the decoded hash is: %s\n", word)
			os.Exit(0)
		}
	}
}
