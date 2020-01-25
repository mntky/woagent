package pkg

import (
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type LxcSpec struct {
	Name		string	`json:"name"`
	Distro	string	`json:"distro"`
	Release	string	`json:"release"`
	Arch		string	`json:"arch"`
}

//TODO viperでmasterのアドレス取ってくる
func ReferSpec(speckey string) error {
	req, err := http.NewRequest(
		"POST",
		"http://localhost:9090/api/refer/spec",
		bytes.NewBuffer([]byte(speckey)),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "text/plain")

	//create client and send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	lxcspec := &LxcSpec{}
	json.Unmarshal(body, lxcspec)
	fmt.Println("debug----")
	fmt.Println(string(body))
	err = Create(*lxcspec)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
