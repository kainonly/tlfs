package tlfs

import (
	"context"
	"fmt"
	"time"

	"github.com/bytedance/sonic"
	"github.com/kainonly/go/help"
)

type PaymentContsQueryDto struct {
	StoreId   string `json:"storeId"`   // 门店唯一标识
	PaymentId string `json:"paymentId"` // 门店收款单唯一标识
}

func NewPaymentContsQueryDto(storeId string, paymentId string) *PaymentContsQueryDto {
	return &PaymentContsQueryDto{StoreId: storeId, PaymentId: paymentId}
}

type PaymentContsQueryResult struct {
	StoreId   string                  `json:"storeId"`   // 门店唯一标识
	PaymentId string                  `json:"paymentId"` // 收款单唯一标识
	Conts     []PaymentContsQueryCont `json:"conts"`     // 收款单信息列表
}

type PaymentContsQueryCont struct {
	Name              string `json:"name"`              // 信息项名称
	Type              string `json:"type"`              // 类型
	Sort              int    `json:"sort"`              // 排序
	Tip               string `json:"tip,omitempty"`     // 提示信息
	Descr             string `json:"descr,omitempty"`   // 字段描述
	Data              string `json:"data,omitempty"`    // 选项值
	Value             string `json:"value,omitempty"`   // 默认值
	Deposit           int64  `json:"deposit,omitempty"` // 商品押金，单位分
	SearchRequired    int    `json:"searchrequired,omitempty"`
	ReceivedRequired  int    `json:"receivedRequired,omitempty"`
	FoldRequired      int    `json:"foldRequired,omitempty"`
	MulPickerRequired int    `json:"mulPickerRequired,omitempty"`
	SmsCodeRequired   int    `json:"smsCodeRequired,omitempty"`
	RequireDeposit    int    `json:"requireDeposit,omitempty"`
	ScanRequired      int    `json:"scanRequired,omitempty"`
	FixAmountRequired int    `json:"fixAmountRequired,omitempty"`
	Desensitized      int    `json:"desensitized,omitempty"`
}

func (x *Tlfs) PaymentContsQuery(ctx context.Context, dto *PaymentContsQueryDto) (_ *PaymentContsQueryResult, err error) {
	now := time.Now()
	var content string
	if content, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var b []byte
	if b, err = x.Request(x.SetNow(ctx, now), `/fp-api/api/payment/conts/query`, content); err != nil {
		return
	}

	var r ResponseBody[*PaymentContsQueryResult]
	if err = sonic.Unmarshal(b, &r); err != nil {
		return
	}

	if r.Code != "200" {
		err = help.E(0, fmt.Sprintf(`第三方请求失败![%s]: %s`, r.Code, r.Msg))
		return
	}

	return r.Data, nil
}
