package riface

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2020/3/18 8:35 上午
 * @description：
 * @modified By：
 * @version    ：$
 */

/*
	路由抽象接口
	路由里的数据都是IRequest
*/
type IRouter interface {
	// 在处理conn业务之前的钩子方法Hook
	PreHandle(request IRequest)
	// 在处理conn业务的主方法Hook
	Handler(request IRequest)
	// 在处理conn业务之后的钩子方法Hook
	PostHandle(request IRequest)
}
