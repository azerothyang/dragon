package handler

import (
	"dragon/core/dragon"
	"dragon/domain/repository"
	"dragon/domain/service"
	"fmt"
	"github.com/go-dragon/validator"
	"log"
	"net/http"
)

type IProductHandler interface {
	Test(ctx *dragon.HttpContext)
}

// 商品处理器
type ProductHandler struct {
}

func (p *ProductHandler) Test(ctx *dragon.HttpContext) {
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

	productSrv := service.NewProductService(repository.GormDB) // 如果是事务处理，这个db可以为gorm，begin的db，只能从头传进去🤷‍
	res, err := productSrv.GetOne()
	log.Println("err:", err)

	//res := dto.TStructToData(product, []string{"product_id", "product_name", "create_time"})

	output := dragon.Output{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: res,
	}
	ctx.Json(&output, http.StatusOK)
	return
}
