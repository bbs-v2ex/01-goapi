package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	post, _ := http.Post("http://127.0.0.1:1234/jsonrpc", "json", strings.NewReader(`{"id":0,"params":["bobby"],"method":"HelloService.Hello"}`))
	all, _ := ioutil.ReadAll(post.Body)
	fmt.Println(string(all))
}
