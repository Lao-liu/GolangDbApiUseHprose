package main

import (
	"GolangDbApiUseHprose/models"
	"GolangDbApiUseHprose/service"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hprose/hprose-golang/rpc"
	// "github.com/xxtea/xxtea-go/xxtea"
)

// 公共变量
var (
	servicePort   = ":4321"
	runDebugModel = true
	log           = logrus.New()
)

// ServerEvent Hprose 事件接口实现
type ServerEvent struct{}

// OnBeforeInvoke 调用前执行
func (e *ServerEvent) OnBeforeInvoke(name string, args []reflect.Value, byref bool, context rpc.Context) {
	log.WithFields(logrus.Fields{
		"FuncName": name,
	}).Info(args)
}

// OnAfterInvoke 调用后执行
func (e *ServerEvent) OnAfterInvoke(name string, args []reflect.Value, byref bool, result []reflect.Value, context rpc.Context) {
	log.WithFields(logrus.Fields{
		"FuncName": name,
	}).Info(result)
}

// OnSendError 服务端出错时执行
func (e *ServerEvent) OnSendError(err error, context rpc.Context) {
	log.Error(err.Error())
}

// ServerFilter Hprose 过滤接口实现
type ServerFilter struct{}

// InputFilter 输入过滤
func (f *ServerFilter) InputFilter(data []byte, context rpc.Context) []byte {
	log.WithFields(logrus.Fields{
		"DATA":    string(data),
		"CONTEXT": context,
	}).Info("INPUT")
	return data
}

// OutputFilter 输出过滤
func (f *ServerFilter) OutputFilter(data []byte, context rpc.Context) []byte {
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

// MissFunctions 迷失方法发布
func MissFunctions(name string, args []reflect.Value, context rpc.Context) (result []reflect.Value) {
	result = make([]reflect.Value, 1)
	switch strings.Title(name) {
	case "Query":
		result[0] = reflect.ValueOf(Query(args[0].Interface().(string)))
	default:
		result[0] = reflect.ValueOf(service.Result(500, "调用"+name+"方法不存在", nil))
	}
	return
}

// Query 执行SQL查询
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
	server := rpc.NewHTTPService()
	// 服务器端工具 ServiceTool Api
	server.AddInstanceMethods(service.ServiceTool{})
	// 工作台接口 Bench Api
	server.AddInstanceMethods(service.Bench{})
	// 添加隐含的迷失方法
	server.AddMissingMethod(MissFunctions, rpc.Options{Mode: rpc.Raw})
	// 加密传输
	// service.SetFilter(XXTEAFilter{"123456790!@#$%^&"});
	// 开发模式下启用调试
	if runDebugModel {
		server.Event = &ServerEvent{}
		server.AddFilter(&ServerFilter{})
		server.Debug = true
	}
	fmt.Println("Run GolangDbApiUseHprose on port http://127.0.0.1" + servicePort + " ...")
	http.ListenAndServe(servicePort, server)
}
