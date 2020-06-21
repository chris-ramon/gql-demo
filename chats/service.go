package chats

type Service interface {
}

type service struct {
	repo *Repository
}

func NewService(repo *Repository) Service {
	srv := service{repo}

	return &srv
}
