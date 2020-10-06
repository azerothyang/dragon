package test

import (
	"dragon/httpclient"
	"fmt"
	"log"
	"testing"
)

// test GET
func TestGET(t *testing.T) {
	srv := &httpclient.Client{}
	res := srv.GET("http://talent.qh-1.cn/pc/httpclient/orgs", map[string]string{
		"cur_code": "440300000000",
		"page":     "0",
		"pageSize": "10",
	}, nil)

	if res.Err != nil {
		log.Println(res.Err)
	}
	fmt.Println(res.Content)
	fmt.Println(res.Status)
}

// test POST
func TestPOST(t *testing.T) {
	srv := &httpclient.Client{}
	res := srv.POST("http://talent.qh-1.cn/pc/httpclient/orgs", map[string]string{
		"cur_code": "440300000000",
		"page":     "0",
		"pageSize": "10",
	}, nil)
	if res.Err != nil {
		log.Println(res.Err)
	}
	fmt.Println(res.Content)
	fmt.Println(res.Status)
}

func TestPUT(t *testing.T) {
	srv := &httpclient.Client{}
	res := srv.PUT("http://talent.qh-1.cn/pc/httpclient/orgs", map[string]string{
		"cur_code": "440300000000",
		"page":     "0",
		"pageSize": "10",
	}, nil)
	if res.Err != nil {
		log.Println(res.Err)
	}
	fmt.Println(res.Content)
	fmt.Println(res.Status)
}

func TestPATCH(t *testing.T) {
	srv := &httpclient.Client{}
	res := srv.PATCH("http://talent.qh-1.cn/pc/httpclient/orgs", map[string]string{
		"cur_code": "440300000000",
		"page":     "0",
		"pageSize": "10",
	}, nil)
	if res.Err != nil {
		log.Println(res.Err)
	}
	fmt.Println(res.Content)
	fmt.Println(res.Status)
}

func TestDELETE(t *testing.T) {
	srv := &httpclient.Client{}
	res := srv.DELETE("http://talent.qh-1.cn/pc/httpclient/orgs", map[string]string{
		"cur_code": "440300000000",
		"page":     "0",
		"pageSize": "10",
	}, nil)
	if res.Err != nil {
		log.Println(res.Err)
	}
	fmt.Println(res.Content)
	fmt.Println(res.Status)
}
