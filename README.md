# tlfs

`tlfs` 是一个 TLFS 的 Go SDK，提供：

- 查询收款单列表
- 查询收款单表单项
- 生成收款跳转页 URL

## 安装

```bash
go get github.com/kainonly/tlfs
```

## 配置

初始化使用 `tlfs.Option`：

| 字段        | 说明                          |
| ----------- | ----------------------------- |
| `base_url`  | TLFS 服务地址                 |
| `org_id`    | 组织 ID，可选                 |
| `app_id`    | 应用 ID                       |
| `key`       | 签名密钥                      |
| `seller_id` | 收款方 ID，用于生成跳转页链接 |

示例配置：

```yaml
base_url: https://fstest.allinpaygd.com
org_id:
app_id: your-app-id
key: your-sign-key
seller_id: your-seller-id
```

## 快速开始

```go
package main

import (
    "context"
    "fmt"

    "github.com/kainonly/tlfs"
)

func main() {
    client, err := tlfs.NewTlfs(tlfs.Option{
        BaseURL:  "https://fstest.allinpaygd.com",
        OrgID:    "",
        AppID:    "your-app-id",
        Key:      "your-sign-key",
        SellerID: "your-seller-id",
    })
    if err != nil {
        panic(err)
    }

    ctx := context.Background()

    result, err := client.PaymentQuery(ctx, tlfs.NewPaymentQueryDto("your-store-id"))
    if err != nil {
        panic(err)
    }

    for _, item := range result.List {
        fmt.Println(item.PaymentId, item.PaymentName)
    }
}
```

## API

### `PaymentQuery`

查询收款单列表。

### `PaymentContsQuery`

查询收款单表单项。

### `PayRedirectURL`

```go
redirectURL, err := client.PayRedirectURL(&tlfs.PayRedirectURLDto{
    Id:        "your-payment-id",
    SellerId:  "your-seller-id",
    EnterType: "1",
    Amount:    "300",
    FontsMap: []*tlfs.FontsMapItem{
        tlfs.NewFontsMapItem("field-id-1", "value-1"),
        tlfs.NewFontsMapItem("field-id-2", "value-2"),
    },
})
```

生成收款页面跳转链接。`Amount` 单位为分，`FontsMap` 会编码到 URL 查询参数中。

## 错误处理

HTTP 状态码不是 `200`，或者上游业务码不是 `200` 时会返回错误。

## 测试

1. 复制 [config/values.example.yml](config/values.example.yml) 为 `config/values.yml`
2. 填入有效的测试环境参数
3. 执行测试：

```bash
go test ./...
```

当前测试覆盖：

- 收款单列表查询
- 收款单表单项查询
- 收款跳转页 URL 生成
