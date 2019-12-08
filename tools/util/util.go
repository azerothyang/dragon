package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"strings"
)

//过滤字段，map中只保留需要的键值对
func OnlyCols(cols []string, data map[string]string) {
	for k := range data {
		//判断k 是否在需要的cols中， 如果不在，则对应的键值对
		have := false //不在
		for _, col := range cols {
			if k == col {
				have = true
				break
			}
		}
		if have == false {
			delete(data, k)
		}
	}
}

//过滤字段，map中只保留需要的键值对
func OnlyColumns(cols []string, data map[string]interface{}) {
	for k := range data {
		//判断k 是否在需要的cols中， 如果不在，则对应的键值对
		have := false //不在
		for _, col := range cols {
			if k == col {
				have = true
				break
			}
		}
		if have == false {
			delete(data, k)
		}
	}
}

// hmac-sha1 input with key
func HmacSha1(input, key string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(input))
	return hex.EncodeToString(mac.Sum(nil))
}

// hmac-md5 input with key
func HmacMD5(input, key string) string {
	mac := hmac.New(md5.New, []byte(key))
	mac.Write([]byte(input))
	return hex.EncodeToString(mac.Sum(nil))
}

// return json string
func ToJsonString(data interface{}) string {
	j, _ := json.Marshal(data)
	return string(j)
}

// slice and trim white space
func SliceAndTrim(str string) []string {
	res := make([]string, 0)
	strs := strings.Split(str, ",")
	for _, s := range strs {
		res = append(res, strings.Trim(s, " "))
	}
	return res
}
