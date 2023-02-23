package main

import (
    "github.com/labstack/echo"
    "src/routing"
)

func main() {
    e := echo.New()
    // routing
    e.GET("/user",routing.BaseAPI_GET())
    e.POST("/user",routing.BaseAPI_POST())

    //e.Logger.Fatal(e.StartTLS(":443", "/crt/live/<あなたのドメイン名>/fullchain.pem", "/crt/live/<あなたのドメイン名>/privkey.pem"))
    e.Logger.Fatal(e.Start(":8080"))
}