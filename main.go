package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var actions = []string{
	"open",
	"close",
	"read",
	"write",
}

type logItem struct {
	action string
	time   time.Time
}
type User struct {
	id    int
	email string
	logs  []logItem
}

func main() {
	time.Sleep(time.Millisecond * 30)
	t := time.Now()

	wg := &sync.WaitGroup{}
	rand.Seed(time.Now().UTC().UnixNano())

	users := generateUsers(1500)
	for _, user := range users {
		wg.Add(1)
		go saveUserInfo(user, wg)
	}
	wg.Wait()
	fmt.Println("Time since last save users:", time.Since(t).String())
}

func (u User) getActivityInfo() string {
	out := fmt.Sprintf("ID: %d | Email: %s\n Activity log:\n", u.id, u.email)

	for i, item := range u.logs {
		out += fmt.Sprintf("%d. [%s] at %s\n", i+1, item.action, item.time)
	}
	return out
}

func generateUsers(count int) []User {
	users := make([]User, count)
	for i := 0; i < count; i++ {
		users[i] = User{
			id:    i + 1,
			email: fmt.Sprintf("user%s@mail.go", i+1),
			logs:  generateLogs(rand.Intn(200)),
		}
	}
	return users
}
func saveUserInfo(user User, wg *sync.WaitGroup) error {
	fmt.Println("Saving user:\n", user.id)
	filename := fmt.Sprintf("logs/user_%d.txt", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = file.WriteString(user.getActivityInfo())
	if err != nil {
		return err
	}
	wg.Done()
	return nil
}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)
	for i := 0; i < count; i++ {
		logs[i] = logItem{
			action: actions[rand.Intn(len(actions)-1)],
			time:   time.Now(),
		}
	}
	return logs
}
