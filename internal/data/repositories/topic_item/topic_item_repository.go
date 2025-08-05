package topic_item_repository

import (
	"fmt"
	data_entities "go-clean-achitech/internal/data/entities"
	data_helpers "go-clean-achitech/internal/data/helpers"
	"strings"
)

func CreateItem(data *data_entities.TopicItemEntity) error {
	if data == nil {
		return fmt.Errorf("topic item entity is nil")
	}

	gorm := data_helpers.GetGormAdapter()

	var count int64
	err := gorm.Model(&data_entities.TopicEntity{}).
		Where("topic_id = ?", data.TopicID).
		Count(&count).Error
	if err != nil {
		return fmt.Errorf("failed to check topic existence: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("topic_id %d not found", data.TopicID)
	}

	return gorm.Create(data).Error
}

func GetItems(limit, offset int) ([]data_entities.TopicItemEntity, error) {
	gorm := data_helpers.GetGormAdapter()

	items := []data_entities.TopicItemEntity{}
	err := gorm.Model(&data_entities.TopicItemEntity{}).
		Where("is_deleted = ?", false).
		Order("sequence ASC").
		Limit(limit).
		Offset(offset).
		Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func GetItemByID(topicID string, limit, offset int) ([]data_entities.TopicItemEntity, error) {
	gorm := data_helpers.GetGormAdapter()

	items := []data_entities.TopicItemEntity{}
	err := gorm.Model(&data_entities.TopicItemEntity{}).
		Where("is_deleted = ? and topic_id = ? ", false, 1).
		Order("sequence ASC").
		Limit(limit).
		Offset(offset).
		Preload("Topic").
		Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}

func UpdateItem(data *data_entities.TopicItemEntity) error {
	if data == nil {
		return fmt.Errorf("topic item entity is nil")
	}

	gorm := data_helpers.GetGormAdapter()

	var existing data_entities.TopicItemEntity
	if err := gorm.First(&existing, "topic_item_id = ?", data.TopicItemID).Error; err != nil {
		return fmt.Errorf("topic item not found: %w", err)
	}

	return gorm.Model(&existing).Updates(map[string]interface{}{
		"topic_item_title": data.TopicItemTitle,
		"update_by":        data.BaseEntity.UpdateBy,
		"update_date":      data.BaseEntity.UpdateDate,
	}).Error
}

func UpdateDeleteStatus(data *data_entities.TopicItemEntity) error {
	if data == nil {
		return fmt.Errorf("topic item entity is nil")
	}

	gorm := data_helpers.GetGormAdapter()

	var existing data_entities.TopicItemEntity
	if err := gorm.First(&existing, "topic_item_id = ?", data.TopicItemID).Error; err != nil {
		return fmt.Errorf("topic item not found: %w", err)
	}

	return gorm.Model(&existing).Updates(map[string]interface{}{
		"is_deleted":  data.BaseEntity.IsDeleted,
		"update_by":   data.BaseEntity.UpdateBy,
		"update_date": data.BaseEntity.UpdateDate,
	}).Error
}

func UpdateSequence(data *data_entities.TopicItemEntity) error {
	if data == nil {
		return fmt.Errorf("topic item entity is nil")
	}

	gorm := data_helpers.GetGormAdapter()

	var existing data_entities.TopicItemEntity
	if err := gorm.First(&existing, "topic_item_id = ?", data.TopicItemID).Error; err != nil {
		return fmt.Errorf("topic item not found: %w", err)
	}

	return gorm.Model(&existing).Updates(map[string]interface{}{
		"is_deleted":  data.BaseEntity.IsDeleted,
		"update_by":   data.BaseEntity.UpdateBy,
		"update_date": data.BaseEntity.UpdateDate,
	}).Error
}

func UpdateSequence2(topics []data_entities.TopicItemEntity) error {
	if len(topics) == 0 {
		return nil
	}

	gorm := data_helpers.GetGormAdapter()

	topicID := topics[0].TopicID

	caseSequence := "CASE topic_item_id "
	caseUpdateBy := "CASE topic_item_id "
	caseUpdateDate := "CASE topic_item_id "

	ids := make([]interface{}, 0, len(topics))

	for _, t := range topics {
		if t.TopicItemID == 0 {
			return fmt.Errorf("topic item ID is required")
		}

		caseSequence += fmt.Sprintf("WHEN %d THEN %d ", t.TopicItemID, t.Sequence)
		caseUpdateBy += fmt.Sprintf("WHEN %d THEN '%s' ", t.TopicItemID, t.BaseEntity.UpdateBy)
		caseUpdateDate += fmt.Sprintf("WHEN %d THEN '%s' ", t.TopicItemID, t.BaseEntity.UpdateDate.Format("2006-01-02 15:04:05"))

		ids = append(ids, t.TopicItemID)
	}

	caseSequence += "ELSE sequence END"
	caseUpdateBy += "ELSE update_by END"
	caseUpdateDate += "ELSE update_date END"

	var idStrs []string
	for _, id := range ids {
		idStrs = append(idStrs, fmt.Sprintf("%v", id))
	}
	inClause := strings.Join(idStrs, ",")

	query := fmt.Sprintf(`
        UPDATE topic_items
        SET sequence = %s,
            update_by = %s,
            update_date = %s
        WHERE topic_item_id IN (%s)
        AND topic_id = %d
    `, caseSequence, caseUpdateBy, caseUpdateDate, inClause, topicID)

	return gorm.Exec(query).Error
}

func DeleteTopicItemByTopicID(itemID int) error {
	gorm := data_helpers.GetGormAdapter()

	var item data_entities.TopicItemEntity
	if err := gorm.First(&item, itemID).Error; err != nil {
		return fmt.Errorf("topic item not found")
	}

	if err := gorm.Model(&data_entities.TopicItemEntity{}).
		Where("topic_item_id = ?", itemID).
		Update("is_deleted", true).Error; err != nil {
		return fmt.Errorf("failed to delete topic item: %w", err)
	}
	return nil
}
