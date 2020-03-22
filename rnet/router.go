package rnet

import "github.com/mrsongindezhou/ratel/riface"

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2020/3/18 8:35 上午
 * @description：
 * @modified By：
 * @version    ：$
 */

/*
	实现router时，先嵌入这个BaseRouter基类，然后根据需要对这个基类的方法进行重写
*/
type BaseRouter struct{}

// 这里之所以BaseRouter的方法都为空
// 是因为有的Router不希望有PreHandle、PostHandle这两个业务
// 所以Router全部继承BaseRouter的好处就是，不需要实现PreHandle，PostHandle
// 在处理conn业务之前的钩子方法Hook
func (br *BaseRouter) PreHandle(request riface.IRequest) {}

// 在处理conn业务的主方法Hook
func (br *BaseRouter) Handler(request riface.IRequest) {}

// 在处理conn业务之后的钩子方法Hook
func (br *BaseRouter) PostHandle(request riface.IRequest) {}
