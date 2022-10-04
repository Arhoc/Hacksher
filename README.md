# Hacksher
This is a very nice hash bruteforcer written in Go

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

## Usage
the usage is very simple, you just need to do
```
hacksher --wlist=/path/to/wordlist --hash=hashType <hash text> or <hash file>
```

## Example
```
hacksher --wlist=/usr/share/wordlists/rockyou.txt --hash=ripemd160 06877917a6451b23e0fd8cd976530ceb7225549f
``` 

```
hacksher --wlist=rockyou.txt --hash=sha512 hash.txt
```
there you go.

<img src="https://i.imgur.com/ymovtcz.png"/>
