package service

// 结果集结构
type ResultData struct {
	Status  int
	Message string
	Data    interface{}
}

// Api 返回结果集
func Result(s int, m string, d interface{}) ResultData {
	var result ResultData
	result.Status = s  // 状态码
	result.Message = m // 消息
	result.Data = d    // 数据集
	return result
}
