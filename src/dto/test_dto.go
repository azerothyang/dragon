package dto

import (
	"log"
	"model"
	"strconv"
)

type TestDTO struct {
	Id int64 `json:"id"`
	Name string `json:"name,omitempty"`
	CreateTime string `json:"create_time"`
}

func TestPToTestS(params map[string]string, testModel *model.Test) {
	id, ok := params["id"]
	if ok {
		intId, err := strconv.Atoi(id)
		log.Println(err)
		testModel.Id = int64(intId)
	}
}

func TestSToTest(testModel *model.Test) *TestDTO {
	return &TestDTO{
		Id: testModel.Id,
		Name: testModel.Name,
		CreateTime: testModel.CreateTime.Format("2006-01-02 15:04:05"),
	}
}
