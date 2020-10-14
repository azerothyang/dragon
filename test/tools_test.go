package test

import (
	"dragon/tools"
	"fmt"
	"log"
	"sync"
	"testing"
)

// test
func TestTOPSign(t *testing.T) {

}

func TestFastJson(t *testing.T) {
	data := map[string]interface{}{
		"x": 1,
		"y": "world",
		"a": "world",
		"b": "world",
		"c": "world",
		"d": "world",
		"e": "world",
		"f": "world",
	}
	var wg sync.WaitGroup
	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go func() {
			tools.FastJson.Marshal(&data)
			wg.Done()
		}()
	}
	wg.Wait()
	log.Println(111)
}

func BenchmarkFastJson(b *testing.B) {
	data := map[string]interface{}{
		"x": 1,
		"y": "world",
		"a": "world",
		"b": "world",
		"c": "world",
		"d": "world",
		"e": "world",
		"f": "world",
	}
	for i := 0; i < b.N; i++ {
		tools.FastJson.Marshal(&data)
	}
}

func BenchmarkFastJsonDecode(b *testing.B) {
	data := `{"x":1, "y":"hello world"}`
	var res map[string]interface{}
	for i := 0; i < b.N; i++ {
		tools.FastJson.Unmarshal([]byte(data), &res)
	}
	log.Println("res", fmt.Sprintf("%+v", res))
}

func TestUUidV4(t *testing.T) {
	const size = 10000000
	car := make(map[string]int, size)
	for i:=0; i<size; i++ {
		uuid := tools.UUidV4()
		if _, ok := car[uuid]; ok {
			// 如果uuid重复则报错
			log.Fatal("uuid repeat", uuid)
		}
		car[uuid] = 1
	}
	log.Println("uuidV4 generate success")
}

func BenchmarkUUidV4(b *testing.B) {
	for i:=0; i<b.N; i++ {
		tools.UUidV4()
	}
}
