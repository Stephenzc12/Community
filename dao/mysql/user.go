package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"user-management/model"
)

const secret = "ZHANG CHENG"

// 把每一步数据库操作封装成函数
// 待logic层根据业务需求调用

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from users where username=?`
	var count int
	if err := db.Raw(sqlStr, username).Scan(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

// Login 在登陆时检查数据库中是否有该用户
func Login(user *model.User) (err error) {
	opassword := user.Password
	sqlStr := `select user_id from users where username=?`
	if err := db.Raw(sqlStr, user.Username).Scan(&user.UserID).Error; err != nil {
		fmt.Printf("check failed, err:%v\n", err)
		return err
	}
	var password string
	sqlStr2 := `select password from users where username=?`
	if err := db.Raw(sqlStr2, user.Username).Scan(&password).Error; err != nil {
		fmt.Printf("check failed, err:%v\n", err)
		return err
	}
	//判断密码是否正确
	if password != encryptPassword(opassword) {
		return errors.New("密码错误")
	}
	return
}

// InsertUser 向数据库中插入一条新的用户信息
func InsertUser(user *model.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	// 执行SQL语句入库
	sqlStr := `insert into users(user_id,username,password) values(?,?,?)`
	err = db.Exec(sqlStr, user.UserID, user.Username, user.Password).Error
	return
}

func CreateTableUser() (err error) {
	createTableUser := `CREATE TABLE IF NOT EXISTS users
	(
		user_id VARCHAR(255),
		username VARCHAR(255),
		password VARCHAR(255)
	)`
	err = db.Exec(createTableUser).Error
	if err != nil {
		fmt.Println("create table users error:", err)
	}
	return
}

func encryptPassword(opassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(opassword)))
}
