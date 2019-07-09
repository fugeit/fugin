package conf

import (
	"fmt"
	"fugin/models"
	"github.com/Unknwon/goconfig"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func New() *gorm.DB{
	return db
}

// 初始化各种配置
func init()  {
	var (
		err  error
		sqlConn string
		config *goconfig.ConfigFile
		confData map[string]string

	)
	// 加载环境变量
	if config, err = goconfig.LoadConfigFile("E:/Go语言开发/Go语言开发练习/project/src/fugin/conf/app.conf"); err != nil {
		log.Println("加载环境配置错误", err)
	}

	if confData, err = config.GetSection(""); err != nil {
		log.Println("获取环境变量错误", err)
	}
	dbname := confData["dbname"]
	// [username[:password]@][protocol[(address)]]/dbname
	sqlConn = fmt.Sprintf("root:mysql@/%s?charset=utf8&parseTime=True&loc=Local", dbname)
	db, err = gorm.Open("mysql", sqlConn)
	if err != nil {
		log.Println("mysql连接错误", err)
		return
	}

	deafult := &models.Default{}
	log.Println(db.HasTable(deafult))
	if !db.HasTable(deafult) {
		db.CreateTable(deafult)
	}
}
