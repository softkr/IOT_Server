package socket

import (
	"encoding/json"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"io"
	"io/ioutil"
	"iot/grpc/client"
	"log"
	"net"
	"os"
	"regexp"
)

type stringDataToJson struct {
	Watchsn             string `json:"Watch_sn"`
	Stepcount           int    `json:"step_count"`
	Camerashootingcount int    `json:"camera_shooting_count"`
}

var (
	Trace *log.Logger
	Error *log.Logger
)

func LogInit() {
	file, err := os.OpenFile("./log/errors.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Trace = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func DoHandler(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			Error.Println(err)
		}
	}(conn)
	Trace.Println("읽기 시작")
	b := make([]byte, 1024)
	for {
		n, err := conn.Read(b)
		if err != nil {
			client.WatchState("", conn.RemoteAddr().String(), 2)
			//client.WatchOnOff("", conn.RemoteAddr().String(), 2)
			log.Println("소켓종료", conn.RemoteAddr().String(), 2)
			Error.Println(err.Error())
			return
		}
		readData := string(b[:n])
		fmt.Println(conn.RemoteAddr().String() + ": 읽어온데이터" + string(b[:n]))
		matched, _ := regexp.MatchString("^2[0-9]+IHPA+", readData)
		data, _ := regexp.MatchString("camera_shooting_count", readData)

		// 워치 SN
		if matched {
			log.Println("소켓 접속", readData, conn.RemoteAddr().String(), 1)
			client.WatchState(readData, conn.RemoteAddr().String(), 1)
			//grpcgo.WatchOnOff(readData, conn.RemoteAddr().String(), 1)
		}
		// 걸음수 카메라 데이터
		if data {
			byt := []byte(readData)
			var jsonData stringDataToJson
			if err := json.Unmarshal(byt, &jsonData); err == nil {
				log.Println(jsonData.Camerashootingcount, "복약걸음수")
				// fmt.Println(jsonData.Watchsn, int32(jsonData.Stepcount), int32(jsonData.Camerashootingcount))
				client.WatchUpdate(jsonData.Watchsn, int32(jsonData.Stepcount), int32(jsonData.Camerashootingcount))
			} else {
				Error.Println(err)
			}
		}
	}
}

func Run() {
	fmt.Println("소켓실행", os.Getenv("SOCKET_HOST"))
	listen, err := net.Listen("tcp", os.Getenv("SOCKET_HOST"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	LogInit()

	Trace.Println("tcp", os.Getenv("SOCKET_HOST"))
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			return
		}
	}(listen)
	// loop 무한 반복
	for {
		// 막힐경우 연결대기
		conn, err := listen.Accept()
		if err != nil {
			Error.Println(err.Error())
			return
		}
		go DoHandler(conn)
	}
}
