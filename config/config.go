package config

type (
	Config struct {
		Version string
		Server struct{
			Name string
		}
		Etcd struct{
			Addrs		[]string
			UserName	string
			Password	string
		}
	}
)

