package logstore

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
	Level   string `gorm:"size:10;not null"`
	Message string `gorm:"type:text;not null"`
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
	// Agrega track_id a una tabla ya existente sin borrar los logs anteriores.
	if err := db.Table("app_logs").AutoMigrate(&Log{}); err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

// Write guarda un log. Table indica el nombre de la tabla que lee Grafana.
func (s *Store) Write(level, message string) {
	s.WriteTrack("", level, message)
}

// WriteTrack guarda un paso del recorrido con el ID recibido en la petición.
func (s *Store) WriteTrack(trackID, level, message string) {
	err := s.db.Table("app_logs").Create(&Log{TrackID: trackID, Level: level, Message: message}).Error
	if err != nil {
		log.Printf("no se pudo guardar el log en MySQL: %v", err)
	}
}
