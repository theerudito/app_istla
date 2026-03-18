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

		host := os.Getenv("ServerDB")
		port := os.Getenv("PortDB")
		user := os.Getenv("UserDB")
		pass := os.Getenv("PasswordBD")
		name := os.Getenv("NameDB")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, pass, name)

		conn, connErr := sql.Open("postgres", dsn)
		if connErr != nil {
			err = fmt.Errorf("error al abrir sql.Open: %w", connErr)
			return
		}

		conn.SetConnMaxLifetime(30 * time.Minute)
		conn.SetMaxOpenConns(10)
		conn.SetMaxIdleConns(5)

		if pingErr := retryPing(conn); pingErr != nil {
			err = pingErr
			conn.Close()
			return
		}

		instance = &DB{conn: conn}
	})

	if err != nil {
		return nil, err
	}

	return instance, nil
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
