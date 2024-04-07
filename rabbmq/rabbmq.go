package rabbmq

import (
	"context"
	"fmt"
	"github.com/jimu-server/logger/logger"
	"github.com/jimu-server/model"
	"github.com/jimu-server/mq/mq_key"
	jsoniter "github.com/json-iterator/go"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

var Client *amqp.Connection

// CreateUserTaskQueue 创建用户任务队列
func CreateUserNotifyQueue(id string) {
	var err error
	var ch *amqp.Channel
	if ch, err = Client.Channel(); err != nil {
		logger.Logger.Error(err.Error())
		return
	}
	defer ch.Close()
	_, err = ch.QueueDeclare(
		id,    // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		logger.Logger.Error(err.Error())
	}
}

// Notify 发布消息
// 根据用户 id 发布消息
// @param sourceId  发送者
// @param targetId	接收者
// @param t		    消息类型
func Notify(body *model.AppNotify) {
	if body == nil {
		logger.Logger.Error("消息体为空")
		return
	}
	// 计算用户消息队列名称
	key := fmt.Sprintf("%s%s", mq_key.Notify, body.SubId)
	var err error
	var ch *amqp.Channel
	var data []byte
	if ch, err = Client.Channel(); err != nil {
		logger.Logger.Error(err.Error())
		return
	}
	defer ch.Close()
	if data, err = jsoniter.Marshal(body); err != nil {
		logger.Logger.Error(err.Error())
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = ch.PublishWithContext(ctx,
		"",    // exchange
		key,   // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		})
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
}
