package util

import (
	"fmt"
	"log"
	"testing"
)

func TestHmacSha1(t *testing.T) {
	fmt.Println(HmacSha1("dasdaf", "123"))
}

func BenchmarkHmacSha1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HmacSha1("dasdaf", "123")
	}
}

func TestHmacMD5(t *testing.T) {
	fmt.Println(HmacMD5("dasdaf", "123"))
	fmt.Println(HmacMD5("dasdaf", "123"))
}

func BenchmarkHmacMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HmacMD5("dasdaf", "123")
	}
}

func TestOnlyCols(t *testing.T) {
	params := map[string]string{
		"p1": "hello",
		"p2": "hi",
		"p3": "good",
	}
	OnlyCols([]string{"p1", "p2", "p5"}, params)
	if params["p1"] != "hello" || params["p2"] != "hi" {
		log.Fatal("TestOnlyCols fail")
	}
	fmt.Println(params)
}

func BenchmarkOnlyCols(b *testing.B) {
	params := map[string]string{
		"p1": "hello",
		"p2": "hi",
		"p3": "good",
	}
	OnlyCols([]string{"p1", "p2", "p5"}, params)
	for i := 0; i < b.N; i++ {
		OnlyCols([]string{"p1", "p2", "p5"}, params)
	}
}
