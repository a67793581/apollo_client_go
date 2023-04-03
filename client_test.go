package openapi

import (
	_context "context"
	"fmt"
	"golang.org/x/net/context"
	_nethttp "net/http"
	"testing"
)

func TestClient(t *testing.T) {
	//
	//const HOST_TEST = 'http://config-admin.test.huajiao.com/openapi/v1';
	//const HOST_ONLINE = 'http://config-admin.huajiao.com/openapi/v1';

	//6f6abd892f7a06f86fd280df8585e2d47df81fb6
	c := NewAPIClient(&Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
		Debug:         false,
		Servers: ServerConfigurations{
			{
				URL:         "http://config-admin.test.huajiao.com/openapi/v1",
				Description: "域名",
			},
		},
		OperationServers: map[string]ServerConfigurations{},
	})
	var (
		f   Field
		res *_nethttp.Response
		err GenericOpenAPIError
	)

	ctx := _context.Background()
	ctx = context.WithValue(ctx, ContextAPIKeys, map[string]APIKey{
		"ApiKeyAuth": {
			Key: "6f6abd892f7a06f86fd280df8585e2d47df81fb6",
		},
	})
	f, res, err = c.DefaultApi.
		GetField(ctx, "hj-live-php", "DEV", "default", "application", "APOLLO_TOKEN").
		Execute()
	fmt.Printf("打印结果 字段：%+v\n 响应：%+v\n 错误：%+v\n", f, res, err)
}
