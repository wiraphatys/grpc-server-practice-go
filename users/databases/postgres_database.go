package databases

import (
	"fmt"
	"user-services/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	Db *gorm.DB
}

func NewPostgresDatabase(cfg *config.Config) Database {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Db.Host,
		cfg.Db.User,
		cfg.Db.Password,
		cfg.Db.Name,
		cfg.Db.Port,
		cfg.Db.SSLMode,
		cfg.Db.Timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		panic("failed to connect database")
	}

	return &postgresDatabase{
		Db: db,
	}
}

func (p *postgresDatabase) GetDb() *gorm.DB {
	return p.Db
}
