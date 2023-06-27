package rocketmq

const (
	DefaultWorkerNum    = 10  // 默认一个topic 消费的协程数量
	DefaultAckNum       = 16  // 默认一次ack的消息数 阿里云 ackMessage 一次最多消费16条 超出会报错
	DefaultAckSecond    = 2   // ackArray 数据不够ackNum时, 一次ack的默认时间间隔
	DefaultMsgGetNum    = 10  // 默认每次获取消息的条数
	DefaultMsgWait      = 10  // 默认获取消息等待时间
	DefaultWatchChanNum = 100 // 默认获取mq消息的管道
)

type (
	TopicHandler map[string]func(data []byte) bool

	RocketMQConfig struct {
		Endpoint   string
		AccessKey  string
		SecretKey  string
		InstanceID string
	}
)
