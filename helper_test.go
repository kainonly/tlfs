package tlfs_test

import (
	"fmt"
	"testing"

	"github.com/kainonly/tlfs"
	"github.com/stretchr/testify/assert"
)

func TestPayRedirectURL(t *testing.T) {
	dto := &tlfs.PayRedirectURLDto{
		Id:        "9de60b130b46427b880104e388a932a6",
		SellerId:  "234",
		EnterType: "1",
		Amount:    "1",
		FontsMap: []*tlfs.FontsMapItem{
			tlfs.NewFontsMapItem("2468", "测试"),
			tlfs.NewFontsMapItem("4568", "测试二"),
		},
	}

	redirectURL, err := x.PayRedirectURL(dto)
	assert.NoError(t, err)

	assert.Equal(t,
		fmt.Sprintf("%s/payment/views/payment.html?id=123&sellerId=234&enterType=1&amount=200&fontsMap=[{\"id\":\"2468\",\"value\":\"测试\"},{\"id\":\"4568\",\"value\":\"测试二\"}]", x.Option.BaseURL),
		redirectURL,
	)
}
