package server

import(
	"time"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/binding"
)

<<<<<<< HEAD
type Post struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
=======

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		var loginPage string = os.Getenv("GOPATH") + "github.com/zwx2000/cloudgo/resource/login.gtpl"
		t, _ := template.ParseFiles(loginPage)
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Fprintf(w, "username: %s\n", r.Form["username"])
		fmt.Fprintf(w, "password: %s\n", r.Form["password"])
	}
>>>>>>> e7c00d7f590522d8e8d696d5129959f870e80c00
}

func Begin(port string) { 

	//默认设置
	defaultConfig := martini.Classic()
	
	//设置资源文件的路径
	defaultConfig.Use(martini.Static("assets"))
	//设置渲染html模板的包
	defaultConfig.Use(render.Renderer())
	
	//GET请求
	defaultConfig.Get("/", func(r render.Render) {
		//获取当前系统时间
		curTime:=time.Now().Format("2006-01-02 15:04:05")
		//渲染index模板，将时间传过去
        r.HTML(200, "index", map[string]interface{}{"Time":curTime})
	})

	//POST请求
	defaultConfig.Post("/", binding.Bind(Post{}), func(post Post, r render.Render) {
		curTime:=time.Now().Format("2006-01-02 15:04:05")
		p := Post{Username: post.Username, Password: post.Password}
		//渲染index模板，将用户信息和时间传过去
        r.HTML(200, "info", map[string]interface{}{"post": p, "Time":curTime})
	})

	//开始运行
	defaultConfig.RunOnAddr(":"+port)
}