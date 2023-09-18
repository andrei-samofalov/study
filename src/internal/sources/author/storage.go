package sources

import (
	"context"
)

type AuthorRepo interface {
	Create(ctx context.Context, author *Author) error
	FindAll(ctx context.Context) (a []Author, err error)
	FindOne(ctx context.Context, id int) (Author, error)
	Update(ctx context.Context, author Author) error
	Delete(ctx context.Context, id int) error
}
