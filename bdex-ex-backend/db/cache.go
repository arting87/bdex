// Copyright (c) 2014, go-xorm
// https://github.com/go-xorm/xorm-redis-cache

package db

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"hash/crc32"
	"reflect"
	"time"
	"unsafe"

	"github.com/garyburd/redigo/redis"
	"github.com/go-xorm/core"
)

const CACHE_FOREVER = time.Duration(0)

type RedisCacher struct {
	pool    *redis.Pool
	expires time.Duration
	logger  core.ILogger
}

func NewRedisCacher(pool *redis.Pool, expires time.Duration, logger core.ILogger) *RedisCacher {
	return &RedisCacher{pool: pool, expires: expires, logger: logger}
}

func exists(conn redis.Conn, key string) bool {
	existed, _ := redis.Bool(conn.Do("EXISTS", key))
	return existed
}

func (c *RedisCacher) getBeanKey(tableName string, id string) string {
	return fmt.Sprintf("xorm:bean:%s:%s", tableName, id)
}

func (c *RedisCacher) getSqlKey(tableName string, sql string) string {
	if sql == "*" {
		return fmt.Sprintf("xorm:sql:%s:%s", tableName, sql)
	}
	// hash sql to minimize key length
	crc := crc32.ChecksumIEEE([]byte(sql))
	return fmt.Sprintf("xorm:sql:%s:%d", tableName, crc)
}

// Delete all xorm cached objects
func (c *RedisCacher) Flush() error {
	return c.delObject("xorm:*")
}

func (c *RedisCacher) getObject(key string) interface{} {
	conn := c.pool.Get()
	defer conn.Close()
	raw, err := conn.Do("GET", key)
	if raw == nil {
		return nil
	}
	item, err := redis.Bytes(raw, err)
	if err != nil {
		c.logger.Errorf("redis.Bytes failed: %s", err)
		return nil
	}

	value, err := c.deserialize(item)

	return value
}

func (c *RedisCacher) GetIds(tableName, sql string) interface{} {
	sqlKey := c.getSqlKey(tableName, sql)
	c.logger.Debugf(" GetIds|tableName:%s|sql:%s|key:%s", tableName, sql, sqlKey)
	return c.getObject(sqlKey)
}

func (c *RedisCacher) GetBean(tableName string, id string) interface{} {
	beanKey := c.getBeanKey(tableName, id)
	c.logger.Debugf("[xorm/redis_cacher] GetBean|tableName:%s|id:%s|key:%s", tableName, id, beanKey)
	return c.getObject(beanKey)
}

func (c *RedisCacher) putObject(key string, value interface{}) error {
	b, err := c.serialize(value)
	if err != nil {
		return err
	}
	conn := c.pool.Get()
	defer conn.Close()
	if c.expires > 0 {
		_, err := conn.Do("SETEX", key, int32(c.expires/time.Second), b)
		return err
	} else {
		_, err := conn.Do("SET", key, b)
		return err
	}
}

func (c *RedisCacher) PutIds(tableName, sql string, ids interface{}) {
	sqlKey := c.getSqlKey(tableName, sql)
	c.logger.Debugf("PutIds|tableName:%s|sql:%s|key:%s|obj:%s|type:%v", tableName, sql, sqlKey, ids, reflect.TypeOf(ids))
	c.putObject(sqlKey, ids)
}

func (c *RedisCacher) PutBean(tableName string, id string, obj interface{}) {
	beanKey := c.getBeanKey(tableName, id)
	c.logger.Debugf("PutBean|tableName:%s|id:%s|key:%s|type:%v", tableName, id, beanKey, reflect.TypeOf(obj))
	c.putObject(beanKey, obj)
}

func (c *RedisCacher) delObject(key string) error {
	c.logger.Debugf("delObject key:[%s]", key)

	conn := c.pool.Get()
	defer conn.Close()
	if !exists(conn, key) {
		c.logger.Errorf("delObject key:[%s] err: %v", key, core.ErrCacheMiss)
		return core.ErrCacheMiss
	}
	_, err := conn.Do("DEL", key)
	return err
}

func (c *RedisCacher) delObjects(key string) error {

	c.logger.Debugf("delObjects key:[%s]", key)

	conn := c.pool.Get()
	defer conn.Close()

	keys, err := conn.Do("KEYS", key)
	c.logger.Debugf("delObjects keys: %v", keys)

	if err == nil {
		for _, key := range keys.([]interface{}) {
			conn.Do("DEL", key)
		}
	}
	return err
}

func (c *RedisCacher) DelIds(tableName, sql string) {
	c.delObject(c.getSqlKey(tableName, sql))
}

func (c *RedisCacher) DelBean(tableName string, id string) {
	c.delObject(c.getBeanKey(tableName, id))
}

func (c *RedisCacher) ClearIds(tableName string) {
	c.delObjects(c.getSqlKey(tableName, "*"))
}

func (c *RedisCacher) ClearBeans(tableName string) {
	c.delObjects(c.getBeanKey(tableName, "*"))
}

func (c *RedisCacher) serialize(value interface{}) ([]byte, error) {

	err := c.registerGobConcreteType(value)
	if err != nil {
		return nil, err
	}

	if reflect.TypeOf(value).Kind() == reflect.Struct {
		return nil, fmt.Errorf("serialize func only take pointer of a struct")
	}

	var b bytes.Buffer
	encoder := gob.NewEncoder(&b)

	c.logger.Debugf("serialize type:%v", reflect.TypeOf(value))
	err = encoder.Encode(&value)
	if err != nil {
		c.logger.Errorf("gob encoding '%s' failed: %s|value:%v", value, err, value)
		return nil, err
	}
	return b.Bytes(), nil
}

func (c *RedisCacher) deserialize(byt []byte) (ptr interface{}, err error) {
	b := bytes.NewBuffer(byt)
	decoder := gob.NewDecoder(b)

	var p interface{}
	err = decoder.Decode(&p)
	if err != nil {
		c.logger.Errorf("decode failed: %v", err)
		return
	}

	v := reflect.ValueOf(p)
	c.logger.Debugf("deserialize type:%v", v.Type())
	if v.Kind() == reflect.Struct {

		var pp interface{} = &p
		datas := reflect.ValueOf(pp).Elem().InterfaceData()

		sp := reflect.NewAt(v.Type(),
			unsafe.Pointer(datas[1])).Interface()
		ptr = sp
		vv := reflect.ValueOf(ptr)
		c.logger.Debugf("deserialize convert ptr type:%v | CanAddr:%t", vv.Type(), vv.CanAddr())
	} else {
		ptr = p
	}
	return
}

func (c *RedisCacher) registerGobConcreteType(value interface{}) error {

	t := reflect.TypeOf(value)

	c.logger.Debugf("registerGobConcreteType:%v", t)

	switch t.Kind() {
	case reflect.Ptr:
		v := reflect.ValueOf(value)
		i := v.Elem().Interface()
		gob.Register(&i)
	case reflect.Struct, reflect.Map, reflect.Slice:
		gob.Register(value)
	case reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Bool, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		// do nothing since already registered known type
	default:
		return fmt.Errorf("unhandled type: %v", t)
	}
	return nil
}
