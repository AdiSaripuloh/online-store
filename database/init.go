package database

import "github.com/AdiSaripuloh/online-store/config"

func InitDatabase(migration, seed *bool) {
	dbConfig := config.BuildDbConfig()
	Connect(dbConfig)
	if *migration {
		Migration()
	}
	if *seed {
		Seed()
	}
}
