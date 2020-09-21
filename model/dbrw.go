package model


import (
	"pkg.deepin.com/service/lib/storage/db"
	"pkg.deepin.com/service/utils/config"
)

func main() {
	// 加载配置文件
	type MyfConfig struct {
		DB *db.ConnConf // DB配置
	}
	var cfg MyfConfig
	err := config.Load("db.toml", &cfg)
	if nil != err {
		panic(err)
	}

	// 初始化默认的DB链接
	db.InitConn(cfg.DB)

	type User struct {
		Id   int
		Age  int
		Name string
	}

	var user User
	db.Default().First(&user, "name = ?", "lsw")      // 读取
	db.Default().Create(&User{Id: 1000, Name: "lsw"}) // 创建
	db.Default().Model(&user).Update("Age", 999)      // 更新
	db.Default().Delete(&user)                        // 删除
}
