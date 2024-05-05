package typemysql

type ConfigMysql struct {
	Dsn          string
	MaxIdleConns int
	MaxOpenConns int
}
