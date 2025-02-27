package log

import (
	"io"
	"log"
	"os"
	"runtime"
)

const (
	flag       = log.Ldate | log.Ltime
	preInfo    = "[INFO]"
	preDebug   = "[DEBUG]"
	preWarning = "[WARNING]"
	preError   = "[ERROR]"
)

var (
	logFile        io.Writer
	infoLogger     *log.Logger
	debugLogger    *log.Logger
	warningLogger  *log.Logger
	errorLogger    *log.Logger
	defaultLogFile = os.Getenv("ERROR_LOG_PATH")
)

func InitLogger() {
	var err error
	logFile, err = os.OpenFile(defaultLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)

	if err != nil {
		defaultLogFile = os.Getenv("ERROR_LOG_PATH")
		logFile, err = os.OpenFile(defaultLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			log.Fatalf("log package:create log file err %+v", err)
		}
	}

	infoLogger = log.New(logFile, preInfo, flag)
	debugLogger = log.New(logFile, preDebug, flag)
	warningLogger = log.New(logFile, preWarning, flag)
	errorLogger = log.New(logFile, preError, flag)
}

func Infof(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	infoLogger.Printf("%s:%d: "+format, append([]interface{}{file, line}, v...)...)
}

func Debugf(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	debugLogger.Printf("%s:%d: "+format, append([]interface{}{file, line}, v...)...)
}

func Warningf(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	warningLogger.Printf("%s:%d: "+format, append([]interface{}{file, line}, v...)...)
}

func Errorf(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	errorLogger.Printf("%s:%d: "+format, append([]interface{}{file, line}, v...)...)
}
