package monitor

import (
	"context"
	"light/internal/config"
)

// Start 启动定时监控服务.
func Start(ctx context.Context) {
	if !config.Config.IsMonitor {
		return
	}
}
