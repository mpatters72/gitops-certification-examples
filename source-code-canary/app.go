package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	//Blue:44B3C2, Yellow: F1A94E, Magenta: e033ff
	// comment out diff version/color to make new images
	version := "1.0"
	color := "#44B3C2"

	// version := "2.0"
	// color := "#F1A94E"

	// version := "3.0"
	// color := "#e033ff"

	http.HandleFunc("/callme", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<div class='pod' style='background:%s'> ver: %s\n </div>", color, version)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["status"] = "ok"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["status"] = version
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Listening now at port 8080")
	http.ListenAndServe(":8080", nil)
}
