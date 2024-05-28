package facade

import (
	"testing"
)

func TestFacadePattern(t *testing.T) {
	tv := &TV{}
	soundSystem := &SoundSystem{}
	lights := &Lights{}

	homeTheater := NewHomeTheaterFacade(tv, soundSystem, lights)

	// 測試開始觀看電影
	homeTheater.WatchMovie()

	// 測試結束觀看電影
	homeTheater.EndMovie()
}
