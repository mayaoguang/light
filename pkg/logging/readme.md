## logging 日志模块

目前支持两种编码模式, `console` 和 `json`; 可在配置文件中设置 `Encoder` 修改;

### console

```
[2023-01-10 15:33:53]	[debug]	[light/pkg/logging.TestNewLogger:8]	debug msg
[2023-01-10 15:33:53]	[debug]	[light/pkg/logging.TestNewLogger:9]	debugf light
[2023-01-10 15:33:53]	[info]	[light/pkg/logging.TestNewLogger:10]	info msg
```

### json

```
{"level":"error","ts":"2023-01-10 15:35:02","caller":"light/pkg/logging.TestFatalf:22","msg":"err msg"}
{"level":"error","ts":"2023-01-10 15:35:02","caller":"light/pkg/logging.TestFatalf:23","msg":"errorf [1 2 3]"}
{"level":"fatal","ts":"2023-01-10 15:35:02","caller":"light/pkg/logging.TestFatalf:24","msg":"fatalf map[name:master]"}
```
