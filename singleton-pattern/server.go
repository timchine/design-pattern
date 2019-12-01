package singleton_pattern

type Conner struct {
	address  string //连接的地址
	userName string //
	password string //
}

func Open(config ConnConfig) *Conner {
	return &Conner{
		address: config.address,
		userName: config.userName,
		password: config.password,
	}
}