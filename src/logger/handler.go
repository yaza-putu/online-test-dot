package logger

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
	"path/filepath"
)

const (
	DEBUG Lvl = iota + 1
	INFO
	WARN
	ERROR
	OFF
	PANIC
	FATAL
)

type (
	optFunc func(opts2 *opts)
	Lvl     uint8
	opts    struct {
		Write bool
		Type  Lvl
	}
)

func defaultOpts() opts {
	return opts{
		Write: true,
		Type:  ERROR,
	}
}

func IsWrite(r bool) optFunc {
	return func(o *opts) {
		o.Write = r
	}
}

func SetType(t Lvl) optFunc {
	return func(o *opts) {
		o.Type = t
	}
}

func New(err error, opts ...optFunc) {
	o := defaultOpts()

	for _, fn := range opts {
		fn(&o)
	}

	if err != nil {
		if o.Write == true {
			writeError(err, o.Type)
		}
		fmt.Printf("%s : %s", getlabel(o.Type), err.Error())
	}
}

func getlabel(l Lvl) string {
	switch l {
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	case FATAL:
		return "FATAL"
	case PANIC:
		return "PANIC"
	default:
		return "ERROR"
	}
}

func writeError(err error, l Lvl) {
	cwd, _ := os.Getwd()

	logFile, er := os.OpenFile("logs/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if er != nil {
		_ = os.Mkdir("logs", os.ModePerm)
		path := filepath.Join(cwd, "logs", "error.log")
		newFilePath := filepath.FromSlash(path)
		logFile, er = os.Create(newFilePath)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	switch l {
	case INFO:
		log.Info(err.Error())
		break
	case DEBUG:
		log.Debug(err.Error())
		break
	case FATAL:
		log.Fatal(err.Error())
		break
	case PANIC:
		log.Panic(err.Error())
		break
	default:
		log.Error(err.Error())
		break
	}
}
