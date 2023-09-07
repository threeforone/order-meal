package initialize

import (
	"fmt"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/threeforone/order-meal/server/global"
)

var Order *viper.Viper

func init() {
	viper.SetConfigFile("./conf/config.yaml")
	viper.AutomaticEnv()       // 读取匹配的环境变量
	viper.SetEnvPrefix("MEAL") // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config change:", e.Name)
	})
	if err := viper.Unmarshal(&global.Conf); err != nil {
		panic(fmt.Errorf("unmarshal to conf failed, err:%v", err))
	}

	//读取时间配置如果没有就创建
	Order = viper.New()
	Order.SetConfigFile("./conf/order-time.yaml")

	// 读取或创建配置文件
	if err := Order.ReadInConfig(); err != nil {
		// 如果配置文件不存在，则创建一个新文件
		if os.IsNotExist(err) {
			ChangeOrderTime("9:00", "16:00")
		} else {
			panic(err)
		}
	} else {
		Order.Unmarshal(&global.OrderT)
	}
}

func ChangeOrderTime(startTime, endTime string) {
	Order.Set("end-time", endTime)
	Order.Set("start-time", startTime)
	Order.Unmarshal(global.OrderT)
	Order.WriteConfig()
}
