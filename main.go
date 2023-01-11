package main

import (
	"emailQue/shared"
	"emailQue/utils"
	"fmt"
	"os"

	"github.com/go-ini/ini"
)

func main() {
	logger := utils.InitLogger()
	defer logger.Sync()

	cfg, err := ini.Load("my.ini")

	if err != nil {
		fmt.Printf("cant find config file my.ini: %s", err)
		os.Exit(1)
	}
	// db := util.InitDB()
	db := utils.InitDB(cfg)
	defer db.Close()

	utils.Env.DB = db
	utils.Env.Log = logger.Sugar()
	utils.Env.Cfg = cfg

	if err := shared.Worker(); err != nil {
		utils.Env.Log.Debug(err)
		os.Exit(1)
	}

}
