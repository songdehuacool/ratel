package rnet

import (
	"fmt"
	"ratel/riface"
	"ratel/utils"
	"strconv"
)

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2020/3/20 6:07 下午
 * @description：消息处理模块实现
 * @modified By：
 * @version    ：$
 */
type MsgHandler struct {
	// 存放每一个MsgID 所对应的处理方法
	Apis map[uint32]riface.IRouter
	// 负责Worker取任务的消息队列
	TaskQueue []chan riface.IRequest
	// 业务工作Worker池的worker数量
	WorkerPoolSize uint32
}

// 初始化/创建MsgHandler方法
func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]riface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize, // 从全局配置中获取
		TaskQueue:      make([]chan riface.IRequest, utils.GlobalObject.WorkerPoolSize),
	}
}

// 调度/执行对应的Router消息处理方法
func (mh *MsgHandler) DoMsgHandler(request riface.IRequest) {
	// 1.从request中找到msgID
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgID = ", request.GetMsgID(), " is NOT FOUND ! Need Register!")
	}
	// 2.根据MsgID调度对应router业务即可
	handler.PreHandle(request)
	handler.Handler(request)
	handler.PostHandle(request)
}

// 为消息添加具体的处理逻辑
func (mh *MsgHandler) AddRouter(msgID uint32, router riface.IRouter) {
	// 1.判断当前msg绑定的API处理方法是否已经存在
	if _, ok := mh.Apis[msgID]; ok {
		// id已经注册了
		panic("repeat api, msgID = " + strconv.Itoa(int(msgID)))
	}

	// 2.添加msg与API的绑定关系
	mh.Apis[msgID] = router
	fmt.Println("Add api MsgID = ", msgID, " succ!")
}

// 启动一个Worker工作池(开启工作池的动作只能发生一次，一个ratel框架只能有一个worker工作池)
func (mh *MsgHandler) StartWorkerPool() {
	// 根据workerPoolSize 分别开启Worker， 每个Worker用一个go来承载
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		// 一个worker被启动

		// 1 要给当前的worker对应的channel 消息队列 开辟空间 第0个worker 就用第0个channel
		mh.TaskQueue[i] = make(chan riface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		// 2 启动当前的Worker， 阻塞等待消息从channel传递过来
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

// 启动一个Worker工作流程
func (mh *MsgHandler) StartOneWorker(workerID int, taskQueue chan riface.IRequest) {
	fmt.Println("Worker ID = ", workerID, " is started ... ")

	// 不断的阻塞等待对应消息队列的消息
	for {
		select {
		// 如果有消息过来，出列就是一个客户端的Request，执行当前Request所绑定的业务
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

// 将消息交给TaskQueue， 由worker进行处理
func (mh *MsgHandler) SendMsgToTaskQueue(request riface.IRequest) {
	// 1 将消息平均分配给不同的worker
	// 根据客户端建立的ConnID来进行分配
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	fmt.Println("Add ConnID = ", request.GetConnection().GetConnID(),
		" request MsgID = ", request.GetMsgID(), " to WorkerID = ", workerID)

	// 2 将消息发送给对应的worker的TaskQueue即可
	mh.TaskQueue[workerID] <- request
}
