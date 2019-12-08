package dto

import (
	"github.com/go-dragon/util"
)

// 单个信息输出格式
type Data map[string]interface{}

// 输出列表
type ListData []Data

func TStructToData(obj interface{}, keys []string) Data {
	res := util.StructJsonTagToMap(obj)
	util.OnlyColumns(keys, res)
	return res
}

func TStructsToListData(objs []interface{}, keys []string) ListData {
	output := ListData{}
	for _, v := range objs {
		output = append(output, TStructToData(v, keys))
	}
	return output
}
