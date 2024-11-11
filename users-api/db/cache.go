// users-api/db/cache.go
package db

import (
	"log"

	"github.com/bradfitz/gomemcache/memcache"
)

var Cache *memcache.Client

func ConnectCache() {
	Cache = memcache.New("localhost:11211")
	err := Cache.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar a Memcached:", err)
	}

	log.Println("Conectado a Memcached")
}
