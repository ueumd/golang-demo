package main

import "github.com/spf13/viper"
import (
	"log"
	"bytes"
	"os"
	"strings"
	"errors"
	"testing"
	"flag"
	"github.com/spf13/pflag"
	"github.com/fsnotify/fsnotify"
	"context"
	"gopkg.in/yaml.v2"
	"encoding/json"
	"fmt"
)

type config struct {
	v *viper.Viper
}

// 加载配置文件
func LoadConfigFromYaml (c *config) error  {
	c.v = viper.New()

	//设置配置文件的名字
	c.v.SetConfigName("config")

	//添加配置文件所在的路径,注意在Linux环境下%GOPATH要替换为$GOPATH
	c.v.AddConfigPath("%GOPATH/src/")
	c.v.AddConfigPath("./demo_config")

	//设置配置文件类型
	c.v.SetConfigType("yaml")

	if err := c.v.ReadInConfig(); err != nil{
		return  err
	}

	log.Printf("age: %s, name: %s n", c.v.Get("information.age"), c.v.Get("information.name"))
	return nil
}

// 由IO读取配置
func ReadConfigFormIo(c *config) error {
	c.v = viper.New()
	if f, err := os.Open("./demo_config/config.yaml"); err != nil{
		log.Printf("filure: %s", err.Error())
		return err
	}else {

		confLength, _ :=f.Seek(0,2)

		//注意，通常写c++的习惯害怕读取字符串的时候越界，都会多留出一个NULL在末尾，但是在这里不行，会报出如下错误：
		//While parsing config: yaml: control characters are not allowed
		//错误参考网址：https://stackoverflow.com/questions/33717799/go-yaml-control-characters-are-not-allowed-error

		configData := make([]byte, confLength)
		f.Seek(0, 0)
		f.Read(configData)
		log.Printf("%sn", string(configData))

		c.v.SetConfigType("yaml")
		if err := c.v.ReadConfig(bytes.NewBuffer(configData)); err != nil{
			log.Fatalf(err.Error())
		}
	}

	log.Printf("age: %s, name: %s n", c.v.Get("information.age"), c.v.Get("information.name"))
	return nil
}


/**
读取本地的环境变量
 SetEnvPrefix 和 AutomaticEnv、BindEnv搭配使用很方便，
比如说我们把当前程序的环境变量都设置为xx_
这样方便我们管理，也避免和其他环境变量冲突
而在读取的时候又很方便的就可以读取。
 */
func EnvConfigPrefix(c *config) error {
	c.v = viper.New()
	//BindEnv($1,$2)
	// 如果只传入一个参数，则会提取指定的环境变量$1，如果设置了前缀，则会自动补全 前缀_$1
	//如果传入两个参数则不会补全前缀，直接获取第二参数中传入的环境变量$2
	os.Setenv("LOG_LEVEL", "INFO")
	if nil == c.v.Get("LOG_LEVEL ") {
		log.Printf("LOG_LEVEL is nil")
	}else {
		// return ErrorNotMacth
		return errors.New("ErrorNotMacth")
	}

	//必须要绑定后才能获取
	c.v.BindEnv("LOG_LEVEL")
	log.Printf("LOG_LEVEL is %s", os.Getenv("log_level"))


	//会获取所有的环境变量，同时如果过设置了前缀则会自动补全前缀名
	c.v.AutomaticEnv()
	//环境变量前缀大小写不区分
	os.Setenv("DEV_ADDONES","none")
	log.Printf("DEV_ADDONES: %s", c.v.Get("dev_addones"))

	//SetEnvPrefix会设置一个环境变量的前缀名
	c.v.SetEnvPrefix("DEV")

	os.Setenv("DEV_MODE", "true")
	//此时会自动补全前缀，实际去获取的是DEV_DEV_MODE
	if nil ==  c.v.Get("dev_mode"){
		log.Printf("DEV_MODE is nil")
	}else {
		return errors.New("ErrorNotMacth")
	}

	//此时我们直接指定了loglevel所对应的环境变量，则不会去补全前缀
	c.v.BindEnv("loglevel", "LOG_LEVEL")
	log.Printf("LOG_LEVEL: %s", c.v.Get("loglevel"))

	return nil
}

/**
方便的替换符
我们有时候需要去替换key中的某些字符，来转化为对应的环境变脸，
比如说例子中将' . '替换为'_' ，由获取api.version变成了api_version，
但是有一点需要注意的，SetEnvPrefix和SetEnvKeyReplacer一起用的时候可能会混淆。
 */
func EnvCongiReplacer(c *config, setPerfix bool) error {
	c.v = viper.New()
	c.v.AutomaticEnv()
	c.v.SetEnvKeyReplacer(strings.NewReplacer(".","_"))

	// 环境变量
	os.Setenv("MYAPI_VERSION","v1.1.9")

	//Replacer和prefix一起使用可能会冲突,比如我下面的例子
	//因为会自动补全前缀最终由获取API_VERSION变成API_API_VERSION
	// if setPerfix{ c.v.SetEnvPrefix("api")}


	if s := c.v.Get("myapi.version"); s==nil{
		return errors.New("ErrorNotMacth")
	}else {
		log.Printf("%s", c.v.Get("myapi.version"))
	}
	return nil
}


//设置重载 和别名
func SetAndAliases(c *config) error {
	c.v = viper.New()
	c.v.Set("Name","wzp")
	c.v.RegisterAlias("id","Name")
	c.v.Set("id","Mr.Wang")

	//我们可以发现当别名对应的值修改之后，原本的key也发生变化
	log.Printf("id %s, name %s",c.v.Get("id"),c.v.Get("name") )
	return nil
}


/**
 超级实惠的一个功能，直接把配置反序列化到一个结构体，
也可以把设置直接序列化为我们想要的类型：yaml、json等等
 */
type favorite struct {
	Sports []string;
	Music []string;
	LuckyNumber int;
}

type information struct {
	Name string
	Age  int
	Alise []string
	Image string
	Public bool
}

type YamlConfig struct {
	TimeStamp string
	Author string
	PassWd string
	Information information
	Favorite favorite
}



//将配置解析为Struct对象
func UmshalStruct(c *config) error  {
	LoadConfigFromYaml(c)
	var cf YamlConfig
	if err := c.v.Unmarshal(&cf); err != nil{
		return err
	}

	return nil
}

func YamlStringSettings(c *config) string {
	c.v = viper.New()
	c.v.Set("name", "wzp")
	c.v.Set("age", 18)
	c.v.Set("aliase",[]string{"one","two","three"})

	cf := c.v.AllSettings()
	bs, err := yaml.Marshal(cf)
	if err != nil {
		log.Fatalf("unable to marshal config to YAML: %v", err)
	}
	return string(bs)
}

func JsonStringSettings(c *config) string {
	c.v = viper.New()
	c.v.Set("name", "wzp")
	c.v.Set("age", 18)
	c.v.Set("aliase",[]string{"one","two","three"})

	cf := c.v.AllSettings()
	bs, err := json.Marshal(cf)
	if err != nil {
		log.Fatalf("unable to marshal config to YAML: %v", err)
	}
	return string(bs)
}

//从command Line中读取配置
func ReadCommandLine()  {
	flag.String("mode","RUN","please input the mode: RUN or DEBUG");
	pflag.Int("port",1080,"please input the listen port");
	pflag.String("ip","127.0.0.1","please input the bind ip");
	//获取标准包的flag
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	//BindFlag
	//在pflag.Init key后面使用
	viper.BindPFlag("port", pflag.Lookup("port"))
	log.Printf("set port: %d", viper.GetInt("port"))

	viper.BindPFlags(pflag.CommandLine)
	log.Printf("set ip: %s", viper.GetString("ip"))
}

//监听配置文件的修改和变动
func WatchConfig(c *config) error {
	if err := LoadConfigFromYaml(c); err !=nil{
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())

	c.v.WatchConfig()

	//监听回调函数
	watch := func(e fsnotify.Event) {
		log.Printf("Config file is changed: %s n", e.String())
		cancel()
	}

	c.v.OnConfigChange(watch)
	<-ctx.Done()
	return nil
}


/**
拷贝子分支
  拷贝一个子分支最大的用途就是我们可以复制一份配置，这样在修改拷贝的时候原配置不会被修改，如果修改的配置出现了问题，我们可以方便的回滚。
 */
func TestSubConfig(t *testing.T)  {
	c := config{};
	LoadConfigFromYaml(&c)
	sc := c.v.Sub("information")
	sc.Set("age", 80)
	scs,_:=yaml.Marshal(sc.AllSettings())
	t.Log(string(scs))
	t.Logf("age: %d", c.v.GetInt("information.age"))
}


/**
获取配置项的方法
如果我们直接用Get获取的返回值都是interface{}类型，这样我们还要手动转化一下，可以直接指定类型去获取，方便快捷
 */

//测试各种get类型
func TestGetValues(t *testing.T)  {
	c := &config{}
	if err := LoadConfigFromYaml(c); err != nil{
		t.Fatalf("%s: %s",t.Name(), err.Error())
	}

	if info := c.v.GetStringMap("information"); info != nil{
		t.Logf("%T", info)
	}

	if aliases := c.v.GetStringSlice("information.aliases"); aliases != nil{
		for _, a := range  aliases{
			t.Logf("%s",a)
		}
	}

	timeStamp := c.v.GetTime("timestamp")
	t.Logf("%s", timeStamp.String())

	if public := c.v.GetBool("information.public"); public{
		t.Logf("the information is public")
	}

	age := c.v.GetInt("information.age")
	t.Logf("%s age  is %d", c.v.GetString("information.name"), age)
}

//https://www.colabug.com/4116973.html
func main()  {
	var c = &config{}
	LoadConfigFromYaml(c)
	// ReadConfigFormIo(c)
	EnvCongiReplacer(c, true)

	s := YamlStringSettings(c)
	fmt.Println(s)

	// UmshalStruct(c)
}