package service

import (
	"GoTest/constant"
	"GoTest/domain"
	"encoding/json"
	"strconv"
)

/**
登录
*/
func (*Action) login(lessonId int64, id int64, userType int) *domain.Result {
	redisUser := &domain.RedisUser{
		UserType:        userType,
		Id:              id,
		ApplicationName: constant.ApplicationNodeName,
	}
	data, _ := json.Marshal(redisUser)
	redisClient.HMSet(constant.RedisRoomKey+strconv.FormatInt(lessonId, 10), map[string]interface{}{
		strconv.Itoa(userType) + "_" + strconv.FormatInt(id, 10): string(data),
	})
	return &domain.Result{Status: 200, Msg: "登陆成功"}
}

