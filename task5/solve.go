package main

import (
	"encoding/json"
	"net/http"
)

var links map[string]string

type request struct {
	Url string `json:"url"`
}

func shortener(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var req request
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if req.Url == "" || err != nil {
			http.Error(w, "", 400)
			return
		}
		ans := make(map[string]string)
		ans["key"] = ""
		len := len(links) + 1
		for ; len > 0; len /= 26 {
			ans["key"] += string('a' + len % 26)
		}
		links[ans["key"]] = req.Url
		output, _ := json.Marshal(ans)
		w.Write(output)
	} else if r.Method == "GET" {
		link, ok := links[r.RequestURI[1:]]
		if !ok {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, link, 301)
	}
}

func main() {
	links = make(map[string]string)
	http.HandleFunc("/", shortener)
	http.ListenAndServe(":8082", nil)
}