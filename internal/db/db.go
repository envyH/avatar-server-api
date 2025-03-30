package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func InitDB(databaseUrl string) {
	var err error
	pool, err = pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalf("Không thể kết nối đến database: %v", err)
	}
	log.Println("✅ Kết nối database thành công")
}

func CloseDB() {
	pool.Close()
}

func GetDB() *pgxpool.Pool {
	return pool
}
