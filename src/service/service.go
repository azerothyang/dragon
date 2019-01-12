package service

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//send get request
func GET(url string, params map[string]string) (rspContent string, rspStatus int, rspErr error){
	paramsStr := ""
	for k, v := range params {
		paramsStr += k + "=" + v + "&"
	}
	if paramsStr != "" {
		// have params
		url += "?" + paramsStr
		url = url[:len(url)-1]
	}
	rsp, err := http.Get(url)
	defer rsp.Body.Close()
	if err != nil {
		log.Println(err)
		rspContent, rspStatus, rspErr = "", 600, err
		return
	}
	content, errR := ioutil.ReadAll(rsp.Body)
	if errR != nil {
		log.Println(errR)
	}
	rspContent, rspStatus, rspErr = string(content), rsp.StatusCode, errR
	return
}

//send post request
func POST(url string, params map[string]string) (string, int, error){
	return sendWWWFormUrlencoded(url, params, "POST")
}

//send put request
func PUT(url string, params map[string]string) (string, int, error){
	return sendWWWFormUrlencoded(url, params, "PUT")
}

//send delete request
func DELETE(url string, params map[string]string) (string, int, error){
	return sendWWWFormUrlencoded(url, params, "DELETE")
}

//send patch request
func PATCH(url string, params map[string]string) (string, int, error){
	return sendWWWFormUrlencoded(url, params, "PATCH")
}

func sendWWWFormUrlencoded(url string, params map[string]string, method string) (rspContent string, rspStatus int, rspErr error) {
	paramsStr := ""
	for k, v := range params {
		paramsStr += k + "=" + v + "&"
	}
	if paramsStr != "" {
		paramsStr = paramsStr[:len(paramsStr)-1]
	}

	req, _ := http.NewRequest(method, url, strings.NewReader(paramsStr))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rsp, err := http.DefaultClient.Do(req)
	defer rsp.Body.Close()
	if err != nil {
		log.Println(err)
		rspContent, rspStatus, rspErr = "", 600, err
		return
	}
	content, errR := ioutil.ReadAll(rsp.Body)
	if errR != nil {
		log.Println(errR)
	}
	rspContent, rspStatus, rspErr = string(content), rsp.StatusCode, errR
	return
}

