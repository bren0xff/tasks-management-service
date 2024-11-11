package notifier

import (
	"fmt"
	"log"
	"tasksManagement/internal/entity"
	"tasksManagement/pkg/queue"
)

type Notifier interface {
	NotifyManager(task *entity.Task)
}

type notifier struct {
	queue queue.Queue
}

func NewNotifier(q queue.Queue) Notifier {
	return &notifier{q}
}

func (n *notifier) NotifyManager(task *entity.Task) {
	message := fmt.Sprintf("The tech %d performed the task %d on date %s", task.UserID, task.ID, task.Date)
	err := n.queue.Publish("notifications", []byte(message))
	if err != nil {
		log.Println("Failed to publish notification:", err)
	}
}
