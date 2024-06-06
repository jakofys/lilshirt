package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type LinkRepository interface {
	Save(ctx context.Context, link Link) (*Link, error)
	Link(ctx context.Context, id string) (*Link, error)
}

type linkRepository struct {
	db *sql.DB
}

type Link struct {
	ID          uuid.UUID
	Long, Short string
	ExpiredAt   time.Time
	CreatedAt   time.Time
}

func NewLinkRepository(db *sql.DB) (LinkRepository, error) {
	if err := db.PingContext(context.TODO()); err != nil {
		return nil, err
	}
	return &linkRepository{db: db}, nil
}

func (r *linkRepository) Save(ctx context.Context, l Link) (*Link, error) {
	_, err := r.db.
		ExecContext(ctx, `INSERT INTO links(long,short,expiredAt) VALUES (?,?,?) ON CONFLICT DO UPDATE SET links.expiredAt=?`, l.Long, l.Short, l.ExpiredAt, l.ExpiredAt)
	if err != nil {
		return nil, err
	}
	err = r.db.
		QueryRowContext(ctx, `SELECT id,long,short,expiredAt,createdAt FROM links WHERE long = ?`, l.Long).
		Scan(&l.ID, &l.Long, &l.Short, &l.ExpiredAt, &l.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &l, nil
}

func (r *linkRepository) Link(ctx context.Context, id string) (*Link, error) {
	l := &Link{}
	err := r.db.
		QueryRowContext(ctx, `SELECT id,long,short,createdAt,expiredAt FROM links WHERE id=? AND expiredAt>?`, id, time.Now()).
		Scan(&l.ID, &l.Long, &l.Short, &l.ExpiredAt, &l.CreatedAt)
	if err != nil {
		return nil, err
	}
	return l, nil
}
