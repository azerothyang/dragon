package dto

import (
	"dragon/tools/util"
	"encoding/json"
)

// 单个信息输出格式
type Data map[string]interface{}

// 输出列表
type ListData []Data

func TStructToData(tStruct interface{}, keys []string) Data {
	bt, _ := json.Marshal(tStruct)
	var res Data
	json.Unmarshal(bt, &res)
	util.OnlyColumns(keys, res)
	return res
}

func TStructsToListData(tStructs []interface{}, keys []string) ListData {
	output := ListData{}
	for v := range tStructs {
		output = append(output, TStructToData(v, keys))
	}
	return output
}
