package ctrl

import (
	"context"
	"dragon/core/dragon/dlogger"
	"dragon/dto"
	"dragon/model"
	"dragon/tools/dmongo"
	"encoding/hex"
	"fmt"
	"github.com/go-dragon/erro"
	"github.com/go-dragon/validator"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	mongoRes, err := dmongo.DefaultDB().Collection("c_device_log").InsertOne(context.Background(), bson.M{
		"device_name": "golang",
	})
	if err != nil {
		fmt.Println("mongoErr", err)
	}
	objectId := mongoRes.InsertedID.(primitive.ObjectID)
	fmt.Println("mongoRes", hex.EncodeToString(objectId[:]))

	// mysql example
	var product model.TProduct
	var conditions = []map[string]interface{}{
		{"product_id = ?": 1},
		//{"create_time <= ?": "2019-08-01"},
	}
	productInfo, err := ProductModel.GetOne(&product,
		conditions, "product_name,create_time", "")
	if err != nil {
		// new client失败
		dlogger.Error(erro.NewError(err))
		output := Output{
			Code: http.StatusBadRequest,
			Msg:  http.StatusText(http.StatusBadRequest),
			Data: err,
		}
		p.Json(&output, http.StatusInternalServerError)
		return
	}

	if productInfo == nil {
		// 如果返回值为空，则直接返回空
		output := Output{
			Code: http.StatusOK,
			Msg:  http.StatusText(http.StatusOK),
		}
		p.Json(&output, http.StatusOK)
		return
	}

	res := dto.TStructToData(product, []string{"product_id", "product_name", "create_time"})

	output := Output{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: res,
	}
	p.Json(&output, http.StatusOK)
	return
}
