// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package queries

type Group struct {
	ID        int64
	Title     string
	IsPrivate int64
}

type GroupsSetting struct {
	ID                      int64
	GroupID                 int64
	IsWelcomeMessageEnabled int64
}

type Message struct {
	ID             int64
	GroupID        int64
	WelcomeMessage string
}

type User struct {
	ID           int64
	LanguageCode string
}