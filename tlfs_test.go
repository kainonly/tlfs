package tlfs_test

import (
	"os"
	"testing"

	"github.com/kainonly/tlfs"
	"gopkg.in/yaml.v3"
)

var x *tlfs.Tlfs
var opt Option

type Option struct {
	BaseURL   string `yaml:"base_url"`
	OrgID     string `yaml:"org_id"`
	AppID     string `yaml:"app_id"`
	Key       string `yaml:"key"`
	SellerID  string `yaml:"seller_id"`
	PaymentID string `yaml:"payment_id"`
}

func TestMain(m *testing.M) {
	var err error
	var b []byte
	if b, err = os.ReadFile("./config/values.yml"); err != nil {
		return
	}
	if err = yaml.Unmarshal(b, &opt); err != nil {
		return
	}
	if x, err = tlfs.NewTlfs(tlfs.Option{
		BaseURL:  opt.BaseURL,
		OrgID:    opt.OrgID,
		AppID:    opt.AppID,
		Key:      opt.Key,
		SellerID: opt.SellerID,
	}); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}
