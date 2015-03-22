package models

import (
	"errors"
	"github.com/go-xorm/xorm"
)

var (
	ErrNotExist = errors.New("not exist")
)

var orm *xorm.Engine

func Init(isProMode bool) {
	var err error
	orm, err = xorm.NewEngine("mysql", "root:123456@/basesystem?charset=utf8")
	if err != nil {
		panic(err)
	}

	if !isProMode {
		orm.ShowSQL = true
	}

	orm.ShowErr = true
	orm.ShowInfo = true
}
