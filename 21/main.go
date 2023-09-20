package main

import "fmt"

// MediaPlayer интерфейс для проигрывания аудиофайлов
type MediaPlayer interface {
	Play(fileType, fileName string)
}

// AdvancedMediaPlayer интерфейс для проигрывания специфичных форматов
type AdvancedMediaPlayer interface {
	PlayVLC(fileName string)
	PlayMP4(fileName string)
}

// VLCPlayer реализация AdvancedMediaPlayer для формата VLC
type VLCPlayer struct{}

func (v *VLCPlayer) PlayVLC(fileName string) {
	fmt.Printf("Проигрывание VLC-файла: %s\n", fileName)
}

func (v *VLCPlayer) PlayMP4(fileName string) {
	// Ничего не делаем
}

// MP4Player реализация AdvancedMediaPlayer для формата MP4
type MP4Player struct{}

func (m *MP4Player) PlayVLC(fileName string) {
	// Ничего не делаем
}

func (m *MP4Player) PlayMP4(fileName string) {
	fmt.Printf("Проигрывание MP4-файла: %s\n", fileName)
}

// MediaAdapter адаптер для проигрывания разных форматов через MediaPlayer
type MediaAdapter struct {
	AdvancedMediaPlayer
}

func (a *MediaAdapter) Play(fileType, fileName string) {
	if fileType == "vlc" {
		a.PlayVLC(fileName)
	} else if fileType == "mp4" {
		a.PlayMP4(fileName)
	} else {
		fmt.Printf("Формат %s не поддерживается\n", fileType)
	}
}

// AudioPlayer реализация MediaPlayer
type AudioPlayer struct {
	adapter MediaAdapter
}

func (a *AudioPlayer) Play(fileType, fileName string) {
	if fileType == "mp3" {
		fmt.Printf("Проигрывание MP3-файла: %s\n", fileName)
	} else if fileType == "vlc" || fileType == "mp4" {
		a.adapter.Play(fileType, fileName)
	} else {
		fmt.Printf("Формат %s не поддерживается\n", fileType)
	}
}

func main() {
	player := &AudioPlayer{}

	player.Play("mp3", "song.mp3")
	player.Play("vlc", "video.vlc")
	player.Play("mp4", "movie.mp4")
	player.Play("avi", "video.avi")
}
