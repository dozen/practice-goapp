package main

import (
	"database/sql"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/go-sessions"
	"github.com/kataras/go-sessions/sessiondb/redis"
	"github.com/kataras/go-sessions/sessiondb/redis/service"
	"github.com/valyala/fasthttp"
	"net/url"
	"strconv"
)

const (
	SessionName = "isucon_session"
	Port        = ":8080"
)

var (
	baseUrl      *url.URL
	db           *sql.DB
	redisSession = redis.New(service.Config{Addr: "192.168.99.100:6379"})
	store        sessions.Sessions
)

func main() {
	var err error
	db, err = sql.Open("mysql", "root:@/sample")
	if err != nil {
		panic(err)
	}
	store = sessions.New(sessions.Config{Cookie:SessionName})
	store.UseDatabase(redisSession)

	r := fasthttprouter.New()
	r.GET("/", initial(index))
	r.GET("/login", initial(login))
	r.POST("/login", initial(postLogin))
	r.GET("/logout", initial(logout))
	r.GET("/signup", initial(signup))
	r.POST("/signup", initial(postSignup))
	r.GET("/category/:id", initial(category))
	r.GET("/theme/:id", initial(theme))
	r.POST("/theme/new", initial(postNewTheme))

	if err := fasthttp.ListenAndServe(Port, r.Handler); err != nil {
		fmt.Println(err.Error())
	}
}

func initial(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(c *fasthttp.RequestCtx) {
		_ = store.StartFasthttp(c)
		handler(c)
	}
}

func index(c *fasthttp.RequestCtx) {
	c.WriteString("Hello, World!")
}

func login(c *fasthttp.RequestCtx) {

}

func postLogin(c *fasthttp.RequestCtx) {

}

func logout(c *fasthttp.RequestCtx) {

}

func register(c *fasthttp.RequestCtx) {

}

func signup(c *fasthttp.RequestCtx) {

}

func postSignup(c *fasthttp.RequestCtx) {

}

func category(c *fasthttp.RequestCtx) {

}

func theme(c *fasthttp.RequestCtx) {
	var id int
	if val, ok := c.UserValue("id").(string); !ok {
		c.NotFound()
	} else {
		if val == "new" {
			newTheme(c)
			return
		} else {
			var err error
			id, err = strconv.Atoi(val)
			if err != nil {
				c.NotFound()
				return
			}
		}
	}

	c.WriteString(strconv.Itoa(id))
}

func newTheme(c *fasthttp.RequestCtx) {
	c.WriteString("New Theme Create.")
}

func postNewTheme(c *fasthttp.RequestCtx) {

}

func joke(c *fasthttp.RequestCtx) {
	var id int
	if val, ok := c.UserValue("id").(string); !ok {
		c.NotFound()
	} else {
		if val == "new" {
			newJoke(c)
			return
		} else {
			var err error
			id, err = strconv.Atoi(val)
			if err != nil {
				c.NotFound()
				return
			}
		}
	}

	c.WriteString(strconv.Itoa(id))
}

func newJoke(c *fasthttp.RequestCtx) {

}

func postNewJoke(c *fasthttp.RequestCtx) {

}

func rate(c *fasthttp.RequestCtx) {
	var id int
	if val, ok := c.UserValue("id").(string); !ok {
		c.NotFound()
	} else {
		if val == "new" {
			newRate(c)
			return
		} else {
			var err error
			id, err = strconv.Atoi(val)
			if err != nil {
				c.NotFound()
				return
			}
		}
	}

	c.WriteString(strconv.Itoa(id))
}

func newRate(c *fasthttp.RequestCtx) {

}

func postRate(c *fasthttp.RequestCtx) {

}
