# Hacksher
This is a very nice hash bruteforcer written in Go

Those are the avaliable hashes that i can crack: <br/>
    - MD4 <br/>
    - MD5<br/>
    - SHA1<br/>
    - SHA224<br/>
    - SHA256<br/>
    - SHA384<br/>
    - SHA512<br/>
    - RIPEMD160<br/>
    - SHA3-224<br/>
    - SHA3-256<br/>
    - SHA3-384<br/>
    - SHA3-512<br/>
    - SHA512-224<br/>
    - SHA512-256<br/>
    - BLAKE2s-256<br/>
    - BLAKE2b-256<br/>
    - BLAKE2b-384<br/>
    - BLAKE2b-512<br/>

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

<img src="https://i.imgur.com/u6D4WNS.png"/>
