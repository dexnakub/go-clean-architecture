package domain_models

type TopicCreateModel struct {
	// TopicID    *int   `json:"topicId"`
	TopicTitle string `json:"topicTitle"`
	// IsDeleted  *bool  `json:"isDeleted"`
}
type TopicCreateDTOModel struct {
	TopicID    *int   `json:"topicId"`
	TopicTitle string `json:"topicTitle"`
	*BaseModel
}

type TopicGetModel struct {
	TopicID    *int   `json:"topicId"`
	TopicTitle string `json:"topicTitle"`
	IsDeleted  *bool  `json:"isDeleted"`
	*BaseModel
}
