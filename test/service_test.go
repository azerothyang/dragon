package test

import (
	"dragon/service"
	"fmt"
	"log"
	"testing"
)

// test GET
func TestGET(t *testing.T) {
	res := service.GET("http://talent.qh-1.cn/pc/service/orgs", map[string]string{
		"cur_code": "440300000000",
		"page":     "0",
		"pageSize": "10",
	})

	if res.Err != nil {
		log.Println(res.Err)
	}
	fmt.Println(res.Content)
	fmt.Println(res.Status)
}

// test POST
func TestPOST(t *testing.T) {
	res := service.POST("http://talent.qh-1.cn/pc/service/orgs", map[string]string{
		"cur_code": "440300000000",
		"page":     "0",
		"pageSize": "10",
	})
	if res.Err != nil {
		log.Println(res.Err)
	}
	fmt.Println(res.Content)
	fmt.Println(res.Status)
}

func TestPUT(t *testing.T) {
	res := service.PUT("http://talent.qh-1.cn/pc/service/orgs", map[string]string{
		"cur_code": "440300000000",
		"page":     "0",
		"pageSize": "10",
	})
	if res.Err != nil {
		log.Println(res.Err)
	}
	fmt.Println(res.Content)
	fmt.Println(res.Status)
}

func TestPATCH(t *testing.T) {
	res := service.PATCH("http://talent.qh-1.cn/pc/service/orgs", map[string]string{
		"cur_code": "440300000000",
		"page":     "0",
		"pageSize": "10",
	})
	if res.Err != nil {
		log.Println(res.Err)
	}
	fmt.Println(res.Content)
	fmt.Println(res.Status)
}

func TestDELETE(t *testing.T) {
	res := service.DELETE("http://talent.qh-1.cn/pc/service/orgs", map[string]string{
		"cur_code": "440300000000",
		"page":     "0",
		"pageSize": "10",
	})
	if res.Err != nil {
		log.Println(res.Err)
	}
	fmt.Println(res.Content)
	fmt.Println(res.Status)
}
