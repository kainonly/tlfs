package tlfs

import (
	"context"
	"crypto/md5"
	"encoding/hex"
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

type ResponseBody[T any] struct {
	Code string `json:"code"`           // 调用结果返回码
	Msg  string `json:"msg"`            // 调用结果返回码描述
	Sign string `json:"sign,omitempty"` // 商户请求参数的签名串
	Data T      `json:"data,omitempty"` // 返回参数的集合
}

func (x *Tlfs) Request(ctx context.Context, path string, content string) (_ []byte, err error) {
	now := x.GetNow(ctx)
	data := map[string]string{
		"version":    "1.0",
		"charset":    "utf-8",
		"timestamp":  now.Format(`20060102150405`),
		"appId":      x.Option.AppID,
		"bizContent": content,
	}

	if x.Option.OrgID != "" {
		data["orgId"] = x.Option.OrgID
	}

	data["sign"], err = x.Sign(data)
	if err != nil {
		return
	}

	var body string
	if body, err = sonic.MarshalString(data); err != nil {
		return
	}

	var resp *resty.Response
	if resp, err = x.Client.R().SetContext(ctx).
		SetHeader("Content-Type", "application/json; charset=utf-8").
		SetBody(body).
		Post(path); err != nil {
		return
	}

	if resp.StatusCode() != 200 {
		err = help.E(0, resp.String())
		return
	}

	return resp.Bytes(), nil
}

func (x *Tlfs) Sign(params map[string]string) (sign string, err error) {
	normalized := make(map[string]string, len(params))
	keys := make([]string, 0, len(params))
	for k, v := range params {
		if v == "" || strings.EqualFold(v, "null") {
			continue
		}
		k = strings.ToLower(k)
		if _, exists := normalized[k]; !exists {
			keys = append(keys, k)
		}
		normalized[k] = v
	}
	sort.Strings(keys)

	var sb strings.Builder
	for i, k := range keys {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(k)
		sb.WriteByte('=')
		sb.WriteString(normalized[k])
	}
	if sb.Len() > 0 {
		sb.WriteByte('&')
	}
	sb.WriteString("key=")
	sb.WriteString(x.Option.Key)

	signContent := sb.String()
	h := md5.Sum([]byte(signContent))
	sign = strings.ToUpper(hex.EncodeToString(h[:]))
	return
}
