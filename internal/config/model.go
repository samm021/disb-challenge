package config

type Config struct {
	Server   Server
	Database Database
}

type Server struct {
	// Host string
	Port string
}

type Database struct {
	Name string
	// User     string
	// Password string
	// Host     string
	// Port     string
}
