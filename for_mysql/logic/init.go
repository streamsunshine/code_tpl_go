package logic

import (
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var DataSQLObj *gorm.DB

func InitMySQL() error {
	var err error
	DataSQLObj, err = gorm.Open("mysql", "name:password@tcp(host:port)/defaultDbname?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		return err
	}

	//todo 连接数数量需要斟酌下 这个
	DataSQLObj.DB().SetMaxOpenConns(100)
	DataSQLObj.DB().SetMaxIdleConns(100)
	DataSQLObj.LogMode(false)


	DataSQLObj.SetLogger(log.New(os.Stdout, "\r\n", 0))

	err = DataSQLObj.DB().Ping()
	if err != nil {
		return err
	}

	return nil
}