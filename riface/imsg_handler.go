package riface

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2020/3/20 6:07 下午
 * @description：消息管理抽象层
 * @modified By：
 * @version    ：$
 */
type IMsgHandler interface {
	// 调度/执行对应的Router消息处理方法
	DoMsgHandler(request IRequest)
	// 为消息添加具体的处理逻辑
	AddRouter(msgID uint32, router IRouter)
	// 启动Worker工作池
	StartWorkerPool()
	// 将消息发送给消息任务队列处理
	SendMsgToTaskQueue(request IRequest)
}
