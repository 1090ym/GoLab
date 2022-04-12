package storage

import (
	"fmt"
	"testing"
)

type User struct {
	id      int
	content Content
}

type Content struct {
	nums   []int
	counts []*Count
}

type Count struct {
	c string
}

func (u *User) GetCount() Content {
	return u.content
}

func (u User) UpdateId() {
	u.id++
}

func PrintContent(content Content) {
	fmt.Println(content)
	for _, count := range content.counts {
		fmt.Println(*count)
	}
}

func TestGetContent(t *testing.T) {
	content := Content{
		nums: []int{1, 2, 3},
		counts: []*Count{
			{"test1"},
			{"test2"},
			{"test3"},
			{"test4"},
			//我的设备可以派上用场了
		},
	}
	user := User{
		id:      0,
		content: content,
	}
	user.UpdateId()
	user.UpdateId()
	fmt.Println(user)
}
