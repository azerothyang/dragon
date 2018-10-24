package dragon

import (
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

	for k, v := range r.Form  {
		if len(v) == 1 {
			requests[k] = v[0]
		}
	}

	return requests
}