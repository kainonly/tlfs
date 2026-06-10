package tlfs

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type PaymentContsQueryDto struct {
	StoreId   string `json:"storeId"`   // 门店唯一标识
	PaymentId string `json:"paymentId"` // 门店收款单唯一标识
}

func NewPaymentContsQueryDto(storeId string, paymentId string) *PaymentContsQueryDto {
	return &PaymentContsQueryDto{StoreId: storeId, PaymentId: paymentId}
}

type PaymentContsQueryResult struct {
	Result string
}

func (x *Tlfs) PaymentContsQuery(ctx context.Context, dto *PaymentContsQueryDto) (_ *PaymentContsQueryResult, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/fp-api/api/payment/conts/query`, data); err != nil {
		return
	}

	var result PaymentContsQueryResult
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
