package tlfs_test

import (
	"testing"

	"github.com/kainonly/tlfs"
	"github.com/stretchr/testify/assert"
)

func TestPayRedirectURL(t *testing.T) {
	dto := &tlfs.PayRedirectURLDto{
		Id:        "c6d30374458c404ab631a69db686c132",
		SellerId:  x.Option.AppID,
		EnterType: "1",
		Amount:    "1",
		FontsMap: []*tlfs.FontsMapItem{
			tlfs.NewFontsMapItem("客户号", "81107"),
			tlfs.NewFontsMapItem("客户名称", "Kain"),
		},
	}

	redirectURL, err := x.PayRedirectURL(dto)
	assert.NoError(t, err)

	t.Log(redirectURL)
}
