package rnet

import (
	"errors"
	"fmt"
	"ratel/riface"
	"sync"
)

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2020/3/21 11:20 上午
 * @description：连接管理模块
 * @modified By：
 * @version    ：$
 */
type ConnManager struct {
	connections map[uint32]riface.IConnection // 管理的连接集合
	connLock    sync.RWMutex                  // 保护连接集合的读写锁
}

// 创建当前连接的方法
func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]riface.IConnection),
	}
}

// 添加连接
func (connMgr *ConnManager) Add(conn riface.IConnection) {
	// 保护共享资源map， 加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	// 将conn加入到ConnManager中
	connMgr.connections[conn.GetConnID()] = conn
	fmt.Println("connID = ", conn.GetConnID(), " add to ConManager successfully: conn num = ", connMgr.Len())
}

// 删除连接
func (connMgr *ConnManager) Remove(conn riface.IConnection) {
	// 保护共享资源map， 加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	// 删除连接资源
	delete(connMgr.connections, conn.GetConnID())
	fmt.Println("connID = ", conn.GetConnID(), " remove from ConManager successfully: conn num = ", connMgr.Len())
}

// 根据connID获取连接
func (connMgr *ConnManager) Get(connID uint32) (riface.IConnection, error) {
	// 保护共享资源map， 加写锁
	connMgr.connLock.RLock()
	defer connMgr.connLock.RUnlock()

	if conn, ok := connMgr.connections[connID]; ok {
		// 找到了
		return conn, nil
	} else {
		return nil, errors.New("connection not Found !")
	}

}

// 得到当前连接总数
func (connMgr *ConnManager) Len() int {
	return len(connMgr.connections)
}

// 清除并终止所有的连接
func (connMgr *ConnManager) ClearConn() {
	// 保护共享资源map， 加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	// 删除conn并停止conn的工作
	for connID, conn := range connMgr.connections {
		// 停止
		conn.Stop()
		// 删除
		delete(connMgr.connections, connID)
	}
	fmt.Println("Clear All connections succ ! conn num = ", len(connMgr.connections))
}
