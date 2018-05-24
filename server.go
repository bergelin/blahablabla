package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
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
	list = "/" + list
	w.Write([]byte(params["list"]))
	log.Println("URL:")
	log.Println(params["list"])
	var files []string
	err := filepath.Walk("/home/temp/TV/", func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		log.Println(file)
	}
}
func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", MainHandler)
	r.HandleFunc("/cmd/{cmd}", CMDHandler).Methods("GET")
	//r.UseEncodedPath()
	r.HandleFunc("/url/{url:.*}", URLHandler).Methods("GET")
	r.HandleFunc("/list/{list:.*}", ListHandler).Methods("GET")
	//r.HandleFunc(url.QueryEscape("/url/{url}"), URLHandler).Methods("GET")
	//r.HandleFunc(html.EscapeString("/url/{url}"), URLHandler).Methods("GET")
	//r.HandleFunc(html.EscapeString("/url/{}"), URLHandler).Methods("GET")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
