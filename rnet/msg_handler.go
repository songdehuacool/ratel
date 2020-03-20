package rnet

import (
	"fmt"
	"ratel/riface"
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
}

// 初始化/创建MsgHandler方法
func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis: make(map[uint32]riface.IRouter),
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
