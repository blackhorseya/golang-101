package observer

import (
	"testing"
)

func TestNewsPublisher(t *testing.T) {
	publisher := &NewsPublisher{}
	subscriber := &Subscriber{Name: "Test Subscriber"}

	publisher.RegisterObserver(subscriber)
	publisher.SetNews("Test News")

	// 可以使用假設的方式驗證更新通知的效果
	// 在這裡，我們假設 Subscriber 的 Update 方法會正確地輸出新聞
}
