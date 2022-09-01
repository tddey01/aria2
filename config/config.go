package config

import (
	"github.com/BurntSushi/toml"
)

type aria2 struct {
	Aria2DownloadDir string `toml:"aria2_download_dir"`
	Aria2Host        string `toml:"aria2_host"`
	Aria2Port        int    `toml:"aria2_port"`
	Aria2Secret      string `toml:"aria2_secret"`
	Aria2Task        int    `toml:"aria2_max_task"`
}

type mysql struct {
	DBType   string `toml:"DbType"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	DbName   string `toml:"dbname"`
	DbUser   string `toml:"users"`
	DbPasswd string `toml:"passwd"`
	Table    string `toml:"table"`
}
type logs struct {
	MaxSize    int    `toml:"maxsize"`
	MaxBackups int    `toml:"backups"`
	MaxAge     int    `toml:"day"`
	Level      string `toml:"level"`
}
type main struct {
	LogName string `toml:"LogName"`
}

type Configuration struct {
	Port    int   `toml:"port"`
	Release bool  `toml:"release"`
	Aria2   aria2 `toml:"aria2"`
	Main    main  `toml:"main"`
	Mysql   mysql `toml:"mysql"`
	Logs    logs  `toml:"logs"`
	//Watch   watch `toml:"watch"`
	Disk disk `toml:"disk"`
}

type disk struct {
	Drive string `toml:"drive"`
}

var config *Configuration

func GetConfig() Configuration {
	if config == nil {
		InitConfig()
	}
	return *config
}

func requiredFieldsAreGiven(metaData toml.MetaData) bool {
	requiredFields := [][]string{
		{"port"},
		{"release"},

		{"aria2"},
		{"main"},
		{"mysql"},
		{"logs"},

		{"disk"},

		{"aria2", "aria2_download_dir"},
		{"aria2", "aria2_host"},
		{"aria2", "aria2_port"},
		{"aria2", "aria2_secret"},
		{"aria2", "aria2_max_task"},

		{"main", "LogName"},

		{"mysql", "DbType"},
		{"mysql", "host"},
		{"mysql", "port"},
		{"mysql", "dbname"},
		{"mysql", "users"},
		{"mysql", "passwd"},
		{"mysql", "table"},

		{"logs", "maxsize"},
		{"logs", "backups"},
		{"logs", "day"},
		{"logs", "level"},

		{"disk", "drive"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			log.Fatal("required conf fields ", v)
		}
	}

	return true
}

func InitConfig() {
	configFile := "./config.toml"
	log.Debug("Your config file is:", configFile)

	if metaData, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal("error:", err)
	} else {
		if !requiredFieldsAreGiven(metaData) {
			log.Fatal("required fields not given")
		}
	}
}
