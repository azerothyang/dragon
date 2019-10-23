package model

import (
	"dragon/core/dragon/conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入mysql驱动
	"github.com/jinzhu/gorm"
	"log"
)

var (
	db *gorm.DB //master db
)

// 模型处理接口
type HandleModel interface {
	Add(data interface{})                                                                                            // 新增
	SoftDelete(model interface{}, conditions map[string]interface{}, field string, val interface{}) bool             // 软删除
	Delete(conditions map[string]interface{}, model interface{}) bool                                                // 真删除
	Updates(conditions map[string]interface{}, data interface{}) bool                                                // 通过编码伪删除
	GetList(list interface{}, conditions map[string]interface{}, orderBy string, offset int, limit int, cols string) // 查询列表
	GetOne(data interface{}, conditions map[string]interface{}, cols string, orderBy string)                         // 查询一条
	GetTotalCount(conditions map[string]interface{}) int64                                                           // 获取总数
}

type BaseModel struct {
	TableName string //表名称
}

// 新增
func (b *BaseModel) Add(data interface{}) {
	db.Create(data)
}

// 软删除通过条件
func (b *BaseModel) SoftDelete(model interface{}, conditions map[string]interface{}, field string, val interface{}) bool {
	res := db.Model(model).Where(conditions).Update(field, val)
	if res.RowsAffected == 0 {
		return false
	}
	return true
}

// 真删除
func (b *BaseModel) Delete(conditions map[string]interface{}, model interface{}) bool {
	res := db.Delete(model, conditions)
	if res.RowsAffected == 0 {
		return false
	}
	return true
}

// 更新通过条件
func (b *BaseModel) Updates(conditions map[string]interface{}, data interface{}) bool {
	res := db.Table(b.TableName).Where(conditions).Updates(data)
	if res.RowsAffected == 0 {
		return false
	}
	return true
}

// 获取列表, limit=-1 拉取全部
func (b *BaseModel) GetList(list interface{}, conditions map[string]interface{}, orderBy string, offset int, limit int, cols string) {
	if limit == -1 {
		db.Select(cols).Where(conditions).Order(orderBy).Offset(offset).Find(list)
		return
	}
	db.Select(cols).Where(conditions).Order(orderBy).Offset(offset).Limit(limit).Find(list)
	return
}

// 获取单个信息
func (b *BaseModel) GetOne(data interface{}, conditions map[string]interface{}, cols string, orderBy string) {
	db.Select(cols).Order(orderBy).First(data, conditions)
}

// 获取总数
func (b *BaseModel) GetTotalCount(conditions map[string]interface{}) int64 {
	var count int64
	db.Table(b.TableName).Where(conditions).Count(&count)
	return count
}

// sql logger
type Logger struct {
}

func (Logger) Print(v ...interface{}) {
	// todo 更好的日志打印方案
	log.Println(v...)
}

//init db
func InitDB() {
	var err error
	var dsnMaster string
	var masterMaxIdle, masterMaxConn int

	//mysql master
	dsnMaster = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&timeout=%s&parseTime=True&loc=Local", //loc set the timezone
		conf.Conf.Database.Mysql.Master.User, conf.Conf.Database.Mysql.Master.Password, conf.Conf.Database.Mysql.Master.Host, conf.Conf.Database.Mysql.Master.Port, conf.Conf.Database.Mysql.Master.Database, conf.Conf.Database.Mysql.Master.Charset, conf.Conf.Database.Mysql.Master.Timeout)

	//gorm realizes mysql reconnect
	db, err = gorm.Open("mysql", dsnMaster)
	if err != nil {
		log.Fatalln(err)
	}

	db.DB().SetMaxIdleConns(masterMaxIdle)
	db.DB().SetMaxOpenConns(masterMaxConn)

	//如果是debug模式则开启彩色sql调试模式, 否则为文本模式
	db.LogMode(true)
	logger := Logger{}
	if conf.Env != "debug" {
		db.SetLogger(logger)
	}
}
