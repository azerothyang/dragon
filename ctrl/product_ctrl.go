package ctrl

import (
	"dragon/service"
	"fmt"
	"github.com/go-dragon/validator"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// 商品控制器
type Product struct {
	Ctrl
}

func (p Product) Test(resp http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// redis ZRangeByScoreWithScores
	//orders, err := dredis.Redis.ZRangeByScoreWithScores("order", redis.ZRangeBy{
	//	Min:    "0",
	//	Max:    "10",
	//	Offset: 0,
	//	Count:  3,
	//}).Result()
	//fmt.Println(err, orders)
	// 初始化req
	(&p).InitReqAndResp(req, resp)
	reqData := p.GetRequestParams()
	fmt.Println("reqData", reqData)
	v := validator.New()
	v.Validate(&reqData, validator.Rules{
		"test": "notEmpty",
	})
	if v.HasErr {
		p.Json(&Output{
			Code: http.StatusBadRequest,
			Msg:  "",
			Data: v.ErrList,
		}, http.StatusBadRequest)
		return
	}
	// mongodb example
	//mongoRes, err := dmongo.DefaultDB().Collection("c_device_log").InsertOne(context.Background(), bson.M{
	//	"device_name": "golang",
	//})
	//if err != nil {
	//	fmt.Println("mongoErr", err)
	//}
	//objectId := mongoRes.InsertedID.(primitive.ObjectID)
	//fmt.Println("mongoRes", hex.EncodeToString(objectId[:]))

	// mysql example

	res := (&service.ProductService{}).GetList()

	//res := dto.TStructToData(product, []string{"product_id", "product_name", "create_time"})

	output := Output{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: res,
	}
	p.Json(&output, http.StatusOK)
	return
}
