package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
	log.Println("Skrev Gorilla")
}

func URLHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("URL\n"))
	params := mux.Vars(r)
	url := params["url"]
	fmt.Println(r.URL.EscapedPath())
	cmd := exec.Command("google-chrome", url)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Done %s", string(out))
}

func CMDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Write([]byte(params["cmd"]))
	log.Println("CMD:")
	log.Println(params["cmd"])
	cmd := exec.Command("vlc", "/home/temp/Downloads/")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Println("Done %s", string(out))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", MainHandler)
	r.HandleFunc("/cmd/{cmd}", CMDHandler).Methods("GET")
	//r.UseEncodedPath()
	r.HandleFunc("/url/{url:.*}", URLHandler).Methods("GET")
	//r.HandleFunc(url.QueryEscape("/url/{url}"), URLHandler).Methods("GET")
	//r.HandleFunc(html.EscapeString("/url/{url}"), URLHandler).Methods("GET")
	//r.HandleFunc(html.EscapeString("/url/{}"), URLHandler).Methods("GET")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
