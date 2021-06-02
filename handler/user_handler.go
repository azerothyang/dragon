package handler

import (
	"dragon/core/dragon"
	"dragon/core/dragon/dlogger"
	"dragon/domain/repository"
	"dragon/domain/service"
	"github.com/go-dragon/erro"
	"github.com/go-dragon/validator"
	"net/http"
)

type IUserHandler interface {
	Test(ctx *dragon.HttpContext)
}

// UserHandler 用户处理器
type UserHandler struct {
}

func (u *UserHandler) Test(ctx *dragon.HttpContext) {
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
		"user_name": "notEmpty",
	})
	if v.HasErr {
		dlogger.Error(erro.NewError("error info"))
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
	//log.Println("reqParams", fmt.Sprintf("%+v", ctx.GetRequestParams()))
	userSrv := service.NewUserService(repository.GormDB) // 如果是事务处理，这个db可以为gorm的begin的db，只能从头传进去🤷‍
	res, _ := userSrv.GetOne()
	//log.Println("err:", err)

	//res := dto.TStructToData(product, []string{"product_id", "product_name", "create_time"})

	output := dragon.Output{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: res,
	}
	ctx.Json(&output, http.StatusOK)
	return
}
