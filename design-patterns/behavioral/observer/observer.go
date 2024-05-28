package observer

import (
	"fmt"
)

// Observer 定義了觀察者的接口
type Observer interface {
	Update(news string)
}

// Subject 定義了主題的接口
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

// NewsPublisher 是具體的主題，實現了 Subject 接口
type NewsPublisher struct {
	observers []Observer
	news      string
}

func (n *NewsPublisher) RegisterObserver(observer Observer) {
	n.observers = append(n.observers, observer)
}

func (n *NewsPublisher) RemoveObserver(observer Observer) {
	for i, obs := range n.observers {
		if obs == observer {
			n.observers = append(n.observers[:i], n.observers[i+1:]...)
			break
		}
	}
}

func (n *NewsPublisher) NotifyObservers() {
	for _, observer := range n.observers {
		observer.Update(n.news)
	}
}

func (n *NewsPublisher) SetNews(news string) {
	n.news = news
	n.NotifyObservers()
}

// Subscriber 是具體的觀察者，實現了 Observer 接口
type Subscriber struct {
	Name string
}

func (s *Subscriber) Update(news string) {
	fmt.Printf("%s received news: %s\n", s.Name, news)
}
