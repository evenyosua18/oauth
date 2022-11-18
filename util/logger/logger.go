package logger

import "log"

const (
	Info  = "info"
	Error = "error"
	Warn  = "warn"
	Fatal = "fatal"
)

func init() {
	log.Default().SetFlags(log.LstdFlags | log.LUTC)
}

func Log(level, desc string, obj interface{}) {
	log.Printf("logger : %s | %s | %v", level, desc, obj)
}
