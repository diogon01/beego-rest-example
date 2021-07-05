package main

import (
	"cripto-moedas/models"
	_ "cripto-moedas/routers"

	"fmt"

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
	o := orm.NewOrm()

	deposit := new(models.Deposit)
	deposit.Id = 12
	deposit.Email = "diogo.henrique@gmail.com"
	deposit.Txid = "1ExXt54FaFfRoJb6AUdimvxbT5GuDHLwT6"
	deposit.Currency = "OPTT"
	deposit.Amount = 456.23
	deposit.Status = "status"

	fmt.Println(o.Insert(deposit))

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		orm.Debug = true
	}
	beego.Run()
}
