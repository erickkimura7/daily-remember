package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// type config struct {
// 	Cfg struct {
// 		Teste  string `mapstructure:"teste"`
// 		Server struct {
// 			Port string `mapstructure:"port"`
// 			Url  string `mapstructure:"url"`
// 		} `mapstructure:"server"`
// 	} `mapstructure:"config"`
// }

type cfg struct {
	Config struct {
		Teste string

		Server struct {
			Abobrinha string `mapstructure:"port"`
			Url       string
		}
	}
}

func main() {
	viper.SetConfigFile("./teste_viper/config.toml")

	viper.AddConfigPath("./teste_viper")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("%v", err)
	}

	conf := &cfg{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	fmt.Println(conf)
}
