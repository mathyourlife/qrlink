package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/skip2/go-qrcode"
)

func main() {
	// Read the url from the command prompt
	// todo: detect string from stdin instead of prompt
	url, err := readURL()
	if err != nil {
		panic(err)
	}

	// create qr code png file
	fname, err := writeTmp(url)
	if err != nil {
		panic(err)
	}
	// open the png file with the default application
	open(fname).Run()
}

// writeTmp take the url string and write it to a temp file
//
// todo: since ioutil.TempFile doesn't take a suffix, there
// is a chance of a file name collision when tmpfile gets
// moved to tmpfile.png
func writeTmp(url string) (string, error) {
	tmpfile, err := ioutil.TempFile("", "url")
	if err != nil {
		return "", err
	}
	tmpfile.Close()
	os.Rename(tmpfile.Name(), tmpfile.Name()+".png")

	err = qrcode.WriteFile(url, qrcode.Medium, 256, tmpfile.Name()+".png")
	return tmpfile.Name() + ".png", err
}

// readURL get a string from user input
func readURL() (string, error) {
	fmt.Print("qr: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	// Add trailing newline after user input to make a cleaner output
	fmt.Println()
	if err != nil {
		return "", err
	}
	return string(bytePassword), nil
}
