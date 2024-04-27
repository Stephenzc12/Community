package logic

import (
	"fmt"
	"user-management/dao/mysql"
	"user-management/model"
	"user-management/pkg/jwt"
	"user-management/pkg/snowflake"
)

// 存放业务逻辑的代码

func SingUp(p *model.ParamSignUp) (err error) {
	// 1.判断用户存不存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		fmt.Println(err)
		return err
	}
	// 2.生成UID
	userID := snowflake.GenID()
	//构造一个User实例
	user := &model.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 3.保存进数据库
	return mysql.InsertUser(user)
}

func Login(p *model.ParamLogin) (token string, err error) {
	user := &model.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针,就能拿到user.userID
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	//生成 JWT
	return jwt.GenToken(user.UserID, user.Username)
}
