package tlfs_test

import (
	"testing"

	"github.com/kainonly/tlfs"
	"github.com/stretchr/testify/assert"
)

func TestPayRedirectURL(t *testing.T) {
	dto := &tlfs.PayRedirectURLDto{
		Id:        "c6d30374458c404ab631a69db686c132",
		SellerId:  x.Option.SellerID,
		EnterType: "1",
		Amount:    "3",
		FontsMap: []*tlfs.FontsMapItem{
			tlfs.NewFontsMapItem("c95ca51a373244aeb1578317d0a9d488", "81107"),
			tlfs.NewFontsMapItem("0b530118efd14b0b9c538127479c859b", "Kain"),
		},
	}

	redirectURL, err := x.PayRedirectURL(dto)
	assert.NoError(t, err)

	t.Log(redirectURL)
}
