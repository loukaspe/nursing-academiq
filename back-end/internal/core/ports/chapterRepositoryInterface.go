package ports

import (
	"context"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
)

type ChapterRepositoryInterface interface {
	GetChapter(ctx context.Context, uid uint32) (*domain.Chapter, error)
	GetChapterByTitle(ctx context.Context, name string) (*domain.Chapter, error)
	//GetExtendedChapter(ctx context.Context, uid uint32) (*domain.Chapter, error)
	CreateChapter(context.Context, *domain.Chapter) (uint, error)
	UpdateChapter(context.Context, uint32, *domain.Chapter) error
	DeleteChapter(ctx context.Context, uid uint32) error
}
