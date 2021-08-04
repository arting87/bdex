package db

import (
	"bdex.in/bdex/bdex-ex-backend/bdexerrors"

	"github.com/GyhzzZ/xorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/rs/xid"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type DB struct {
	*xorm.Engine

	MaxQueryLimit  uint32
	MaxQueryOffset uint32
}

func NewDB(dsn string) *DB {
	cfg := viper.Sub("db")

	var err error
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	engine.NewSession()

	engine.SetMapper(core.SameMapper{})
	engine.SetMaxIdleConns(cfg.GetInt("maxIdleConns"))
	engine.SetMaxOpenConns(cfg.GetInt("maxOpenConns"))
	engine.ShowSQL(cfg.GetBool("showSQL"))
	engine.Logger().SetLevel(core.LogLevel(cfg.GetInt("logLevel")))

	if Redis != nil {
		cacheExpiration, err := cast.ToDurationE(cfg.GetString("cacheExpiration"))
		if err != nil {
			panic(err.Error())
		}

		// Redis server may be configured with --maxmemory <size> and --maxmemory-policy volatile-lru
		// Refer to https://redis.io/topics/lru-cache
		engine.SetDefaultCacher(NewRedisCacher(Redis, cacheExpiration, engine.Logger()))
	}

	maxQueryLimit := uint32(cfg.GetInt64("maxQueryLimit"))
	maxQueryOffset := uint32(cfg.GetInt64("maxQueryOffset"))
	if maxQueryLimit <= 0 && maxQueryOffset <= 0 {
		panic("max db query limit and offset must be greater than 0")
	}

	return &DB{
		Engine:         engine,
		MaxQueryLimit:  maxQueryLimit,
		MaxQueryOffset: maxQueryOffset,
	}
}

func (db *DB) NewID() string {
	return xid.New().String()
}

type Session struct {
	*xorm.Session
}

func (s *Session) Commit(noPanic ...bool) error {
	err := s.Session.Commit()
	if err != nil {
		if len(noPanic) > 0 && noPanic[0] {
			return err
		}
		panic(err.Error())
	}
	return nil
}

func (s *Session) Close() {
	s.Session.Close()
}

func (db *DB) Session() *Session {
	session := db.NewSession()
	err := session.Begin()
	if err != nil {
		panic(err.Error())
	}
	return &Session{session}
}

func (db *DB) Upsert(id interface{}, bean interface{}) (bool, error) {
	has, err := db.ID(id).Exist()
	if err != nil {
		return false, err
	}
	var n int64
	if has {
		n, err = db.ID(id).Update(bean)
	} else {
		n, err = db.Insert(bean)
	}
	return n > 0, err
}

func (db *DB) InsertIfNotExists(id interface{}, bean interface{}) (bool, error) {
	has, err := db.ID(id).Exist()
	if err != nil {
		return false, err
	}
	var n int64
	if !has {
		n, err = db.Insert(bean)
	}
	return n > 0, err
}

func (db *DB) Limit(session *xorm.Session, limit, offset uint32) (*xorm.Session, error) {
	if limit > db.MaxQueryLimit {
		return nil, bdexerrors.ErrInvalidArg
	} else if limit == 0 {
		limit = db.MaxQueryLimit
	}
	if offset > db.MaxQueryOffset {
		return nil, bdexerrors.ErrInvalidArg
	}
	return session.Limit(int(limit), int(offset)), nil
}
