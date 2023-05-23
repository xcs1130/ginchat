package models

import (
	"fmt"
	"ginchat/utils"

	"gorm.io/gorm"
)

//人员关系
type Contact struct {
	gorm.Model
	OwnerId  uint //谁的关系信息
	TargetId uint //对应的谁
	Type     int  //对应的类型 0 1好友 2群 3
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}
func SearchFriend(userId uint) []UserBasic {
	contacts := make([]Contact, 0)
	objIds := make([]uint64, 0)
	utils.DB.Where("owner_id = ? and type = 1", userId).Find(&contacts)
	for _, v := range contacts {
		fmt.Println("..................", v)
		objIds = append(objIds, uint64(v.TargetId))

	}
	users := make([]UserBasic, 0)
	utils.DB.Where("id in ?", objIds).Find(&users)
	return users

}

//添加好友
func AddFriend(userId uint, targetId uint) (int, string) {
	user := UserBasic{}
	if targetId != 0 {
		user = FindByID(targetId)
		if user.Salt != "" {
			if userId == user.ID {
				//自己加自己为好友的情况
				return -1, "不能添加自己为好友"
			}
			contact0 := Contact{}
			utils.DB.Where("owner_id=? and target_id=? and type=1", userId, targetId).Find(&contact0)
			if contact0.ID != 0 {
				return -1, "不能重复添加用户（该用户已被添加）"
			}

			tx := utils.DB.Begin()
			//事务一旦开始，无论什么异常最终都会Rollback
			defer func() {
				if r := recover(); r != nil {
					tx.Rollback()
				}
			}()

			contact := Contact{}
			contact.OwnerId = userId
			contact.TargetId = targetId
			contact.Type = 1
			if err := utils.DB.Create(&contact).Error; err != nil {
				tx.Rollback()
				return -1, "添加好友失败"
			}
			//双向好友，好友的好友里面也有你
			contact1 := Contact{}
			contact1.OwnerId = targetId
			contact1.TargetId = userId
			contact1.Type = 1
			if err := utils.DB.Create(&contact1).Error; err != nil {
				tx.Rollback()
				return -1, "添加好友失败"
			}
			tx.Commit()
			return 0, "添加好友成功"
		}
		return -1, "没有找到该用户"

	}
	return -1, "用户ID不能为空"

}
