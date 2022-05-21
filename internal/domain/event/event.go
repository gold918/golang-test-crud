package event

type Event struct {
	Id    int64  `db:"id,omitempty"`
	Title string `db:"title"`
}
