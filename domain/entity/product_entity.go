package entity

const (
	ProductStatusDelete = 0
	ProductStatusOK     = 1
)

type ProductEntity struct {
	ProductId     int64 `gorm:"primaryKey;AUTO_INCREMENT"`
	ProductCode   string
	ProductName   string
	BrandId       int64
	ProductStatus int8 `gorm:"default:1"` // 商品状态: 0删除, 1正常
	CreateTime    string
	UpdateTime    string
}

//set orm table name
func (ProductEntity) TableName() string {
	return "t_product"
}
