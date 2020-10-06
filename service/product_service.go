package service

import (
	"dragon/repository"
	"log"
)

type Product struct {
}

func (p *Product) GetList() interface{} {
	tx := repository.NewDefaultTx()
	productRepo := repository.ProductRepository{BaseRepository: repository.BaseRepository{
		TableName: repository.TProduct{}.TableName(),
		Tx:        tx,
	}}
	var product repository.TProduct
	var conditions = []map[string]interface{}{
		{"product_id = ?": 1},
		//{"create_time <= ?": "2019-08-01"},
	}
	productInfo, _ := productRepo.GetOne(&product,
		conditions, "product_id,product_name,create_time", "")
	log.Println(productInfo)

	var product2 repository.TProduct
	var conditions2 = []map[string]interface{}{
		{"product_id = ?": 2},
		//{"create_time <= ?": "2019-08-01"},
	}
	productInfo2, _ := productRepo.GetOne(&product2,
		conditions2, "product_id,product_name,create_time", "")

	return productInfo2

}
