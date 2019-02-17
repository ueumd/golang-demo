package config
import (
	"os"
	"path/filepath"
	"fmt"
	"github.com/spf13/viper"
	"github.com/lexkong/log"
)

func Init()  {
	initConfig()
	initLog()
}

func initConfig() (e error) {
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

func initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}

	log.InitWithConfig(&passLagerCfg)
}