package main

import (
	"net/http"
	"encoding/json"
	"log"
	"html/template"
	"net"
	"github.com/skratchdot/open-golang/open"
)

func ui__json(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data,err := json.Marshal(PERCENT_DOWNLOADED_NW)
	if err != nil {
		log.Panic("Cannot serve json",PERCENT_DOWNLOADED_NW)
	}
	w.Write(data)
}
func ui__static(w http.ResponseWriter, r *http.Request) {
   log.Println(r.URL.Path[1:])
   http.ServeFile(w, r, r.URL.Path[1:])
}
func ui__index(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("nw-sulfur-ui/index.html")
	t.Execute(w,nil)
//	http.ServeFile(w,r,"./ui/index.html")
}

func startServer(){
	fs := http.FileServer(http.Dir("nw-sulfur-ui/static/"))
	http.HandleFunc("/json",ui__json)
	http.Handle("/static/",http.StripPrefix("/static/", fs))
	http.HandleFunc("/", ui__index)

	listner, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal("Listen:", err)
	}
	log.Println(listner.Addr().String())
    // Open local server in a browser
	open.Start("http://" + listner.Addr().String())

	err2 := http.Serve(listner, nil)

	if err2 != nil {
		log.Fatal("http.Serve:", err2)
	}
//	http.ListenAndServe(":5000", nil)
}





