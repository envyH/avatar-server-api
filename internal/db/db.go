package db

import (
	"context"
	"log"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func InitDB(databaseUrl string) {
	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		log.Fatalf("❌ Lỗi khi parse database URL: %v", err)
	}

	config.ConnConfig.DialFunc = func(ctx context.Context, network, addr string) (net.Conn, error) {
		// Chỉ dùng IPv4
		return net.Dial("tcp4", addr)
	}

	pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("❌ Không thể kết nối đến database: %v", err)
	}
	log.Println("✅ Kết nối database thành công")
}

func CloseDB() {
	pool.Close()
}

func GetDB() *pgxpool.Pool {
	return pool
}
