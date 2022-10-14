package src

// type Topic struct {
// 	TopicID    int    `json:"id" form:"id"`
// 	TopicTitle string `json:"title" form:"title" binding:"required"`
// }

type Topic struct {
	TopicID         int    `json:"id" form:"id"`
	TopicTitle      string `json:"title" form:"title" binding:"min=4,max=20"`
	TopicShortTitle string `json:"stitle" form:"stitle" binding:"required,nefield=TopicTitle"`
	UserIP          string `json:"ip" form:"ip" binding:"ipv4"`
	TopicScore      int    `json:"score" form:"score"`
}

func CreateTopic(id int, title string) Topic {
	return Topic{id, title}
}

type TopicQuery struct {
	UserName string `json:"username" form:"username"`
	Page     int    `json:"page" form:"page" binding:"required"`
	PageSize int    `json:"pagesize" form:"pagesize"`
}
