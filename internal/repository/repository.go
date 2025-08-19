package repository

import (
	"context"

	"github.com/nandoalvarado022/comments-play/internal/models"
)

type Repository interface {
	//posts
	InsertPost(ctx context.Context, comment *models.Comment) error
	GetPostById(ctx context.Context, id string) (*models.Comment, error)
	UpdatePost(ctx context.Context, post *models.Comment) error
	DeletePost(ctx context.Context, id string, userId string) error
	ListPost(ctx context.Context, page uint64) ([]*models.Comment, error)
	Close() error
}

var implementacion Repository

func SetRepository(repositorio Repository) {
	implementacion = repositorio
}

func Close() error {
	return implementacion.Close()
}

func InsertPost(ctx context.Context, post *models.Comment) error {
	return implementacion.InsertPost(ctx, post)
}

func GetPostById(ctx context.Context, id string) (*models.Comment, error) {
	return implementacion.GetPostById(ctx, id)
}

func UpdatePost(ctx context.Context, post *models.Comment) error {
	return implementacion.UpdatePost(ctx, post)
}

func DeletePost(ctx context.Context, id string, userId string) error {
	return implementacion.DeletePost(ctx, id, userId)
}
func ListPost(ctx context.Context, page uint64) ([]*models.Comment, error) {
	return implementacion.ListPost(ctx, page)
}
