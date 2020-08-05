package main

import (
	"encoding/json"
	"net/http"
	"path"
)

func find(x string) int {
	for k, v := range records {
		if x == v[0] {
			return k
		}
	}
	return -1
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := find(id)
	if i == -1 {
		return
	}
	dataJson, err := json.Marshal(records[i])
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var data []string
	json.Unmarshal(body, &data)
	records = append(records, data)
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	w.WriteHeader(200)
	return
}
