package repository

const (
	ProductStatusDelete = 0
	ProductStatusOK     = 1
)

// 商品表
type ProductRepository struct {
	BaseRepository
}

type TProduct struct {
	ProductId     int64 `gorm:"primaryKey;AUTO_INCREMENT"`
	ProductCode   string
	ProductName   string
	BrandCode     string
	ProductPrice  int64
	StockNum      int64
	ProductStatus int8 `gorm:"default:1"` // 商品状态: 0删除, 1正常
	CreateTime    string
	UpdateTime    string
}

//set orm table name
func (TProduct) TableName() string {
	return "t_product"
}
