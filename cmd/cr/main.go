package main

import (
	"flag"
	"strings"
	"os"
	"fmt"
	"io"
	"bufio"
	"path/filepath"
)

//const passPhrase = "Cette faucille d'or dans le champ des etoiles"

func main() {
	encryptCmd := flag.NewFlagSet("encrypt", flag.ExitOnError)
	decryptCmd := flag.NewFlagSet("decrypt", flag.ExitOnError)
	if len(os.Args) < 2 {
		Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "encrypt":
		encryptCmd.Parse(os.Args[2:])
		if len(os.Args) < 4 {
			Usage()
			os.Exit(1)
		}
		args := encryptCmd.Args()
		path := args[0]
		passPhrase := args[1]
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		input, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}
		fmt.Println(input)
		cipher := encrypt(input, passPhrase)
		fmt.Println(cipher)
		fileName := filepath.Base(path)
		w, err := os.Create(fileName + ".enc")
		if err != nil {
			panic(err)
		}
		defer w.Close()
		writer := bufio.NewWriter(w)
		writer.Write(cipher)
		writer.Flush()

	case "decrypt":
		decryptCmd.Parse(os.Args[2:])
		if len(os.Args) < 4 {
			Usage()
			os.Exit(1)
		}
		args := decryptCmd.Args()
		path := args[0]
		passPhrase := args[1]
		fileName := filepath.Base(path)
		if !strings.HasSuffix(fileName, ".enc") {
			fmt.Println("File suffix should be .enc")
		}
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		input, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}
		decipher := decrypt(input, passPhrase)
		w, err := os.Create(fileName[:len(args[0]) - 4])
		if err != nil {
			panic(err)
		}
		defer w.Close()
		writer := bufio.NewWriter(w)
		writer.Write(decipher)
		writer.Flush()

	default:
		Usage()
	}
	fmt.Println("done")
}

func Usage() {
	fmt.Println("USAGE:")
	fmt.Println("cr encrypt <file> <passphrase>")
	fmt.Println("cr decrypt <file.enc> <passphrase>")
}

