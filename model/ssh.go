package model

type SSHConfig struct {
	Hostname     string
	User         string
	Port         int
	KeyFile      string
	ForwardAgent bool
}

type Connection struct {
	Hostname     string
	User         string
	Port         int
	KeyFile      string
	ForwardAgent bool
}
