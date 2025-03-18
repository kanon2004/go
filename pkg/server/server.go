package server

import (
	"log"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

	"go-college/pkg/http/middleware"
	"go-college/pkg/server/handler"
)

// Serve HTTPサーバを起動する
func Serve(addr string) {
	e := echo.New()
	// panicが発生した場合の処理
	e.Use(echomiddleware.Recover())
	// CORSの設定
	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		Skipper:      echomiddleware.DefaultCORSConfig.Skipper,
		AllowOrigins: echomiddleware.DefaultCORSConfig.AllowOrigins,
		AllowMethods: echomiddleware.DefaultCORSConfig.AllowMethods,
		AllowHeaders: []string{"Content-Type,Accept,Origin,x-token"},
	}))

	/* ===== URLマッピングを行う ===== */
	// 認証を必要としないAPI
	e.GET("/setting/get", handler.HandleSettingGet())
	e.POST("/user/create", handler.HandleUserCreate())

	// 認証を必要とするAPI
	// AuthenticateMiddlewareによってx-tokenヘッダのチェックとユーザーの特定が行われる
	authAPI := e.Group("", middleware.AuthenticateMiddleware())
	authAPI.GET("/user/get", handler.HandleUserGet())
	authAPI.POST("/user/update", handler.HandleUserUpdate())
	authAPI.GET("/collection/list", handler.HandleCollectionGet())
	authAPI.GET("/ranking/list", handler.HandleRankingGet())
	authAPI.POST("/game/finish", handler.HandleScoreGet())
	
	/* ===== サーバの起動 ===== */
	log.Println("Server running...")
	if err := e.Start(addr); err != nil {
		log.Fatalf("failed to start server. %+v", err)
	}
}
