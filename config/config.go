package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
	"time"
)

type aria2 struct {
	Aria2DownloadDir string `toml:"aria2_download_dir"`
	Aria2Host        string `toml:"aria2_host"`
	Aria2Port        int    `toml:"aria2_port"`
	Aria2Secret      string `toml:"aria2_secret"`
	Aria2Task        int    `toml:"aria2_max_task"`
}

type lotus struct {
	ClientApiUrl      string `toml:"client_api_url"`
	MarketApiUrl      string `toml:"market_api_url"`
	MarketAccessToken string `toml:"market_access_token"`
}

type mysql struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	DbName   string `toml:"dbname"`
	DbUser   string `toml:"users"`
	DbPasswd string `toml:"passwd"`
}

type main struct {
	SwanApiUrl               string        `toml:"api_url"`
	SwanApiKey               string        `toml:"api_key"`
	SwanAccessToken          string        `toml:"access_token"`
	SwanApiHeartbeatInterval time.Duration `toml:"api_heartbeat_interval"`
	MinerFid                 string        `toml:"miner_fid"`
	LotusImportInterval      time.Duration `toml:"import_interval"`
	LotusScanInterval        time.Duration `toml:"scan_interval"`
	LogName                  string        `toml:"aria2"`
}

type Configuration struct {
	Port    int   `toml:"port"`
	Release bool  `toml:"release"`
	Lotus   lotus `toml:"lotus"`
	Aria2   aria2 `toml:"aria2"`
	Main    main  `toml:"main"`
	Mysql   mysql `toml:"mysql"`
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

		{"lotus"},
		{"aria2"},
		{"main"},
		{"bid"},

		{"lotus", "client_api_url"},
		{"lotus", "market_api_url"},
		{"lotus", "market_access_token"},

		{"aria2", "aria2_download_dir"},
		{"aria2", "aria2_host"},
		{"aria2", "aria2_port"},
		{"aria2", "aria2_secret"},
		{"aria2", "aria2_max_task"},

		{"main", "api_url"},
		{"main", "miner_fid"},
		{"main", "import_interval"},
		{"main", "scan_interval"},
		{"main", "api_key"},
		{"main", "access_token"},
		{"main", "api_heartbeat_interval"},

		{"mysql", "host"},
		{"mysql", "port"},
		{"mysql", "dbname"},
		{"mysql", "users"},
		{"mysql", "passwd"},

		{"bid", "bid_mode"},
		{"bid", "expected_sealing_time"},
		{"bid", "start_epoch"},
		{"bid", "auto_bid_deal_per_day"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			log.Fatal("required conf fields ", v)
		}
	}

	return true
}

func InitConfig() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Cannot get home directory.")
	}

	configFile := filepath.Join(homedir, "./config.toml")

	log.Info("Your config file is:", configFile)

	if metaData, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal("error:", err)
	} else {
		if !requiredFieldsAreGiven(metaData) {
			log.Fatal("required fields not given")
		}
	}
}
