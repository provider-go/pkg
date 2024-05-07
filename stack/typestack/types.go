package typestack

type ConfigStack struct {
	Addr     string // redis 连接地址
	Password string // redis 密码
	DB       int    // redis 库
}
