package db

import (
	"log"
	"os"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

var Cache *memcache.Client

func ConnectCache() {
	memcachedHost := os.Getenv("MEMCACHED_HOST")
	memcachedPort := os.Getenv("MEMCACHED_PORT")
	memcachedAddress := memcachedHost + ":" + memcachedPort

	Cache = memcache.New(memcachedAddress)
	err := Cache.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar a Memcached:", err)
	}

	log.Println("Conectado a Memcached")
}

// Función para obtener un valor de la caché
func GetCache(key string) (string, error) {
	item, err := Cache.Get(key)
	if err != nil {
		return "", err
	}
	return string(item.Value), nil
}

// Función para guardar un valor en la caché
func SetCache(key string, value string, duration time.Duration) error {
	return Cache.Set(&memcache.Item{Key: key, Value: []byte(value), Expiration: int32(duration.Seconds())})
}
