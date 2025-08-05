package data_entities

type TopicEntity struct {
	TopicID    int    `gorm:"primaryKey;autoIncrement"`
	TopicTitle string `gorm:"type:nvarchar(255);not null;unique"`
	*BaseEntity
}

func (*TopicEntity) TableName() string {
	return "topics"
}
