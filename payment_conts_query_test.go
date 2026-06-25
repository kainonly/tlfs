package tlfs_test

import (
	"context"
	"testing"

	"github.com/kainonly/tlfs"
	"github.com/stretchr/testify/assert"
)

func TestPaymentContsQuery(t *testing.T) {
	ctx := context.TODO()
	dto := tlfs.NewPaymentContsQueryDto(x.Option.AppID, opt.PaymentID)

	r, err := x.PaymentContsQuery(ctx, dto)
	assert.NoError(t, err)

	t.Log(r.StoreId)
	t.Log(r.PaymentId)

	for i, v := range r.Conts {
		t.Logf(`==== %d ====`, i)
		t.Log(`id:`, v.ID)
		t.Log(`name:`, v.Name)
		t.Log(`type:`, v.Type)
		t.Log(`sort:`, v.Sort)
		t.Log(`tip:`, v.Tip)
		t.Log(`descr:`, v.Descr)
		t.Log(`data:`, v.Data)
		t.Log(`value:`, v.Value)
		t.Log(`deposit:`, v.Deposit)
		t.Log(`searchRequired:`, v.SearchRequired)
		t.Log(`receivedRequired:`, v.ReceivedRequired)
		t.Log(`foldRequired:`, v.FoldRequired)
		t.Log(`mulPickerRequired:`, v.MulPickerRequired)
		t.Log(`smsCodeRequired:`, v.SmsCodeRequired)
		t.Log(`requireDeposit:`, v.RequireDeposit)
		t.Log(`scanRequired:`, v.ScanRequired)
		t.Log(`fixAmountRequired:`, v.FixAmountRequired)
		t.Log(`desensitized:`, v.Desensitized)
	}
}
