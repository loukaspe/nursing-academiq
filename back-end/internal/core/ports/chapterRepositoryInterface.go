package ports

import (
	"context"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
)

type ChapterRepositoryInterface interface {
	GetChapter(ctx context.Context, uid uint32) (*domain.Chapter, error)
	//GetExtendedChapter(ctx context.Context, uid uint32) (*domain.Chapter, error)
}
