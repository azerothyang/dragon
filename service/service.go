package service

import (
	"dragon/core/dragon/conf"
	"dragon/core/dragon/dlogger"
	"dragon/core/dragon/trace"
	"io/ioutil"
	"net/http"
	"strings"
)

// service response struct
type Response struct {
	Content string
	Status  int
	Err     error
	Rsp     *http.Response
}

//send get request
func GET(url string, params map[string]string, calleeService ...string) *Response {
	return send(url, params, "GET", calleeService...)
}

//send post request
func POST(url string, params map[string]string, calleeService ...string) *Response {
	return send(url, params, "POST", calleeService...)
}

//send put request
func PUT(url string, params map[string]string, calleeService ...string) *Response {
	return send(url, params, "PUT", calleeService...)
}

//send delete request
func DELETE(url string, params map[string]string, calleeService ...string) *Response {
	return send(url, params, "DELETE", calleeService...)
}

//send patch request
func PATCH(url string, params map[string]string, calleeService ...string) *Response {
	return send(url, params, "PATCH", calleeService...)
}

func send(url string, params map[string]string, method string, calleeService ...string) (resp *Response) {
	paramsStr := ""
	for k, v := range params {
		paramsStr += k + "=" + v + "&"
	}
	if paramsStr != "" {
		paramsStr = paramsStr[:len(paramsStr)-1]
	}
	var req *http.Request
	if method == "GET" {
		if paramsStr != "" {
			url += "?" + paramsStr
		}
		req, _ = http.NewRequest(method, url, nil)
	} else {
		req, _ = http.NewRequest(method, url, strings.NewReader(paramsStr))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		dlogger.SugarLogger.Error(err)
		resp = &Response{
			"",
			http.StatusInternalServerError,
			err,
			rsp,
		}
		return
	}
	defer rsp.Body.Close()

	// if enable zipkin,
	if conf.Conf.Zipkin.Enable {
		go chainZipkin(req, calleeService...)
	}

	content, errR := ioutil.ReadAll(rsp.Body)
	if errR != nil {
		dlogger.SugarLogger.Error(errR)
		resp = &Response{
			string(content),
			http.StatusInternalServerError,
			errR,
			rsp,
		}
		return
	}
	resp = &Response{
		string(content),
		rsp.StatusCode,
		errR,
		rsp,
	}
	return
}

// chain zipkin monitor
func chainZipkin(req *http.Request, calleeService ...string) {
	// args[0] for serviceName
	res, err := trace.Client.DoWithAppSpan(req, calleeService[0])
	if err != nil {
		dlogger.SugarLogger.Error(err)
	}
	res.Body.Close()
}
