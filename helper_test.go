package tlfs_test

import (
	"testing"

	"github.com/kainonly/tlfs"
	"github.com/stretchr/testify/assert"
)

func TestPayRedirectURL(t *testing.T) {
	dto := &tlfs.PayRedirectURLDto{
		Id:        opt.PaymentID,
		SellerId:  x.Option.SellerID,
		EnterType: "1",
		Amount:    "3",
		FontsMap: []*tlfs.FontsMapItem{
			tlfs.NewFontsMapItem("2af57f6c7a8e4da9be5a9a4099cf5384", "81107"),
			tlfs.NewFontsMapItem("3a5fa1073920469096542e59b6e42a36", "Kain"),
		},
	}

	redirectURL, err := x.PayRedirectURL(dto)
	assert.NoError(t, err)

	t.Log(redirectURL)
}
