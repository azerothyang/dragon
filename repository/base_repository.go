package repository

import (
	"dragon/core/dragon/conf"
	"dragon/core/dragon/dlogger"
	"errors"
	"fmt"
	"gorm.io/driver/mysql" //导入mysql驱动
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"regexp"
	"sync"
	"time"
)

var (
	GormDB *gorm.DB //master db
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

// new default tx
func NewDefaultTx() *gorm.DB {
	return GormDB.Session(&gorm.Session{
		PrepareStmt:            true,
		WithConditions:         true,
		SkipDefaultTransaction: true,
		Context:                GormDB.Statement.Context,
	})
}

// 判断是否有致命错误，致命错误 不包含 (查询结果为空)， 错误不要进行log.Fatal处理，这个会让进程挂掉
func HasSeriousError(res *gorm.DB) bool {
	if res.Error != nil && (!errors.Is(res.Error, gorm.ErrRecordNotFound)) {
		return true
	}
	return false
}

type BaseRepository struct {
	TableName string //表名称
	Tx        *gorm.DB
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
func (b *BaseRepository) Add(data interface{}) error {
	return b.Tx.Create(data).Error
}

// 软删除通过条件
func (b *BaseRepository) SoftDelete(conditions []map[string]interface{}, field string, val interface{}) *gorm.DB {
	queryDb := b.Tx.Table(b.TableName)
	for _, condition := range conditions {
		for cond, val := range condition {
			queryDb = queryDb.Where(cond, val)
		}
	}
	return queryDb.Update(field, val)
}

// 真删除
func (b *BaseRepository) Delete(conditions []map[string]interface{}, model interface{}) *gorm.DB {
	queryDb := b.Tx
	for _, condition := range conditions {
		for cond, val := range condition {
			queryDb = queryDb.Where(cond, val)
		}
	}
	return queryDb.Delete(model)
}

// 更新通过条件
// https://gorm.io/zh_CN/docs/update.html
func (b *BaseRepository) Updates(conditions []map[string]interface{}, data map[string]interface{}) *gorm.DB {
	queryDb := b.Tx.Table(b.TableName)
	for _, condition := range conditions {
		for cond, val := range condition {
			queryDb = queryDb.Where(cond, val)
		}
	}
	return queryDb.Updates(data)
}

// 获取列表, limit=-1 拉取全部。 res返回为 可以根据
func (b *BaseRepository) GetList(list interface{}, conditions []map[string]interface{}, orderBy string, offset int, limit int, cols string) *gorm.DB {
	queryDb := b.Tx.Select(cols)
	for _, condition := range conditions {
		for cond, val := range condition {
			queryDb = queryDb.Where(cond, val)
		}
	}
	if orderBy != "" {
		queryDb = queryDb.Order(orderBy)
	}
	queryDb = queryDb.Offset(offset)
	if limit == -1 {
		return queryDb.Find(list)
	}
	return queryDb.Limit(limit).Find(list)
}

// 获取单个信息
func (b *BaseRepository) GetOne(data interface{}, conditions []map[string]interface{}, cols string, orderBy string) *gorm.DB {
	queryDb := b.Tx.Select(cols)
	for _, condition := range conditions {
		for cond, val := range condition {
			queryDb = queryDb.Where(cond, val)
		}
	}
	// orderBy空字符，则无需加入Order条件
	if orderBy == "" {
		return queryDb.First(data)
	}

	return queryDb.Order(orderBy).First(data)
}

// 获取总数
func (b *BaseRepository) GetCount(conditions []map[string]interface{}) (count int64, res *gorm.DB) {
	queryDb := b.Tx.Table(b.TableName)
	for _, condition := range conditions {
		for cond, val := range condition {
			queryDb = queryDb.Where(cond, val)
		}
	}
	res = queryDb.Count(&count)
	return
}

// 获取列表和总数，此方法不能用于事务
func (b *BaseRepository) GetListAndCount(list interface{}, conditions []map[string]interface{}, orderBy string, offset int, limit int, cols string) (count int64, listRes *gorm.DB, countRes *gorm.DB) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		repo := BaseRepository{
			TableName: b.TableName,
			Tx:        NewDefaultTx(),
		}
		listRes = repo.GetList(list, conditions, orderBy, offset, limit, cols)
		wg.Done()
	}()
	go func() {
		repo := BaseRepository{
			TableName: b.TableName,
			Tx:        NewDefaultTx(),
		}
		count, countRes = repo.GetCount(conditions)
		wg.Done()
	}()
	wg.Wait()
	return
}

// sql logger
type Logger struct {
	logger.Writer
}

func (l Logger) Printf(s string, i ...interface{}) {
	s = fmt.Sprintf(s, i...)
	// 日志打印
	res, _ := regexp.MatchString("Error", s)
	// if sql error
	if res {
		dlogger.SqlError(s)
	} else {
		dlogger.SqlInfo(s)
	}
}

//init db
func InitDB() {
	var err error
	var dsnMaster string
	var logHandler logger.Interface
	if conf.Env == "dev" {
		logHandler = logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: 100 * time.Millisecond,
			Colorful:      true,
			LogLevel:      logger.Info,
		})
	} else {
		// other env write log
		logHandler = logger.New(Logger{}, logger.Config{
			SlowThreshold: 100 * time.Millisecond,
			Colorful:      false,
			LogLevel:      logger.Info,
		})
	}

	//mysql master
	dsnMaster = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&timeout=%s&loc=Local", //loc set the timezone
		conf.Conf.Database.Mysql.Master.User, conf.Conf.Database.Mysql.Master.Password, conf.Conf.Database.Mysql.Master.Host, conf.Conf.Database.Mysql.Master.Port, conf.Conf.Database.Mysql.Master.Database, conf.Conf.Database.Mysql.Master.Charset, conf.Conf.Database.Mysql.Master.Timeout)

	//gorm realizes mysql reconnect
	GormDB, err = gorm.Open(mysql.New(mysql.Config{
		DriverName:                "mysql",
		DSN:                       dsnMaster,
		Conn:                      nil,
		SkipInitializeWithVersion: false,
		DefaultStringSize:         0,
		DisableDatetimePrecision:  false,
		DontSupportRenameIndex:    false,
		DontSupportRenameColumn:   false,
	}), &gorm.Config{
		SkipDefaultTransaction:                   true,
		NamingStrategy:                           nil,
		Logger:                                   logHandler,
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              true,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		AllowGlobalUpdate:                        false,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	})
	if err != nil {
		log.Fatalln(err)
	}
	sqlDb, err := GormDB.DB()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("mysql maxIdle conns:", conf.Conf.Database.Mysql.Master.Maxidle)
	log.Println("mysql maxOpenConn conns:", conf.Conf.Database.Mysql.Master.Maxconn)
	sqlDb.SetMaxIdleConns(conf.Conf.Database.Mysql.Master.Maxidle)
	sqlDb.SetMaxOpenConns(conf.Conf.Database.Mysql.Master.Maxconn)
	sqlDb.SetConnMaxIdleTime(time.Hour)
	sqlDb.SetConnMaxLifetime(24 * time.Hour)
}
