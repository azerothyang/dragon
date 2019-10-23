package dragon

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//parse request params to map exclude file/binary upload or raw data
func Parse(r *http.Request) map[string]string {

	requests := make(map[string]string)
	var err error
	err = r.ParseForm()
	if err != nil {
		log.Println(err)
		return requests
	}

	for k, v := range r.Form {
		if len(v) == 1 {
			requests[k] = v[0]
		}
	}

	return requests
}

//parse raw json
func ParseRawJson(r *http.Request, data interface{}) error {
	var body []byte
	var err error
	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	err = json.Unmarshal(body, data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
