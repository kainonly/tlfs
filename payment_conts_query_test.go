package tlfs_test

import (
	"context"
	"testing"

	"github.com/kainonly/tlfs"
	"github.com/stretchr/testify/assert"
)

func TestPaymentContsQuery(t *testing.T) {
	ctx := context.TODO()
	dto := tlfs.NewPaymentContsQueryDto(x.Option.StoreId, ``)

	r, err := x.PaymentContsQuery(ctx, dto)
	assert.NoError(t, err)

	t.Log(r.Result)
}
