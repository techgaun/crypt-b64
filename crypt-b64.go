package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
)

func usage(progname string) {
	fmt.Fprintf(os.Stderr, "Usage: %s <crypt_encoded_string>", progname)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usage(os.Args[0])
	}

	cryptCharset := "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	cryptStr := os.Args[1]

	cryptEncoder := base64.NewDecoder(base64.NewEncoding(cryptCharset).WithPadding(base64.NoPadding), strings.NewReader(cryptStr))
	decodedBuf := new(strings.Builder)
	_, err := io.Copy(decodedBuf, cryptEncoder)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred while decoding")
		os.Exit(2)
	}

	decodedStr := decodedBuf.String()
	stdBase64Str := base64.StdEncoding.EncodeToString([]byte(decodedStr))

	fmt.Fprintf(os.Stdout, "%s", stdBase64Str)
}
