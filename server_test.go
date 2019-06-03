package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"testing"
	"time"
	"../proto"
)

type agent struct{
	conn                             net.Conn
	r                                *bufio.Reader
	w                                *bufio.Writer
}





//var driver = mysql.InitMysqlDriver()
//
//var strTemp = heroesList()

func TestServer(t *testing.T) {
	//建立socket端口监听
	netListen, err := net.Listen("tcp", "192.168.1.225:1025")
	CheckError(err)

	defer netListen.Close()

	Log("Waiting for clients ...")

	//等待客户端访问
	for{
		conn, err := netListen.Accept()     //监听接收
		if err != nil{
			continue        //如果发生错误，继续下一个循环。
		}

		Log(conn.RemoteAddr().String(), "tcp connect success")  //tcp连接成功

		go func() {
			agent := new(agent)
			agent.OnInit(conn)
			agent.handleConnection()
		}()
	}
}

func (this *agent) OnInit( conn net.Conn) error {
	this.conn = conn
	this.r = bufio.NewReaderSize(conn, 256)
	this.w = bufio.NewWriterSize(conn, 256)
	return nil
}

//处理连接
func (this *agent) handleConnection() {
	//buffer := make([]byte, 2048)        //建立一个slice
	for{
		this.conn.SetReadDeadline(time.Now().Add(time.Second * 5))
		err :=proto.UnPacket(this.r)
		if err != nil {
			this.conn.Close()
		}

	}
}

func(this *agent) unPacket()  {

	//b,err :=this.r.ReadByte()
	//fmt.Println(b,"&&&&&")
	//log.Println(string(b),err)
	//buf := make([]byte,4)
	//io.ReadFull(this.r,buf)
}

//日志处理
func Log(v ...interface{}) {
	log.Println(v...)
}

//错误处理
func CheckError(err error) {
	if err != nil{
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
}


// 全服榜
//func  heroesList()  map[string]interface{}{
//	res := make(map[string]interface{})
//	res["cmd"] = "HD_HeroesList"
//
//	// attr, ok := msg["attribute"].(float64)
//	// roleId, rok := msg["roleId"].(string)
//	// if !ok || !rok {
//	// 	res["data"] = ErrorOfParameters
//	// 	w.sendMsg(session, res)
//	// 	return
//	// }
//
//	attr := 1003.0
//	roleId := "10000000"
//	cond := fmt.Sprintf("attribute = %f",attr)
//	order := fmt.Sprintf(" difficulty desc,damage desc,end_time asc,role_id asc")
//	list,err := driver.User.SelectHeroRankList(cond,"*",order,50)
//	if err != nil {
//		logger.LogInfo(fmt.Sprintf("heroesList: SelectHeroRankList error:%v",err))
//	}
//
//	data := make(map[string]interface{})
//	data["rankWars"] = list
//
//	cond = fmt.Sprintf("role_id = %s and attribute = %f ", roleId, attr)
//	rank, err :=driver.User.SelectRoleRankList(cond,"*")
//	if err == nil && rank.RoleId != 0 && rank.Difficulty != 0 {
//		cond = fmt.Sprintf("attribute = %f and (difficulty >%d or (difficulty = %d and (damage >%d or (damage =%d and end_time <= %d))))",attr,rank.Difficulty,rank.Difficulty,rank.Damage,rank.Damage,rank.EndTime)
//		num, err := driver.User.GetRoleRankNum(cond)
//		if err != nil {
//			logger.LogInfo("heroesList: GetRoleRankNum error:%v",err)
//		}
//		data["selfRank"] = num
//		data["selfWar"] = rank
//	}
//	res["data"] = data
//
//	return res
//}