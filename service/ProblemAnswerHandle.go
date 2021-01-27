package service

import "GoTest/domain"

/**
提交问题
*/
func (*Action) pushProblem(problemInfo domain.Problem) (result *domain.Result) {
	//取到问题 存入数据库

	//分发给所有的学生
	return &domain.Result{}
}

/**
分发问题
*/
func (*Action) distributionProblem(problemInfo domain.Problem) (result *domain.Result) {

	return &domain.Result{}
}

/**
提交答案
*/
func (*Action) pushAnswer(answerInfo domain.Answer) (result *domain.Result) {

	return &domain.Result{}
}

/**
推送结果
*/
func (*Action) pushResult(lessonId int64, userType int8) (result *domain.Result) {

	return &domain.Result{}
}
