package repository

import (
	"errors"

	"github.com/sawalreverr/recything/internal/article/entity"
	"github.com/sawalreverr/recything/internal/database"
	"github.com/sawalreverr/recything/pkg"
)

type articleRepositoryImpl struct {
	DB database.Database
}

func NewArticleRepository(DB database.Database) ArticleRepository {
	return &articleRepositoryImpl{DB: db}
}

func (r *articleRepositoryImpl) CreateArticleRepository(*entity.Article, error) error {
	query := "INSERT INTO articles (id, title, content, image_url) VALUES (?, ?, ?, ?)"
	_, err := r.DB.GetDB().Exec(query, article.ID, article.Title, article.Content, article.ImageUrl)
	if err != nil {
		return err
	}
	return nil
}

func (r *articleRepositoryImpl) GetAllArticleRepository(limit, offset int) ([]entity.Article, int, error) {
	query := "SELECT id, title, content, image_url FROM articles LIMIT ? OFFSET ?"
	rows, err := r.DB.GetDB().Query(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var articles []entity.Article
	for rows.Next() {
		var article *entity.Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.ImageUrl); err != nil {
			return nil, 0, err
		}
		articles = append(articles, article)
	}

	countQuery := "SELECT COUNT(*) FROM articles"
	var total int
	err = r.DB.GetDB().QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}

func (r *articleRepositoryImpl) GetByIDArticleRepository(id string) (entity.Article, error) {
	query := "SELECT id, title, content, image_url FROM articles WHERE id = ?"
	row := r.DB.GetDB().QueryRow(query, id)

	var article entity.Article
	if err := row.Scan(&article.ID, &article.Title, &article.Content, &article.ImageUrl); err != nil {
		if errors.Is(err, r.DB.GetDB().ErrNoRows) {
			return article, pkg.ErrArticleNotFound
		}
		return article, err
	}
	return article, nil
}

func (r *articleRepositoryImpl) UpdateArticleRepository(article entity.Article) error {
	query := "UPDATE articles SET title = ?, content = ?, image_url = ? WHERE id = ?"
	_, err := r.DB.GetDB().Exec(query, article.Title, article.Content, article.ImageUrl, article.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *articleRepositoryImpl) DeleteArticleRepository(id string) error {
	query := "DELETE FROM articles WHERE id = ?"
	_, err := r.DB.GetDB().Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
