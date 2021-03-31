package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

func init() {
	fmt.Println("set log")
	log.Printf("init log config")
	// logFile := "test.log"
	hostname, err := os.Hostname()
	fmt.Println(hostname, err)
	if err != nil {
		log.Fatal("get host name error")
	}
	logFile := fmt.Sprintf("/data/log/%s.log", hostname)
	logF, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// logF, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	mv := io.MultiWriter(os.Stdout, logF)
	log.SetOutput(mv) //写入控制台和文件
	// log.SetOutput(logF)
	log.SetPrefix(hostname)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate | log.Lshortfile)
}
