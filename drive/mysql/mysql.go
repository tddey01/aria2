package mysql

import (
	"bytes"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	config2  "github.com/tddey01/aria2/config"
	"strconv"
)

var (
	Eloquent *gorm.DB
)

func init() {
	initAdmin()
}

//初始化管理系统数据库链接
func initAdmin() {

	var err error
	conn, dbType := Mysqlconn("admin")
	log.Println("管理系统数据库链接：" + conn)
	var db Database
	if dbType == "mysql" {
		db = new(Mysql)
	} else {
		panic("db type unknow")
	}

	Eloquent, err = db.Open(dbType, conn)
	Eloquent.LogMode(true)
	if err != nil {
		log.Fatalln("mysql admin connect error %v", err)
	} else {
		log.Println("mysql admin connect success!")
	}
	if Eloquent.Error != nil {
		log.Fatalln("database error %v", Eloquent.Error)
	}
	//config2.AdminBeegoOrmJoinMysql() //初始化beego 数据库链接

}


//数据库链接
func Mysqlconn(typesql string) (conns string, dbType string) {
	var host, database, username, password string
	var port int

	switch typesql {
	case "center":
		dbType = config2.GetConfig().Mysql.DBType
		host = config2.GetConfig().Mysql.Host
		port = config2.GetConfig().Mysql.Port
		database = config2.GetConfig().Mysql.DbName
		username = config2.GetConfig().Mysql.DbUser
		password = config2.GetConfig().Mysql.DbPasswd
	case "admin":
		dbType = config2.GetConfig().Mysql.DBType
		host = config2.GetConfig().Mysql.Host
		port = config2.GetConfig().Mysql.Port
		database = config2.GetConfig().Mysql.DbName
		username = config2.GetConfig().Mysql.DbUser
		password = config2.GetConfig().Mysql.DbPasswd
	}

	if dbType != "mysql" {
		fmt.Println("db type unknow")
	}

	var conn bytes.Buffer
	conn.WriteString(username)
	conn.WriteString(":")
	conn.WriteString(password)
	conn.WriteString("@tcp(")
	conn.WriteString(host)
	conn.WriteString(":")
	conn.WriteString(strconv.Itoa(port))
	conn.WriteString(")")
	conn.WriteString("/")
	conn.WriteString(database)
	conn.WriteString("?charset=utf8&parseTime=true&loc=Local&timeout=5s")
	conns = conn.String()
	return
}

type Database interface {
	Open(dbType string, conn string) (db *gorm.DB, err error)
}

type Mysql struct {
}

func (*Mysql) Open(dbType string, conn string) (db *gorm.DB, err error) {
	eloquent, err := gorm.Open(dbType, conn)
	return eloquent, err
}

type SqlLite struct {
}

func (*SqlLite) Open(dbType string, conn string) (db *gorm.DB, err error) {
	eloquent, err := gorm.Open(dbType, conn)
	return eloquent, err
}
