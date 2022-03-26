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

func (cli *DBClient) List(models interface{}) error {
	tx := cli.db.DB

	err := tx.Find(models).Error
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (cli *DBClient) Create(model interface{}) {
	tx := cli.db.DB
	tx.Create(model)
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
