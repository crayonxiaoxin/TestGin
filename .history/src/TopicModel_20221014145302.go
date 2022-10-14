package src

// type Topic struct {
// 	TopicID    int    `json:"id" form:"id"`
// 	TopicTitle string `json:"title" form:"title" binding:"required"`
// }

type Topic struct {
	TopicID         int    `json:"id" form:"id"`
	TopicTitle      string `json:"title" form:"title" binding:"min=4,max=20"`                  // 在4-20个字符之间
	TopicShortTitle string `json:"stitle" form:"stitle" binding:"required,nefield=TopicTitle"` // 不能与title相同
	UserIP          string `json:"ip" form:"ip" binding:"ipv4"`                                // 必须是ipv4
	TopicScore      int    `json:"score" form:"score" binding:"omitempty,gt=5"`                // omitempty可不填，有先后顺序。gt=5：如果填了，必须大于5
	TopicUrl        string `json:"url" form:"url" binding:"omitempty,topicurl"`
}

type Topics struct {
	TopicList     []Topic `json:"topics" binding:"gt=0,lt=3,dive"` // dive可以验证子item，写在最后
	TopicListSize int     `json:"size"`
}

func CreateTopic(id int, title string) Topic {
	return Topic{TopicID: id, TopicTitle: title}
}

type TopicQuery struct {
	UserName string `json:"username" form:"username"` // json：格式化时用username。form：传递参数时用username
	Page     int    `json:"page" form:"page" binding:"required"`
	PageSize int    `json:"pagesize" form:"pagesize"`
}
