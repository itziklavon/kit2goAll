package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../log"
)

type GenericHttpResponse struct {
	httpResponse string
	httpBody     string
	httpHeaders  map[string]string
}

func GET(url string, headers map[string]string) GenericHttpResponse {
	req, err := http.NewRequest("GET", url, nil)
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.ErrorException(":GET: couldn't send request", err)
	}
	defer resp.Body.Close()
	log.Info("response Status:" + resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Info("response Body:" + string(body))
	var responseHeaders = make(map[string]string)
	for key, value := range resp.Header {
		responseHeaders[key] = value[0]
		log.Info(":GET: response Headers: key: " + key + ", value: " + value[0])
	}
	return GenericHttpResponse{httpResponse: resp.Status, httpBody: string(body), httpHeaders: responseHeaders}
}

func POST(url string, body map[string]string, headers map[string]string) GenericHttpResponse {
	jsonValue, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.ErrorException(":POST: couldn't send request", err)
	}
	defer resp.Body.Close()
	log.Info("response Status:" + resp.Status)
	responseBody, _ := ioutil.ReadAll(resp.Body)
	log.Info("response Body:" + string(responseBody))
	var responseHeaders = make(map[string]string)
	for key, value := range resp.Header {
		responseHeaders[key] = value[0]
		log.Info(":POST: response Headers: key: " + key + ", value: " + value[0])
	}
	return GenericHttpResponse{httpResponse: resp.Status, httpBody: string(responseBody), httpHeaders: responseHeaders}
}

func PUT(url string, body map[string]string, headers map[string]string) GenericHttpResponse {
	jsonValue, _ := json.Marshal(body)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.ErrorException(":PUT: couldn't send request", err)
	}
	defer resp.Body.Close()
	log.Info("response Status:" + resp.Status)
	responseBody, _ := ioutil.ReadAll(resp.Body)
	log.Info("response Body:" + string(responseBody))
	var responseHeaders = make(map[string]string)
	for key, value := range resp.Header {
		responseHeaders[key] = value[0]
		log.Info(":PUT: response Headers: key: " + key + ", value: " + value[0])
	}
	return GenericHttpResponse{httpResponse: resp.Status, httpBody: string(responseBody), httpHeaders: responseHeaders}
}

func DELETE(url string, body map[string]string, headers map[string]string) GenericHttpResponse {
	jsonValue, _ := json.Marshal(body)
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonValue))
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.ErrorException(":DELETEGET: couldn't send request", err)
	}
	defer resp.Body.Close()
	log.Info("response Status:" + resp.Status)
	responseBody, _ := ioutil.ReadAll(resp.Body)
	log.Info("response Body:" + string(responseBody))
	var responseHeaders = make(map[string]string)
	for key, value := range resp.Header {
		responseHeaders[key] = value[0]
		log.Info(":DELETE: response Headers: key: " + key + ", value: " + value[0])
	}
	return GenericHttpResponse{httpResponse: resp.Status, httpBody: string(responseBody), httpHeaders: responseHeaders}
}
