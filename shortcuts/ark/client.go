package ark

import (
	"github.com/example/gocli/internal/config"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
	"github.com/volcengine/volcengine-go-sdk/volcengine/universal"
)

const (
	defaultRegion = "cn-beijing"
	serviceName   = "ark"
	apiVersion    = "2024-01-01"
)

// newUniversalClient 从 gocli config 创建 volcengine Universal 客户端
func newUniversalClient(cfg *config.Config) (*universal.Universal, error) {
	volcConfig := volcengine.NewConfig().
		WithRegion(defaultRegion).
		WithCredentials(credentials.NewStaticCredentials(cfg.APIKey, cfg.APISecret, ""))
	sess, err := session.NewSession(volcConfig)
	if err != nil {
		return nil, err
	}
	return universal.New(sess), nil
}
