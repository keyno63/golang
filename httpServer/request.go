package httpServer

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Request do get.
func Request(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray)) // htmlをstringで取得
}
