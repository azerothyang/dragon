package test

import (
	"core/dragon/util"
	"fmt"
	"log"
	"testing"
)

// test generate uuid
func TestNewUUID(t *testing.T) {
	log.Println(util.NewUUID())
	log.Println(util.NewUUID())
}

//test rand generate string
func TestRandomStr(t *testing.T) {
	var key string
	pressure := 10000000
	mp := make(map[string]int)
	for i := 0; i < pressure; i++ {
		key = util.RandomStr(16)
		_, ok := mp[key]
		if ok {
			//repeat random str
			log.Fatal("repeat random string")
		}
		mp[key] = 1
	}
}

func BenchmarkRandomStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		util.RandomStr(16)
	}
}

func TestHmacSha1(t *testing.T)  {
	fmt.Println(util.HmacSha1("dasdaf", "123"))
}
