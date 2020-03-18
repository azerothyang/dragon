package service

import (
	"dragon/core/dragon/tracker"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"strings"
)

type Service struct {
	TrackWriter *http.Request
}

// service response struct
type Response struct {
	Content string
	Status  int
	Err     error
}

//send get request
func (s *Service) GET(url string, params map[string]string, headers map[string]string) *Response {
	return s.send(url, params, "GET", headers)
}

//send post request
func (s *Service) POST(url string, params map[string]string, headers map[string]string) *Response {
	return s.send(url, params, "POST", headers)
}

//send put request
func (s *Service) PUT(url string, params map[string]string, headers map[string]string) *Response {
	return s.send(url, params, "PUT", headers)
}

//send delete request
func (s *Service) DELETE(url string, params map[string]string, headers map[string]string) *Response {
	return s.send(url, params, "DELETE", headers)
}

//send patch request
func (s *Service) PATCH(url string, params map[string]string, headers map[string]string) *Response {
	return s.send(url, params, "PATCH", headers)
}

func (s *Service) send(url string, params map[string]string, method string, headers map[string]string) (resp *Response) {
	// 跟踪器
	var trackMan *tracker.Tracker
	if s.TrackWriter != nil {
		trackInfo := s.TrackWriter.Header.Get(tracker.TrackKey)
		trackMan = tracker.UnMarshal(trackInfo)
	}

	paramsStr := ""
	for k, v := range params {
		paramsStr += k + "=" + url2.QueryEscape(v) + "&"
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
	// add req headers
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	if trackMan != nil {
		//trackMan.Service.Req = req todo req直接结构体不行
		trackMan.Service.Req.Uri = req.URL.String()
		trackMan.Service.Req.Body = paramsStr // 记录请求内容
	}

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		resp = &Response{
			"",
			http.StatusInternalServerError,
			err,
		}
		if trackMan != nil {
			trackMan.Error = err
			s.TrackWriter.Header.Set(tracker.TrackKey, trackMan.Marshal()) // 最后写日志跟踪
		} // todo 继续完成
		return
	}
	defer rsp.Body.Close()

	content, errR := ioutil.ReadAll(rsp.Body)
	contentStr := string(content)
	if trackMan != nil {
		trackMan.Service.Resp = contentStr
	}

	if errR != nil {
		resp = &Response{
			contentStr,
			http.StatusInternalServerError,
			errR,
		}

		if trackMan != nil {
			trackMan.Error = errR
			s.TrackWriter.Header.Set(tracker.TrackKey, trackMan.Marshal()) // 最后写日志跟踪
		}
		return
	}
	// service返回

	resp = &Response{
		contentStr,
		rsp.StatusCode,
		errR,
	}
	if trackMan != nil {
		s.TrackWriter.Header.Set(tracker.TrackKey, trackMan.Marshal()) // 最后写日志跟踪
	}
	return
}

//send postJson
func (s *Service) POSTJson(url string, paramsStr string, spanId string, calleeService ...string) (resp *Response) {
	var req *http.Request
	req, _ = http.NewRequest("POST", url, strings.NewReader(paramsStr))
	req.Header.Add("Content-Type", "application/json")
	// 跟踪器
	trackInfo := s.TrackWriter.Header.Get(tracker.TrackKey)
	trackMan := tracker.UnMarshal(trackInfo)

	rsp, err := http.DefaultClient.Do(req)

	if err != nil {
		resp = &Response{
			"",
			http.StatusInternalServerError,
			err,
		}
		trackMan.Error = err
		s.TrackWriter.Header.Set(tracker.TrackKey, trackMan.Marshal()) // 最后写日志跟踪
		return
	}
	defer rsp.Body.Close()

	content, errR := ioutil.ReadAll(rsp.Body)
	// 写入返回content
	trackMan.Service.Resp = string(content)

	if errR != nil {
		log.Println(err)
		resp = &Response{
			trackMan.Service.Resp,
			http.StatusInternalServerError,
			errR,
		}

		trackMan.Error = errR
		s.TrackWriter.Header.Set(tracker.TrackKey, trackMan.Marshal()) // 最后写日志跟踪
		return
	}

	resp = &Response{
		trackMan.Service.Resp,
		rsp.StatusCode,
		errR,
	}
	s.TrackWriter.Header.Set(tracker.TrackKey, trackMan.Marshal()) // 最后写日志跟踪
	return
}
