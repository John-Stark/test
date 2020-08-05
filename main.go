package main

import (
	"fmt"
	"loginimpl/dao"
	"loginimpl/model"
	"loginimpl/routers"
)

func main() {
	err := dao.InitDB()
	if err != nil {
		fmt.Println(err)
	}
	dao.Db.AutoMigrate(&model.User{})
	defer dao.CloseDB()

	r := routers.SetupRouter()
	r.Run(":8080")
}
