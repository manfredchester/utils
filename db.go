package main

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var orm *xorm.Engine

func CloudprojectEngine() *xorm.Engine {
	orm, err := GetEngine()
	if err != nil {
		fmt.Println(err)
	}
	return orm
}

func GetEngine() (*xorm.Engine, error) {
	if orm == nil {
		var err error
		orm, err = mysqlEngine()
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}
	if orm == nil {
		err := errors.New("database init error")
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}

	orm.ShowSQL()
	return orm, nil
}

func mysqlEngine() (*xorm.Engine, error) {
	Host := "10.128.0.180"
	Port := "3306"
	Name := "cloudproject"
	User := "root"
	Password := "Connext@0101"
	dburl := User + ":" + Password + "@tcp(" + Host + ":" + Port + ")/" + Name + "?charset=utf8"
	fmt.Println("dburl:", dburl)
	return xorm.NewEngine("mysql", dburl)
}
