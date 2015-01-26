package main

import (
	"fmt"
	"github.com/bsm/openrtb"
	"net/http"
)

func bidder(w http.ResponseWriter, r *http.Request) {
	if _, err := openrtb.ParseRequest(r.Body); err != nil {
		fmt.Println(err.Error())
	} else {
		//fmt.Printf("%q - %q\n", *req.Id, *req.Device.Ua)
	}
	w.WriteHeader(204)
}

func bidderServer() {
	http.HandleFunc("/bid/openrtb", bidder)
	http.ListenAndServe(":8080", nil)
}
