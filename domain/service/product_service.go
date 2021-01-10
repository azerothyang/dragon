package service

import (
	"dragon/domain/entity"
	"dragon/domain/repository"
	"errors"
	"github.com/go-dragon/util"
	"gorm.io/gorm"
	"time"
)

// ProductService interface
type IProductService interface {
	GetOne() (*entity.ProductEntity, error)
	TransactionTest() error
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

// test transaction easy demo
func (p *ProductService) TransactionTest() error {
	productInfo := entity.ProductEntity{
		ProductCode:   util.RandomStr(10),
		ProductName:   "",
		BrandId:       0,
		ProductStatus: 0,
		CreateTime:    time.Now().Format("2006-01-02 15:04:05"),
		UpdateTime:    time.Now().Format("2006-01-02 15:04:05"),
	}
	err1 := p.ProductRepository.Add(&productInfo)
	productInfo = entity.ProductEntity{
		ProductCode:   util.RandomStr(10),
		ProductName:   "",
		BrandId:       0,
		ProductStatus: 0,
		CreateTime:    time.Now().Format("2006-01-02 15:04:05"),
		UpdateTime:    time.Now().Format("2006-01-02 15:04:05"),
	}
	err2 := p.ProductRepository.Add(&productInfo)
	err1 = errors.New("manual error")
	if err1 != nil || err2 != nil {
		p.TxConnDB.Rollback()
		return errors.New("data write fail")
	}
	p.TxConnDB.Commit()
	return nil
}
