package di

import (
	"fmt"
	"github.com/Summer2021-Project/cli"
	"github.com/Summer2021-Project/di"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	obj := di.Object{
		Name: "logrus",
		New: func() (i interface{}, e error) {
			logger := logrus.New()
			logger.ReportCaller = true // 显示调用信息
			formatter := new(logrus.TextFormatter)
			formatter.FullTimestamp = true
			formatter.TimestampFormat = "2006-01-02 15:04:05.000"
			formatter.DisableQuote = true // 不转义换行符，为了保存错误堆栈到日志文件
			formatter.CallerPrettyfier = func(frame *runtime.Frame) (function string, file string) {
				return "", fmt.Sprintf("%s:%d", filepath.Base(frame.File), frame.Line)
			}
			logger.Formatter = formatter
			filename := fmt.Sprintf("%s/../runtime/logs/cli.log", cli.App().BasePath)
			fileRotate := &lumberjack.Logger{
				Filename:   filename,
				MaxBackups: 7,
			}
			writer := io.MultiWriter(os.Stdout, fileRotate)
			logger.SetOutput(writer)
			if cli.App().Debug {
				logger.SetLevel(logrus.DebugLevel)
			}
			return logger, nil
		},
	}
	if err := di.Provide(&obj); err != nil {
		panic(err)
	}
}

func Logrus() (logger *logrus.Logger) {
	if err := di.Populate("logrus", &logger); err != nil {
		panic(err)
	}
	return
}
