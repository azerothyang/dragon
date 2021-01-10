package test

import (
	"dragon/core/dragon"
	"dragon/domain/repository"
	"dragon/domain/service"
	"testing"
)

func TestProductService_TransactionTest(t *testing.T) {
	dragon.AppInit()

	tx := repository.GormDB.Begin()
	productSrv := service.NewProductService(tx)
	productSrv.TransactionTest()

}
