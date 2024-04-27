package main

import (
	"fmt"
	"user-management/dao/mysql"
	"user-management/pkg/snowflake"
	"user-management/router"
)

func main() {

	// 初始化数据库
	if err := mysql.InitDB(); err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
		return
	}
	//初始化雪花算法生成UID
	if err := snowflake.Init("2020-07-01", 1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	// 注册路由
	r := router.SetUpRoute()
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("run server failed, err:%v", err.Error())
		return
	}

}
