package model

const (
	ProductStatusDelete = 0
	ProductStatusOK     = 1
)

// 商品表
type ProductModel struct {
	BaseModel
}

type TProduct struct {
	ProductId     int64  `gorm:"primaryKey;AUTO_INCREMENT" json:"product_id"`
	ProductCode   string `json:"product_code"`
	ProductName   string `json:"product_name"`
	BrandCode     string `json:"brand_code"`
	ProductPrice  int64  `json:"product_price"`
	StockNum      int64  `json:"stock_num"`
	ProductStatus int8   `gorm:"default:1" json:"product_status"` // 商品状态: 0删除, 1正常
	CreateTime    string `json:"create_time"`
	UpdateTime    string `json:"update_time"`
}

//set orm table name
func (TProduct) TableName() string {
	return "t_product"
}
