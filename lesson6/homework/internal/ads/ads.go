package ads

type Ad struct {
	ID        int64
	Title     string `validate:"range:1,99"`
	Text      string `validate:"range:1,499"`
	AuthorID  int64
	Published bool
}
