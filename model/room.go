package model

import "gorilla/config"

type Room struct {
	Number int      `json:"number"`
	Member []string `json:"member"`
}

func (r *Room) CheckRoom(gorilla *config.ViperConfig) bool {
	return true
}
