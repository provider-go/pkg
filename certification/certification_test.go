package certification

import (
	"github.com/provider-go/pkg/certification/typecert"
	"testing"
)

func TestName(t *testing.T) {
	c := typecert.ConfigCertification{
		Appcode:  "**************",
		Endpoint: "https://eid.shumaidata.com/eid/check",
	}

	cert := NewCertification("ali", c)
	res := cert.Send("111111111111111111", "biwow")
	t.Log(res)
}
