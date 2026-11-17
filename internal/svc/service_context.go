// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package svc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"example.com/iam-service/internal/config"

	"github.com/go-sql-driver/mysql"
)

type ServiceContext struct {
	Config config.Config
	DB     *sql.DB
}

func NewServiceContext(ctx context.Context, c config.Config) (*ServiceContext, error) {
	service := &ServiceContext{Config: c}
	if c.MySQL == nil {
		return service, nil
	}

	password := os.Getenv("IAM_MYSQL_PASSWORD")
	if password == "" {
		return nil, errors.New("MySQL requires IAM_MYSQL_PASSWORD")
	}

	cfg := mysql.NewConfig()
	cfg.User = c.MySQL.User
	cfg.Passwd = password
	cfg.Net = "tcp"
	cfg.Addr = net.JoinHostPort(c.MySQL.Host, strconv.Itoa(c.MySQL.Port))
	cfg.DBName = c.MySQL.Database
	cfg.ParseTime = true
	cfg.TLSConfig = "true"
	if c.MySQL.TLSSkipVerify {
		cfg.TLSConfig = "skip-verify"
	}

	connector, err := mysql.NewConnector(cfg)
	if err != nil {
		return nil, fmt.Errorf("configure MySQL connector: %w", err)
	}

	db := sql.OpenDB(connector)
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := db.PingContext(pingCtx); err != nil {
		db.Close()
		return nil, fmt.Errorf("connect to MySQL: %w", err)
	}

	service.DB = db
	return service, nil
}
