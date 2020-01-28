package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gotest/common"
	"gotest/entity"
)
// 定义全局的engine
var dbm *xorm.Engine

func init() {
	//此处不能使用:=，若使用，定义的dbm是局部变量，没有实例化xrom.Engine
	// 虽然想要的意思此处只有异常是局部变量，实际都为局部变量
	var err error
	dbm, err = xorm.NewEngine("mysql", "root:123@/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	err = dbm.Sync2(new(entity.Member))
	if err != nil {
		common.Logger().Info("Sync2 db failed", err)
	}
	dbm.SetMaxIdleConns(10)
	dbm.SetMaxOpenConns(200)
	dbm.ShowSQL(true)
	dbm.ShowExecTime(true)
}