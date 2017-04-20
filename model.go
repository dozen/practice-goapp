package main

import "time"

type User struct {
	ID        int
	Account   string
	PassHash  string
	CreatedAt time.Time

	Jokes  []*Joke
	Rates  []*Rate
	Themes []*Theme
}

type Joke struct {
	ID        int
	UserID    int
	ThemeID   int
	Text      string
	CreatedAt time.Time

	Author *User
	Theme  *Theme
	Rates  []*Rate
}

type Rate struct {
	ID        int
	UserID    int
	JokeID    int
	Star      int
	CreatedAt time.Time

	Author *User
	Joke   *Joke
}

type Theme struct {
	ID         int
	ImageID    int
	CategoryId int
	CreatedAt  time.Time

	Category string

	Author *User
	Jokes  []*Joke
}

type Image struct {
	ID        int
	ImageID   int
	CreatedAt time.Time
	UserID    int

	Author *User
	Theme  *Theme
}
