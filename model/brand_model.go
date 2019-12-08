package model

const (
	BrandStatusDelete = 0
	BrandStatusOK     = 1
)

// 品牌表
type BrandModel struct {
	BaseModel
}

type TBrand struct {
	BrandId      int64  `gorm:"primary_key;AUTO_INCREMENT"`
	BrandCode    string `gorm:"unique_index"`
	BrandName    string
	BrandProfile string
	AppId        string
	AppSecret    string
	BrandStatus  int8 `gorm:"default:1"` //品牌状态: 0删除, 1正常
	CreateTime   string
	UpdateTime   string
}

//set orm table name
func (TBrand) TableName() string {
	return "t_brand"
}
