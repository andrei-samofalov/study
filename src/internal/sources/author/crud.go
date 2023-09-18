package sources

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"log"
	"study/database/postgreSQL"
)

type repository struct {
	client postgresql.Client
}

func (r *repository) Create(ctx context.Context, author *Author) error {
	q := `
		INSERT INTO author 
		    (name) 
		VALUES 
		       ($1) 
		RETURNING id
	`

	if err := r.client.QueryRow(ctx, q, author.Name).Scan(&author.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			log.Println(newErr)
			return newErr
		}
		return err
	}

	return nil
}

func (r *repository) FindAll(ctx context.Context) (authors []Author, err error) {
	q := `
		SELECT id, name FROM public.author;
	`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	authors = make([]Author, 0)

	for rows.Next() {
		var ath Author

		err = rows.Scan(&ath.ID, &ath.Name)
		if err != nil {
			return nil, err
		}

		authors = append(authors, ath)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

func (r *repository) FindOne(ctx context.Context, id int) (Author, error) {
	q := `
		SELECT id, name FROM public.author WHERE id = $1
	`

	var ath Author
	err := r.client.QueryRow(ctx, q, id).Scan(&ath.ID, &ath.Name)
	if err != nil {
		return Author{}, err
	}

	return ath, nil
}

func (r *repository) Update(ctx context.Context, author Author) error {
	q := `
		UPDATE author 
		SET name = $1, is_deleted = $2
		WHERE id = $3
	`
	cmd, err := r.client.Exec(ctx, q, &author.Name, &author.IsDeleted, &author.ID)
	if err != nil {
		log.Println(err)
	}
	cmd.Update()

	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	q := `
		UPDATE author SET is_deleted = true WHERE id = $1
	`
	cmd, err := r.client.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	cmd.Update()

	return nil
}

func NewRepository(client postgresql.Client) AuthorRepo {
	return &repository{
		client: client,
	}
}
