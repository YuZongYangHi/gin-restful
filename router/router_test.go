package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UsersViews struct {
	RouterEngine *gin.RouterGroup
}

type GroupsViews struct {
	RouterEngine *gin.RouterGroup
}

func (u *UsersViews) Delete(ctx *gin.Context) {
	ctx.JSON(200, "delete")
}
func (u *UsersViews) Update(ctx *gin.Context) {
	ctx.JSON(200, "update")
}
func (u *UsersViews) Retrieve(ctx *gin.Context) {
	ctx.JSON(200, "get")
}
func (u *UsersViews) List(ctx *gin.Context) {
	ctx.JSON(200, "LIST")
}
func (u *UsersViews) Create(ctx *gin.Context) {
	ctx.JSON(200, "11111")
}
func (u *UsersViews) ExtraURI(fullPath string, ctx *gin.Context) {
	fmt.Println(ctx.FullPath())
	if ctx.Request.RequestURI == fullPath+"/heihei/heihei" {
		ctx.JSON(200, "other")
	}
	ctx.JSON(200, "2222")
}

func (u *GroupsViews) Delete(ctx *gin.Context) {
	ctx.JSON(200, "delete")
}
func (u *GroupsViews) Update(ctx *gin.Context) {
	ctx.JSON(200, "update")
}
func (u *GroupsViews) Retrieve(ctx *gin.Context) {
	ctx.JSON(200, "get")
}
func (u *GroupsViews) List(ctx *gin.Context) {
	ctx.JSON(200, "list")
}
func (u *GroupsViews) Create(ctx *gin.Context) {
	ctx.JSON(200, "create")
}
func (u *GroupsViews) ExtraURI(fullPath string, ctx *gin.Context) {
	ctx.JSON(200, "extends")
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	v1.Any("/users/*resource", (&RouterController{&UsersViews{v1}}).Dispath)
	v1.Any("/groups/*resource", (&RouterController{&GroupsViews{v1}}).Dispath)

	err := router.Run(":8080")

	if err != nil {
		fmt.Println("start Server error")
		return
	}
}
