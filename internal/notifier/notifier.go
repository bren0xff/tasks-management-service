package notifier

import (
	"fmt"
	"log"
	"tasksManagement/internal/entity"
	"tasksManagement/pkg/queue"
)

type Notifier interface {
	NotifyManager(user *entity.User, task *entity.Task)
}

type notifier struct {
	queue queue.Queue
}

func NewNotifier(q queue.Queue) Notifier {
	return &notifier{q}
}

func (n *notifier) NotifyManager(user *entity.User, task *entity.Task) {
	message := fmt.Sprintf("The tech \"%s\" performed the task \"%s\" on date %s", user.Name, task.Summary, task.Date)
	err := n.queue.Publish("notifications", []byte(message))
	if err != nil {
		log.Println("Failed to publish notification:", err)
	}
}
