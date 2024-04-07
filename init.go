package mq

import (
	"embed"
	"github.com/jimu-server/db"
	"github.com/jimu-server/mq/rabbmq"
)

//go:embed mapper/file/*.xml
var mapperFile embed.FS

func init() {
	db.GoBatis.LoadByRootPath("mapper", mapperFile)
	db.GoBatis.ScanMappers(rabbmq.RabbitMQMapper)
}
