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
	BrandId      int64  `gorm:"primary_key;AUTO_INCREMENT" json:"brand_id"`
	BrandCode    string `gorm:"unique_index" json:"brand_code"`
	BrandName    string `json:"brand_name"`
	BrandProfile string `json:"brand_profile"`
	AppId        string `json:"app_id"`
	AppSecret    string `json:"app_secret"`
	BrandStatus  int8   `gorm:"default:1" json:"brand_status"` //品牌状态: 0删除, 1正常
	CreateTime   string `json:"create_time"`
	UpdateTime   string `json:"update_time"`
}

//set orm table name
func (TBrand) TableName() string {
	return "t_brand"
}
