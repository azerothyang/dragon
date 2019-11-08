package ctrl

import (
	"dragon/core/dragon"
	"dragon/core/dragon/util"
	"dragon/core/dragon/util/validate"
	"dragon/dto"
	"dragon/model"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
)

// 商品控制器
type Product struct {
	Ctrl
}

func (p *Product) Test(resp http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var products []model.TProduct
	var conditions []map[string]interface{}
	var cond = map[string]interface{}{
		"product_id = ?":   1,
		"create_time <= ?": "2019-08-01",
	}
	conditions = append(conditions, cond)
	count := ProductModel.GetListAndCount(&products,
		nil, "product_id desc", 0, 10, "*")
	output := Output{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: map[string]interface{}{
			"list":  products,
			"count": count,
		},
	}
	p.Json(&output, resp)
}

// 新增商品
func (p *Product) Add(resp http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	reqData := dragon.Parse(req)
	// 数据校验
	validator := validate.New()
	validator.Validate(&reqData, validate.Rules{
		"product_name":  "notEmpty|maxLength:64",
		"brand_code":    "notEmpty|maxLength:32",
		"product_price": "notEmpty|int64|min:0",
		"stock_num":     "notEmpty|int64|min:0",
	})
	if validator.HasErr {
		// 校验不通过，直接返回
		res := Output{
			Code:   http.StatusBadRequest,
			Msg:    "参数校验错误",
			Data:   validator.ErrList,
			SpanId: reqData["SpanId"],
		}
		p.Json(&res, resp)
		return
	}
	// 验证品牌是否有效
	//brandInfo := model.TBrand{}
	//BrandModel.GetOne(&brandInfo, {map[string]interface{}{
	//	"brand_code":   reqData["brand_code"],
	//	"brand_status": model.BrandStatusOK,
	//}}, "brand_id", "brand_id ASC")
	//if brandInfo.BrandId == 0 {
	//	// 如果没有此品牌，提示品牌编码无效
	//	res := Output{
	//		Code:   http.StatusBadRequest,
	//		Msg:    "品牌编码无效",
	//		SpanId: reqData["SpanId"],
	//	}
	//	p.Json(&res, resp)
	//	return
	//}

	productPrice, _ := strconv.ParseInt(reqData["product_price"], 10, 32)
	stockNum, _ := strconv.ParseInt(reqData["stock_num"], 10, 64)
	product := model.TProduct{
		ProductCode:  util.RandomStr(32),
		ProductName:  reqData["product_name"],
		BrandCode:    reqData["brand_code"],
		ProductPrice: productPrice,
		StockNum:     stockNum,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}
	ProductModel.Add(&product)
	if product.ProductId == 0 {
		// 商品新建失败
		res := Output{
			Code:   http.StatusInternalServerError,
			Msg:    "商品新建失败",
			SpanId: reqData["SpanId"],
		}
		p.Json(&res, resp)
		return
	}

	res := Output{
		Code:   http.StatusOK,
		Msg:    http.StatusText(http.StatusOK),
		Data:   dto.TProduct2ProductOutput(&product),
		SpanId: reqData["SpanId"],
	}
	p.Json(&res, resp)
	return
}
