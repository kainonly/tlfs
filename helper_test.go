package tlfs_test

import (
	"testing"

	"github.com/kainonly/tlfs"
	"github.com/stretchr/testify/assert"
)

func TestPayRedirectURL(t *testing.T) {
	dto := &tlfs.PayRedirectURLDto{
		Id:        "123",
		SellerId:  "234",
		EnterType: "1",
		Amount:    "200",
		FontsMap: []*tlfs.FontsMapItem{
			tlfs.NewFontsMapItem("2468", "测试"),
			tlfs.NewFontsMapItem("4568", "测试二"),
		},
	}

	redirectURL, err := x.PayRedirectURL(dto)
	assert.NoError(t, err)

	t.Log(redirectURL)
}
