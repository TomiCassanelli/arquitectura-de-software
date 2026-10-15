package logstore

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	// El driver registra el nombre "mysql" que usa database/sql.
	_ "github.com/go-sql-driver/mysql"
)

// Store guarda los mensajes de la app en la tabla app_logs.
type Store struct{ db *sql.DB }

// Open usa la dirección interna de Compose y comprueba que MySQL responda.
func Open() (*Store, error) {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "demo:demo@tcp(localhost:3306)/observabilidad?parseTime=true"
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, err
	}
	return &Store{db: db}, nil
}

// Write inserta un mensaje con nivel info, warn o error.
// El tiempo límite evita que un problema de MySQL deje la petición esperando indefinidamente.
func (s *Store) Write(ctx context.Context, level, message string) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	// La tabla completa created_at automáticamente al insertar la fila.
	_, err := s.db.ExecContext(ctx, "INSERT INTO app_logs (level, message) VALUES (?, ?)", level, message)
	if err != nil {
		log.Printf("no se pudo guardar el log en MySQL: %v; nivel=%s mensaje=%s", err, level, message)
	}
}

// Close libera las conexiones a MySQL al terminar la aplicación.
func (s *Store) Close() error { return s.db.Close() }
