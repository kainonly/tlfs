package tlfs

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	"github.com/kainonly/go/help"
	"resty.dev/v3"
)

type Tlfs struct {
	Option *Option
	Client *resty.Client
}

type Option struct {
	BaseURL string `yaml:"base_url"`
	OrgID   string `yaml:"org_id"`
	AppID   string `yaml:"app_id"`
	Key     string `yaml:"key"`
	StoreId string `yaml:"store_id"`
}

func NewTlfs(opt Option) (x *Tlfs, err error) {
	x = &Tlfs{
		Option: &opt,
		Client: resty.New().SetBaseURL(opt.BaseURL),
	}
	return
}

type M map[string]any

func (x *Tlfs) SetNow(ctx context.Context, ts time.Time) context.Context {
	return context.WithValue(ctx, "now", ts)
}

func (x *Tlfs) GetNow(ctx context.Context) time.Time {
	return ctx.Value("now").(time.Time)
}

type ResponseBody struct {
	Code string `json:"code"`           // 调用结果返回码
	Msg  string `json:"msg"`            // 调用结果返回码描述
	Sign string `json:"sign,omitempty"` // 商户请求参数的签名串
	Data string `json:"data,omitempty"` // 返回参数的集合
}

func (x *Tlfs) Request(ctx context.Context, path string, data string) (_ string, err error) {
	now := x.GetNow(ctx)
	body := map[string]string{
		"version":   "1.0",
		"charset":   "utf-8",
		"timestamp": now.Format(`20060102150405`),
		"appId":     x.Option.AppID,
		"signType":  "MD5",
		"bizData":   data,
	}

	if x.Option.OrgID != "" {
		body["orgId"] = x.Option.OrgID
	}

	body["sign"], err = x.Sign(body)

	var resp *resty.Response
	if resp, err = x.Client.R().
		SetContext(ctx).
		SetBody(body).
		Post(path); err != nil {
		return
	}

	if resp.StatusCode() != 200 {
		err = help.E(0, resp.String())
		return
	}
	var content M
	if err = sonic.Unmarshal(resp.Bytes(), &content); err != nil {
		return
	}

	if content["code"] != "00000" {
		err = help.E(0, fmt.Sprintf(`第三方请求失败![%s]: %s`, content["code"], content["msg"]))
		return
	}

	delete(content, "sign")
	delete(content, "signType")

	return content["bizData"].(string), nil
}

func (x *Tlfs) Sign(params map[string]string) (sign string, err error) {
	keys := make([]string, 0, len(params))
	for k, _ := range params {
		k = strings.ToLower(k)
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sb strings.Builder
	for i, k := range keys {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(k)
		sb.WriteByte('=')
		sb.WriteString(params[k])
	}
	sb.WriteString("&key=")
	sb.WriteString(x.Option.Key)

	signContent := sb.String()
	h := md5.Sum([]byte(signContent))
	sign = strings.ToUpper(hex.EncodeToString(h[:]))
	return
}
