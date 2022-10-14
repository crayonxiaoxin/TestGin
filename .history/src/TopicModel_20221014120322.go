package src

type Topic struct {
	TopicID    int    `json:"id" form:"id"`
	TopicTitle string `json:"title" form:"title" binding:"required"`
}

func CreateTopic(id int, title string) Topic {
	return Topic{id, title}
}

type TopicQuery struct {
	UserName string `json:"username" form:"username"`
	Page     int    `json:"page" form:"page" binding:"required"`
	PageSize int    `json:"pagesize" form:"pagesize"`
}
