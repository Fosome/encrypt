// encrypt.go
package main

import (
  "fmt"
  "flag"
  "os"
  "os/exec"
)

var filepath, password string

func init() {
  flag.StringVar(&filepath, "f", "", "file to encrypt")
  flag.StringVar(&password, "p", "", "password to decrypt file")

  flag.Parse()
}

func main() {
  if filepath == "" || password == "" {
    flag.Usage()
    os.Exit(1)
  }

  opensslCmd := exec.Command("openssl",
                              "aes-256-cbc",
                              "-v",
                              "-a",
                              "-salt",
                              "-pass",
                              "pass:" + password,
                              "-in",
                              filepath,
                              "-out",
                              filepath + ".enc")

  _, err := opensslCmd.Output()

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("openssl aes-256-cbc -d -a -pass \"pass:XXX\" -in", filepath + ".enc", "-out", filepath)
  }
}
