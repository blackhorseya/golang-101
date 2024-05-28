package proxy

import (
	"testing"
)

func TestProxyPattern(t *testing.T) {
	image := NewProxyImage("test_image.jpg")

	// 第一次顯示應該會加載圖像
	expectedFirstLoad := "Displaying test_image.jpg"
	resultFirstLoad := image.Display()
	if resultFirstLoad != expectedFirstLoad {
		t.Errorf("Expected %s, but got %s", expectedFirstLoad, resultFirstLoad)
	}

	// 第二次顯示不應該重新加載圖像
	resultSecondLoad := image.Display()
	if resultSecondLoad != expectedFirstLoad {
		t.Errorf("Expected %s, but got %s", expectedFirstLoad, resultSecondLoad)
	}
}
