package campaign

import "time"


type Contact struct {
	ID        string `json:"id"`
	Name			string `json:"name"`
	Email			string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Campaign struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	Contacts  []Contact `json:"contacts"`
	CreatedAt time.Time `json:"created_at"`
}