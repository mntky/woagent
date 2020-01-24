package cmd

import (
	"fmt"
	"bytes"
	"net/http"
	"encoding/json"
	"woagent/pkg"
)

type notice struct {
	Name	string
}

func startAgent(url string) error {
	fmt.Println("start agent")
	http.HandleFunc("/api/notice", notice_handle)

	err := http.ListenAndServe(url, nil)
	return err
}

func notice_handle(w http.ResponseWriter, r *http.Request) {
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)

	var noticedata notice
	err := json.Unmarshal(bufbody.Bytes(), &noticedata)
	if err != nil {
		fmt.Println(err)
	}

	//TODO 受け取ったnoticeのコンテナ名のspecを取得する
	fmt.Println(noticedata.Name)
	pkg.ReferSpec(noticedata.Name)

	w.Write([]byte("notice receive ok"))
}
