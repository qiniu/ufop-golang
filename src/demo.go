package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type ReqArgs struct {
	Cmd  string `json:"cmd"`
	Mode uint32 `json:"mode"`
	Src  struct {
		Url      string `json:"url"`
		Mimetype string `json:"mimetype"`
		Fsize    int32  `json:"fsize"`
		Bucket   string `json:"bucket"`
		Key      string `json:"key"`
	} `json: "src"`
}

func demoHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("start")

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(500)
		log.Println("fail to read req body")
		return
	}

	var args ReqArgs
	err = json.Unmarshal(body, &args)
	if err != nil {
		w.WriteHeader(500)
		log.Println("fail to unmarshal")
		return
	}

	resp, err := http.Get(args.Src.Url)
	if err != nil {
		w.WriteHeader(500)
		log.Println("fail to fetch file")
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 512)
	io.ReadFull(resp.Body, buf)
	contentType := http.DetectContentType(buf)
	w.Write([]byte("Hello World!\n"))
	w.Write([]byte("The file's mime type is " + contentType))

	log.Println("end")
}

func main() {
	http.HandleFunc("/uop", demoHandler)
	err := http.ListenAndServe(":9100", nil)
	if err != nil {
		log.Fatal("Demo server failed to start:", err)
	}
}
