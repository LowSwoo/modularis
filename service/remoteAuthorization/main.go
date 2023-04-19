package remoteAuthorization

import (
	"log"
	"modularis/models"
	"net/http"
	"net/url"
)

type Service struct {
	Server   AuthServerConfig
	Endpoint AuthServerEndpointConfig
}

type AuthServerConfig struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
}

type AuthServerEndpointConfig struct {
	SignUp string `yaml:"SignUp"`
	SignIn string `yaml:"SignIn"`
}

func NewService(serverConfig AuthServerConfig, endpointConfig AuthServerEndpointConfig) *Service {
	return &Service{
		Server:   serverConfig,
		Endpoint: endpointConfig,
	}
}

func (a Service) SignUp(user *models.User) (string, error) {
	var data url.Values
	data.Add("login", user.Login)
	data.Add("password", user.Password)
	data.Add("username", user.Username)
	resp, err := http.PostForm("http://"+a.Server.Host+":"+a.Server.Port+"/"+a.Endpoint.SignUp, data)
	if err != nil {
		return "", err
	}
	var bytes []byte
	_, err = resp.Body.Read(bytes)
	if err != nil {
		return "", err
	}
	log.Default().Println("Status code: " + resp.Status)
	log.Default().Println("Body: " + string(bytes))
	return string(bytes), err
}

func (a Service) SignIn(login, password string) (string, error) {
	data := url.Values{}
	data.Add("login", login)
	data.Add("password", password)
	resp, err := http.PostForm("http://"+a.Server.Host+":"+a.Server.Port+"/"+a.Endpoint.SignIn, data)
	if err != nil {
		return "", err
	}
	var bytes []byte
	_, err = resp.Body.Read(bytes)
	log.Default().Println("Status code: " + resp.Status)
	log.Default().Println("Body: " + string(bytes))
	return string(bytes), err
}
