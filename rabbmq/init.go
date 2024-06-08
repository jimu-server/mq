package rabbmq

import (
	"fmt"
	"github.com/jimu-server/config"
	"github.com/jimu-server/logger"
	"github.com/jimu-server/mq/mapper"
	"github.com/jimu-server/mq/mq_key"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

var RabbitMQMapper = &mapper.RabbitMQMapper{}

func init() {
	var err error
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.Evn.App.RabbitMQ.User, config.Evn.App.RabbitMQ.Password, config.Evn.App.RabbitMQ.Host, config.Evn.App.RabbitMQ.Port)
	Client, err = amqp.Dial(url)
	if err != nil {
		logger.Logger.Panic("connect rabbitmq failed", zap.Error(err))
	}
	// 初始化一个 root 角色的队列
	key := fmt.Sprintf("%s%s", mq_key.Notify, "1")
	CreateUserNotifyQueue(key)
}
