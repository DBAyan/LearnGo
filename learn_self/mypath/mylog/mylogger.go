package mylog123

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
)

// 定义日志级别类型

type LogLevel int

//日志级别常量

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)
const timeFormat = "2006-01-02 15:04:05 "

// 定义logger结构体

type Logger struct {
	level      LogLevel
	showFile   bool
	showLine   bool
	showTime   bool
	showScreen bool
	fileName   string
	file       *os.File
	logger     *log.Logger
}

func NewLogger(level LogLevel, showFile bool, showLine bool, showTime bool, showScreen bool, fileName string) (*Logger, error) {
	// 结构体初始化 ，将函数的参数赋值给结构体
	logger := &Logger{
		level:      level,
		showFile:   showFile,
		showLine:   showLine,
		showTime:   showTime,
		showScreen: showScreen,
		fileName:   fileName,
	}

	if fileName != "" {
		//fmt.Printf("打开文件%v",fileName)
		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			return nil, err
		}
		logger.file = file
		logger.logger = log.New(file, "", log.Ldate|log.Ltime)
		//fmt.Println("打开文件成功")
	} else {
		logger.logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	}
	return logger, nil
}

func (l *Logger) writeLog(level string, format string, v ...interface{}) {
	var logMessage string
	//if l.showTime {
	//	logMessage += time.Now().Format(timeFormat)
	//}
	logMessage += level + ": "
	if l.showFile || l.showLine {
		_, file, line, ok := runtime.Caller(2)
		if ok {
			if l.showFile {
				logMessage += "(" + file + ")"
			}
			if l.showLine {
				logMessage += "[" + strconv.Itoa(line) + "]"
			}
		}
	}
	logMessage += " " + fmt.Sprintf(format, v...)
	l.logger.Println(logMessage)
	if l.showScreen {
		fmt.Println(logMessage)
	}
}
func (l *Logger) Debug(format string, v ...interface{}) {
	if l.level <= LogLevelDebug {
		l.writeLog("DEBUG", format, v...)
	}
}

func (l *Logger) Info(format string, v ...interface{}) {
	if l.level <= LogLevelDebug {
		l.writeLog("INFO", format, v...)
	}
}

func (l *Logger) Warn(format string, v ...interface{}) {
	if l.level <= LogLevelDebug {
		l.writeLog("WARN", format, v...)
	}
}

func (l *Logger) Error(format string, v ...interface{}) {
	if l.level <= LogLevelDebug {
		l.writeLog("ERROE", format, v...)
	}
}

func (l *Logger) Close() {
	if l.file != nil {
		l.file.Close()
	}
}


//func main()  {
//	myLogger,err := NewLogger(LogLevelDebug,true,true,true,true,"20230428.log")
//	if err != nil{
//		fmt.Printf("初始化日志报错：%v",err)
//		return
//	}
//	defer myLogger.Close()
//	myLogger.Debug("测试日志123")
//	myLogger.Info("info 测试日志")
//}