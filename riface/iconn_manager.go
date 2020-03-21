package riface

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2020/3/21 11:20 上午
 * @description：连接管理模块抽象层
 * @modified By：
 * @version    ：$
 */
type IConneManager interface {
	// 添加连接
	Add(conn IConnection)
	// 删除连接
	Remove(conn IConnection)
	// 根据connID获取连接
	Get(connID uint32) (IConnection, error)
	// 得到当前连接总数
	Len() int
	// 清除并终止所有的连接
	ClearConn()
}
