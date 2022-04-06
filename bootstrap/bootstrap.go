package bootstrap

import (
	"fileshare/config"
	"log"
	"os"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Application struct {
    ConfigViper *viper.Viper
    Config config.Configuration
}

var App = new(Application)

func InitializeConfig() *viper.Viper {
    config := "config.yaml"

    if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
        config = configEnv
    }
    
    v := viper.New()
    v.SetConfigFile(config)
    v.SetConfigType("yaml")
    if err := v.ReadInConfig(); err != nil {
        log.Panic("read config failed:", err)
    }
   
    v.WatchConfig()
    v.OnConfigChange(func(in fsnotify.Event) {
        log.Println("config file changed:", in.Name)
        if err := v.Unmarshal(&App.Config); err != nil {
            log.Println(err)
        }
    })
    
    if err := v.Unmarshal(&App.Config); err != nil {
       log.Println(err)
    }
    return v
}