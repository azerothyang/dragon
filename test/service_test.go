package test

import (
	"dragon/service"
	"fmt"
	"log"
	"testing"
)

// test GET
func TestGET(t *testing.T) {
	rsp, status, err := service.GET("http://talent.qh-1.cn/pc/service/orgs", map[string]string{
		"cur_code": "440300000000",
		"page":     "0",
		"pageSize": "10",
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(rsp)
	fmt.Println(status)

	rsp, status, err = service.GET("http://talent.qh-1.cn/pc/service/orgs?cur_code=440300000000&page=0", nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(rsp)
	fmt.Println(status)
}

// test POST
func TestPOST(t *testing.T) {
	rsp, status, err := service.POST("http://talent.qh-1.cn/pc/service/orgs", map[string]string{
		"cur_code": "440300000000",
		"page":     "0",
		"pageSize": "10",
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(rsp)
	fmt.Println(status)
}

func TestPUT(t *testing.T) {
	rsp, status, err := service.PUT("http://talent.qh-1.cn/pc/service/orgs", map[string]string{
		"cur_code": "440300000000",
		"page":     "0",
		"pageSize": "10",
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(rsp)
	fmt.Println(status)
}

func TestPATCH(t *testing.T) {
	rsp, status, err := service.PATCH("http://talent.qh-1.cn/pc/service/orgs", map[string]string{
		"cur_code": "440300000000",
		"page":     "0",
		"pageSize": "10",
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(rsp)
	fmt.Println(status)
}

func TestDELETE(t *testing.T) {
	rsp, status, err := service.DELETE("http://talent.qh-1.cn/pc/service/orgs", map[string]string{
		"cur_code": "440300000000",
		"page":     "0",
		"pageSize": "10",
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(rsp)
	fmt.Println(status)
}
