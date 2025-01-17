package model

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/goextension/log"

	"github.com/godcong/fate/config"
	"github.com/godcong/fate/ent"
)

const dsn = "%s:%s@tcp(%s)/%s?loc=%s&charset=utf8mb4&parseTime=true"

type Model struct {
	*ent.Client
	cfg config.DBConfig
}

func Open(cfg config.DBConfig, debug bool) (*Model, error) {
	var options []ent.Option
	if debug {
		options = append(options, ent.Debug())
	}

	if cfg.Log != "" {
		options = append(options, ent.Log(func(i ...interface{}) {
			log.Debug(i...)
		}))
	}

	open, err := ent.Open(cfg.Driver, cfg.DSN, options...)
	if err != nil {
		return nil, err
	}
	return &Model{
		Client: open,
		cfg:    cfg,
	}, nil
}

// ID ...
func ID(name string) string {
	s := sha256.New()
	return hex.EncodeToString(s.Sum([]byte(name)))
}
