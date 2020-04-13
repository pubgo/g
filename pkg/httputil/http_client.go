package httputil

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var defaultTimeout = time.Duration(15 * time.Second)

func Client_Get(url string) (string, int, error) {

	client := http.Client{
		Timeout: defaultTimeout,
	}
	if strings.Contains(url, "oss") {
		client.Timeout = time.Duration(60 * time.Second)
	}

	response, err := client.Get(url)
	if err != nil {
		log.Printf("client.Get(url), %s", err)
		return "", http.StatusNotFound, err
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll: %s", err)
		return "", http.StatusNotFound, nil
	}

	if response != nil {
		return string(contents), response.StatusCode, nil
	}

	return string(contents), http.StatusNotFound, nil
}

func Client_Post(url, jsonData string, session string) (string, error) {
	var jsonStr = []byte(jsonData)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	if session != "" {
		req.Header.Set("SessionKey", session)
	}

	client := &http.Client{
		Timeout: defaultTimeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = resp.Body.Close()
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func Client_Post_Time(url string, jsonData string, timeout time.Duration) (string, error) {
	var jsonStr = []byte(jsonData)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = resp.Body.Close()
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func Client_Put(url string, jsonData string) (string, error) {
	var jsonStr = []byte(jsonData)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: defaultTimeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = resp.Body.Close()
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func ClientSelfPost(client *http.Client, url string, jsonData string) (string, error) {
	if client == nil {
		return "", errors.New("param client is nil")
	}

	var jsonStr = []byte(jsonData)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = resp.Body.Close()
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func ClientSelfGet(client *http.Client, url string) (string, error) {
	if client == nil {
		return "", errors.New("param client is nil")
	}

	response, err := client.Get(url)
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("%s", err)
		return "", nil
	}

	err = response.Body.Close()
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	return string(contents), nil
}
