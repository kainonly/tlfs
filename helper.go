package tlfs

import (
	"net/url"

	"github.com/bytedance/sonic"
)

type PayRedirectURLDto struct {
	Id        string          `json:"id"`                 // 收款单id
	SellerId  string          `json:"sellerId"`           // 收款人Id
	EnterType string          `json:"enterType"`          // 类型
	Amount    string          `json:"amount"`             // 金额（分为单位）
	FontsMap  []*FontsMapItem `json:"fontsMap,omitempty"` // 表单数据
}

type FontsMapItem struct {
	Id       string `json:"id"`                 // 表单项id
	Value    string `json:"value"`              // 值
	Disabled string `json:"disabled,omitempty"` // 禁用属性
}

func NewFontsMapItem(id, value string) *FontsMapItem {
	return &FontsMapItem{Id: id, Value: value}
}

func (x *FontsMapItem) SetDisabled() *FontsMapItem {
	x.Disabled = "1"
	return x
}

func (x *Tlfs) PayRedirectURL(dto *PayRedirectURLDto) (_ string, err error) {
	var fontsMap string
	if fontsMap, err = sonic.MarshalString(dto.FontsMap); err != nil {
		return
	}

	return x.Option.BaseURL + "/payment/views/payment.html?id=" + dto.Id +
		"&sellerId=" + dto.SellerId +
		"&enterType=" + dto.EnterType +
		"&amount=" + dto.Amount +
		"&fontsMap=" + url.QueryEscape(fontsMap), nil
}
