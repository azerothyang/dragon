package ctrl

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)



type Ctrl struct {
	
}

func init()  {

}

//return with json
func (*Ctrl)Json(data interface{}, resp http.ResponseWriter) {
	resp.Header().Set("content-type", "application/json; charset=utf-8")
	resp.Header().Set("server", "dragon")
	js, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(resp, "")
		return
	}
	_, err = resp.Write(js)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(resp, "")
		return
	}
}