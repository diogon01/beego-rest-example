package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["cripto-moedas/controllers:DepositController"] = append(beego.GlobalControllerRouter["cripto-moedas/controllers:DepositController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cripto-moedas/controllers:DepositController"] = append(beego.GlobalControllerRouter["cripto-moedas/controllers:DepositController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cripto-moedas/controllers:DepositController"] = append(beego.GlobalControllerRouter["cripto-moedas/controllers:DepositController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cripto-moedas/controllers:DepositController"] = append(beego.GlobalControllerRouter["cripto-moedas/controllers:DepositController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cripto-moedas/controllers:DepositController"] = append(beego.GlobalControllerRouter["cripto-moedas/controllers:DepositController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cripto-moedas/controllers:DepositController"] = append(beego.GlobalControllerRouter["cripto-moedas/controllers:DepositController"],
        beego.ControllerComments{
            Method: "SeedDeposit",
            Router: "/seed",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cripto-moedas/controllers:DepositController"] = append(beego.GlobalControllerRouter["cripto-moedas/controllers:DepositController"],
        beego.ControllerComments{
            Method: "UpdateStatus",
            Router: "/status/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
