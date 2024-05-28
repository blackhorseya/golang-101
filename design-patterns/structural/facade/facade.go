package facade

import (
	"fmt"
)

// TV 表示家庭影院中的電視
type TV struct{}

func (t *TV) On() {
	fmt.Println("Turning on the TV")
}

func (t *TV) Off() {
	fmt.Println("Turning off the TV")
}

// SoundSystem 表示家庭影院中的音響系統
type SoundSystem struct{}

func (s *SoundSystem) On() {
	fmt.Println("Turning on the sound system")
}

func (s *SoundSystem) Off() {
	fmt.Println("Turning off the sound system")
}

// Lights 表示家庭影院中的燈光系統
type Lights struct{}

func (l *Lights) Dim() {
	fmt.Println("Dimming the lights")
}

func (l *Lights) Brighten() {
	fmt.Println("Brightening the lights")
}

// HomeTheaterFacade 提供了一個簡單的接口來控制家庭影院系統
type HomeTheaterFacade struct {
	tv          *TV
	soundSystem *SoundSystem
	lights      *Lights
}

// NewHomeTheaterFacade 用於創建 HomeTheaterFacade 對象.
func NewHomeTheaterFacade(tv *TV, soundSystem *SoundSystem, lights *Lights) *HomeTheaterFacade {
	return &HomeTheaterFacade{
		tv:          tv,
		soundSystem: soundSystem,
		lights:      lights,
	}
}

func (ht *HomeTheaterFacade) WatchMovie() {
	ht.lights.Dim()
	ht.tv.On()
	ht.soundSystem.On()
	fmt.Println("Movie is starting...")
}

func (ht *HomeTheaterFacade) EndMovie() {
	ht.tv.Off()
	ht.soundSystem.Off()
	ht.lights.Brighten()
	fmt.Println("Movie is finished...")
}
