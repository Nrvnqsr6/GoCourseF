package service

import (
	"final/internal/action"
	"final/internal/consts"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/task/"+consts.SHIFT, func(ctx *gin.Context) {
		res, err := action.CircularShift()
		if err != nil {
			ctx.Writer.Write([]byte(err.Error()))
			return
		}
		ctx.IndentedJSON(http.StatusOK, res)
	})

	r.GET("/task/"+consts.SUBSEQUENCE, func(ctx *gin.Context) {
		res, err := action.Subsequence()
		if err != nil {
			ctx.Writer.Write([]byte(err.Error()))
			return
		}
		ctx.IndentedJSON(http.StatusOK, res)
	})

	r.GET("/task/"+consts.ABSENTELEMENT, func(ctx *gin.Context) {
		res, err := action.AbsentElement()
		if err != nil {
			ctx.Writer.Write([]byte(err.Error()))
			return
		}
		ctx.IndentedJSON(http.StatusOK, res)
	})

	r.GET("/task/"+consts.ARRAYENTRY, func(ctx *gin.Context) {
		res, err := action.ArrayEntry()
		if err != nil {
			ctx.Writer.Write([]byte(err.Error()))
			return
		}
		ctx.IndentedJSON(http.StatusOK, res)
	})

	r.GET("/tasks", func(ctx *gin.Context) {
		res, err := action.SolveAndApproveAll()
		if err != nil {
			ctx.Writer.Write([]byte(err.Error()))
			return
		}
		ctx.IndentedJSON(http.StatusOK, res)
	})

	return r
}
