package main

import (
	"github.com/spf13/viper"
	"fmt"
	"os"
	"path/filepath"
	"context"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
)

/**
https://blog.csdn.net/sd653159/article/details/83143760
 */



type CompanyInfomation struct{
	Name string
	MarketCapitalization int64
	EmployeeNum int64
	Department []interface{}
	IsOpen bool
}

type YamlSetting struct{
	TimeStamp string
	Address string
	Postcode int64
	CompanyInfomation CompanyInfomation
}

// 反序列化
func parseYaml(v *viper.Viper)  {
	var yamlObj YamlSetting
	if err := v.Unmarshal(&yamlObj); err != nil {
		fmt.Printf("err: %s", err)
	}

	fmt.Println(yamlObj)
}


/**
读取命令行
Golang | flag & pflag介绍
https://blog.csdn.net/liuyh73/article/details/83025435
 */
func readCommandLine()  {
	pflag.String("hostAddress", "127.0.0.1", "Server running address")
	pflag.Int64("port", 8080, "Server running port")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	fmt.Printf("hostAddress :%s , port:%s", viper.GetString("hostAddress"), viper.GetString("port"))
	//hostAddress :192.192.1.10 , port:9000
	/*
	example:
	go run main2.go --hostAddress=192.192.1.10 --port=9000
	help:
	Usage of /tmp/go-build183981952/b001/exe/main:
     --hostAddress string   Server running address (default "127.0.0.1")
     --port int             Server running port (default 8080)
	*/
}

func main()  {

	// 多个viper, 也可以直接使用viper.xxx
	v := viper.New()

	// 设置读取的配置文件
	v.SetConfigName("linux_config")

	// 添加读取的配置文件路径
	v.AddConfigPath("./config/")

	gopath := os.Getenv("GOPATH")
	for _, p := range filepath.SplitList(gopath) {
		peerpath := filepath.Join(p, "src/")
		v.AddConfigPath(peerpath)
	}

	// windows环境下为%GOPATH，linux环境下为$GOPATH
	//v.AddConfigPath("%GOPATH/src/")

	// 设置配置文件类型
	v.SetConfigType("yaml")

	// viper解析配置文件
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
	}


	fmt.Printf(
		`
		TimeStamp:%s
		CompanyInfomation.Name:%s
		CompanyInfomation.Department:%s `,
		v.Get("TimeStamp"),
		v.Get("CompanyInfomation.Name"),
		v.Get("CompanyInfomation.Department"),
	)

	fmt.Println("\n")

	parseYaml(v)
	//{2018-10-18 10:09:23 Shenzhen 518000 {Sunny 50000000 200 [Finance Design Program Sales] false}}

	// 读取命令行
	 readCommandLine()

	/**
	很多时候，我们服务器启动之后，如果临时想修改某些配置参数，需要重启服务器才能生效，
	但是viper提供了监听函数，可以免重启修改配置参数，非常的实用：
	 */

	//创建一个信道等待关闭（模拟服务器环境）
	 ctx, _ := context.WithCancel(context.Background())

	//cancel可以关闭信道
	// ctx, cancel := context.WithCancel(context.Background())
	//设置监听回调函数
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("config is change : %s \n", in.String())
		//cancel()
	})

	 // 开始监听
	 v.WatchConfig()

	//信道不会主动关闭，可以主动调用cancel关闭
	<-ctx.Done()

}