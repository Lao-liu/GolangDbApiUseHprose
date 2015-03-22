package service

import (
	"GolangDbApiUseHprose/models"
)

type Bench struct{}

// 根据机构ID，获取结构信息。
func (Bench) GetOrgInfo(orgId string) ResultData {
	if orgId == "" {
		return Result(500, "人员机构ID不可为空", nil)
	}
	return new(Base).GetTbByFieldString("Dict_Organization", "Orgid", orgId)
}
