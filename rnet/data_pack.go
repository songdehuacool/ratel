package rnet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"ratel/riface"
	"ratel/utils"
)

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2020/3/19 6:13 下午
 * @description：封包、拆包 具体模块
 * @modified By：
 * @version    ：$
 */
type DataPack struct {
}

// 拆包 封包的一个初始化方法
func NewDataPack() *DataPack {
	return &DataPack{}
}

// 获取包的头的长度方法
func (d *DataPack) GetHeadLen() uint32 {
	// Datalen uint32(4字节) + ID uint32(4字节)
	return 8
}

// 封包方法
func (d *DataPack) Pack(msg riface.IMessage) ([]byte, error) {
	// 创建一个存放byte字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	// 将dataLen写入到dataBuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	// 将MsgId 写进dataBuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}
	// 将data数据写入dataBuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

// 拆包方法 (将包的Head信息读出来) 之后再根据head信息里的data长度，再进行一次读
func (d *DataPack) Unpack(binaryData []byte) (riface.IMessage, error) {
	// 创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)

	// 只解压head信息，得到dataLen和msgID
	msg := &Message{}

	// 读DataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Datalen); err != nil {
		return nil, err
	}
	// 读MsgID
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}
	// 判断datalen是否已经超出了允许的最大包长度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.Datalen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too Large msg data recv!")
	}
	return msg, nil
}
