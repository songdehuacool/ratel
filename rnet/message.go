package rnet

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2020/3/19 5:04 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
type Message struct {
	ID      uint32 // 消息的ID
	Datalen uint32 // 消息的长度
	Data    []byte // 消息的内容
}

// 提供一个Message消息包
func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		ID:      id,
		Datalen: uint32(len(data)),
		Data:    data,
	}
}

// 获取消息的ID
func (m *Message) GetMsgID() uint32 {
	return m.ID
}

// 获取消息的长度
func (m *Message) GetMsgLen() uint32 {
	return m.Datalen
}

// 获取消息的内容
func (m *Message) GetData() []byte {
	return m.Data
}

// 设置消息的ID
func (m *Message) SetMsgID(id uint32) {
	m.ID = id
}

// 设置消息的内容
func (m *Message) SetData(data []byte) {
	m.Data = data
}

// 设置消息的长度
func (m *Message) SetDataLen(len uint32) {
	m.Datalen = len
}
