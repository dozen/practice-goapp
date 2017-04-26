package main

import (
	"crypto/sha512"
	"database/sql"
	"encoding/hex"
	"github.com/kataras/go-sessions"
	"github.com/valyala/fasthttp"
	"regexp"
	"strconv"
	"fmt"
	marshal "github.com/dozen/ruby-marshal"
	"bytes"
)

func Must(stmt *sql.Stmt, err error) *sql.Stmt {
	if err != nil {
		panic(err)
	}
	return stmt
}

type Session struct {
	CsrfToken string `ruby:"csrf_token"`
	User `ruby:"user"`
	Flash `ruby:"__FLASH__"`
}

type Flash struct {}

func GetSession(c *fasthttp.RequestCtx) Session {
	cookie := string(c.Request.Header.Cookie("rack.session"))
	b, e := redi.Get("rack:session:" + cookie).Bytes()
	fmt.Printf("%#v", cookie)
	if e != nil {
		fmt.Printf("GetSession Error: %#v\n", e)
		return Session{}
	}
	s := Session{}
	marshal.NewDecoder(bytes.NewReader(b)).Decode(&s)
	return s
}

func ParamID(c *fasthttp.RequestCtx) (int, bool) {
	id, err := strconv.Atoi(c.UserValue("id").(string))
	if err != nil {
		return 0, false
	}
	return id, true
}

func ParamNew(c *fasthttp.RequestCtx) (bool, int, bool) {
	userValueID := c.UserValue("id").(string)
	if userValueID == "new" {
		return true, 0, true
	}
	id, err := strconv.Atoi(userValueID)
	if err != nil {
		return false, 0, false
	}
	return false, id, true
}

func Authenticate(c *fasthttp.RequestCtx, s sessions.Session) func() {
	if s.GetString("account") == "" {
		return func() {
			c.Redirect("/login", 302)
		}
	}
	return nil
}

func TryLogin(account, password string) *User {
	stmt := Must(
		db.Prepare("SELECT id, account, passhash FROM users WHERE account = ?"))
	defer func() { stmt.Close() }()
	u := &User{}
	stmt.QueryRow(account).Scan(&u.ID, &u.Account, &u.PassHash)

	if u.Account != "" && CalcPassHash(u.Account, password) == u.PassHash {
		return u
	}

	return nil
}

var validUserNameRe = regexp.MustCompile(`\A[0-9a-zA-Z_]{3,}\z`)
var validUserPassRe = regexp.MustCompile(`\A[0-9a-zA-Z_]{6,}\z`)

func ValidateUser(account, password string) bool {
	if !(validUserNameRe.MatchString(account) && validUserPassRe.MatchString(password)) {
		return false
	}
	return true
}

func CalcPassHash(account, password string) string {
	//OpenSSLから剥がす
	return Digest(password + ":" + CalcSalt(account))
}

func CalcSalt(str string) string {
	return Digest(str)
}

func Digest(str string) string {
	hash := sha512.Sum512([]byte(str))
	return hex.EncodeToString(hash[:])
}

func GetSessionUser(s sessions.Session) *User {
	if s.GetString("account") != "" {
		stmt := Must(db.Prepare("SELECT id, account, passhash WHERE id = ?"))
		defer func() { stmt.Close() }()
		u := &User{}
		stmt.QueryRow(&u.ID, &u.Account, &u.PassHash)
		return u
	}
	return nil
}

func GetTheme(id int) *Theme {
	stmt := Must(db.Prepare(
		"SELECT t.id, t.category_id, i.id, i.image " +
			"FROM themes AS t " +
			"LEFT JOIN images AS i ON t.image_id = i.id " +
			"WHERE t.id = ?",
	))
	defer func() { stmt.Close() }()
	t := &Theme{Image: &Image{}}
	stmt.QueryRow(id).Scan(&t.ID, &t.CategoryId, &t.Image.ID, &t.Image.File)
	return t
}

func CountJoke(themeID int) int {
	stmt := Must(db.Prepare(
		"SELECT COUNT(*) AS count FROM jokes WHERE theme_id = ?",
	))
	defer func() { stmt.Close() }()
	var count int
	stmt.QueryRow(themeID).Scan(&count)
	return count
}

func GetLatestJokeByThemeIDCategoryID(themeID, categoryID int) *Joke {
	stmt := Must(db.Prepare("SELECT id, content FROM " +
		"jokes WHERE theme_id = ? AND category_id = ? " +
		"ORDER BY created_at DESC LIMIT 1"))
	defer func() { stmt.Close() }()
	j := &Joke{}
	stmt.QueryRow(themeID, categoryID).Scan(&j.ID, &j.Content)
	return j
}

func GetLatestJokeByThemeID(themeID int) *Joke {
	stmt := Must(db.Prepare("SELECT id, content FROM " +
		"jokes WHERE theme_id = ? ORDER BY created_at DESC LIMIT 1"))
	defer func() { stmt.Close() }()
	j := &Joke{}
	stmt.QueryRow(themeID).Scan(&j.ID, &j.Content)
	return j
}

func ImgDir() string {
	return "/uploads"
}

func Star(rate int) string {
	white := "<img src=\"/img/white_star_16x16.png\">"
	black := "<img src=\"/img/black_star_16x16.png\">"
	switch rate {
	case 1:
		return black + white + white
	case 2:
		return black + black + white
	}
	//case 3:
	return black + black + black
}

func GetThemeCategories() []string {
	return ThemeCategories
}

func GetJokeCategories() []string {
	return JokeCategories
}

func CalcOffset(page int) int {
	if page > 1 {
		return (page - 1) * ContentsPerPage
	}
	return 0
}

func CreatePage(url string, page, contentsSize int) (string, string) {
	prev := ""
	next := ""

	if page == 0 {
		page = 1
	}

	if page > 1 {
		prev = url + "?page=" + strconv.Itoa(page-1)
	}

	if contentsSize == ContentsPerPage {
		next = url + "?page=" + strconv.Itoa(page+1)
	}

	return prev, next
}

func ValidateCsrfToken(csrfToken string, s sessions.Session) bool {
	return s.GetString("csrf_token") == csrfToken
}
