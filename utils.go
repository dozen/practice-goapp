package main

import (
	"github.com/valyala/fasthttp"
	"strconv"
	"github.com/kataras/go-sessions"
	"database/sql"
	"crypto/sha512"
	"encoding/hex"
)

func MustPrepare(stmt *sql.Stmt, err error) *sql.Stmt {
	if err != nil {
		panic(err)
	}
	return stmt
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
	stmt := MustPrepare(
		db.Prepare("SELECT id, account, passhash FROM users WHERE account = ?",
	))
	u := &User{}
	stmt.QueryRow(account).Scan(u.ID, u.Account, u.PassHash)

	if u.Account != "" && CalcPassHash(u.Account, password) == u.PassHash {
		return u
	}

	return nil
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