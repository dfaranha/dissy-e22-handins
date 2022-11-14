package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func testSystemsSecurity(baseURL string) {
	res, _ := http.Get(baseURL + "/")
	ciphertextHex := res.Cookies()[0].Value
	fmt.Printf("[+] received ciphertext: %s\n", ciphertextHex)

	// To convert hex to bytes:
	// ciphertextBytes, _ := hex.DecodeString(ciphertextHex)

	req, _ := http.NewRequest(http.MethodGet, baseURL+"/check/", nil)
	req.AddCookie(&http.Cookie{Name: "authtoken", Value: ciphertextHex})
	res, _ = http.DefaultClient.Do(req)
	resBody, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("[+] done:\n%s\n", resBody)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: go run example.go <base url>")
		return
	}
	testSystemsSecurity(os.Args[1])
}
