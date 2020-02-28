package logInit

import (
	"log"
	"monitorCore/common"
	"os"
	"time"
)

func init() {
	var dir = "log/"
	_ = common.EnsurePath(dir)
	filename := dir + time.Now().Format("20060102") + ".log"
	f, _ := os.OpenFile(filename,  os.O_RDWR| os.O_APPEND | os.O_CREATE, 0666)
	log.SetOutput(f)

	log.Println("程序已启动...")
}
