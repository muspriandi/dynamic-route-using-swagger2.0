package domain

type Route struct {
	Path   string
	Method string
}

type ServiceConfig struct {
	BaseURL string
	Prefix  string
	Routes  []Route
}
