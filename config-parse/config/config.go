package config

import (
	"fmt"
	"os"

	"github.com/coreos/pkg/capnslog"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	//"github.com/BurntSushi/toml"
)

var log = capnslog.NewPackageLogger("config-parse", "config")

// Config 配置参数
type Config struct {
	RunMode string `toml:"run_mode"`
	Log     Log    `toml:"log"`
}

// Log 日志配置参数
type Log struct {
	Level      int    `toml:"level"`
	Format     string `toml:"format"`
	Output     string `toml:"output"`
	OutputFile string `toml:"output_file"`
}

func Init(cfgFile string) {
	capnslog.SetGlobalLogLevel(capnslog.WARNING)
	capnslog.SetFormatter(capnslog.NewPrettyFormatter(os.Stdout, false))

	viper.SetEnvPrefix("cfg")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.BindEnv("log.level", "CFG_LOG_LEVEL")
	viper.AutomaticEnv()
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
	err := viper.ReadInConfig()
	if err != nil {
		e, ok := err.(viper.ConfigParseError)
		if ok {
			log.Fatalf("error parsing config file: %v", e)
		}
		log.Debugf("No config file used")
	} else {
		log.Debugf("Using config file: %v", viper.ConfigFileUsed())
	}
}

func values() Config {
	return Config{
		RunMode: viper.GetString("run_mode"),
		Log: Log{
			Level:      viper.GetInt("log.level"),
			Format:     viper.GetString("log.format"),
			Output:     viper.GetString("log.output"),
			OutputFile: viper.GetString("log.output_file"),
		},
	}
}

func Print() {
	cfg := values()
	cfgBytes, err := yaml.Marshal(cfg)
	if err != nil {
		log.Fatalf("marshalling configuration: %v", err)
	}

	fmt.Println("Configuration")
	fmt.Printf("%v", string(cfgBytes))
}
