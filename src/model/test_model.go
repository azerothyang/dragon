package model

type TestModel struct {
	baseModel
}

type Test struct {
	Id int64 `gorm:"primary_key"`
}

func (*TestModel) Get() *Test{
	test := Test{}
	readDB.First(&test)
	return &test
}