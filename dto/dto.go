package dto

import (
	"encoding/json"
)

// 单个信息输出格式
type Data map[string]interface{}

// 输出列表
type ListData []Data

func TStructToData(tStruct interface{}, keys []string) Data {
	bt, _ := json.Marshal(tStruct)
	var res Data
	output := make(Data)
	json.Unmarshal(bt, &res)
	for _, key := range keys {
		if v, ok := res[key]; ok {
			// 如果数据库查询值在搜索字段中，则写入输出
			output[key] = v
		}
	}
	return output
}

func TStructsToListData(tStructs []interface{}, keys []string) ListData {
	output := ListData{}
	for v := range tStructs {
		output = append(output, TStructToData(v, keys))
	}
	return output
}
