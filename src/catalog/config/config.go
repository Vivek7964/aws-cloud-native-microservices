package config

// Configuration exported
type AppConfiguration struct {
	Port      int    `env:"PORT,default=8080"`
	ImagePath string `env:"IMAGE_PATH,default=./images/"`
	Database  DatabaseConfiguration
}

type DatabaseConfiguration struct {
	Type string `env:"DB_TYPE,default=json"`
}