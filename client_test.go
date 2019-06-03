package tcp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"testing"
	"time"

	"../proto"
)

const (
	header ="header"
	headerLen = 5
	saveMsgLen = 4
)
//发送信息
func sender(conn net.Conn) {
	msg :=proto.Packet()

	log.Println(msg,"11111111111111111")
	conn.Write(msg)
	go recvMsg(conn)

}

func encodeBytes(field []byte) []byte {
	fieldLength := make([]byte, 2)
	fmt.Println(fieldLength,"111")
	binary.BigEndian.PutUint16(fieldLength, uint16(len(field)))
	fmt.Println(fieldLength,"2222")
	return append(fieldLength, field...)
}

func recvMsg(conn net.Conn)  {
	//接收服务端反馈
	buffer := make([]byte, 2048)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Log1(conn.RemoteAddr().String(), "waiting server back msg error: ", err)
			return
		}
		Log1(conn.RemoteAddr().String(), "receive server back msg: ", string(buffer[:n]),len(buffer[:n]))
		time.Sleep(0.5e9)
	}

}

func stringToByte(header string)  []byte{
	head := []byte(header)
	return head
}

func intToByte(n int64) ([]byte,error) {

	buf := new(bytes.Buffer)
	err := binary.Write(buf,binary.BigEndian,n)
	return buf.Bytes(),err
}


//日志
func Log1(v ...interface{}) {
	log.Println(v...)
}

func TestClient(t *testing.T) {
	server := "192.168.1.225:1025"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connection success")
	sender(conn)
	ch := make(chan int)
	<-ch
}


