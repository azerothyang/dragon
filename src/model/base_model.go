package model

import (
	"core/dragon"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入mysql驱动
	"github.com/jinzhu/gorm"
	"log"
)
var (
	db  *gorm.DB 	 //master db
	readDB *gorm.DB  //read db
)

type baseModel struct {

}

//init db
func InitDB() {
	var errM, errS error
	var dsnMaster, dsnSlave string
	var masterMaxIdle, masterMaxConn, slaveMaxIdle, slaveMaxConn int

	//mysql master
	dsnMaster = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&timeout=%s&parseTime=True&loc=Local", //loc set the timezone
		dragon.Conf.Database.Mysql.Master.User, dragon.Conf.Database.Mysql.Master.Password, dragon.Conf.Database.Mysql.Master.Host, dragon.Conf.Database.Mysql.Master.Port, dragon.Conf.Database.Mysql.Master.Database, dragon.Conf.Database.Mysql.Master.Charset, dragon.Conf.Database.Mysql.Master.Timeout)

	//mysql slave
	dsnSlave = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&timeout=%s&parseTime=True&loc=Local",
		dragon.Conf.Database.Mysql.Slave.User, dragon.Conf.Database.Mysql.Slave.Password, dragon.Conf.Database.Mysql.Slave.Host, dragon.Conf.Database.Mysql.Slave.Port, dragon.Conf.Database.Mysql.Slave.Database, dragon.Conf.Database.Mysql.Slave.Charset, dragon.Conf.Database.Mysql.Slave.Timeout)
	masterMaxIdle = dragon.Conf.Database.Mysql.Master.MaxIdle
	masterMaxConn = dragon.Conf.Database.Mysql.Master.MaxConn
	slaveMaxIdle = dragon.Conf.Database.Mysql.Slave.MaxIdle
	slaveMaxConn = dragon.Conf.Database.Mysql.Slave.MaxConn

	//gorm realizes mysql reconnect
	db, errM = gorm.Open("mysql", dsnMaster)
	readDB, errS = gorm.Open("mysql", dsnSlave)
	if errM != nil || errS != nil {
		log.Fatalln(errM, errS)
	}

	db.DB().SetMaxIdleConns(masterMaxIdle)
	db.DB().SetMaxOpenConns(masterMaxConn)
	readDB.DB().SetMaxIdleConns(slaveMaxIdle)
	readDB.DB().SetMaxOpenConns(slaveMaxConn)
}