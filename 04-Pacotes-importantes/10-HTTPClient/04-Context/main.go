package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	//ctx, cancel := context.WithDeadline(ctx) //Inverso de timeout (como assim??)
	//ctx, cancel := context.WithCancel(ctx) //O critério de cancelamento é somente se alguem executar o método cancel()
	ctx, cancel := context.WithTimeout(ctx, time.Second) //passar de 1 seg, cancela
	defer cancel()                                       //cancel, cancela a rotina de forma natural
	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
