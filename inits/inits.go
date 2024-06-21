package inits

import (
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"role-catalog-go/model"
)

var db *gorm.DB

func LoadConfig(config string, viper *viper.Viper) {
	// Set the file name of the configurations file
	viper.SetConfigName(config)

	// Set the path to look for the configurations file
	viper.AddConfigPath("./config")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Error reading config file", err)
	}

}

func LoadDBConfig() {
	db, err := gorm.Open(sqlite.Open("roleCatalog.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Unable to load DB ", err)
	}
	log.Fatalln(db.AutoMigrate(&model.RoleCatalogTable{}, &model.RoleByLobTable{}))
}

func GetDBInstance() *gorm.DB {
	return db
}
