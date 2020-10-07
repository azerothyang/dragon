package serviceimpl

import (
	"dragon/repository"
	"fmt"
	"log"
)

type ProductServiceImpl struct {
}

func (p *ProductServiceImpl) GetList() interface{} {
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
	res := productRepo.GetOne(&product,
		conditions, "product_id,product_name,create_time", "")
	if repository.HasFatalError(res) {
		log.Fatal(res.Error)
	}
	log.Println("product", fmt.Sprintf("%+v", product))

	return product

}
