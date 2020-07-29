package main

import (
	"fmt"
	"gin-restful/router"
	"github.com/gin-gonic/gin"
)

type UsersViews struct{}

type GroupsViews struct{}

func (u *UsersViews) ExtendsAction() map[string]gin.HandlerFunc {
	m := make(map[string]gin.HandlerFunc, 0)
	m["d/callback/"] = u.Callback
	return m

}

func (u *GroupsViews) ExtendsAction() map[string]gin.HandlerFunc {
	m := make(map[string]gin.HandlerFunc, 0)

	m["d/callback/"] = u.Callback
	return m

}

func (u *GroupsViews) Callback(ctx *gin.Context) {
	ctx.JSON(200, "callback")
}

func (u *UsersViews) Callback(ctx *gin.Context) {
	ctx.JSON(200, "callback")
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

func (u *GroupsViews) Delete(ctx *gin.Context) {
	ctx.JSON(200, "delete1")
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


func UserNew() *UsersViews {
	return &UsersViews{}
}

func GroupNew() *GroupsViews {
	return &GroupsViews{}
}

func ControllerNew() *router.RouterController {
	return &router.RouterController{}
}


func main() {

	router := gin.Default()
	controller := ControllerNew()

	// views
	vu := UserNew()
	vg := GroupNew()


	v1 := router.Group("/api/v1")

	v1.Any("/users/*resource", controller.Register(vu))
	v1.Any("/groups/*resource", controller.Register(vg))

	err := router.Run(":8080")

	if err != nil {
		fmt.Println("start Server error")
		return
	}
}