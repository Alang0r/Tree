package log

import (
	"go.uber.org/zap"
)	

type Logger struct {
	zl zap.SugaredLogger
	Service string
}


func (logger *Logger) Init(service string)  {
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
func (logger *Logger) Info(args...interface{})  {
	logger.zl.Info(args...)
}

func (logger *Logger) Fatalf(template string, args...interface{})  {
	logger.zl.Fatalf(template, args...)
}

func (logger *Logger) Fatal(args...interface{})  {
	logger.zl.Fatal(args...)
}
