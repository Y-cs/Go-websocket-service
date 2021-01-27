package domain

type RedisUser struct {
	UserType        int    `json:"userType"`
	Id              int64  `json:"id"`
	ApplicationName string `json:"applicationName"`
}

type LoginUser struct {
	CourseId int64 `json:"courseId"`
	LessonId int64 `json:"lessonId"`
	UserId   int64 `json:"userId"`
	UserType int   `json:"userType"`
}

type Problem struct {
	Id      int64    `json:"id"`
	Title   string   `json:"title"`
	Body    string   `json:"body"`
	Answers []Answer `json:"answers"`
}

type Answer struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	IsTrue bool   `json:"isTrue"`
}
