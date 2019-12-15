// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package meetup

type Meetup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	User        *User  `json:"user"`
}

type User struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Meetups  []*Meetup `json:"meetups"`
}
