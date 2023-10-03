package main

import (
	"flag"
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	url := flag.String("url", "https://go.dev/", "Url that will be converted to qrcode.")
    size := flag.Int("size", 256, "Size of the png file.")
	name := flag.String("filename", "qrcode.png", "Name of the png file.")

	flag.Parse()
    
	generateQR(*url, *size, *name)
}

func generateQR(url string, size int, name string)  {
	err := qrcode.WriteFile(url, qrcode.Medium, size, name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Successfully generated QR code and saved as %v\n", name)
}