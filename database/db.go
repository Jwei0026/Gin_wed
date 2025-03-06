package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//数据库初始化操作
//1.定义自己的数据对象结构
//2.完成数据库连接初始化
//3.制作数据库的相关方法

//连接认证

const (
	DRIVER   = "mysql"
	HOST     = "127.0.0.1"
	Port     = "3306"
	USERNAME = "root"
	PASSWORD = "123asd!@#"
	PROTPCPL = "tcp"
	DATABASE = "shop"
	CHARSET  = "utf8"
)

// 使用gorm进行数据库连接认证
var Gdb *gorm.DB

func init() {
	//配置DSN
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USERNAME,
		PASSWORD, HOST, Port, DATABASE, CHARSET)

	//数据库连接
	var err error
	Gdb, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		//log.Fatalf("failed to conect database: %v",err)  生产环境写入日志
		panic("数据库连接认证是失败" + err.Error())
	}

}
