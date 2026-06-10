package tlfs_test

import (
	"context"
	"testing"

	"github.com/kainonly/tlfs"
	"github.com/stretchr/testify/assert"
)

func TestPaymentQuery(t *testing.T) {
	ctx := context.TODO()
	dto := tlfs.NewPaymentQueryDto(x.Option.StoreId)

	r, err := x.PaymentQuery(ctx, dto)
	assert.NoError(t, err)

	t.Log(r.Result)
}
