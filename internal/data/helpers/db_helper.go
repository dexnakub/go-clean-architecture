package data_helpers

import (
	"fmt"
	"go-clean-achitech/internal/adapters"
)

func MigrateIfNotExists(gorm *adapters.GormAdapter, tableName string, model interface{}) error {
	exists, err := TableExists(gorm, tableName)
	if err != nil {
		return err
	}
	if !exists {
		if err := gorm.AutoMigrate(model); err != nil {
			return err
		}
	}
	return nil
}

func TableExists(db *adapters.GormAdapter, tableName string) (bool, error) {
	var count int64
	err := db.Raw("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = ?", tableName).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func GetGormAdapter() *adapters.GormAdapter {
	gorm, err := adapters.NewGormAdapter()

	if err != nil {
		panic(err)
	}

	return gorm
}

func GetNextSequenceOFchildrenTable[Table any](primaryKey string, parentID int) (int, error) {
	var maxSequence int
	var model Table
	gorm := GetGormAdapter()

	err := gorm.Model(model).
		Where(primaryKey+"= ?", parentID).
		Select("COALESCE(MAX(sequence), 0)").
		Scan(&maxSequence).Error

	if err != nil {
		return 0, err
	}
	return maxSequence + 1, nil
}

type Sequencable interface {
	GetID() int
	GetSequence() int
	SetSequence(int)
	GetIDColumnName() string
}

func ReorderTopicItemSequence[T Sequencable](primaryKey string, parentID int) error {
	var model T
	gorm := GetGormAdapter()

	var items []T
	if err := gorm.Model(&model).
		Where(primaryKey+" = ? AND is_deleted = ?", parentID, false).
		Order("sequence ASC").
		Find(&items).Error; err != nil {
		return fmt.Errorf("failed to fetch items: %w", err)
	}

	for i := range items {
		newSeq := i + 1

		if items[i].GetSequence() != newSeq {
			items[i].SetSequence(newSeq)

			err := gorm.Model(&model).
				Where(items[i].GetIDColumnName()+" = ?", items[i].GetID()).
				Update("sequence", newSeq).Error
			if err != nil {
				return fmt.Errorf("failed to update sequence for item ID %d: %w", items[i].GetID(), err)
			}
		}
	}

	return nil
}
