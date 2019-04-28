package model

import (
	"time"
)

type TestModel struct {
	baseModel
}

type Test struct {
	Id         int64 `gorm:"primary_key"`
	Name       string
	CreateTime time.Time
}

//set orm table name
func (Test) TableName() string {
	return "test"
}

func (*TestModel) Get() *Test {
	test := Test{}
	readDB.First(&test)
	return &test
}

func (*TestModel) Create() *Test {
	test := Test{}
	readDB.Create(&Test{
		Name:       "holmes",
		CreateTime: time.Now(),
	})
	return &test
}

func (*TestModel) Update() *Test {
	test := Test{}
	readDB.Model(&test).Updates(map[string]interface{}{
		"name": "33",
	})
	return &test
}
