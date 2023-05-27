package models

import (
	"fmt"
	"ginchat/utils"

	"gorm.io/gorm"
)

type Community struct {
	gorm.Model
	Name    string //群名称
	OwnerId uint   //群归属于
	Img     string //群图片
	Desc    string //描述信息
}

func CreateCommunity(community Community) (int, string) {
	if len(community.Name) == 0 {
		return -1, "群名称不能为空"
	}
	if community.OwnerId == 0 {
		return -1, "请先登录"
	}
	if err := utils.DB.Create(&community).Error; err != nil {
		fmt.Println(err)
		return -1, "建群失败"

	}
	return 0, "建群成功"

}
func LoadCommunity(ownerId uint) ([]*Community, string) {
	data := make([]*Community, 10)
	utils.DB.Where("owner_id = ?", ownerId).Find(&data)
	for _, v := range data {
		fmt.Println(v)

	}

	return data, "查询成功"

}
