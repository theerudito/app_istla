package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	conn *sql.DB
}

var (
	instance *DB
	once     sync.Once
)

func ConnectarDB() (*DB, error) {

	var err error

	once.Do(func() {

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("ServerDB"), os.Getenv("PortDB"), os.Getenv("UserDB"), os.Getenv("PasswordBD"), os.Getenv("NameDB"))

		conn, connErr := sql.Open("postgres", dsn)

		if connErr != nil {
			err = fmt.Errorf("%s: %w", "Error conectar a la db", connErr)
			return
		}

		conn.SetConnMaxLifetime(30 * time.Minute)
		conn.SetMaxOpenConns(10)
		conn.SetMaxIdleConns(5)

		if err = retryPing(conn); err != nil {
			err = fmt.Errorf("%s: %w", "Error al cerrar la connecion", err)
			return
		}

		instance = &DB{conn: conn}
	})

	return instance, err
}

func retryPing(conn *sql.DB) error {
	const maxRetries = 3
	const delay = 2 * time.Second

	for i := 0; i < maxRetries; i++ {
		if err := conn.Ping(); err == nil {
			return nil
		}
		log.Printf("la coneccion fallo (intentos %d/%d), reintentando...", i+1, maxRetries)
		time.Sleep(delay)
	}
	return fmt.Errorf("no es posible estableser una connecion con la base de datos despues de  %d intentos", maxRetries)
}

func (d *DB) GetConn() *sql.DB {
	return d.conn
}

func (d *DB) Close() {
	if d.conn != nil {
		if err := d.conn.Close(); err != nil {
			log.Printf("%s: %v", "error verificando la base de datos", err)
		}
	}
}
