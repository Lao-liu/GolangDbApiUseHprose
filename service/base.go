package service

import (
	"GolangDbApiUseHprose/models"
)

type Base struct{}

// 通用根据单字段String为条件进行赋值查询
func (Base) GetTbByFieldString(tb, field, value string) ResultData {
	var err error
	sql := `SELECT * FROM ` + tb + ` WHERE ` + field + `=?`
	data, err := models.ORM().Query(sql, value)
	if err != nil {
		return Result(500, err.Error(), nil)
	}
	if len(data) == 0 {
		return Result(404, "未找到数据", nil)
	}
	return Result(200, "找到记录", &data)
}

// 通用根据单字段Int为条件进行赋值查询
func (Base) GetTbByFieldInt(tb, field string, value int) ResultData {
	var err error
	sql := `SELECT * FROM ` + tb + ` WHERE ` + field + `=?`
	data, err := models.ORM().Query(sql, value)
	if err != nil {
		return Result(500, err.Error(), nil)
	}
	if len(data) == 0 {
		return Result(404, "未找到数据", nil)
	}
	return Result(200, "找到记录", &data)
}
