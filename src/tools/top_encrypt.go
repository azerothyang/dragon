package tools

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
)

// get TOP sign string with md5
func TOPSign(params map[string]string, key string) string {
	var kArr sort.StringSlice
	for k := range params {
		kArr = append(kArr, k)
	}

	// sort string with ascii
	sort.Sort(kArr)
	str := key
	// concat a string md5(key+bar2foo1foo_bar3foobar4+key)
	for _, k := range kArr {
		str += k + params[k]
	}
	str += key
	enc := md5.Sum([]byte(str))
	return hex.EncodeToString(enc[:])
}
