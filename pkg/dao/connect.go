/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package dao

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db_Orm *gorm.DB //连接池对象
func GetDBConn() *gorm.DB {
	if db_Orm != nil {
		return db_Orm
	}
	dsn := "root:123456@tcp(127.0.0.1:3306)/publicX?charset=utf8mb4&parseTime=True&loc=Local" //MO
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// get connection
	if err != nil {
		fmt.Println("Database Connection Failed") //Connection failed
	} else {
		fmt.Println("Database Connection Succeed") //Connection succeed
		db_Orm = db
	}
	return db_Orm
}
