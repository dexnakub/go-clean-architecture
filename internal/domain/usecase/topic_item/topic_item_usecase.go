package topic_item_usecase

import (
	"fmt"
	data_entities "go-clean-achitech/internal/data/entities"
	data_helpers "go-clean-achitech/internal/data/helpers"
	topic_item_repository "go-clean-achitech/internal/data/repositories/topic_item"
	domain_models "go-clean-achitech/internal/domain/models"
	"strconv"
	"time"
)

// func CreateItems(topic domain_models.TopicItemCreateModel) error {
// 	if topic.TopicID == 0 {
// 		return fmt.Errorf("topic title is empty")
// 	}
// 	now := time.Now()

// 	return topic_item_repository.CreateItems(&data_entities.TopicItemEntity{
// 		TopicID: topic.TopicID,
// 		TopicItemTitle: topic.TopicItemTitle,
// 		BaseEntity: &data_entities.BaseEntity{
// 			CreateBy:   "System",
// 			CreateDate: &now,
// 		},
// 	})
// }

func CreateItem(topic domain_models.TopicItemCreateModel) error {
	if topic.TopicID == 0 {
		return fmt.Errorf("topic ID is required")
	}
	if topic.TopicItemTitle == "" {
		return fmt.Errorf("topic item title is required")
	}

	sequence, err := data_helpers.GetNextSequenceOFchildrenTable[data_entities.TopicItemEntity]("topic_id", topic.TopicID)
	if err != nil {
		return fmt.Errorf("failed to get next sequence: %v", err)
	}

	now := time.Now()

	return topic_item_repository.CreateItem(&data_entities.TopicItemEntity{
		TopicID:        topic.TopicID,
		TopicItemTitle: topic.TopicItemTitle,
		Sequence:       sequence,
		BaseEntity: &data_entities.BaseEntity{
			CreateBy:   "System",
			CreateDate: &now,
		},
	})
}

func GetItems(topicID string, limit, offset int) ([]domain_models.TopicItemGetModel, error) {
	if limit <= 0 || limit > 100 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	var result []domain_models.TopicItemGetModel

	if topicID != "" {
		items, err := topic_item_repository.GetItemByID(topicID, limit, offset)
		if err != nil {
			return nil, err
		}

		if len(items) == 0 {
			return []domain_models.TopicItemGetModel{}, nil
		}

		result = make([]domain_models.TopicItemGetModel, len(items))
		for i, item := range items {
			result[i] = domain_models.TopicItemGetModel{
				TopicItemID:    item.TopicItemID,
				TopicID:        item.TopicID,
				TopicItemTitle: item.TopicItemTitle,
				Sequence:       item.Sequence,
				IsDeleted:      &item.IsDeleted,
			}
		}
	}

	entities, err := topic_item_repository.GetItems(limit, offset)
	if err != nil {
		return nil, err
	}

	result = make([]domain_models.TopicItemGetModel, len(entities))
	for i, e := range entities {
		result[i] = domain_models.TopicItemGetModel{
			TopicItemID:    e.TopicItemID,
			TopicID:        e.TopicID,
			TopicItemTitle: e.TopicItemTitle,
			Sequence:       e.Sequence,
			IsDeleted:      &e.IsDeleted,
		}
	}

	return result, nil
}

func Updateitem(topic domain_models.TopicItemUpdateModel) error {
	if topic.TopicItemID == 0 {
		return fmt.Errorf("topic item ID is required")
	}

	if topic.TopicItemTitle == "" {
		return fmt.Errorf("topic item title is required")
	}

	now := time.Now()

	return topic_item_repository.UpdateItem(&data_entities.TopicItemEntity{
		TopicItemID:    topic.TopicItemID,
		TopicID:        topic.TopicID,
		TopicItemTitle: topic.TopicItemTitle,
		BaseEntity: &data_entities.BaseEntity{
			UpdateBy:   "System Update Item",
			UpdateDate: &now,
		},
	})
}

func UpdateDeleteStatus(topic domain_models.TopicItemUpdateDeleteStatusModel) error {
	if topic.TopicItemID == 0 {
		return fmt.Errorf("topic item ID is required")
	}

	now := time.Now()

	err := topic_item_repository.UpdateDeleteStatus(&data_entities.TopicItemEntity{
		TopicItemID: topic.TopicItemID,
		BaseEntity: &data_entities.BaseEntity{
			IsDeleted:  topic.IsDeleted,
			UpdateBy:   "System Delete Status",
			UpdateDate: &now,
		},
	})

	if err != nil {
		return err
	}

	err = data_helpers.ReorderTopicItemSequence[*data_entities.TopicItemEntity]("topic_id", topic.TopicID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateSequence(topics []domain_models.TopicItemUpdateSequenceModel) error {
	if len(topics) == 0 {
		return nil
	}

	dataEntities := make([]data_entities.TopicItemEntity, len(topics))

	now := time.Now()

	for index, topic := range topics {
		if topic.TopicItemID == 0 {
			return fmt.Errorf("topic item ID is required")
		}
		dataEntities[index] = data_entities.TopicItemEntity{
			TopicItemID: topic.TopicItemID,
			TopicID:     topic.TopicID,
			Sequence:    topic.Sequence,
			BaseEntity: &data_entities.BaseEntity{
				UpdateBy:   "System Update Sequence",
				UpdateDate: &now,
			},
		}
	}

	return topic_item_repository.UpdateSequence2(dataEntities)
}

func DeleteItem(topicID string, itemID string) error {
	if itemID == "" || topicID == "" {
		return fmt.Errorf("topic id and topic item id title is required")
	}

	itemIDInt, err := strconv.Atoi(itemID)
	if err != nil {
		return fmt.Errorf("invalid item id: %w", err)
	}

	err = topic_item_repository.DeleteTopicItemByTopicID(itemIDInt)
	if err != nil {
		return err
	}

	topicIDInt, err := strconv.Atoi(topicID)
	if err != nil {
		return fmt.Errorf("invalid topic id: %w", err)
	}

	err = data_helpers.ReorderTopicItemSequence[*data_entities.TopicItemEntity]("topic_id", topicIDInt)
	if err != nil {
		return err
	}
	return nil
}
