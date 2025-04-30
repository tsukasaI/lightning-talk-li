package cache

import (
	"errors"
	"time"
)

type (
	StringValue struct {
		Value     string
		ExpiresAt time.Time
	}
	StringStore map[string]StringValue
)

var NotFoundErr = errors.New("not found or expired")

func NewStringStore() StringStore {
	return make(StringStore)
}

func (s *StringStore) Save(key, value string, ttl time.Duration) {
	(*s)[key] = StringValue{
		Value:     value,
		ExpiresAt: time.Now().Add(ttl),
	}
}

func (s *StringStore) Get(key string) (string, error) {
	v, ok := (*s)[key]
	if !ok {
		return "", NotFoundErr
	}
	if v.ExpiresAt.After(time.Now()) {
		return v.Value, nil
	}
	return "", NotFoundErr
}
