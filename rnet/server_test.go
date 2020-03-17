package rnet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2020/3/17 4:17 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
func TestServer_Start(t *testing.T) {
	// 1 创建一个server句柄，使用Ratel的api
	s := NewServer("[ratel v0.1]")
	// 2 启动server
	s.Server()
}

func Test_Client(t *testing.T) {
	fmt.Println("client start...")

	time.Sleep(1 * time.Second)

	// 1 直接链接远程服务器，得到一个conn链接
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}
	for {
		// 链接调用Write 写数据
		_, err := conn.Write([]byte("Hello Ratel V0.1..."))
		if err != nil {
			fmt.Println("write conn err", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}
		fmt.Printf("server cal back: %s, cnt = %d\n", buf, cnt)

		// cpu阻塞
		time.Sleep(1 * time.Second)
	}
}
