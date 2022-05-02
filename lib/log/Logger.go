package log

import (
	"go.uber.org/zap"
)	

type Logger struct {
	zl zap.SugaredLogger
}


func (logger *Logger) Init()  {
	log, _ := zap.NewProduction()
	defer log.Sync()
	logger.zl = *log.Sugar()
}

func (logger *Logger) Errorf(template string, args...interface{})  {
	logger.zl.Errorf(template,args...)
}

func (logger *Logger) Infof(template string, args...interface{})  {
	logger.zl.Infof(template, args...)
}

func (logger *Logger) Fatalf(template string, args...interface{})  {
	logger.zl.Fatalf(template, args...)
}