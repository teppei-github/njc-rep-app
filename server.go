package main

import (
	"net/http"

	"github.com/labstack/echo"

	"app/setting"

	"app/withApis"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	// 設定取得
	conf := setting.GetConfig()

	e := echo.New()
	handle(e)
	e.Logger.Fatal(e.Start(":" + conf.Port))
}

func handle(e *echo.Echo) {
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })

	// e.GET("/foo", func(c echo.Context) error {
	// 	ret, err := withApis.GetFoo()
	// 	if err != nil {
	// 		return c.JSON(http.StatusInternalServerError, err.Error)
	// 	}
	// 	return c.JSON(http.StatusOK, ret)
	// })

	// 受注 (SARECO)
	e.GET("/SARECO", func(c echo.Context) error {
		torihikiCd := c.QueryParam("CSMCSTCD")
		juchuNo := c.QueryParam("RORNU")
		saleYoteiYmdFrom := c.QueryParam("SALSCDDEFR")
		saleYoteiYmdTo := c.QueryParam("SALSCDDETO")
		nonyuName := c.QueryParam("CSMCSTNM")
		telNo := c.QueryParam("TELNU")
		isKanryoChk := c.QueryParam("KANRYOCHK")

		c.Response().Header().Set("Access-Control-Allow-Origin", "*")

		ret, err := withApis.GetJuchu(torihikiCd, juchuNo, saleYoteiYmdFrom, saleYoteiYmdTo, nonyuName, telNo, isKanryoChk)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error)
		}
		// 取得件数が0の場合、エラーとする
		if len(ret.Views) == 0 {
			message := Message{"data is nothing."}
			return c.JSON(http.StatusNotFound, message)
		}
		return c.JSON(http.StatusOK, ret)
	})

	// 受注明細 (SARECD)
	e.GET("/SARECD", func(c echo.Context) error {
		juchuNo := c.QueryParam("RORNU")
		juchuKobanFrom := c.QueryParam("EROARNNUFR")
		juchuKobanTo := c.QueryParam("EROARNNUTO")
		isKanryoChk := c.QueryParam("KANRYOCHK")

		c.Response().Header().Set("Access-Control-Allow-Origin", "*")

		if len(juchuNo) == 0 {
			message := Message{"Param(RORNU) is nothing."}
			return c.JSON(http.StatusBadRequest, message)
		}
		if len(juchuKobanFrom) == 0 {
			message := Message{"Param(EROARNNUFR) is nothing."}
			return c.JSON(http.StatusBadRequest, message)
		}
		if len(juchuKobanTo) == 0 {
			message := Message{"Param(EROARNNUTO) is nothing."}
			return c.JSON(http.StatusBadRequest, message)
		}
		ret, err := withApis.GetJuchuMeisai(juchuNo, juchuKobanFrom, juchuKobanTo, isKanryoChk)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error)
		}
		// 取得件数が0の場合、エラーとする
		if len(ret.Views) == 0 {
			message := Message{"data is nothing."}
			return c.JSON(http.StatusNotFound, message)
		}
		return c.JSON(http.StatusOK, ret)
	})
}
