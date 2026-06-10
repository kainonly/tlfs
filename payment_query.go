package tlfs

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

type PaymentQueryDto struct {
	StoreId string `json:"storeId"` // 门店唯一标识
}

func NewPaymentQueryDto(storeId string) *PaymentQueryDto {
	return &PaymentQueryDto{StoreId: storeId}
}

type PaymentQueryResult struct {
	Result string
}

func (x *Tlfs) PaymentQuery(ctx context.Context, dto *PaymentQueryDto) (_ *PaymentQueryResult, err error) {
	now := time.Now()
	var data string
	if data, err = sonic.MarshalString(*dto); err != nil {
		return
	}

	var bizData string
	if bizData, err = x.Request(x.SetNow(ctx, now), `/fp-api/api/payment/query`, data); err != nil {
		return
	}

	var result PaymentQueryResult
	if err = sonic.UnmarshalString(bizData, &result); err != nil {
		return
	}
	return &result, nil
}
