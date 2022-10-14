package src

type Topic struct {
	TopicID    int
	TopicTitle string
}

func CreateTopic(id int, title string) Topic {
	return Topic{id, title}
}
