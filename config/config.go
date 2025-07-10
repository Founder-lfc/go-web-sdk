package config

import "log"

var (
	ExtendConfig interface{}
	_cfg         *Settings
)

type Settings struct {
	Settings  Config `yaml:"settings"`
	callbacks []func()
}

func (e *Settings) runCallback() {
	for i := range e.callbacks {
		e.callbacks[i]()
	}
}

func (e *Settings) OnChange() {
	e.init()
	log.Println("config change and reload")
}

// Config 配置集合
type Config struct {
	Application *Application          `yaml:"application"`
	Jwt         *Jwt                  `yaml:"jwt"`
	Database    *Database             `yaml:"database"`
	Databases   *map[string]*Database `yaml:"databases"`
	Gen         *Gen                  `yaml:"gen"`
	Extend      interface{}           `yaml:"extend"`
}
