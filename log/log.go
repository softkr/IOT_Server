package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	logs = [...]string{
		"error",
		"trace",
	}
	Trace *log.Logger // 모든로그 저장
	Error *log.Logger // 에러로그 저장
)
var now = time.Now()

func init() {
	for _, i := range logs {
		os.MkdirAll(fmt.Sprintf("log/%v", i), os.ModePerm)
	}

	now := time.Now()
	errorLog := fmt.Sprintf("./log/error/%v-error.log", now.Local().Format("2006-01-02"))
	errorLogFile, err := os.OpenFile(errorLog,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	traceLog := fmt.Sprintf("./log/trace/%v-trace.log", now.Local().Format("2006-01-02"))
	traceLogFile, err := os.OpenFile(traceLog,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Trace = log.New(io.MultiWriter(traceLogFile, os.Stdout),
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(errorLogFile, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
