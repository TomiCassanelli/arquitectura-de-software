package logs

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Log contiene los datos que la app guarda por cada evento.
// MySQL completa id y created_at automáticamente.
type Log struct {
	TrackID string `gorm:"size:32;not null;default:'';index:idx_app_logs_track_id"`
	Stage   string `gorm:"size:20;not null;default:'api';index:idx_app_logs_stage"`
	Level   string `gorm:"size:10;not null"`
	Message string `gorm:"type:text;not null"`
}

// Store guarda los logs usando GORM.
type Store struct{ db *gorm.DB }

// Connect conecta GORM a MySQL. En Docker, MYSQL_DSN apunta al servicio mysql.
func Connect() (*Store, error) {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "demo:demo@tcp(localhost:3306)/observabilidad?parseTime=true"
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Agrega stage a una tabla ya existente sin borrar los logs anteriores.
	if err := db.Table("app_logs").AutoMigrate(&Log{}); err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

// Write guarda un log en la tabla que lee Grafana, con el track_id y la etapa cuando corresponda.
func (store *Store) Write(trackID, stage, level, message string) {
	err := store.db.Table("app_logs").Create(&Log{TrackID: trackID, Stage: stage, Level: level, Message: message}).Error
	if err != nil {
		log.Printf("no se pudo guardar el log en MySQL: %v", err)
	}
}
