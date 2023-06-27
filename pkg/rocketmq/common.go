package rocketmq

type (
	ConsumerCommon struct {
		group         string
		topic         string
		tag           string
		handler       TopicHandler
		workerNum     int
		ackNum        int
		ackSecond     int   // 每个consumer 的配置
		numOfMessages int32 // 每次获取消息的条数
		waitSeconds   int64 // 获取消息等待时间
		chanNum       int   // 获取消息的管道的数量
	}
)

func NewConsumerConf() ConsumerCommon {
	conf := ConsumerCommon{
		numOfMessages: DefaultMsgGetNum,
		waitSeconds:   DefaultMsgWait,
		chanNum:       DefaultWatchChanNum,
		workerNum:     DefaultWorkerNum,
		ackNum:        DefaultAckNum,
		ackSecond:     DefaultAckSecond,
		handler:       make(TopicHandler, 0),
		tag:           "*",
	}
	return conf
}

func (s *ConsumerCommon) SetGroup(group string) *ConsumerCommon {
	s.group = group
	return s
}

func (s *ConsumerCommon) SetTopicAndHandler(topic string, h TopicHandler) *ConsumerCommon {
	s.topic = topic
	s.handler = h
	return s
}

func (s *ConsumerCommon) SetTag(tag string) *ConsumerCommon {
	s.group = tag
	return s
}

func (s *ConsumerCommon) SetWorkNum(num int) *ConsumerCommon {
	s.workerNum = num
	return s
}

func (s *ConsumerCommon) SetAckNum(num int) *ConsumerCommon {
	s.ackNum = num
	return s
}

func (s *ConsumerCommon) SetAckSecond(second int) *ConsumerCommon {
	s.ackSecond = second
	return s
}

func (s *ConsumerCommon) SetGetMsgNum(num int32) *ConsumerCommon {
	s.numOfMessages = num
	return s
}

func (s *ConsumerCommon) SetWatchWaitSecond(num int64) *ConsumerCommon {
	s.waitSeconds = num
	return s
}

func (s *ConsumerCommon) SetWatchChanNum(num int) *ConsumerCommon {
	s.chanNum = num
	return s
}
