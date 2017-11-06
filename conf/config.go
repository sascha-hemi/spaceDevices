package conf

import (
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

func LoadConfig(configFile string) TomlConfig {
	log.WithField("configFile", configFile).Info("Loading config.")
	config := &TomlConfig{}
	if _, err := toml.DecodeFile(configFile, config); err != nil {
		log.WithError(err).Fatal("Could not read config file.")
	}

	return *config
}

type TomlConfig struct {
	Server ServerConf
	MacDb  MacDbConf
	Mqtt   MqttConf
}

type ServerConf struct {
	Host     string
	Port     int
	Https    bool
	KeyFile  string
	CertFile string
}

type MacDbConf struct {
	MasterFile string
	UserFile   string
}

type MqttConf struct {
	Url      string
	Username string
	Password string
	// if empty, the system certificates are used
	CertFile string
}
