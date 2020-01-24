package pkg

import (
	"os"
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type joinmeta struct {
	Name	string	`yaml:"name"`
	Addr	string	`yaml:"addr"`
}

func Join(master, agent string) error {
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}
	metadata := &joinmeta{
		Name:	hostname,
		Addr:	agent,
	}

	metabyte, _ := json.Marshal(metadata)

	fmt.Println(string(metabyte))
	//create request
	req, err := http.NewRequest(
		"POST",
		"http://" + master+"/api/node/add",
		bytes.NewBuffer(metabyte),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	//create client and send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return nil
}


