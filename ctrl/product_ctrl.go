package ctrl

import (
	"dragon/core/dragon/dlogger"
	"dragon/dto"
	"dragon/model"
	"fmt"
	"github.com/go-dragon/erro"
	"github.com/go-dragon/validator"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// 商品控制器
type Product struct {
	Ctrl
}

func (p Product) Test(resp http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// 初始化req
	(&p).InitReqAndResp(req, resp)
	reqData := p.GetRequestParams()
	fmt.Println(reqData)
	v := validator.New()
	v.Validate(&reqData, validator.Rules{
		"test": "notEmpty",
	})
	if v.HasErr {
		p.Json(&Output{
			Code: http.StatusBadRequest,
			Msg:  "",
			Data: v.ErrList,
		})
		return
	}
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
		p.Json(&output)
		return
	}

	if productInfo == nil {
		// 如果返回值为空，则直接返回空
		output := Output{
			Code: http.StatusOK,
			Msg:  http.StatusText(http.StatusOK),
		}
		p.Json(&output)
		return
	}

	res := dto.TStructToData(product, []string{"product_id", "product_name", "create_time"})

	output := Output{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: res,
	}
	p.Json(&output)
	return
}
