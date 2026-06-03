package tests

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

type Task struct {
	ID      int64  `db:"id"`
	Date    string `db:"date"`
	Title   string `db:"title"`
	Comment string `db:"comment"`
	Repeat  string `db:"repeat"`
}

func count(db *sqlx.DB) (int, error) {
	var count int
	return count, db.Get(&count, `SELECT count(id) FROM scheduler`)
}

func openDB(t *testing.T) *sqlx.DB {
	// Без этого тест не видит переменные окружения и не читает TODO_DBFILE на чистой машине
	_ = godotenv.Load("../.env")

	dbfile := DBFile

	// Тесты запускаются из директории tests, а TODO_DBFILE содержит относительный путь
	// Поэтому относительные пути дополнительно приводятся через "..".
	// Иначе тесты просто не могут найти файл базы данных при прямом подключении, как тут.
	// ЕСЛИ ЭТОТ ФАЙЛ НЕЛЬЗЯ МЕНЯТЬ НИКАК, ТОГДА ПРИДЁТСЯ ПЕРЕД ЗАПУСКОМ ТЕСТА
	// ПРОПИСЫВАТЬ: export TODO_DBFILE=$(pwd)/путь к БД из env файла, Я ДРУГОГО ПУТИ НЕ НАШЁЛ
	envFile := os.Getenv("TODO_DBFILE")
	if len(envFile) > 0 {
		if !filepath.IsAbs(envFile) {
			dbfile = filepath.Join("..", envFile)
		} else {
			dbfile = envFile
		}
	}
	db, err := sqlx.Connect("sqlite", dbfile)
	assert.NoError(t, err)
	return db
}

func TestDB(t *testing.T) {
	db := openDB(t)
	defer db.Close()

	before, err := count(db)
	assert.NoError(t, err)

	today := time.Now().Format(`20060102`)

	res, err := db.Exec(`INSERT INTO scheduler (date, title, comment, repeat) 
	VALUES (?, 'Todo', 'Комментарий', '')`, today)
	assert.NoError(t, err)

	id, err := res.LastInsertId()

	var task Task
	err = db.Get(&task, `SELECT * FROM scheduler WHERE id=?`, id)
	assert.NoError(t, err)
	assert.Equal(t, id, task.ID)
	assert.Equal(t, `Todo`, task.Title)
	assert.Equal(t, `Комментарий`, task.Comment)

	_, err = db.Exec(`DELETE FROM scheduler WHERE id = ?`, id)
	assert.NoError(t, err)

	after, err := count(db)
	assert.NoError(t, err)

	assert.Equal(t, before, after)
}
