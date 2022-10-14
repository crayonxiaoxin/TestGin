package src

type Topic struct {
	TopicID    int    `json:"id"`
	TopicTitle string `json:"title"`
}

func CreateTopic(id int, title string) Topic {
	return Topic{id, title}
}

type TopicQuery struct {
	UserName string `json:"username" form:"username"`
	Page     int    `json:"page"`
	PageSize int    `json:"pagesize"`
}
