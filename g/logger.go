package g

import log "github.com/Sirupsen/logrus"

//InitLog 初始化日志配置
func InitLog(level string) {
	switch level {
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	default:
		log.Fatal("log conf only allow [info, debug, warn], please check your configure")
	}
}
