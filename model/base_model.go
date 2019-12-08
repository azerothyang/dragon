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
	Add(data interface{}) error                                                                                                                                                                  // 新增
	SoftDelete(conditions []map[string]interface{}, field string, val interface{}) (err error)                                                                                                   // 软删除
	Delete(conditions []map[string]interface{}, model interface{}) (err error)                                                                                                                   // 真删除
	Updates(conditions []map[string]interface{}, data interface{}) (err error)                                                                                                                   // 通过编码伪删除
	GetList(list interface{}, conditions []map[string]interface{}, orderBy string, offset int, limit int, cols string) (resData interface{}, err error)                                          // 查询列表
	GetOne(data interface{}, conditions []map[string]interface{}, cols string, orderBy string) (resData interface{}, err error)                                                                  // 查询一条
	GetCount(conditions []map[string]interface{}) (count int64, err error)                                                                                                                       // 获取总数
	GetListAndCount(list interface{}, conditions []map[string]interface{}, orderBy string, offset int, limit int, cols string) (resData interface{}, count int64, listErr error, countErr error) // 获取总数
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
func (b BaseModel) Add(data interface{}) error {
	return db.Create(data).Error
}

// 软删除通过条件
func (b BaseModel) SoftDelete(conditions []map[string]interface{}, field string, val interface{}) (err error) {
	queryDb := db.Table(b.TableName)
	for _, condition := range conditions {
		for cond, val := range condition {
			queryDb = queryDb.Where(cond, val)
		}
	}
	res := db.Update(field, val)
	return res.Error
}

// 真删除
func (b BaseModel) Delete(conditions []map[string]interface{}, model interface{}) (err error) {
	queryDb := db
	for _, condition := range conditions {
		for cond, val := range condition {
			queryDb = queryDb.Where(cond, val)
		}
	}
	res := queryDb.Delete(model)
	return res.Error
}

// 更新通过条件
func (b BaseModel) Updates(conditions []map[string]interface{}, data interface{}) (err error) {
	queryDb := db.Table(b.TableName)
	for _, condition := range conditions {
		for cond, val := range condition {
			queryDb = queryDb.Where(cond, val)
		}
	}
	res := queryDb.Updates(data)
	return res.Error
}

// 获取列表, limit=-1 拉取全部
func (b BaseModel) GetList(list interface{}, conditions []map[string]interface{}, orderBy string, offset int, limit int, cols string) (resData interface{}, err error) {
	queryDb := db.Select(cols)
	for _, condition := range conditions {
		for cond, val := range condition {
			queryDb = queryDb.Where(cond, val)
		}
	}
	queryDb = queryDb.Order(orderBy).Offset(offset)
	if limit == -1 {
		res := queryDb.Find(list)
		if res.Error == gorm.ErrRecordNotFound {
			// 如果是记录未找到，则直接返回空,而不返回错误,list也改为nil
			return nil, nil
		}
		return list, res.Error
	}
	res := queryDb.Limit(limit).Find(list)
	if res.Error == gorm.ErrRecordNotFound {
		// 如果是记录未找到，则直接返回空,而不返回错误,list也改为nil
		return nil, nil
	}
	return list, res.Error
}

// 获取单个信息
func (b BaseModel) GetOne(data interface{}, conditions []map[string]interface{}, cols string, orderBy string) (resData interface{}, err error) {
	queryDb := db.Select(cols)
	for _, condition := range conditions {
		for cond, val := range condition {
			queryDb = queryDb.Where(cond, val)
		}
	}
	// orderBy空字符，则无需加入Order条件
	if orderBy == "" {
		res := queryDb.Find(data)
		if res.Error == gorm.ErrRecordNotFound {
			// 如果是记录未找到，则直接返回空,而不返回错误,list也改为nil
			return nil, nil
		}
		return data, res.Error
	}

	res := queryDb.Order(orderBy).First(data)
	if res.Error == gorm.ErrRecordNotFound {
		// 如果是记录未找到，则直接返回空,而不返回错误,list也改为nil
		return nil, nil
	}
	return data, res.Error
}

// 获取总数
func (b BaseModel) GetCount(conditions []map[string]interface{}) (count int64, err error) {
	queryDb := db.Table(b.TableName)
	for _, condition := range conditions {
		for cond, val := range condition {
			queryDb = queryDb.Where(cond, val)
		}
	}
	res := queryDb.Count(&count)
	err = res.Error
	return
}

// 获取列表和总数
func (b BaseModel) GetListAndCount(list interface{}, conditions []map[string]interface{}, orderBy string, offset int, limit int, cols string) (resData interface{}, count int64, listErr error, countErr error) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		resData, listErr = b.GetList(list, conditions, orderBy, offset, limit, cols)
		wg.Done()
	}()
	go func() {
		count, countErr = b.GetCount(conditions)
		wg.Done()
	}()
	wg.Wait()
	return
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

	//mysql master
	dsnMaster = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&timeout=%s&parseTime=True&loc=Local", //loc set the timezone
		conf.Conf.Database.Mysql.Master.User, conf.Conf.Database.Mysql.Master.Password, conf.Conf.Database.Mysql.Master.Host, conf.Conf.Database.Mysql.Master.Port, conf.Conf.Database.Mysql.Master.Database, conf.Conf.Database.Mysql.Master.Charset, conf.Conf.Database.Mysql.Master.Timeout)

	//gorm realizes mysql reconnect
	db, err = gorm.Open("mysql", dsnMaster)
	if err != nil {
		log.Fatalln(err)
	}

	db.DB().SetMaxIdleConns(conf.Conf.Database.Mysql.Master.MaxIdle)
	db.DB().SetMaxOpenConns(conf.Conf.Database.Mysql.Master.MaxConn)

	//如果是debug模式则开启彩色sql调试模式, 否则为文本模式
	db.LogMode(true)
	logger := Logger{}
	if conf.Env != "debug" {
		db.SetLogger(logger)
	}
}
