package model

import (
	"fmt"
	"os"
	"pkg.deepin.com/service/lib/storage/db"
	"pkg.deepin.com/service/utils/config"
)

func GetRwcfg() {
	// 加载配置文件
	type MyfConfig struct {
		DB *db.Conf // DB Manager配置
	}
	var cfg MyfConfig
	err := config.Load("config/db.toml", &cfg)
	if nil != err {
		panic(err)
	}

	fmt.Println(cfg.DB,989898,os.Getenv("SERVER_ROOT"))
	//return
	// 初始化默认的DBManager
	db.InitManager(cfg.DB)

	type User struct {
		Id   int
		Age  int
		Name string
	}

	var user User
	db.Master("test").First(&user, "name = ?", "lsw")      // 从Master读取
	fmt.Println(user)
	db.Slave("test").Where("Name =? and id=?","lsw",23).First(&User{}) // 从Master创建
	//db.Master("test").Model(&user).Update("Age", 999)      // 从Master更新
	//db.Master("test").Delete(&user)                        // 从Master删除
	//db.Slave("test").First(&user, "name = ?", "lsw")       // 从Slave读取
}
