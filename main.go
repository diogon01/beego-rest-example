package main

import (
	_ "cripto-moedas/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)

	orm.RegisterDataBase("default",
		"postgres",
		"user=dev password=dev host=127.0.0.1 port=5432 dbname=gotest sslmode=disable")

	orm.RunSyncdb("default", false, true)

}

func main() {

	// Using default, you can use other database

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		orm.Debug = true
	}
	beego.Run()
}
