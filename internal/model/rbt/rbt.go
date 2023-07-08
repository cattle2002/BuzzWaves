package rbt

import (
	"BuzzWaves/pkkg"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strconv"
)

var RabbitClient *amqp.Connection
var RabbitChannel *amqp.Channel

func init() {
	RabbitClient = CreateConn()
	RabbitChannel = CreateChannel(RabbitClient)
	DeclareDirectExchange(RabbitChannel)
}

//封装rabbitmq等函数
func CreateConn() *amqp.Connection {
	dsn := "amqp://" + pkkg.GetRabbitmqUser() + ":" + strconv.Itoa(pkkg.GetRabbitmqPassword()) + "@" + pkkg.GetRabbitmqIP() + ":" + strconv.Itoa(pkkg.GetRabbitmqPort())
	//fmt.Println(dsn)
	conn, err := amqp.Dial(dsn)
	//conn, err := amqp.Dial("amqp://admin:123@localhost:5672/")
	if err != nil {
		fmt.Println("连接到rabbitmq失败", err)
	} else {

		fmt.Println("连接rabbbitmq server success")
	}
	//defer conn.Close()
	return conn
}
func CreateChannel(conn *amqp.Connection) *amqp.Channel {
	channel, err := conn.Channel()
	if err != nil {
		fmt.Println("创建信道失败", err)
	}

	return channel
}
func DeclareDirectExchange(ch *amqp.Channel) {
	err := ch.ExchangeDeclare(
		"BuzzWaves_offline_exchange", // exchange name
		"direct",                     // exchange type
		true,                         // durable
		false,                        // auto-deleted
		false,                        // internal
		false,                        // no-wait
		nil,                          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}
}
func DeclareQueue(ch *amqp.Channel, user string, friend string) amqp.Queue {
	queueName := user + ":" + friend
	queue, err := ch.QueueDeclare(
		queueName, // queue name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}
	return queue
}
func QueryBind(ch *amqp.Channel, q amqp.Queue) {
	err := ch.QueueBind(
		q.Name,                       // queue name
		q.Name,                       // routing key
		"BuzzWaves_offline_exchange", // exchange name
		false,                        // no-wait
		nil,                          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to bind the queue to the exchange: %v", err)
	}
}
func SendMessage(ch *amqp.Channel, message string, q amqp.Queue) {
	err := ch.Publish(
		"BuzzWaves_offline_exchange", // exchange name
		q.Name,                       // routing key
		false,                        // mandatory
		false,                        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}

}
func ConsumeMessage(ch *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {

	msgs, err := ch.Consume(
		queueName, // 队列名称（留空表示由 RabbitMQ 自动生成）
		"",        // 消费者标识符
		true,      // 自动应答
		false,     // 非独占消费者
		false,     // 不等待
		false,     // 无阻塞
		nil,       // 参数
	)
	return msgs, err
}
