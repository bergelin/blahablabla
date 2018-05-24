package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
	log.Println("Skrev Gorilla")
}

func URLHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("URL\n"))
	params := mux.Vars(r)
	url := params["url"]
	if strings.Contains(url, "http") == false {
		url = "https://" + url
	}
	fmt.Println(r.URL.EscapedPath())
	cmd := exec.Command("google-chrome-unstable", url)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Opened '%s' \n", url)
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
	fmt.Printf("Done %s", string(out))
}
func ListHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	list := params["list"]
	w.Write([]byte(params["list"]))
	log.Println("URL:")
	log.Println(params["list"])
	var files []string
	fileInfo, _ := ioutil.ReadDir("/home/temp/" + list)
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	log.Println(files)
	io.WriteString(w, strings.Join(files, "\n"))
}
func PlayHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	play := params["play"]
	play = "/" + play
	play = play + "/*"
	w.Write([]byte(play))
	cmd := exec.Command("vlc", play)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("Done %s", string(out))
}
func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", MainHandler)
	r.HandleFunc("/cmd/{cmd}", CMDHandler).Methods("GET")
	//r.UseEncodedPath()
	r.HandleFunc("/url/{url:.*}", URLHandler).Methods("GET")
	r.HandleFunc("/list/{list:.*}", ListHandler).Methods("GET")
	r.HandleFunc("/play/{play:.*}", PlayHandler).Methods("GET")
	//r.HandleFunc(url.QueryEscape("/url/{url}"), URLHandler).Methods("GET")
	//r.HandleFunc(html.EscapeString("/url/{url}"), URLHandler).Methods("GET")
	//r.HandleFunc(html.EscapeString("/url/{}"), URLHandler).Methods("GET")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
