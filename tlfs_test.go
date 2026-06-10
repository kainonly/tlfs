package tlfs_test

import (
	"os"
	"testing"

	"github.com/kainonly/tlfs"
	"gopkg.in/yaml.v3"
)

var x *tlfs.Tlfs
var v *tlfs.Option

func TestMain(m *testing.M) {
	var err error
	var b []byte
	if b, err = os.ReadFile("./config/values.yml"); err != nil {
		return
	}
	if err = yaml.Unmarshal(b, &v); err != nil {
		return
	}
	if x, err = tlfs.NewTlfs(*v); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}
