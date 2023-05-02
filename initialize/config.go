package initialize

import (
	"flag"
	"fmt"

	"qsgo-web-templete/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Config() {
	var config string

	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()

	if config != "" {
		config = fmt.Sprintf("%s-config.yaml", config)
	} else {
		config = "config.yaml"
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.Conf); err != nil {
			panic(err)
		}
	})
	if err = v.Unmarshal(&global.Conf); err != nil {
		panic(err)
	}

	global.ZapS.Info("config init success")
}
