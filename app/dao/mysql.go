package dao

import (
	"nes/app/model"
	"nes/config"
	"nes/pkg/database"
	"nes/pkg/log"
)

type DBClient struct {
	db *database.DBEngine
}

func (cli *DBClient) FindAll(models interface{}) error {
	tx := cli.db.DB
	err := tx.Find(models).Error
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (cli *DBClient) Save(entity model.Entity) {
	tx := cli.db.DB
	id := entity.GetId()
	if id == "" {
		tx.Create(entity)
	}
	err := tx.Find(&entity, id).Error
	if err != nil {
		tx.Create(entity)
		return
	}
	tx.Save(entity)
}

func NewDbClient(opts *config.Database) *DBClient {
	db := database.NewDbEngine(opts)

	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8")
	registerModel(db)
	return &DBClient{db: db}

}

func registerModel(db *database.DBEngine) {
	err := db.AutoMigrate(
		&model.Customer{},
	)
	if err != nil {
		panic(err)
	}
}
