package models

import (
	"fmt"
	"ginchat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model

	Name          string
	PassWord      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string
	ClientIp      string //客户IP
	ClientPort    string //客户端口
	Salt          string
	LoginTime     *time.Time
	HeartbeatTime *time.Time
	LoginOutTime  *time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)

	}

	return data

}

//登录
func FindUserByNameAndPwd(name string, password string) UserBasic {
	user := UserBasic{}

	utils.DB.Where("name = ? and pass_word = ?", name, password).First(&user)
	//token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.MD5Encode(str)
	utils.DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp)

	return user

}

//通过名字查找用户
func FindUserByName(name string) UserBasic {
	user := UserBasic{}

	utils.DB.Where("name = ?", name).First(&user)

	return user

}

//通过电话号码查找用户
func FindUserByPhone(phone string) *gorm.DB {
	user := UserBasic{}

	return utils.DB.Where("phone = ?", phone).First(&user)

}

//通过邮箱查找用户
func FindUserByEmail(email string) *gorm.DB {
	user := UserBasic{}

	return utils.DB.Where("email = ?", email).First(&user)

}

//新增用户
func CreateUser(user UserBasic) *gorm.DB {

	return utils.DB.Create(&user)

}

//删除用户
func DeleteUser(user UserBasic) *gorm.DB {

	return utils.DB.Delete(&user)

}

//修改用户
func UpdateUser(user UserBasic) *gorm.DB {

	return utils.DB.Model(&user).Updates(UserBasic{Name: user.Name, PassWord: user.PassWord, Phone: user.Phone, Email: user.Email})

}

//查找某个用户
func FindByID(id uint) UserBasic {
	user := UserBasic{}
	utils.DB.Where("id = ?", id).First(&user)
	return user

}
