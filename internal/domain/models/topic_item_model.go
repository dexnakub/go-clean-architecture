package domain_models

type TopicItemCreateModel struct {
	TopicItemID    int    `json:"topicItemId"`
	TopicID        int    `json:"topicId"`
	TopicItemTitle string `json:"topicItemTitle"`
	Sequence       int    `json:"sequence"`
	IsDeleted      *bool  `json:"isDeleted"`
}

type TopicItemGetModel struct {
	TopicItemID    int    `json:"topicItemId"`
	TopicID        int    `json:"topicId"`
	TopicItemTitle string `json:"topicItemTitle"`
	Sequence       int    `json:"sequence"`
	IsDeleted      *bool  `json:"isDeleted"`
	*BaseModel
}

type TopicItemUpdateModel struct {
	TopicItemID    int    `json:"topicItemId"`
	TopicID        int    `json:"topicId"`
	TopicItemTitle string `json:"topicItemTitle"`
	*BaseModel
}

type TopicItemUpdateDeleteStatusModel struct {
	TopicItemID int `json:"topicItemId"`
	TopicID     int `json:"topicId"`
	*BaseModel
}

type TopicItemUpdateSequenceModel struct {
	TopicItemID int `json:"topicItemId"`
	TopicID     int `json:"topicId"`
	Sequence    int `json:"sequence"`
	*BaseModel
}
