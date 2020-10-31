package models

// AppConfig - AppConfig
type AppConfig struct {
	AppPort      string
	JWTSecret    string
	MinJWTLength int
	MongoServer  string
}
