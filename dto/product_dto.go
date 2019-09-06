package dto

import (
	"dragon/model"
)

// 商品信息输出格式
type ProductOutput struct {
	ProductId    int64  `json:"product_id"`
	ProductCode  string `json:"product_code"`
	ProductName  string `json:"product_name"`
	BrandCode    string `json:"brand_code"`
	ProductPrice int64  `json:"product_price"`
	StockNum     int64  `json:"stock_num"`
	CreateTime   string `json:"create_time"`
	UpdateTime   string `json:"update_time"`
}

// TProduct转ProductOutput
func TProduct2ProductOutput(tProduct *model.TProduct) *ProductOutput {
	output := ProductOutput{
		ProductId:    tProduct.ProductId,
		ProductCode:  tProduct.ProductCode,
		ProductName:  tProduct.ProductName,
		BrandCode:    tProduct.BrandCode,
		ProductPrice: tProduct.ProductPrice,
		StockNum:     tProduct.StockNum,
		CreateTime:   tProduct.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:   tProduct.UpdateTime.Format("2006-01-02 15:04:05"),
	}
	return &output
}

// []TProduct转 productList 输出
func TProductList2ProductOutput(tProductList []model.TProduct) []ProductOutput {
	var productList []ProductOutput
	for _, v := range tProductList {
		output := TProduct2ProductOutput(&v)
		productList = append(productList, *output)
	}
	return productList
}
