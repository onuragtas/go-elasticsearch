package v7

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func (t *ElasticSearchV7) search(message []byte) ([]byte, error) {
	if t.Host == "" {
		t.Host = os.Getenv("ELASTIC_HOST")
	}
	req, err := http.NewRequest("GET", t.Host+"/"+t.Index+"/"+t.Type+"/_search?pretty", bytes.NewBuffer(message))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		var res []byte
		return res, errors.New("error")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil
}

func (t *ElasticSearchV7) scroll(message []byte) ([]byte, error) {
	if t.Host == "" {
		t.Host = os.Getenv("ELASTIC_HOST")
	}
	req, err := http.NewRequest("GET", t.Host+"/"+t.Index+"/"+t.Type+"/_search?scroll=10m&pretty", bytes.NewBuffer(message))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		var res []byte
		return res, errors.New("error")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil
}

func (t *ElasticSearchV7) request(method, url string, message []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(message))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		var res []byte
		return res, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil
}
