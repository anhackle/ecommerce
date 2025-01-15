package initialize

import (
	"fmt"
	"os"

	"github.com/anle/codebase/global"
	"github.com/spf13/viper"
)

func getConfigName() string {
	mode := "dev"
	fmt.Println(os.Getenv("STAGE"))
	switch os.Getenv("STAGE") {
	case "local":
		mode = "local"
	case "uat":
		mode = "uat"
	case "prod":
		mode = "production"
	}
	return mode
}

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")

	viper.SetConfigName("production")
	configName := getConfigName()
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic("Error reading config")
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to unmarshal config")
	}

}
