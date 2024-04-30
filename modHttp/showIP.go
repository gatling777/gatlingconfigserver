package modHttp

import (
	"fmt"
	"io"
	"net/http"
)

func showIP() {
	res, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return
	}
	text, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("get ip error: ", err)
	}
	fmt.Println("get ip: ", string(text))
}
