package data_entities

type TopicItemEntity struct {
	TopicItemID    int    `gorm:"primaryKey;autoIncrement"`
	TopicID        int    `gorm:"not null;index:idx_topicid_title"`
	TopicItemTitle string `gorm:"type:nvarchar(255);not null;index:idx_topicid_title,unique"`
	Sequence       int    `gorm:"not null"`
	*BaseEntity

	Topic TopicEntity `gorm:"foreignKey:TopicID;references:TopicID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (*TopicItemEntity) TableName() string {
	return "topic_items"
}

func (t TopicItemEntity) GetID() int {
	return t.TopicItemID
}
func (t *TopicItemEntity) GetSequence() int {
	return t.Sequence
}
func (t *TopicItemEntity) SetSequence(seq int) {
	t.Sequence = seq
}
func (t TopicItemEntity) GetIDColumnName() string {
	return "topic_item_id"
}
