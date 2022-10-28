package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/golang_api_crud_sample/entity"
)

type database struct {
	connection *gorm.DB
}

func NewConnection(db_host, db_user, db_password, db_name, db_port string) *gorm.DB {
	dsn := "host=" + db_host + " user=" + db_user + " password=" + db_password + " dbname=" + db_name + " port=" + db_port

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	  }), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entity.User{}, &entity.Patient{})
	return db
}