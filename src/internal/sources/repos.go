package sources

import (
	"context"
	"log"
	psql "study/database/postgreSQL"
	"study/internal/config"
	"sync"

	author "study/internal/sources/author"
)

type Repo struct {
	Author author.AuthorRepo
}

var once sync.Once
var repo *Repo

func GetRepo() *Repo {
	once.Do(func() {
		cfg := config.GetConfig()
		postgreSQLClient, err := psql.NewClient(context.TODO(), 5, cfg.Storage)
		if err != nil {
			log.Printf("%v\n", err)
		}
		repo = &Repo{}
		repo.Author = author.NewRepository(postgreSQLClient)
	})
	return repo
}
