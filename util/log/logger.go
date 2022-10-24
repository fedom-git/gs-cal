package log

import (
	"log"
	"os"
	"time"
)

const (
	LogFlag = log.Lshortfile | log.Ldate | log.Lmicroseconds
	Indent1 = " "
	Indent2 = "  "
)

var Log = log.Logger{}

func InitLog() {
	fpath := time.Now().Format("2006-01-02 15:04:05")
	fpath += ".log"
	Log.SetFlags(LogFlag)
	logFile, err := os.OpenFile(fpath, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	Log.SetOutput(logFile)
}
