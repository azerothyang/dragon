package dto

import (
	"log"
	"model"
	"strconv"
)

func TestPToTestS(params map[string]string, testModel *model.Test) {
	id, ok := params["id"]
	if ok {
		intId, err := strconv.Atoi(id)
		log.Println(err)
		testModel.Id = int64(intId)
	}
}
