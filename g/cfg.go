package g

import (
	"encoding/json"

	"log"

	"sync"

	"github.com/toolkits/file"
)

//GlobalConfig 全局配置
type GlobalConfig struct {
	LogLevel      string      `json:"log_level"`
	Sqlite        string      `json:"sqlite"`
	Mysql         string      `json:"mysql"`
	HTTP          *HTTPConfig `json:"http"`
	AccessControl *[]ACLUser  `json:"access_control"`
}

//ACLUser 用户的角色配置
type ACLUser struct {
	User    string `json:"user"`
	XAPIKEY string `json:"x-api-key"`
	Role    int64  `json:"role"`
}

//HTTPConfig http 配置
type HTTPConfig struct {
	Listen string `json:"listen"`
}

var (
	//ConfigFile 配置文件
	ConfigFile string
	config     *GlobalConfig
	lock       = new(sync.RWMutex)
)

//Config 给其他包调用的配置加载方法，带锁
func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

//ParseConfig 从配置文件加载配置
func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent. maybe you need `mv cfg.example.json cfg.json`")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	lock.Lock()
	defer lock.Unlock()

	config = &c

}
