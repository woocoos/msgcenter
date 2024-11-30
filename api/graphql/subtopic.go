package graphql

type SubTopic string

var (
	SubTopicMessage SubTopic = "message"
)

func SubTopics() []string {
	return []string{string(SubTopicMessage)}
}
