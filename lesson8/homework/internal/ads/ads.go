package ads

import "time"

type Ad struct { //добавить в структуру объявления дату создания, обновления объявления и поддержать ее в существующих методах
	ID        int64
	Title     string `validate:"range:1,99"`
	Text      string `validate:"range:1,499"`
	AuthorID  int64
	Published bool
	Created   time.Time
	Updated   time.Time
}
