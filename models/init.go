package models

import (
	"errors"

	"github.com/go-xorm/xorm"
)

var (
	// ErrNotExist 不存在
	ErrNotExist = errors.New("not exist")
)

var orm *xorm.Engine

// Init 初始化数据库
func Init(isProMode bool) {
	var err error
	orm, err = xorm.NewEngine("mysql", "root:123456@/basesystem?charset=utf8")
	if err != nil {
		panic(err)
	}

	if !isProMode {
		orm.ShowSQL(true)
	}
}
