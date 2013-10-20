package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

var logger *Logger
var dest string
var Print = fmt.Println

type Logger struct {
	// 1 ERR
	// 2 WARNING
	// 3 INFO
	// 4 DEBUG
	Level int
}

var Level map[int]string = map[int]string{
	1: "ERROR",
	2: "WARN",
	3: "INFO",
	4: "DEBUG",
}

func Init(level int, dst ...string) {
	logger = &Logger{Level: level}
	if len(dst) == 1 {
		dest = dst[0]
	}
}

func getPath() (dir, path string, line int) {
	_, fullpath, line, _ := runtime.Caller(3)
	f := strings.Split(fullpath, "/")
	dir = f[len(f)-2]
	file := f[len(f)-1]
	path = dir + "/" + file
	return
}

func out(level int, a interface{}) {
	_, path, line := getPath()
	a = fmt.Sprintf("[%s][%s] [%v:%v] %v\n", Level[level], time.Now().Format("2006/01/02 15:04:05"), path, line, a)
	if dest != "" {
		f, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err == nil {
			fmt.Fprintf(f, "%v", a)
			f.Close()
		} else {
			Warn("Can not open log file")
		}
	}
	if runtime.GOOS == "linux" {
		switch level {
		case 1:
			a = Red(fmt.Sprintf("%v", a))
		case 2:
			a = Cyan(fmt.Sprintf("%v", a))
		case 4:
			a = Green(fmt.Sprintf("%v", a))
		}
	}
	fmt.Fprintf(os.Stderr, "%v", a)

}

func Error(a interface{}) {
	if logger.Level >= 1 {
		out(1, a)
	}
}

func Warn(a interface{}) {
	if logger.Level >= 2 {
		out(2, a)
	}
}

func Info(a interface{}) {
	if logger.Level >= 3 {
		out(3, a)
	}
}

func Debug(a interface{}) {
	if logger.Level >= 4 {
		out(4, a)
	}
}
