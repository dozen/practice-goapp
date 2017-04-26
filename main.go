package main

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/buaazp/fasthttprouter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/valyala/fasthttp"
	"html/template"
	"github.com/go-redis/redis"
	"strconv"
	"bytes"
)

const (
	SessionName     = "rack.session"
	Port            = ":8081"
	ContentsPerPage = 200
	TplDir 		= "tpl/"
)

var (
	baseUrl      *url.URL
	db           *sql.DB
	redi = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",

	})

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
	tpl      *template.Template
)

func InitTemplates() *template.Template {
	       return template.Must(template.New("tmpl").Funcs(template.FuncMap{
		       "joke_count": func(themeID int) string {
				return redi.Get("joke_count:" + strconv.Itoa(themeID)).String()
		       },
		       }).ParseGlob(TplDir + "*.tpl"))
	}

func main() {
	var e error
	if db, e = sql.Open("mysql", "root:@/joker2"); e != nil {
		panic(e)
	}

	tpl = InitTemplates()

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
	c.SetContentType("text/html")

	buf := bytes.NewBuffer([]byte{})

	var page = 0
	if pageString, ok := c.UserValue("page").(int); ok {
		page = pageString
	}

	query :=
	"SELECT t1.id, " +
		"i.image, " +
		"iu.account AS iu_account, " +
		"tu.account AS tu_account, " +
		"(SELECT content FROM jokes AS j WHERE j.theme_id = t1.id ORDER BY created_at DESC LIMIT 1) AS content " +
		"FROM " +
		"( SELECT t.id, t.user_id, t.image_id FROM themes AS t ORDER BY t.created_at, id DESC LIMIT ? OFFSET ? ) AS t1 " +
		" JOIN images AS i ON t1.image_id = i.id " +
		" JOIN users AS tu ON t1.user_id = tu.id" +
		" JOIN users AS iu ON i.user_id = iu.id"

	result, e := Must(db.Prepare(query)).Query(ContentsPerPage, page)
	defer func() {
		result.Close()
	}()
	if e != nil {
		fmt.Errorf(e.Error())
		return
	}

	var t []map[string]interface{}

	for i:= 0;; i++ {
		if !result.Next() {
			break
		}
		var (
			id int
			iuAccount string
			image string
			tuAccount string
			content string
		)
		result.Scan(
			&id,
			&image,
			&iuAccount,
			&tuAccount,
			&content,
		)
		t = append(t, map[string]interface{}{
			"id": id, "image": image, "iu_account": iuAccount, "tu_account": tuAccount, "content": content,
		})
	}

	fmt.Printf("%#v\n", t)

	prev, next := CreatePage(string(c.Referer()), page, ContentsPerPage)
	tpl.ExecuteTemplate(buf, "index", map[string]interface{}{
		"Themes": t,
		"Prev": prev,
		"Next": next,
	})
	c.Write(buf.Bytes())
}

func getLogin(c *fasthttp.RequestCtx) {

}

func postLogin(c *fasthttp.RequestCtx) {

	fmt.Printf("%#v\n", string(c.FormValue("account")))
	fmt.Printf("%#v\n", string(c.FormValue("password")))
}

func getLogout(c *fasthttp.RequestCtx) {
	//s := store.StartFasthttp(c)

}

func getSignup(c *fasthttp.RequestCtx) {

}

func postSignup(c *fasthttp.RequestCtx) {

}

func getCategory(c *fasthttp.RequestCtx) {

	//id, ok := ParamID(c)

}

func getThemeIDNew(c *fasthttp.RequestCtx) {
}

func postThemeNew(c *fasthttp.RequestCtx) {
}

func getJokeIDNew(c *fasthttp.RequestCtx) {

	isNew, id, ok := ParamNew(c)

	if ok == false {
		c.WriteString(fmt.Sprintf("%#v\n", string(c.Referer())))
	}

	c.WriteString(fmt.Sprintf("isNew: %#v\n", isNew))
	c.WriteString(fmt.Sprintf("id: %#v\n", id))

	//c.WriteString(strconv.Itoa(id))
}

func postJokeNew(c *fasthttp.RequestCtx) {
}

func postRating(c *fasthttp.RequestCtx) {
}
