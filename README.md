## Go Encrypt

1. Encrypt your files with a pass phrase.
2. Share file over first channel with instructions to decrypt.
3. Share pass phrase over second channel.

### Requirements

- openssl

### Build

```go
go build
```

### Use

```
> echo "secret things" > secret_file.txt
> ./encrypt -f=secret_file.txt -p="dont forget the nickels"
openssl aes-256-cbc -d -a -pass "pass:XXX" -in secret_file.txt.enc -out secret_file.txt
```

Share the output and encrypted file on one channel.  Share the pass phrase on a second.

### Options

```
> ./encrypt
Usage of ./encrypt:
  -f="": file to encrypt
  -p="": password to decrypt
```
