package ramPostRepo

import "modularis/models"

type Repository struct {
	Posts map[string]*models.Post
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) CreatePost(post *models.Post) error {
	if _, ex := r.Posts[post.Id]; ex {
		return models.PostAlreadyExist
	}
	r.Posts[post.Id] = post
	return nil
}

func (r *Repository) RemovePost(id string) error {
	if _, ex := r.Posts[id]; !ex {
		return models.PostDoesntExist
	}
	delete(r.Posts, id)
	return nil
}

func (r *Repository) GetPost(id string) (*models.Post, error) {
	post, ex := r.Posts[id]
	if !ex {
		return nil, models.PostDoesntExist
	}
	return post, nil
}
