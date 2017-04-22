package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	"github.com/buaazp/fasthttprouter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/go-sessions"
	"github.com/kataras/go-sessions/sessiondb/redis"
	"github.com/kataras/go-sessions/sessiondb/redis/service"
	"github.com/valyala/fasthttp"
)

const (
	SessionName     = "rack.session"
	Port            = ":8080"
	ContentsPerPage = 20
)

var (
	baseUrl      *url.URL
	db           *sql.DB
	redisSession = redis.New(service.Config{
		Addr:     ":6379",
		Database: "0",
	})
	store sessions.Sessions

	ThemeCategories = []string{
		"人物",
		"人物2人以上",
		"動物",
		"風景",
		"無機物",
		"イラスト",
		"その他",
	}

	JokeCategories = []string{
		"バカ",
		"シュール",
		"ブラック",
		"身内",
		"例え",
		"その他",
	}
)

func main() {
	var e error
	if db, e = sql.Open("mysql", "root:@/joker2"); e != nil {
		panic(e)
	}
	store = sessions.New(sessions.Config{
		Cookie: SessionName,
	})
	store.UseDatabase(redisSession)

	r := fasthttprouter.New()
	r.GET("/", index)

	r.GET("/login", getLogin)
	r.POST("/login", postLogin)

	r.GET("/logout", getLogout)

	r.GET("/signup", getSignup)
	r.POST("/signup", postSignup)

	r.GET("/category/:id", getCategory)

	r.GET("/theme/:id" /*new*/, getThemeIDNew)
	r.POST("/theme/new", postThemeNew)

	r.GET("/joke/:id" /*new*/, getJokeIDNew)
	r.POST("/joke/new", postJokeNew)

	r.POST("/rating", postRating)

	if e := fasthttp.ListenAndServe(Port, r.Handler); e != nil {
		fmt.Println(e.Error())
	}
}

func index(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	_ = s
	for i := 0; i < 1000; i++ {
		c.WriteString("Hello, World!")
	}
}

func getLogin(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	_ = s

}

func postLogin(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	_ = s

	fmt.Printf("%#v\n", string(c.FormValue("account")))
	fmt.Printf("%#v\n", string(c.FormValue("password")))
}

func getLogout(c *fasthttp.RequestCtx) {
	//s := store.StartFasthttp(c)

	store.DestroyFasthttp(c)
	log.Printf("%#v", store.StartFasthttp(c))
}

func getSignup(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	_ = s

}

func postSignup(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	_ = s

}

func getCategory(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	_ = s

	//id, ok := ParamID(c)

}

func getThemeIDNew(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	if jump := Authenticate(c, s); jump != nil {
		jump()
		return
	}
}

func postThemeNew(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	if jump := Authenticate(c, s); jump != nil {
		jump()
		return
	}
}

func getJokeIDNew(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	_ = s
	if jump := Authenticate(c, s); jump != nil {
		jump()
		return
	}

	isNew, id, ok := ParamNew(c)

	if ok == false {
		c.WriteString(fmt.Sprintf("%#v\n", string(c.Referer())))
	}

	c.WriteString(fmt.Sprintf("isNew: %#v\n", isNew))
	c.WriteString(fmt.Sprintf("id: %#v\n", id))

	//c.WriteString(strconv.Itoa(id))
}

func postJokeNew(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	_ = s
	if jump := Authenticate(c, s); jump != nil {
		jump()
		return
	}

}

func postRating(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	_ = s
	if jump := Authenticate(c, s); jump != nil {
		jump()
		return
	}

}
