package utils

import (
	"fmt"
	//mysql 驱动
	_ "github.com/go-sql-driver/mysql"

	"github.com/fsnotify/fsnotify"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	//sqlite3 驱动
	_ "github.com/mattn/go-sqlite3"
)

//Config 配置
type Config struct {
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	IP       string `mapstructure:"ip"`
	Dbport   string `mapstructure:"dbport"`
	Dbname   string `mapstructure:"dbname"`
}

//Conf 初始化配置变量
var Conf = new(Config)

//Db 数据量连接对象
var (
	Db  *sqlx.DB
	err error //一定要加err
)

func init() {
	//调用读取配置方法
	ReadConf()

	// dsn := "root:root@tcp(192.168.1.4:3306)/poetry"

	// // 也可以使用MustConnect连接不成功就panic
	// Db, err = sqlx.Connect("mysql", dsn)

	// //设置连接可以重用的最长时间。
	// Db.SetConnMaxLifetime(5)
	// //设置连接池最大连接数
	// Db.SetMaxOpenConns(100)
	// //设置连接池最大空闲连接数
	// Db.SetMaxIdleConns(100)

	if Conf.Database == "sqlite" {
		//如果配置设置为sqlite执行代码
		Db, err = sqlx.Connect("sqlite3", "./DB.db")
		if err != nil {
			panic(err)
		}
	} else {
		//如果配置设置为mysql执行代码
		// dsn := "root:root@tcp(192.168.1.4:3306)/poetry"
		dsn := Conf.User + ":" + Conf.Password + "@tcp(" + Conf.IP + ":" + Conf.Dbport + ")/" + Conf.Dbname

		// 也可以使用MustConnect连接不成功就panic
		Db, err = sqlx.Connect("mysql", dsn)

		//设置连接可以重用的最长时间。
		Db.SetConnMaxLifetime(5)
		//设置连接池最大连接数
		Db.SetMaxOpenConns(100)
		//设置连接池最大空闲连接数
		Db.SetMaxIdleConns(100)
	}

}

//ReadConf 读取配置
func ReadConf() {
	viper.SetConfigName("config")   // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")     // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("./config") // 查找配置文件所在的路径

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		fmt.Println("读取配置错误", err)
	}

	//反序列化至变量 写在监控之前
	err = viper.Unmarshal(&Conf)
	if err != nil {
		fmt.Println("配置饭序列化错误", err)
	}

	// 监控配置文件变化(变化以后会自动重新读取)
	viper.WatchConfig()
	// 监控配置回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		//反序列化至变量
		err = viper.Unmarshal(&Conf)
		if err != nil {
			fmt.Println("重新配置饭序列化错误", err)
		}
	})
}
