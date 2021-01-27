package service

import (
	_ "GoTest/constant"
	"GoTest/domain"
	_ "GoTest/domain"
	"GoTest/redis"
	_ "GoTest/redis"
	"encoding/json"
)

type Action struct {
}

const (
	Login       = "Login"
	PushProblem = "PushProblem"
)

var (
	redisClient = redis.RC
)

func (actionHandle *Action) ActionHandle(jsonBytes []byte) *domain.Result {
	action := &domain.ActionInfo{}
	_ = json.Unmarshal(jsonBytes, action)
	switch action.Action {
	case Login:
		if loginUser, ok := action.Data.(domain.LoginUser); ok {
			return actionHandle.login(loginUser.LessonId, loginUser.UserId, loginUser.UserType)
		} else {
			goto ERR
		}
	case PushProblem:
		if problem, ok := action.Data.(domain.Problem); ok {
			return actionHandle.pushProblem(problem)
		} else {
			goto ERR
		}
	default:
		return &domain.Result{Status: 501, Msg: "Action Not Found"}
	}
ERR:
	return &domain.Result{Status: 400, Msg: "参数错误"}
}



