package main

import (
	"log"
	"os"

	"github.com/go-redis/store"
)

type LoggingMiddleware struct {
	inner  store.Storer
	logger *log.Logger
}

func NewLoggerMiddleware(s store.Storer) *LoggingMiddleware {
	return &LoggingMiddleware{
		inner:  s,
		logger: log.New(os.Stdout, "[log]", log.Ltime|log.Lmicroseconds),
	}
}

func (l *LoggingMiddleware) Get(key string) (string, error) {
	l.logger.Printf("GET %q", key)
	val, err := l.inner.Get(key)
	if err != nil {
		l.logger.Printf("GET %q -> miss", key)
	} else {
		l.logger.Printf("get %q -> hit", key)
	}
	return val, err
}

func (l *LoggingMiddleware) Set(key, value string) error {
	l.logger.Printf("SET %q", key)
	err := l.inner.Set(key, value)
	return err
}

func (l *LoggingMiddleware) Keys() []string {
	return l.inner.Keys()
}

func (l *LoggingMiddleware) Del(key string) {
	l.logger.Printf("DEL %q", key)
	l.inner.Del(key)
}

func (l *LoggingMiddleware) Len() int {

	got := l.inner.Len()
	l.logger.Printf("Len %q", got)
	return got
}
