package repository

import (
	"dragon/domain/entity"
	"gorm.io/gorm"
)

// 商品仓库
type IProductRepository interface {
	IBaseRepository
}

type ProductRepository struct {
	BaseRepository
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{
		BaseRepository: BaseRepository{
			TableName: entity.ProductEntity{}.TableName(),
			MysqlDB:   db,
		},
	}
}
