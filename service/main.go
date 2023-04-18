package service

type Service struct {
	Authorization
	Posting
}

func NewService(authorization Authorization, posting Posting) *Service {
	return &Service{Authorization: authorization, Posting: posting}
}
