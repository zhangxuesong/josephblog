package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

/**
 * 使用级别，参照一下
 * - Fatal：网站挂了，或者极度不正常
 * - Error：跟遇到的用户说对不起，可能有bug
 * - Warn：记录一下，某事又发生了
 * - Info：提示一切正常
 * - debug：没问题，就看看堆栈
 **/

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPreFix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}) {
	setPreFix(INFO)
	logger.Println(v)
}
func Warn(v ...interface{}) {
	setPreFix(WARNING)
	logger.Println(v)
}
func Error(v ...interface{}) {
	setPreFix(ERROR)
	logger.Println(v)
}
func Fatal(v ...interface{}) {
	setPreFix(FATAL)
	logger.Fatalln(v)
}

func setPreFix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
