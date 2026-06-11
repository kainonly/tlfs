package tlfs_test

import (
	"context"
	"testing"

	"github.com/kainonly/tlfs"
	"github.com/stretchr/testify/assert"
)

func TestPaymentQuery(t *testing.T) {
	ctx := context.TODO()
	dto := tlfs.NewPaymentQueryDto(x.Option.AppID)

	r, err := x.PaymentQuery(ctx, dto)
	assert.NoError(t, err)

	for i, v := range r.List {
		t.Logf(`==== %d ====`, i)
		t.Log(`storeId:`, v.StoreId)
		t.Log(`paymentId:`, v.PaymentId)
		t.Log(`paymentName:`, v.PaymentName)
		t.Log(`dataStatus:`, v.DataStatus)
		t.Log(`type:`, v.Type)
		t.Log(`createdDate:`, v.CreatedDate)
	}
}
