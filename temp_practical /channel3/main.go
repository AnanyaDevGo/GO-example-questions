package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/getdata", getdata)
	http.ListenAndServe(":8080", nil)
}

func getdata(w http.ResponseWriter, r *http.Request) {
	log.Print("before")
	timechan := time.After(10 * time.Second)

	select {
	case <-r.Context().Done():
		log.Print("req cancel")
		return
	case <-timechan:

	}
	log.Print("after")
	w.Write([]byte("hello"))

}
