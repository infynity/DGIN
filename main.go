package main

import (
	. "dcgin/config"
	. "dcgin/model"
	"dcgin/router"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"log"
	"os"
	"reflect"
	"time"
)

//ex
type ClockSyncConfig struct {
	SyncSwitch int `json:"sync_switch"`

	SyncScript string `json:"sync_script"`
}

type abc struct {
}

func (thi *abc) sad(config *ClockSyncConfig) {
	config.SyncScript = "ooopasdppp"
}

func main() {

	s1 := "{\n\"sync_switch\":1,\"sync_period\":[1,7,30],\"checked_period\":1,\"sync_script\":\"echo %s\",\"ntp_server\":\"\"}"

	s2 := "{\n\"sync_script\":\"echo %s\",\"sync_switch\":2,\"sync_period\":[1,7,30],\"checked_period\":1,\"ntp_server\":\"\"}"
	cc := ClockSyncConfig{}
	cc2 := ClockSyncConfig{}
	err := json.Unmarshal([]byte(s1), &cc)
	er2 := json.Unmarshal([]byte(s2), &cc2)
	fmt.Println(err, cc, cc2, er2)
	fmt.Println(reflect.DeepEqual(cc, cc2), 3333333)

	oldPcConfigByte, err := json.Marshal(cc)
	if err != nil {
		//log.Error("Marshal ClockSyncConfig Json error:", err.Error())
		//return false
	}
	currPcConfigs, err := json.Marshal(cc2)
	if err != nil {

	}
	if string(oldPcConfigByte) == string(currPcConfigs) {
		fmt.Println("oldClockSyncConfig == curClockSyncConfig")
	}

	sd := ClockSyncConfig{SyncScript: "scasd", SyncSwitch: 1}

	f := abc{}
	f.sad(&sd)
	fmt.Println(sd)
	//return

	//加载配置
	InitConfig()
	//日志配置
	InitLog()
	//	split := strings.Split("2,3,41,123", ",")
	//	fmt.Println(split)
	//return
	//GetRwcfg()
	//return
	//数据库配置
	InitDb()
	//run
	router.Init()

}

func InitLog() {
	ss := &zap.Config{
		OutputPaths: []string{"zap.log"},
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		InitialFields: map[string]interface{}{
			"app": "test",
		},
		Encoding: "json"}
	logger, err := ss.Build()
	if err != nil {
		log.Println(err, 12222)
	}
	//logger, _ := zap.NewProduction()

	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "123",
		"attempt", 3,
		"backoff", time.Second,
	)
	//sugar.Infof("Failed to fetch URL: %s", "123")
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "faster?"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	//hex
	fmt.Println(0x8400, 19191919, 0x8, 0x9, 0x10, 0xa, 0xB,0xe,0xf)
	if AppConfig.Env == "online" {
		//使用daily加入日期
		f, _ := os.Create(fmt.Sprintf("gin_%s.log", time.Now().Format("2006-01-02")))
		// 同时将日志写入文件和控制台
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
		Logger = log.New(f, "pre", 1)
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

}
