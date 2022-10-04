package main

import (
	"crypto"
	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	_ "golang.org/x/crypto/md4"
	_ "golang.org/x/crypto/ripemd160"
	_ "golang.org/x/crypto/sha3"
	_ "golang.org/x/crypto/blake2s"
	_ "golang.org/x/crypto/blake2b"

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
	hashType string
	wordlistFile string
	encoded string

	help bool
	listHashes bool
	sout = os.Stdout
	asciiart = `
  _   _    __    ___  _  _  ___  _   _  ____  ____
 ( )_( )  /__\  / __)( )/ )/ __)( )_( )( ___)(  _ \
  ) _ (  /(__)\( (__  )  ( \__ \ ) _ (  )__)  )   /
 (_) (_)(__)(__)\___)(_)\_)(___/(_) (_)(____)(_)\_)

	`
	helptext = `
    Usage:
        |  --hash  : The hash type (MD5, SHA1, etc.)
        |  --wlist : The wordlist to bruteforce
				|  --list  : List compatible hashes

    Examples:
        $ hacksher --wlist=/usr/share/wordlists/rockyou.txt --hash=sha256 4980b1f29fa32ff18c95d0ed931fd48e1ad43a729251d6eddb3cece705ed4d05
        $ hacksher --wlist /usr/share/wordlists/rockyou.txt --hash md5 63e780c3f321d13109c71bf81805476e
        $ hacksher --wlist=myWordlist.txt --hash=sha1 myHashFile.txt
      	`

	avaliableHashes = `
    Those are the avaliable hashes that i can crack:
        - MD4
        - MD5
        - SHA1
        - SHA224
        - SHA256
        - SHA384
        - SHA512
        - RIPEMD160
        - SHA3-224
        - SHA3-256
        - SHA3-384
        - SHA3-512
        - SHA512-224
        - SHA512-256
        - BLAKE2s-256
        - BLAKE2b-256
        - BLAKE2b-384
        - BLAKE2b-512
	`
)

func main() {
	//hash, encoded, wordlistFile := strings.ToLower(os.Args[1]), os.Args[2], os.Args[3]

	flagset := flag.NewFlagSet("Hacksher", flag.ExitOnError)

	flagset.StringVar(&hashType, "hash", "", "The hash type to bruteforce")
	flagset.StringVar(&wordlistFile, "wlist", "", "The wordlist to bforce hash")

	flagset.BoolVar(&listHashes, "l", false, "List compatible hashes")
	flagset.BoolVar(&listHashes, "list", false, "List compatible hashes")

	flagset.BoolVar(&help, "h", false, "Hacksher help")
	flagset.BoolVar(&help, "help", false, "Hacksher help")

	flagset.Parse(os.Args[1:])

	if len(os.Args) == 1 {
		fmt.Println(asciiart + helptext)
		os.Exit(0)
	} else {
		encoded = os.Args[len(os.Args) - 1]
	}

	if help {
		fmt.Println(asciiart + helptext)
		os.Exit(0)
	}

	if listHashes {
		fmt.Println(asciiart + avaliableHashes)
		os.Exit(0)
	}

	if !(strings.Contains(strings.Join(os.Args[1:], " "), "type")) && !(strings.Contains(strings.Join(os.Args[1:], " "), "wlist")) {
		fmt.Println(asciiart + helptext)
		os.Exit(1)
	}

	fmt.Println(asciiart)

	if _, err := os.Stat(encoded); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("[*] Hmmm, %s does not look like a file, so i will assume it's text\n", encoded)
	}

	if _, err := os.Stat(wordlistFile); errors.Is(err, os.ErrNotExist) {
		panic("You must provide a valid wordlist")
	}

	wordlist, err := ioutil.ReadFile(wordlistFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("[-] Hang on, we're breaking your hash...")
	fmt.Println("[-] Trying to break", encoded, "with", len(strings.Split(string(wordlist), "\n")), "Words...")

		for _, word := range strings.Split(string(wordlist), "\n") {
			var hasher hash.Hash

			switch hashType {
			case "md4":
				hasher = crypto.MD4.New()
			case "md5":
				hasher = crypto.MD5.New()
			case "sha1":
				hasher = crypto.SHA1.New()
			case "sha224":
				hasher = crypto.SHA224.New()
			case "sha256":
				hasher = crypto.SHA256.New()
			case "sha384":
				hasher = crypto.SHA384.New()
			case "sha512":
				hasher = crypto.SHA512.New()
			case "ripemd160":
				hasher = crypto.RIPEMD160.New()
			case "sha3_224":
				hasher = crypto.SHA3_224.New()
			case "sha3-256":
				hasher = crypto.SHA3_256.New()
			case "sha3-384":
				hasher = crypto.SHA3_384.New()
			case "sha3-512":
				hasher = crypto.SHA3_512.New()
			case "sha512-224":
				hasher = crypto.SHA512_224.New()
			case "sha512-256":
				hasher = crypto.SHA512_256.New()
			case "blake2s-256":
				hasher = crypto.BLAKE2s_256.New()
			case "blake2b-256":
				hasher = crypto.BLAKE2b_256.New()
			case "blake2b-384":
				hasher = crypto.BLAKE2b_384.New()
			case "blake2b-512":
				hasher = crypto.BLAKE2b_512.New()
			default:
				panic(fmt.Sprintf("I do not know '%s'", strings.ToUpper(hashType)))
			}

			if _, err := hasher.Write([]byte(word)); err != nil {
				panic(err)
			}

			hashed := hex.EncodeToString(hasher.Sum([]byte{}))
			if hashed == encoded {
				fmt.Println("[!] The decoded hash is:", word)
				break
			}
		}
}
