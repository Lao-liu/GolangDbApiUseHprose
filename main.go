package main

import (
	"GolangDbApiUseHprose/models"
	"GolangDbApiUseHprose/service"
	"fmt"
	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hprose/hprose-go/hprose"
	"net/http"
	"reflect"
	"strings"
	// "github.com/xxtea/xxtea-go/xxtea"
)

// 公共变量
var (
	servicePort   = ":4321"
	runDebugModel = true
	log           = logrus.New()
)

// Hprose 事件接口实现
type ServerEvent struct{}

func (e *ServerEvent) OnBeforeInvoke(name string, args []reflect.Value, byref bool, context hprose.Context) {
	log.WithFields(logrus.Fields{
		"FuncName": name,
	}).Info(args)
}
func (e *ServerEvent) OnAfterInvoke(name string, args []reflect.Value, byref bool, result []reflect.Value, context hprose.Context) {
	log.WithFields(logrus.Fields{
		"FuncName": name,
	}).Info(result)
}
func (e *ServerEvent) OnSendError(err error, context hprose.Context) {
	log.Error(err.Error())
}

// Hprose 过滤接口实现
type ServerFilter struct{}

func (f *ServerFilter) InputFilter(data []byte, context hprose.Context) []byte {
	log.WithFields(logrus.Fields{
		"DATA":    string(data),
		"CONTEXT": context,
	}).Info("INPUT")
	return data
}
func (f *ServerFilter) OutputFilter(data []byte, context hprose.Context) []byte {
	log.WithFields(logrus.Fields{
		"DATA":    string(data),
		"CONTEXT": context,
	}).Info("OUTPUT")
	return data
}

// 使用XXTEA 加密传输数据

// service.SetFilter(XXTEAFilter{"123456790!@#$%^&"});

// type XXTEAFilter struct {
//     Key string
// }

// func (filter XXTEAFilter) InputFilter(data []byte, context Context) []byte {
//     return xxtea.Decrypt(data, []byte(filter.Key));
// }

// func (filter XXTEAFilter) OutputFilter(data []byte, context Context) []byte {
//     return xxtea.Encrypt(data, []byte(filter.Key));
// }

// 迷失方法发布
func MissFunctions(name string, args []reflect.Value) (result []reflect.Value) {
	result = make([]reflect.Value, 1)
	switch strings.Title(name) {
	case "Query":
		result[0] = reflect.ValueOf(Query(args[0].Interface().(string)))
	default:
		result[0] = reflect.ValueOf(service.Result(500, "调用"+name+"方法不存在", nil))
	}
	return
}

// 执行SQL查询
func Query(sql string) service.ResultData {
	if sql == "" {
		return service.Result(500, "查询语句不可为空", nil)
	}
	results, err := models.ORM().Query(sql)
	if err != nil {
		return service.Result(500, err.Error(), nil)
	}
	if len(results) == 0 {
		return service.Result(404, "未找到数据", nil)
	}
	return service.Result(200, "找到数据", results)
}

func main() {
	// 初始化Model
	models.Init(false)
	// 初始化Hprose
	server := hprose.NewHttpService()
	// 服务器端工具 ServerTool Api
	server.AddMethods(service.ServerTool{})
	// 工作台接口 Bench Api
	server.AddMethods(service.Bench{})
	// 添加隐含的迷失方法
	server.AddMissingMethod(MissFunctions, true)
	// 加密传输
	// service.SetFilter(XXTEAFilter{"123456790!@#$%^&"});
	// 开发模式下启用调试
	if runDebugModel {
		server.ServiceEvent = &ServerEvent{}
		server.AddFilter(&ServerFilter{})
		server.DebugEnabled = true
	}
	fmt.Println("Run GolangDbApiUseHprose on port http://127.0.0.1" + servicePort + " ...")
	http.ListenAndServe(servicePort, server)
}
