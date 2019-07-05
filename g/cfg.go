package g

import (
	"encoding/json"

	"log"

	"sync"

	"github.com/toolkits/file"
)

type GlobalConfig struct {
	LogLevel      string      `json:"log_level"`
	Sqlite        string      `json:"sqlite"`
	Mysql         string      `json:"mysql"`
	Http          *HttpConfig `json:"http"`
	AccessControl *[]ACLUser  `json:"access_control"`
}

type ACLUser struct {
	User      string `json:"user"`
	X_API_KEY string `json:"x-api-key"`
	Role      int64  `json:"role"`
}

type HttpConfig struct {
	Listen string `json:"listen"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	lock       = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

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
