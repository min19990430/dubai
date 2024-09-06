package bigcache

import (
	"bytes"
	"context"
	"encoding/gob"
	"log"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/spf13/viper"
)

type Option struct {
	LifeWindow time.Duration
}

func NewOption(conf *viper.Viper) Option {
	return Option{
		LifeWindow: conf.GetDuration("bigcache.life_window"),
	}
}

func NewBigcache(bigcacheOption Option) *bigcache.BigCache {
	bigcache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(bigcacheOption.LifeWindow))
	if err != nil {
		log.Fatalf("bigcache error: %v", err)
	}

	return bigcache
}

func Set(c *bigcache.BigCache, key string, value interface{}) error {
	// Serialize the value into bytes
	valueBytes, err := serialize(value)
	if err != nil {
		return err
	}

	return c.Set(key, valueBytes)
}

func Get(c *bigcache.BigCache, key string) (interface{}, error) {
	valueBytes, err := c.Get(key)
	if err != nil {
		return nil, err
	}
	// Deserialize the bytes of the value
	value, err := deserialize(valueBytes)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func serialize(value interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	gob.Register(value)

	err := enc.Encode(&value)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func deserialize(valueBytes []byte) (interface{}, error) {
	var value interface{}
	buf := bytes.NewBuffer(valueBytes)
	dec := gob.NewDecoder(buf)

	err := dec.Decode(&value)
	if err != nil {
		return nil, err
	}

	return value, nil
}
