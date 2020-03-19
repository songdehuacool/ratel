package utils

import (
	"encoding/json"
	"io/ioutil"
	"ratel/riface"
)

/**
* @author     ：songdehua
* @emall      ：200637086@qq.com
* @date       ：Created in 2020/3/19 4:09 下午
* @description：存储一切有关Ratel框架的全局参数，供其它模块使用
                一些参数可以通过ratel.json由用户进行配置
* @modified By：
* @version    ：$
*/
type GlobalObj struct {
	/*
		Server
	*/
	TcpServer riface.IServer // 当前Ratel全局的Server对象
	Host      string         // 当前服务器主机监听的IP
	TcpPort   int            // 当前服务器主机监听的端口号
	Name      string         // 当前服务器的名称

	/*
		Ratel
	*/
	Version        string // 当前Ratel的版本号
	MaxCount       int    // 当前服务器主机允许的最大链接数
	MaxPackageSize uint32 // 当前Ratel框架数据包的最大值
}

/*
	定义一个全局的对外GlobalObj
*/
var GlobalObject *GlobalObj

/*
	从ratel.json取加载用于自定义的参数
*/
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/ratel.json")
	// 将json文件数据解析到struct中
	json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

/*
	提供一个init方法，初始化当前的GlobalObject
*/
func init() {
	// 如果配置文件没有加载，默认的值
	GlobalObject = &GlobalObj{
		Host:           "0.0.0.0",
		TcpPort:        8999,
		Name:           "RatelServer",
		Version:        "V0.4",
		MaxCount:       1000,
		MaxPackageSize: 4096,
	}

	// 应该尝试从conf/ratel.json取加载一些用户自定义的参数
	GlobalObject.Reload()
}
