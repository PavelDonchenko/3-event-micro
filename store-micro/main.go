package main

import "github.com/PavelDonchenko/3-event-micro/common"

func main() {
	common.Consumer(common.AmqpUrl, common.RabbitQueueName)
}
