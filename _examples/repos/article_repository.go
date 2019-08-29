package repos

import (
	"sort"
	"time"

	"github.com/fivestar/resozyme/_examples/model"
)

// NewArticleRepository creates an ArticleRepository.
func NewArticleRepository() *ArticleRepo {
	return &ArticleRepo{
		Data: []*model.Article{
			{
				ID:      1,
				Title:   "Foo",
				PubDate: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
			{
				ID:      2,
				Title:   "Bar",
				PubDate: time.Date(2019, time.February, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}
}

// ArticleRepository is a repository of the Article.
type ArticleRepository interface {
	Find(id int64) (*model.Article, error)
	FindLatest(limit int) ([]*model.Article, error)
	Add(article *model.Article) error
}

// ArticleRepo implements ArticleRepository.
type ArticleRepo struct {
	Data []*model.Article
}

// Find finds an Article by id.
func (repo *ArticleRepo) Find(id int64) (*model.Article, error) {
	for _, article := range repo.Data {
		if article.ID == id {
			return article, nil
		}
	}
	return nil, nil
}

// FindLatest finds latest articles.
func (repo *ArticleRepo) FindLatest(limit int) ([]*model.Article, error) {
	data := append(repo.Data[:0:0], repo.Data...)
	sort.SliceStable(data, func(i, j int) bool {
		if data[i].PubDate.Equal(data[j].PubDate) {
			return data[i].ID > data[j].ID
		}
		return data[i].PubDate.After(data[j].PubDate)
	})
	if limit >= len(data) {
		return data, nil
	}
	return data[:(limit - 1):limit], nil
}

// Add adds an Article to the repository.
func (repo *ArticleRepo) Add(article *model.Article) error {
	repo.Data = append(repo.Data, article)
	article.ID = int64(len(repo.Data))
	return nil
}
