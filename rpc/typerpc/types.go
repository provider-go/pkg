package typerpc

type ConfigRPC struct {
	Method string
	URL    string
	Data   string
	Header map[string]string
}
