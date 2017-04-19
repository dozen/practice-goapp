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
	var e error
	if db, e = sql.Open("mysql", "root:@/sample"); e != nil {
		panic(e)
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
	r.GET("/theme/:id", initial(theme))		//new含む
	r.POST("/theme/new", initial(postTheme))
	r.GET("/joke/:id", initial(joke))			//new含む
	r.POST("/joke/new", initial(postJoke))
	r.GET("/rate/new", initial(newRate))
	r.POST("/rate/new", initial(postRate))

	if e := fasthttp.ListenAndServe(Port, r.Handler); e != nil {
		fmt.Println(e.Error())
	}
}

func initial(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(c *fasthttp.RequestCtx) {
		handler(c)
	}
}

func index(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	_ = s

	c.WriteString("Hello, World!")
}

func login(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	_ = s

}

func postLogin(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	_ = s

	fmt.Printf("%#v\n", string(c.FormValue("account")))
	fmt.Printf("%#v\n", string(c.FormValue("password")))
}

func logout(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	_ = s

}

func signup(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	_ = s

}

func postSignup(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	_ = s

}

func category(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	_ = s
}

func theme(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	_ = s

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
	s := store.StartFasthttp(c)
	_ = s

	c.WriteString("New Theme Create.")
}

func postTheme(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	_ = s

}

func joke(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	_ = s

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
	s := store.StartFasthttp(c)
	_ = s

}

func postJoke(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	_ = s

}

func rate(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	_ = s

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
	s := store.StartFasthttp(c)
	_ = s

}

func postRate(c *fasthttp.RequestCtx) {
	s := store.StartFasthttp(c)
	_ = s

}
