package topic_usecase

import (
	"fmt"
	data_entities "go-clean-achitech/internal/data/entities"
	topic_repository "go-clean-achitech/internal/data/repositories/topic"
	domain_models "go-clean-achitech/internal/domain/models"
	"time"
)

func CreateTopic(topic domain_models.TopicCreateModel) error {
	if topic.TopicTitle == "" {
		return fmt.Errorf("topic title is empty")
	}
	now := time.Now()

	return topic_repository.CreateTopic(&data_entities.TopicEntity{
		TopicTitle: topic.TopicTitle,
		BaseEntity: &data_entities.BaseEntity{
			CreateBy:   "System",
			CreateDate: &now,
		},
	})
}

func GetTopics(limit, offset int) ([]domain_models.TopicGetModel, error) {
	if limit <= 0 || limit > 100 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	entities, err := topic_repository.GetTopics(limit, offset)
	if err != nil {
		return nil, err
	}

	result := make([]domain_models.TopicGetModel, len(entities))
	for i, e := range entities {
		result[i] = domain_models.TopicGetModel{
			TopicID:    &e.TopicID,
			TopicTitle: e.TopicTitle,
			IsDeleted:  &e.IsDeleted,
		}
	}

	return result, nil
}
