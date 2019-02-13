package config
import (
	"os"
	"path/filepath"
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() (e error) {
	viper.SetConfigName("config")

	// 添加读取的配置文件路径
	viper.AddConfigPath(".")

	gopath := os.Getenv("GOPATH")
	for _, p := range filepath.SplitList(gopath) {
		peerpath := filepath.Join(p, "src/")
		viper.AddConfigPath(peerpath)
	}

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
		return err
	}

	fmt.Printf("appName:%s\n", viper.GetString("appName"))
	
	return e
}