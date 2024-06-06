package services

import (
	"context"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/jakofys/fluid/internal/pkg/random"
	"github.com/jakofys/fluid/internal/repositories"
)

const MaxSlugSize = 6

type LinkService struct {
	rand     random.StringRandom
	linkRepo repositories.LinkRepository
}

type Link struct {
	ID          uuid.UUID
	Long, Short *url.URL
	ExpiredAt   time.Time
	CreatedAt   time.Time
}

func buildShort(domain, slug string) string {
	return "https://fluid.com/" + domain + "/" + slug
}

func NewLinkService(linkRepo repositories.LinkRepository) *LinkService {
	return &LinkService{
		linkRepo: linkRepo,
		rand:     random.NewStringRandom(),
	}
}

func (s *LinkService) GetLink(ctx context.Context, id uuid.UUID) (*Link, error) {
	link, err := s.linkRepo.Link(ctx, id.String())
	if err != nil {
		return nil, ErrGetLink(err.Error())
	}
	ulong, err := url.Parse(link.Long)
	if err != nil {
		return nil, ErrInvalidURL(err.Error())
	}
	ushort, err := url.Parse(link.Short)
	if err != nil {
		return nil, ErrInvalidURL(err.Error())
	}
	return &Link{
		ID:        link.ID,
		Long:      ulong,
		Short:     ushort,
		CreatedAt: link.CreatedAt,
		ExpiredAt: link.ExpiredAt,
	}, nil
}

func (s *LinkService) GenerateLink(ctx context.Context, long string) (*Link, error) {
	ulong, err := url.Parse(long)
	if err != nil {
		return nil, ErrInvalidURL(err.Error())
	}

	now := time.Now()
	expirationDate := time.Date(now.Year(), now.Month()+1, now.Day(), 0, 0, 0, 0, now.Location())
	slug := s.rand.StringN(MaxSlugSize)
	ushort, err := url.Parse(buildShort(ulong.Hostname(), slug))
	if err != nil {
		return nil, ErrInvalidURL(err.Error())
	}

	link, err := s.linkRepo.Save(ctx, repositories.Link{
		Long:      long,
		Short:     ushort.String(),
		ExpiredAt: expirationDate,
	})
	if err != nil {
		return nil, ErrSaveLink(err.Error())
	}

	return &Link{
		ID:        link.ID,
		Long:      ulong,
		Short:     ushort,
		CreatedAt: link.CreatedAt,
		ExpiredAt: link.ExpiredAt,
	}, nil
}
