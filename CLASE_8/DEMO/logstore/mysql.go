package logstore

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Log contiene los dos datos que la app guarda por cada evento.
// MySQL completa id y created_at automáticamente.
type Log struct {
	Level   string
	Message string
}

// Store guarda los logs usando GORM.
type Store struct{ db *gorm.DB }

// Open conecta GORM a MySQL. En Docker, MYSQL_DSN apunta al servicio mysql.
func Open() (*Store, error) {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "demo:demo@tcp(localhost:3306)/observabilidad?parseTime=true"
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

// Write guarda un log. Table indica el nombre de la tabla que lee Grafana.
func (s *Store) Write(level, message string) {
	err := s.db.Table("app_logs").Create(&Log{Level: level, Message: message}).Error
	if err != nil {
		log.Printf("no se pudo guardar el log en MySQL: %v", err)
	}
}
