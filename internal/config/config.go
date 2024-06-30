package config

type Config struct {
	ServerHost  string `env:"SERVER_HOST" envDefault:"localhost"`
	ServerPort  string `env:"SERVER_PORT" envDefault:"80"`
	StoragePath string `env:"STORAGE_PATH" envDefault:"./storage/"`

	DBDriver string `env:"DB_DRIVER"`
	DBName   string `env:"DB_NAME"`
}
