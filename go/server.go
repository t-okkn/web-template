package main

import (
	"github.com/gin-gonic/gin"
)

// <summary>: 待ち受けるサーバのルーターを定義します
// <remark>: httpHandlerを受け取る関数にそのまま渡せる
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", sample)

	 return router
}

// <summary>: サンプル
func sample(c *gin.Context) {
	//途中で処理を中断
	//c.Abort()
	//return
}

