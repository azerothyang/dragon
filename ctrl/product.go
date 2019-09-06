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
	"sync"
	"time"
)

// 商品控制器
type Product struct {
	Ctrl
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
			Code: http.StatusBadRequest,
			Msg:  "参数校验错误",
			Data: validator.ErrList,
		}
		p.Json(res, resp, reqData["SpanId"])
		return
	}
	// 验证品牌是否有效
	brandInfo := model.TBrand{}
	BrandModel.GetOne(&brandInfo, map[string]interface{}{
		"brand_code":   reqData["brand_code"],
		"brand_status": model.BrandStatusOK,
	}, "brand_id", "brand_id ASC")
	if brandInfo.BrandId == 0 {
		// 如果没有此品牌，提示品牌编码无效
		res := Output{
			Code: http.StatusBadRequest,
			Msg:  "品牌编码无效",
		}
		p.Json(res, resp, reqData["SpanId"])
		return
	}

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
			Code: http.StatusInternalServerError,
			Msg:  "商品新建失败",
		}
		p.Json(res, resp, reqData["SpanId"])
		return
	}

	res := Output{
		Code: http.StatusOK,
		Msg:  http.StatusText(http.StatusOK),
		Data: dto.TProduct2ProductOutput(&product),
	}
	p.Json(res, resp, reqData["SpanId"])
	return
}

// 伪删除单个商品
func (p *Product) Delete(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	reqData := dragon.Parse(req)
	productCode := params.ByName("product_code")
	reqData["product_code"] = productCode
	// 数据校验
	validator := validate.New()
	validator.Validate(&reqData, validate.Rules{
		"product_code": "notEmpty|maxLength:32",
	})
	if validator.HasErr {
		// 校验不通过，直接返回
		res := Output{
			Code: http.StatusBadRequest,
			Msg:  "参数校验错误",
			Data: validator.ErrList,
		}
		p.Json(res, resp, reqData["SpanId"])
		return
	}
	updateRes := ProductModel.SoftDelete(&model.TProduct{}, map[string]interface{}{
		"product_code": reqData["product_code"],
	}, "product_status", model.ProductStatusDelete)
	if updateRes == false {
		// 删除失败
		res := Output{
			Code: http.StatusInternalServerError,
			Msg:  "商品删除失败",
		}
		p.Json(res, resp, reqData["SpanId"])
		return
	}
	res := Output{
		Code: http.StatusOK,
		Msg:  http.StatusText(http.StatusOK),
	}
	p.Json(res, resp, reqData["SpanId"])
}

// 拉取商品列表
func (p *Product) GetList(resp http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	reqData := dragon.Parse(req)
	// 数据校验
	validator := validate.New()
	validator.Validate(&reqData, validate.Rules{
		"brand_code": "notEmpty|maxLength:32",
		"page":       "notEmpty|int32|min:0",
		"page_size":  "notEmpty|int32|min:0",
	})
	if validator.HasErr {
		// 校验不通过，直接返回
		res := Output{
			Code: http.StatusBadRequest,
			Msg:  "参数校验错误",
			Data: validator.ErrList,
		}
		p.Json(res, resp, reqData["SpanId"])
		return
	}

	// 初始化分页
	offset, pageSize := util.InitPageAndPageSize(reqData["page"], reqData["page_size"])

	// 协程查询商品列表和总数

	// 初始化通道和waitgroup
	wg := sync.WaitGroup{}
	resMap := sync.Map{}

	wg.Add(2)
	go func() {
		var productList []model.TProduct
		ProductModel.GetList(&productList, map[string]interface{}{
			"brand_code":     reqData["brand_code"],
			"product_status": model.ProductStatusOK,
		}, "product_id DESC", offset, pageSize, "*")
		list := dto.TProductList2ProductOutput(productList)
		resMap.Store("list", list)
		wg.Done()
	}()
	go func() {
		count := ProductModel.GetTotalCount(map[string]interface{}{
			"brand_code":     reqData["brand_code"],
			"product_status": model.ProductStatusOK,
		})
		resMap.Store("count", count)
		wg.Done()
	}()
	wg.Wait()

	resData := make(map[string]interface{}, 2)
	resData["list"], _ = resMap.Load("list")
	resData["count"], _ = resMap.Load("count")
	res := Output{
		Code: http.StatusOK,
		Msg:  http.StatusText(http.StatusOK),
		Data: resData,
	}
	p.Json(res, resp, reqData["SpanId"])
}

// 更新单个商品信息
func (p *Product) Update(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	reqData := dragon.Parse(req)
	code := params.ByName("product_code")
	reqData["product_code"] = code
	// 数据校验
	validator := validate.New()
	validator.Validate(&reqData, validate.Rules{
		"product_code":  "notEmpty|maxLength:32",
		"product_name":  "notEmpty|maxLength:64",
		"brand_code":    "notEmpty|maxLength:32",
		"product_price": "notEmpty|int64|min:0",
		"stock_num":     "notEmpty|int64|min:0",
	})
	if validator.HasErr {
		// 校验不通过，直接返回
		res := Output{
			Code: http.StatusBadRequest,
			Msg:  "参数校验错误",
			Data: validator.ErrList,
		}
		p.Json(res, resp, reqData["SpanId"])
		return
	}
	// 生成要更新的设备struct
	productPrice, _ := strconv.ParseInt(reqData["product_price"], 10, 64)
	stockNum, _ := strconv.ParseInt(reqData["stock_num"], 10, 64)
	updateData := map[string]interface{}{
		"product_name":  reqData["product_name"],
		"brand_code":    reqData["brand_code"],
		"product_price": productPrice,
		"stock_num":     stockNum,
	}
	updateRes := ProductModel.Updates(map[string]interface{}{
		"product_code":   reqData["product_code"],
		"product_status": model.ProductStatusOK,
	}, updateData)
	if !updateRes {
		// 更新失败
		res := Output{
			Code: http.StatusInternalServerError,
			Msg:  "商品更新失败",
		}
		p.Json(res, resp, reqData["SpanId"])
		return
	}
	res := Output{
		Code: http.StatusOK,
		Msg:  http.StatusText(http.StatusOK),
	}
	p.Json(res, resp, reqData["SpanId"])
}

// 获取单个商品详情
func (p *Product) GetOne(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	reqData := dragon.Parse(req)
	code := params.ByName("product_code")
	reqData["product_code"] = code
	// 数据校验
	validator := validate.New()
	validator.Validate(&reqData, validate.Rules{
		"product_code": "notEmpty|maxLength:32",
	})
	if validator.HasErr {
		// 校验不通过，直接返回
		res := Output{
			Code: http.StatusBadRequest,
			Msg:  "参数校验错误",
			Data: validator.ErrList,
		}
		p.Json(res, resp, reqData["SpanId"])
		return
	}

	product := model.TProduct{}
	// 定义查询条件
	ProductModel.GetOne(&product, map[string]interface {
	}{
		"product_code":   reqData["product_code"],
		"product_status": model.ProductStatusOK,
	}, "*", "product_id ASC")
	if product.ProductId == 0 {
		// 如果不存在则直接返回空
		res := Output{
			Code: http.StatusOK,
			Msg:  "商品不存在",
		}
		p.Json(res, resp, reqData["SpanId"])
		return
	}
	resData := dto.TProduct2ProductOutput(&product)
	res := Output{
		Code: http.StatusOK,
		Msg:  http.StatusText(http.StatusOK),
		Data: resData,
	}
	p.Json(res, resp, reqData["SpanId"])
}
