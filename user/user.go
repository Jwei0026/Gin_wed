package user

import (
	"2024-2025/database"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

//用户表的数据操作
//1.结构
//2.函数

type Users struct {
	User_id      int    `gorm:"primarykey"`
	Username     string `gorm:"size:30"`
	Password     string `gorm:"size:128"`
	Email        string `gorm:"size:50"`
	Phone_number string `gorm:"size:11"`
}

// 定义一个结构，用来接收用户提交的注册数据（表单提交form）
type UserInfoRegister struct {
	Username     string `gorm:"size:30" form:"username" binding:"required,min=6,max=30"`
	Password     string `gorm:"size:20" form:"password" binding:"required,min=6,max=30"`
	Password2    string `form:"password2" binding:"eqfield=Password"`
	Email        string `gorm:"size:50" form:"email" binding:"email"`
	Phone_number string `gorm:"size:11" form:"phone" binding:"required"`
	Captcha      string `form:"captcha" binding:"required"` //form适配表单数据，gorm适配gorm框架，binding适配数据验证
}

// 验证用户名和密码
// 分两部分，分别验证用户名、密码
func CheckLogin(username string, password string) (Users, error) {
	//通过用户名查询数据
	var user Users
	// fmt.Println("这是查询函数输出：", username, password)
	database.Gdb.Where("username = ?", username).Find(&user)

	// fmt.Println(user)
	//若找到说明用户存在，若找不到则返回用户不存在
	if user.Username == "" {
		//查不到结果
		return Users{}, errors.New("该用户不存在！")
	}

	//密码验证
	//这里密码用来md5加密，所以得将原数据进行同样加密再比对，md5单向不可逆
	hash := md5.Sum([]byte(password))
	if user.Password != hex.EncodeToString(hash[:]) {
		//密码不正确
		return Users{}, errors.New("密码不正确")
	}

	return user, nil
}

func CheckRegister(userinfo UserInfoRegister) error {
	var user Users
	//合理性验证：用户名不能重复、邮箱不能重复、手机号不能重复
	database.Gdb.Where("username = ?", userinfo.Username).Find((&user))
	if user.User_id != 0 {
		return errors.New("用户名已存在！")
	}

	database.Gdb.Where("email = ?", userinfo.Email).Find((&user))
	if user.User_id != 0 {
		return errors.New("邮箱已注册，请换一个邮箱注册")
	}

	database.Gdb.Where("phone_number = ?", userinfo.Phone_number).Find((&user))
	if user.User_id != 0 {
		return errors.New("手机号已注册！请换一个手机号注册")
	}

	//账号真正注册入库
	user.Username = userinfo.Username
	hash := md5.Sum([]byte(userinfo.Password))
	user.Password = hex.EncodeToString(hash[:])
	user.Phone_number = userinfo.Phone_number
	user.Email = userinfo.Email
	database.Gdb.Save(&user)

	return nil
}
