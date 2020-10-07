package ctrl

import (
	"dragon/core/dragon"
	"dragon/service/serviceimpl"
	"fmt"
	"github.com/go-dragon/validator"
	"log"
	"net/http"
)

// 商品控制器
type Product struct {
}

func (p Product) Test(ctx *dragon.HttpContext) {
	// redis ZRangeByScoreWithScores
	//orders, err := dredis.Redis.ZRangeByScoreWithScores("order", redis.ZRangeBy{
	//	Min:    "0",
	//	Max:    "10",
	//	Offset: 0,
	//	Count:  3,
	//}).Result()
	//fmt.Println(err, orders)
	// 初始化req
	reqData := ctx.GetRequestParams()
	//fmt.Println("reqData", reqData)
	v := validator.New()
	v.Validate(&reqData, validator.Rules{
		"test": "notEmpty",
	})
	if v.HasErr {
		ctx.Json(&dragon.Output{
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
	log.Println("reqParams", fmt.Sprintf("%+v", ctx.GetRequestParams()))

	res := (&serviceimpl.ProductServiceImpl{}).GetList()

	//res := dto.TStructToData(product, []string{"product_id", "product_name", "create_time"})

	output := dragon.Output{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: res,
	}
	ctx.Json(&output, http.StatusOK)
	return
}
