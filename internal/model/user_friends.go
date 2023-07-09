package model

import "fmt"

//todo 这里应该 从队列里面取出离线消息 (用户登录之后先从biz_user这张表查询出自己的user_id 然后拿到所有好友的username)
func GetOwnerUserDFromBizUser(username string) int {
	var bu BizUser
	err := DB.Where("username = ?", username).First(&bu).Error
	if err != nil {
		fmt.Println(" (用户登录之后先从biz_user这张表查询出自己的user_id 然后拿到所有好友的username) 失败", err)

	}
	return int(bu.ID)
}

//todo 查询一对多的所有数据
func FindOneToMore(bizUserID int) []Friend {
	var friends []Friend
	err := DB.Where("biz_user_id = ?", bizUserID).Find(&friends).Error
	if err != nil {
		fmt.Println("查询一对多数据失败", err)
	}
	return friends
}
