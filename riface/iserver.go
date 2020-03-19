package riface

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2020/3/16 11:22 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
// 定义一个服务器接口
type IServer interface {
	// 启动服务器
	Start()

	// 停止服务器
	Stop()

	// 运行服务器
	Server()

	//  路由功能，给当前的服务注册一个路由方法，供客户端链接处理使用
	AddRouter(router IRouter)
}
