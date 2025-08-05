package data_migrations

import (
	"go-clean-achitech/internal/adapters"
	data_entities "go-clean-achitech/internal/data/entities"
	data_helpers "go-clean-achitech/internal/data/helpers"
)

func StartMigration() error {
	gorm, err := adapters.NewGormAdapter()
	if err != nil {
		return err
	}

	if err := data_helpers.MigrateIfNotExists(gorm, "topics", &data_entities.TopicEntity{}); err != nil {
		return err
	}

	if err := data_helpers.MigrateIfNotExists(gorm, "topic_items", &data_entities.TopicItemEntity{}); err != nil {
		return err
	}

	return nil
}
