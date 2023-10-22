package config

type Config struct {
	Mongo  Mongo
	Server Server
	Token  Token
}

type Server struct {
	Port string
}

type Mongo struct {
	Uri string
}

type Token struct {
	Access_Token string
}
