package service

import (
	"dragon/domain/entity"
	"dragon/domain/repository"
	"gorm.io/gorm"
)

// ProductService interface
type IProductService interface {
	GetOne() (*entity.ProductEntity, error)
}

type ProductService struct {
	ProductRepository repository.IProductRepository
	TxConnDB          *gorm.DB //当前服务所用的一条DB连接，用于事务处理
}

func NewProductService(txConnDB *gorm.DB) IProductService {
	return &ProductService{
		ProductRepository: repository.NewProductRepository(txConnDB),
		TxConnDB:          txConnDB,
	}
}

// 获取一条 todo 测试事务
func (p *ProductService) GetOne() (*entity.ProductEntity, error) {
	var product entity.ProductEntity
	var conditions = []map[string]interface{}{
		{"product_id = ?": 1},
		//{"create_time <= ?": "2019-08-01"},
	}
	res := p.ProductRepository.GetOne(&product, conditions, "product_id,product_name,create_time", "")
	if repository.HasSeriousError(res) {
		return &product, res.Error
	}
	return &product, nil
}
