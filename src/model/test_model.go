package model

type TestModel struct {
	baseModel
}

type Test struct {
	Id int64 `gorm:"primary_key"`
}

//set orm table name
func (Test)TableName() string {
	return "test"
}

func (*TestModel) Get() *Test {
	test := Test{}
	readDB.First(&test)
	return &test
}
