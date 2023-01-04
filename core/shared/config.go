package shared

import "github.com/spf13/viper"

type Config struct {
	APIToken string `mapstructure:"API_TOKEN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(path + "/.env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
