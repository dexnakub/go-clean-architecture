package topic_repository

import (
	"fmt"
	data_entities "go-clean-achitech/internal/data/entities"
	data_helpers "go-clean-achitech/internal/data/helpers"
)

func CreateTopic(data *data_entities.TopicEntity) error {
	if data == nil {
		return fmt.Errorf("topic entity is nil")
	}

	gorm := data_helpers.GetGormAdapter()

	return gorm.Model(&data_entities.TopicEntity{}).Create(data).Error
}

func GetTopics(limit, offset int) ([]data_entities.TopicEntity, error) {
	gorm := data_helpers.GetGormAdapter()

	items := []data_entities.TopicEntity{}
	err := gorm.Model(&data_entities.TopicEntity{}).
		Where("is_deleted = ?", false).
		Order("topic_id ASC").
		Limit(limit).
		Offset(offset).
		Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}
