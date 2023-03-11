package main

import (
	"fmt"
	"log"
	"net/http"

	contracts "kxu913/demo/nft_contracts"
	ntx_tx "kxu913/demo/nft_tx"

	sc "kxu913/demo/nft_scripts"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main(){
	port := ":8888"
	log.Println("Start a server in ", port)
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	  }))
	contractGroup(e)
	transactionsGroup(e)
	e.Logger.Fatal(e.Start(port))
}

func contractGroup(e *echo.Echo) {
	t := e.Group("/contract")

	t.POST("", func(c echo.Context) error {
		address,_ :=contracts.DeployContract(true)
		return c.String(http.StatusOK, fmt.Sprintf("Deployed contract in %s",address))
	})

	t.PUT("", func(c echo.Context) error {
		address,_ := contracts.DeployContract(false)
		return c.String(http.StatusOK, fmt.Sprintf("ReDeployed contract in %s",address))
	})
}

func transactionsGroup(e *echo.Echo) {
	t := e.Group("/usage")

	t.POST("/:asset", func(c echo.Context) error {
		asset:= c.Param("asset")
		txId :=ntx_tx.ExecuteTransacation("InitAssetUsage","",asset)
		return c.String(http.StatusOK, fmt.Sprintf("Asset %s had been init in %s",asset,txId))
	})

	type Action struct{
		Like bool
		Collect bool
	}

	t.POST("/usage/:asset", func(c echo.Context) error {
		asset:= c.Param("asset")
		address := c.Request().Header.Get("address")
		privateKey := c.Request().Header.Get("pkey")
		ntx_tx.ExecuteTransacationUsingMultipleSig("AssignAssetUsage",address,privateKey,asset)
		ntx_tx.ExecuteUserTransacation("RecordUsage",address,privateKey,asset)
		return c.String(http.StatusOK, fmt.Sprintf("ReDeployed contract in %s",address))
	})

	t.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, sc.GetUsage())
	})

}
