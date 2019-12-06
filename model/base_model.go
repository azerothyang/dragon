package model

import (
	"dragon/core/dragon/conf"
	"dragon/core/dragon/dlogger"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入mysql驱动
	"github.com/jinzhu/gorm"
	"log"
	"sync"
)

var (
	db *gorm.DB //master db
)

const (
	StatusDelete = 0 //0表示删除
	StatusOK     = 1 //1表示正常
)

// 模型处理接口
type HandleModel interface {
	Add(data interface{})                                                                                                            // 新增
	SoftDelete(conditions []map[string]interface{}, field string, val interface{}) bool                                              // 软删除
	Delete(conditions []map[string]interface{}, model interface{}) bool                                                              // 真删除
	Updates(conditions []map[string]interface{}, data interface{}) bool                                                              // 通过编码伪删除
	GetList(list interface{}, conditions []map[string]interface{}, orderBy string, offset int, limit int, cols string)               // 查询列表
	GetOne(data interface{}, conditions []map[string]interface{}, cols string, orderBy string)                                       // 查询一条
	GetCount(conditions []map[string]interface{}) int64                                                                              // 获取总数
	GetListAndCount(list interface{}, conditions []map[string]interface{}, orderBy string, offset int, limit int, cols string) int64 // 获取总数
}

type BaseModel struct {
	TableName string //表名称
}

// todo 拼接sql的insert方法，看是否采用。 优点：可以极大的缩减代码量，将请求的参数直接写入数据库。 缺点，需要将请求的参数做参数过滤，同时前端传入的参数直接写入数据库，安全性没有结构体struct检测后更可靠。
//func (b *BaseModel) Add(data map[string]string) error {
//	var values []string
//	var fields []string
//	for field, value := range data {
//		fields = append(fields, field)
//		values = append(values, value)
//	}
//	sql := "INSERT INTO " + b.TableName + " ("
//	var sql2 = " VALUES (?)"
//	for _, field := range fields {
//		sql += field + ","
//	}
//	sql = sql[:len(sql)-1] + ")" + sql2
//	err := db.Exec(sql, values).Error
//	return err
//}

// 新增
func (b BaseModel) Add(data interface{}) {
	db.Create(data)
}

// 软删除通过条件
func (b BaseModel) SoftDelete(conditions []map[string]interface{}, field string, val interface{}) bool {
	queryDb := db.Table(b.TableName)
	for _, condition := range conditions {
		for cond, val := range condition {
			queryDb = queryDb.Where(cond, val)
		}
	}
	res := db.Update(field, val)
	if res.RowsAffected == 0 {
		return false
	}
	return true
}

// 真删除
func (b BaseModel) Delete(conditions []map[string]interface{}, model interface{}) bool {
	queryDb := db
	for _, condition := range conditions {
		for cond, val := range condition {
			queryDb = queryDb.Where(cond, val)
		}
	}
	res := queryDb.Delete(model)
	if res.RowsAffected == 0 {
		return false
	}
	return true
}

// 更新通过条件
func (b BaseModel) Updates(conditions []map[string]interface{}, data interface{}) bool {
	queryDb := db.Table(b.TableName)
	for _, condition := range conditions {
		for cond, val := range condition {
			queryDb = queryDb.Where(cond, val)
		}
	}
	res := queryDb.Updates(data)
	if res.RowsAffected == 0 {
		return false
	}
	return true
}

// 获取列表, limit=-1 拉取全部
func (b BaseModel) GetList(list interface{}, conditions []map[string]interface{}, orderBy string, offset int, limit int, cols string) {
	queryDb := db.Select(cols)
	for _, condition := range conditions {
		for cond, val := range condition {
			queryDb = queryDb.Where(cond, val)
		}
	}
	queryDb = queryDb.Order(orderBy).Offset(offset)
	if limit == -1 {
		queryDb.Find(list)
		return
	}
	queryDb.Limit(limit).Find(list)
	return
}

// 获取单个信息
func (b BaseModel) GetOne(data interface{}, conditions []map[string]interface{}, cols string, orderBy string) {
	queryDb := db.Select(cols)
	for _, condition := range conditions {
		for cond, val := range condition {
			queryDb = queryDb.Where(cond, val)
		}
	}
	// orderBy空字符，则无需加入Order条件
	if orderBy == "" {
		queryDb.Find(data)
		return
	}

	queryDb.Order(orderBy).First(data)
}

// 获取总数
func (b BaseModel) GetCount(conditions []map[string]interface{}) int64 {
	var count int64
	queryDb := db.Table(b.TableName)
	for _, condition := range conditions {
		for cond, val := range condition {
			queryDb = queryDb.Where(cond, val)
		}
	}
	queryDb.Count(&count)
	return count
}

// 获取列表和总数
func (b BaseModel) GetListAndCount(list interface{}, conditions []map[string]interface{}, orderBy string, offset int, limit int, cols string) int64 {
	var wg sync.WaitGroup
	var count int64
	wg.Add(2)
	go func() {
		b.GetList(list, conditions, orderBy, offset, limit, cols)
		wg.Done()
	}()
	go func() {
		count = b.GetCount(conditions)
		wg.Done()
	}()
	wg.Wait()
	return count
}

// sql logger
type Logger struct {
}

func (Logger) Print(values ...interface{}) {
	// todo 更好的日志打印方案
	logInfo := fmt.Sprint(values)
	dlogger.Info(logInfo)
	log.Println(values...)
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
