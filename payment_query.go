package tlfs

import (
	"context"
	"fmt"
	"time"

	"github.com/bytedance/sonic"
	"github.com/kainonly/go/help"
)

type PaymentQueryDto struct {
	StoreId string `json:"storeId"` // 门店唯一标识
}

func NewPaymentQueryDto(storeId string) *PaymentQueryDto {
	return &PaymentQueryDto{StoreId: storeId}
}

type PaymentQueryResult struct {
	List []PaymentQueryList `json:"list"`
}

type PaymentQueryList struct {
	StoreId     string `json:"storeId"`     // 门店唯一标识
	PaymentId   string `json:"paymentId"`   // 收款单唯一标识
	PaymentName string `json:"paymentName"` // 收款单名称
	DataStatus  string `json:"dataStatus"`  // 收款单状态
	Type        string `json:"type"`        // 收款单类型
	CreatedDate string `json:"createdDate"` // 创建时间
}

func (x *Tlfs) PaymentQuery(ctx context.Context, dto *PaymentQueryDto) (_ *PaymentQueryResult, err error) {
	now := time.Now()
	var content string
	if content, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var b []byte
	if b, err = x.Request(x.SetNow(ctx, now), `/fp-api/api/payment/query`, content); err != nil {
		return
	}

	var r ResponseBody[*PaymentQueryResult]
	if err = sonic.Unmarshal(b, &r); err != nil {
		return
	}

	if r.Code != "200" {
		err = help.E(0, fmt.Sprintf(`第三方请求失败![%s]: %s`, r.Code, r.Msg))
		return
	}

	return r.Data, nil
}
