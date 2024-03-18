package config

type Config struct {
	Server   Server
	Database Database
	Xendit   Xendit
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

type Xendit struct {
	XApiKey        string
	xCallbackToken string
}
