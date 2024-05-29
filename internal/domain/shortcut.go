package domain

import "time"

type Shortcut struct {
	LongURL   string    `json:"longURL"`
	ShortURL  string    `json:"shortURL"`
	Domain    string    `json:"domain"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiredAt time.Time `json:"expiredAt"`
}

type RequestedShortcut struct {
	URL string `json:"url"`
}
