package ramUserRepo

import "modularis/models"

type Repository struct {
	Users map[string]*models.User
}

func NewRepository() *Repository {
	return &Repository{
		Users: make(map[string]*models.User),
	}
}

func (r *Repository) CreateUser(user *models.User) error {
	if _, ex := r.Users[user.Login]; ex {
		return models.UserAlreadyExist
	}
	r.Users[user.Login] = user
	return nil
}

func (r *Repository) RemoveUSer(login string) error {
	if _, ex := r.Users[login]; !ex {
		return models.UserDoesntExist
	}
	delete(r.Users, login)
	return nil
}

func (r *Repository) GetUserInfo(login string) (*models.User, error) {
	user, ex := r.Users[login]
	if !ex {
		return nil, models.UserDoesntExist
	}
	return user, nil
}
