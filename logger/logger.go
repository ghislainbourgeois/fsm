package logger

import (
	"os"
	"time"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"

	"github.com/omec-project/logger_conf"
	"github.com/omec-project/logger_util"
)

var log *logrus.Logger
var FsmLog *logrus.Entry

func init() {
	log = logrus.New()
	log.SetReportCaller(false)

	log.Formatter = &formatter.Formatter{
		TimestampFormat: time.RFC3339,
		TrimMessages:    true,
		NoFieldsSpace:   true,
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
	}

	free5gcLogHook, err := logger_util.NewFileHook(logger_conf.Free5gcLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err == nil {
		log.Hooks.Add(free5gcLogHook)
	}

	selfLogHook, err := logger_util.NewFileHook(logger_conf.LibLogDir+"fsm.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err == nil {
		log.Hooks.Add(selfLogHook)
	}

	FsmLog = log.WithFields(logrus.Fields{"component": "LIB", "category": "FSM"})
}

func SetLogLevel(level logrus.Level) {
	FsmLog.Infoln("set log level :", level)
	log.SetLevel(level)
}

func SetReportCaller(bool bool) {
	FsmLog.Infoln("set report call :", bool)
	log.SetReportCaller(bool)
}
