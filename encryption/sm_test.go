package encryption

import (
	"github.com/provider-go/pkg/util"
	"strconv"
	"testing"
)

func TestJianRong(t *testing.T) {
	pub := "021f5949aa53241d0dd9723cf7b1b3cbfe9807bc599c5128dd1ec0be36ac459f6c"
	sign := "304502210082102686fb4059b442a4e7bad3d841edf9c8dfd3ad4bb8e98617aee8c308f1f0022025b8a9a641d405f722252bdc951266cafb79e804a3f78c88594820965545c455"
	t.Log(SM2Verify(pub, "star", sign))
}

func TestName(t *testing.T) {
	pri := "e2292fb9867804a713bb7866690a6ad9aa07c9fcc8280de945e73d27c5181c22"
	timestamp := util.CurrentSecond()
	sign := SM2Sign(pri, strconv.FormatInt(timestamp, 10))
	t.Log(timestamp)
	t.Log(sign)
}
