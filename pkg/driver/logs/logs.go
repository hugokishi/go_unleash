package logs

import (
	"fmt"
	"path"
	"runtime"

	"github.com/hugokishi/go_unleash/pkg/domain/memory"
	"github.com/sirupsen/logrus"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
)

func Init() {
	InitLogrus()
}

func InitLogrus() {
	var formater = &zt_formatter.ZtFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}

	l := logrus.WithFields(logrus.Fields{})

	l.Logger.SetReportCaller(true)
	l.Logger.SetFormatter(formater)
	l.Logger.SetLevel(getLoggerLevel())
}

func getLoggerLevel() logrus.Level {

	switch memory.APP_ENV_LoggingLevel {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}
