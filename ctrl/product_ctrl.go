package ctrl

import (
	"dragon/dto"
	"dragon/model"
	"fmt"
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

	var product model.TProduct
	var conditions = []map[string]interface{}{
		{"product_id = ?": 1},
		//{"create_time <= ?": "2019-08-01"},
	}
	productInfo, err := ProductModel.GetOne(&product,
		conditions, "product_name,create_time", "")
	if err != nil {
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
