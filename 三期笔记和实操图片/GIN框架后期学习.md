# $\color{#5AF}{||GIN框架学习||}$

---

前部分学习以图片形式上传在仓库，后半截寒假学习如下。

md笔记基本全部手打。

---

学习的gin课程链接：【最新Go Web开发教程】基于gin框架和gorm的web开发实战 (七米出品)

[lesson01_内容介绍_哔哩哔哩_bilibili](https://www.bilibili.com/video/BV1gJ411p7xC?p=1&vd_source=16d869afecffa8da2167060c3de0b7c6)

资料[【置顶】Go语言学习之路/Go语言教程 | 李文周的博客 (liwenzhou.com)](https://www.liwenzhou.com/posts/Go/golang-menu/)

---

## 一、模板补充

### （1）修改默认的标识符

```go
package main

import(
    "fmt"
    "net/http"
    "html/template"
)

func index(w http.ResponseWriter,r *http.Request){
    //定义模板
    //解析模板
    t,err := template.New("index.tmpl").
    Delims("{[","]}").
    ParseFiles("./web08/index.tmpl")

    if err != nil{
        fmt.Printf("parse template failed,err:%v\n",err)
    }
    //渲染模板
    name := "寒假第一战"
    err = t.Execute(w,name)
    if err != nil{
        fmt.Printf("execute template failed,err:%v\n",err)
        return
    }


}

func main(){
    http.HandleFunc("/index", index)
    err := http.ListenAndServe(":9000", nil)
    if err != nil {
        fmt.Printf("http server failed, err:%v\n", err)
        return
    }
}
```

```html
<!DOCTYPE html>
<html lang="en">
<head>
   <title>修改模板引擎的标识符</title>
</head>
<body>
    <div>Hello {[ . ]}</div>
</body>
</html>
```

放在一个文件夹下运行两个文件，可以得到单独一行的输出

```
Hello 寒假第一战
```

这里是在修改默认的标识符

> Go标准库的模板引擎使用的花括号`{{`和`}}`作为标识，而许多前端框架（如`Vue`和 `AngularJS`）也使用`{{`和`}}`作为标识符，所以当我们同时使用Go语言模板引擎和以上前端框架时就会出现冲突，这个时候我们需要修改标识符，修改前端的或者修改Go语言的。这里演示如何修改Go语言模板引擎默认的标识符：
> 
> > ```go
> > template.New("test").Delims("{[", "]}").ParseFiles("./t.tmpl")
> > ```
> > 
> > 与原来比较多了New和Delims
> > 
> > 对应的我代码里其实是把{{}}变成了{[]}，以另一种形式解析模板
> > 
> > Delims是一个函数

---

### （2）text/template与html/tempalte的区别

`html/template`针对的是需要返回HTML内容的场景，在模板渲染过程中会对一些有风险的内容进行<mark>转义</mark>，以此来防范跨站脚本攻击。

其中常用就是`<script>alert(123)</script>`

这个时候传入一段JS代码并使用`html/template`去渲染该文件，会在页面上显示出<mark>转义后的JS内容</mark>。 `<script>alert('嘿嘿嘿')</script>` 这就是`html/template`为我们做的事。防止被无限循环等破坏服务器。

如果要相信用户输入的内容，不想被一起转义掉，就自己编写safe函数，手动返回一个`template.HTML`类型的内容。

```go
package main

import (
    "fmt"
    "html/template"
    "net/http"
)



func xss(w http.ResponseWriter, r *http.Request) {
    //定义模板

    //解析模板
    //解析模板之前定义一个自定义的函数safe再解析
    t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
        "safe": func(str string) template.HTML {
            return template.HTML(str)
        },
    }).ParseFiles("./web08/xss.tmpl")

    if err != nil {
        fmt.Printf("parse template failed,err:%v\n", err)
        return
    }
    //渲染模板
    str1 := "<script>alert(123);</script>"
    str2 := "<a href='http://hdufhq.cn:8888/'>hdu孵化器地址</a>"
    t.Execute(w, map[string]string{
        "str1": str1,
        "str2": str2,
    })
}

func main() {
    http.HandleFunc("/xss", xss)
    err := http.ListenAndServe(":9000", nil)
    if err != nil {
        fmt.Printf("http server failed, err:%v\n", err)
        return
    }
}
```

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>xss</title>
</head>
<body>
    用户的评论是：{{.str1}}
    用户的评论是：{{.str2 | safe}}
</body>
</html>
```

这里的safe操作`{{. | safe}}`就是让用户输入不被转义进行调用自己编写的函数的写法。好像也可以`{{safe .}}`这么写

解析模板的地方先新建了xss.tmpl的模板，然后Funcs调用函数，然后定义一个函数映射集合，里面定义了safe自定义函数，用来返回我们str内容变成HTML的string，在最后才解析模板。

渲染的时候使用了映射，字符串对应字符串，可以多个数据传输过去。

这个第二个part主要讲的是go的包template自带防止被脚本攻击的转义功能，如果想要不转义，就自己编写函数并在tmpl文件里调用它。

最后的结果就可以变成如下：

**用户的评论是：<script>alert(123);</script> 用户的评论是：[hdu孵化器地址](http://hdufhq.cn:8888/)**

前面是转义结果，后面是避免了转移，自带了链接（夹带私货证明自己学的bushi）。

---

---

## 二、gin框架模板渲染

#### （1）首个实例

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main(){
    //创建一个默认路由
    r := gin.Default()

    //解析模板
    r.LoadHTMLFiles("./web09/templates/index.tmpl")

    //渲染模板
    r.GET("/index",func(c *gin.Context){

        //HTTP请求要有状态码的响应被返回 三次握手 渲染上数据
        c.HTML(http.StatusOK,"index.tmpl",gin.H{
            "title" : "http://hdufhq.cn:8888/",
        })

    })

    //启动server 这里浏览器就是客户端 本机是服务器
    r.Run(":9090")
}
```

所有笔记都在注释里了，一步步走。

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>posts/index</title>
</head>
<body>
    {{.title}}
</body>
</html>
```

搭配的模板里只需要在点号传入title数据就可以了。

Gin框架中使用`LoadHTMLGlob()`或者`LoadHTMLFiles()`方法进行HTML模板渲染。

得到的页面结果是`http://hdufhq.cn:8888/`

---

### （2）多个模板一起渲染的情况

首先要先在templates文件夹下新建posts和users文件夹，都各自有一个index.tmpl文件  

所以要区分这两个相同的文件，可以在开头加入语句`{{define "users/index.tmpl"}}`(若是posts文件夹就换一下)，在结尾加上`{{end}}`

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main(){
    //创建一个默认路由
    r := gin.Default()

    //解析模板
    //旧方法解析r.LoadHTMLFiles("./web09/templates/index.tmpl")

    r.LoadHTMLGlob("web09/***/**/*") //意思是直接加载这个web09下面的所有文件 不用具体指定路径了


    //渲染模板  在这里写上上面星号的相对路径
    r.GET("posts/index",func(c *gin.Context){

        //HTTP请求要有状态码的响应被返回 三次握手 渲染上数据
        c.HTML(http.StatusOK,"index.tmpl",gin.H{
            "title" : "http://hdufhq.cn:8888/",
        })//不可行的写法，无法展示title内容，应该是找不到文件

    })
    r.GET("users/index",func(c *gin.Context){

        //HTTP请求要有状态码的响应被返回 三次握手 渲染上数据
        c.HTML(http.StatusOK,"users/index.tmpl",gin.H{
            "title" : "https://www.hdu.edu.cn/main.htm",
        })//可行写法

    })

    //启动server 这里浏览器就是客户端 本机是服务器
    r.Run(":9090")
}
```

这里出现了一个很好玩的情况：

视频里是只解析到二级目录，因为仓库关系我没有新建文件夹，所以在原来文件夹里写，多了一层文件壳，我就尝试了一下用三个*来解析web09下的所有文件，结果是可行的！！虽然本来gpt跟我说最多到两级。

有个问题是，我调试的时候发现post这种HTML的name写法是读不出title的，虽然还能GET但是没有内容。但是下面那种写法是可以正常展现title的，我觉得可能是为了区分相同文件名也就是index这个文件需要至少外写一层文件夹。（后来发现可能是因为文件首行里手写define了

`https://www.hdu.edu.cn/main.htm` 的title内容在页面路由为`http://127.0.0.1:9090/users/index`上正常展现

GET那个路径纯粹是路由相对路径，下面c.HTML函数才是需要找到自己文件的。

> 改进写法
> 
> ```go
> r.GET("users/index", func(c *gin.Context) {
> 
>         //HTTP请求要有状态码的响应被返回 三次握手 渲染上数据
>         c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
>             "title": "<a href='https://www.hdu.edu.cn/main.htm'>杭电地址</a>",
>         })
> 
>      })
> ```

> 这个写法的显示结果是单纯文字，显然是被当成风险转义了，不想转义，让它变成可点击链接就得继续改写。

改写过程遇到一个问题就是引包。

要区分这两个，否则template.HTML会显示是undefined。

> `html/template`和`text/template`是Go语言中用于处理模板的两个包，它们之间的区别在于对待输出的方式。
> 
> 1. `text/template`：这个包主要用于处理纯文本模板，不会对输出进行特殊处理。所有的变量在输出时都会被转义，以确保安全性。这意味着如果您在模板中使用类似`{{ .Content }}`的表达式，其中`.Content`包含HTML代码，那么输出将会把HTML标签转义为实体字符，而不会被当作HTML代码渲染。
> 
> 2. `html/template`：相比于`text/template`，`html/template`提供了更多的HTML相关功能。它会自动对输出进行HTML转义，以避免XSS（跨站脚本攻击）等安全问题。在`html/template`中，通过使用`template.HTML`类型来指示某个输出是安全的HTML内容，不需要进行转义。您可以使用`template.FuncMap`注册自定义函数，并在函数中返回`template.HTML`类型的值。

在避免转移时候用的是第二个加上自定义函数，前面编写用的是第一个。

```go
package main

import (
    "net/http"

    "html/template"
    "github.com/gin-gonic/gin"
)

func main() {
    //创建一个默认路由
    r := gin.Default()

    //解析模板
    //先给gin框架中模板添加自定义函数
    r.SetFuncMap(template.FuncMap{
        "safe" : func(str string) template.HTML{
            return template.HTML(str)
        },
    })

    //旧方法解析r.LoadHTMLFiles("./web09/templates/index.tmpl")

    r.LoadHTMLGlob("web09/***/**/*") //意思是直接加载这个web09下面的所有文件 不用具体指定路径了

    //渲染模板  在这里写上上面星号的相对路径
    r.GET("posts/index", func(c *gin.Context) {

        //HTTP请求要有状态码的响应被返回 三次握手 渲染上数据
        c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
            "title": "<a href='http://hdufhq.cn:8888/'>孵化器地址</a>",
        })

    })
    r.GET("users/index", func(c *gin.Context) {

        //HTTP请求要有状态码的响应被返回 三次握手 渲染上数据
        c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
            "title": "<a href='https://www.hdu.edu.cn/main.htm'>杭电地址</a>",
        })

    })

    //启动server 这里浏览器就是客户端 本机是服务器
    r.Run(":9090")
}
na
```

那边文件里写成`{{.title | safe}}`可以展示为两个链接

---

### （3）静态文件处理

静态文件：`html页面上用到的样式文件 如css js文件 图片`

当我们渲染的HTML文件中引用了静态文件时，我们只需要按照以下方式在渲染页面前调用`gin.Static`方法即可。

在解析之前插入

```go
//加载静态文件 意思是直接去右边这个目录找文件 
//然后在浏览器f12打开的代码是左边这个链接，右边是被指代隐藏了吧
    r.Static("/xxx","./web09/statics")
```

```html
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="/xxx/index.css">
    <title>users/index</title>
</head>
```

用link来插入连接的样式表

我感觉/xxx这个就是在指代index的长长的文件夹路径，方便前端链接书写。

> `r` 是一个 Gin 的实例，用于处理 HTTP 请求和路由。
> 
> `.Static()` 是 Gin 框架提供的方法之一，用于配置静态文件服务。
> 
> `"/xxx"` 是在服务器上<mark>暴露的 URL 路径，即访问静态文件时需要使用的路径</mark>。例如，如果设置为 `/xxx`，那么访问静态文件时的 URL 就是 `http://yourdomain/xxx/filename.ext`。
> 
> `"./web09/statics"` 是本地文件系统上的静态文件目录路径。在这个例子中，静态文件存储在项目根目录下的 "web09/statics" 文件夹中。
> 
> 通过以上代码，当用户在浏览器中请求 "/xxx/filename.ext" 时，Gin 框架会自动去 "./web09/statics/filename.ext" 路径下寻找对应的静态文件，并将其返回给客户端。
> 
> 请注意，`/xxx` 可以自定义为其他路径，而 `./web09/statics` 则应该根据实际情况修改为您的静态文件目录路径。

.

> 其中：终端
> 
> [GIN] 2024/02/15 - 22:55:28 | 200 |      1.0184ms |       127.0.0.1 | GET      "/users/index"
> [GIN] 2024/02/15 - 22:55:28 | 200 |     74.0016ms |       127.0.0.1 | GET      "/xxx/index.css"
> 
> 浏览器访问的时候发送了两次请求，在渲染的时候额外需要一个css，浏览器会又发一次请求获得css。

若添加js文件

```js
alert(123);
```

```html
<body>
    {{.title | safe}}

    <script src="/xxx/index.js"></script>

</body>
```

会在页面上先弹窗出123的响应，可见上传附带图片4。

> [GIN] 2024/02/16 - 21:02:37 | 200 |      2.2963ms |       127.0.0.1 | GET      "/users/index"
> [GIN] 2024/02/16 - 21:02:38 | 200 |     77.6227ms |       127.0.0.1 | GET      "/xxx/index.js"
> [GIN] 2024/02/16 - 21:02:38 | 200 |     93.6283ms |       127.0.0.1 | GET      "/xxx/index.css"

---

### （4）使用模板继承

> Gin框架默认都是使用单模板，如果需要使用`block template`功能，可以通过`"github.com/gin-contrib/multitemplate"`库实现，具体示例如下：
> 
> 首先，假设我们项目目录下的templates文件夹下有以下模板文件，其中`home.tmpl`和`index.tmpl`继承了`base.tmpl`：
> 
> ```bash
> templates
> ├── includes
> │   ├── home.tmpl
> │   └── index.tmpl
> ├── layouts
> │   └── base.tmpl
> └── scripts.tmpl
> ```

实名感谢学长萌的指正！！！！i死！！st哥太帅了！

犯了通配符的路径的错误，搞了半天搜不到解决方法，找了哥，最后可能还是glob函数使用不能三个*的原因。

```go
package main

import (
    "net/http"

    "html/template"
    "github.com/gin-gonic/gin"
)

func main() {
    gin.SetMode(gin.ReleaseMode)
    //创建一个默认路由
    r := gin.Default()

    //解析模板

       //加载静态文件
    r.Static("statics","/statics")

    //先给gin框架中模板添加自定义函数
    r.SetFuncMap(template.FuncMap{
        "safe" : func(str string) template.HTML{
            return template.HTML(str)
        },
    })


    //旧方法解析r.LoadHTMLFiles("./web09/templates/index.tmpl")

    r.LoadHTMLGlob("templates/**/*") //意思是直接加载这个web09下面的所有文件 不用具体指定路径了

    //渲染模板  在这里写上浏览器上url后小截
    r.GET("/posts/index", func(c *gin.Context) {

        //HTTP请求要有状态码的响应被返回 三次握手 渲染上数据
        c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
            "title": "<a href='http://hdufhq.cn:8888/'>孵化器地址</a>",
        })

    })
    r.GET("/users/index", func(c *gin.Context) {

        //HTTP请求要有状态码的响应被返回 三次握手 渲染上数据
        c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
            "title": "<a href='https://www.hdu.edu.cn/main.htm'>杭电地址</a>",
        })

    })

    //返回从网上下载的模板,没有需要上传的数据，所以是nil
    r.GET("/home",func(c *gin.Context){
        c.HTML(http.StatusOK,"posts/home.html",nil)
    })

    //启动server 这里浏览器就是客户端 本机是服务器
    r.Run(":9090")
}
```

把这个整个文件夹拿出来单独运行，保证最多就两级目录，这样就可以运行了，gin框架安装和初始没有问题。

这样就可以实现前端模板的一个全栈应用了。

---

---

## 三、gin框架返回json

方法一：映射和gin.H

```go
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main(){

    r:= gin.Default()

    r.GET("/json",func(c *gin.Context){
        //方法一：使用map
        data:= map[string]interface{}{
            "name":"Leeson",
            "message":"from Cambridge",
            "age":24,
        }
        c.JSON(http.StatusOK,data)
    })

    r.Run(":9090")
}
```

结果反应在图片5，就是返回一串json格式的数据

```go
//方法1其他写法：gin.H
        data := gin.H{
            "name":"Leeson",
            "message":"from Cambridge",
            "age":24,
        }
```

这个H是一个type，就是上面方法1一样的结果。

方法二：**定义结构体（项目常用）**

```go
//方法二：结构体
    type msg struct {
        Name    string
        Message string
        Age     int
    }

    r.GET("/json2", func(c *gin.Context) {
        data := msg{
            "fhq",
            "lab",
            999999,
        }

        c.JSON(http.StatusOK, data)
    })
```

返回结果：

```json
{
    "Name": "fhq",
    "Message": "lab",
    "Age": 999999
}
```

最后一个没有逗号。

而且结构体里面的变量必须必须是首字母大写的！！！！否则会丢失！

因为json的序列化，默认是使用go语言标准库的，反射取值，go里面函数方法啥的有首字母小写会导致无法导出，相当于没有这个字段。

如果一定要返回小写，那么就要用tag，进行自定义，json的包做反射操作。

```go
//方法二：结构体 灵活使用tag来做定制化操作
    type msg struct {
        Name    string `json:"name"`
        Message string
        Age     int
    }

    r.GET("/json2", func(c *gin.Context) {
        data := msg{
            "fhq",
            "lab",
            999999,
        }

        c.JSON(http.StatusOK, data)
    })
```

注意，这里是反引号哦，单引号是不对的。

在Go语言中，结构体字段的标签需要使用反引号（`）而不是单引号（'）。

> 在Go语言中，反引号（`）是一种用于表示原始字符串的字符。使用反引号定义的字符串会保留其原始格式，其中包含的任何特殊字符和转义序列都会被直接输出。例如：
> 
> ``str := `Hello\nWorld``` 
> 
> `fmt.Println(str)`
> 
> 以上代码输出的结果是：
> 
> `Hello\nWorld`
> 
> 可以看到，使用反引号定义的字符串中的`\n`并没有被转义成一个换行符，而是直接输出了它本身。
> 
> 在Go语言中，反引号还被广泛用于定义结构体字段标签，如我之前所提到的。反引号可以让开发者自定义结构体字段的元数据，例如字段的名称、长度、类型等等，这些元数据可以在后续的代码中被读取和使用。

---

---

## 四、gin获取querystring参数

`https://www.sogou.com/web?query=%E9%87%8C%E6%A3%AE`

搜索是额外附加信息。问号前面是我们的url，问号+什么等于什么的格式

```go
package main

//querystring
import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main(){
    r := gin.Default()

    r.GET("/web",func(c *gin.Context){
        //获取浏览器那边发请求携带的 query string 参数
        c.JSON(http.StatusOK,"ok")
    })

    r.Run(":9090")
}
```

得到：

`"ok"`

如果写入为`127.0.0.1:9090/web?query=里森`返回还是ok，这个时候我们需要的到里森的相关结果

修改：

```go
    r.GET("/web",func(c *gin.Context){
        //获取浏览器那边发请求携带的 query string 参数
        name := c.Query("query")  
        //通过Query获取请求中携带的querystring参数
        c.JSON(http.StatusOK,gin.H{
            "name": name,

        })
    })
```

可以得到的是格式下各种输入对应的H形式，比如：

```json
{
    "name": "里森"
}
```

```json
{
    "name": "泥电"
}
```

还可以添加默认值，如果查不到，就显示替换的默认值，这里是somebody。

```go
name := c.DefaultQuery("query","somebody")  //娶不到就用指定的默认值
```

第三种就是用getquery，因为这个函数会返回一个布尔值，就用ok来接收，如果接收不到布尔值，就false然后做下面if，返回指定的默认值。

```go
name,ok := c.GetQuery("query")
if !ok {
    //取不到
    name ="somebody"
}
```

> 开始叠加(运用英文的and符号也就是&)
> 
> 如`9090/web?query=里森&age=23`
> 
> 代码修改为：
> 
> ```go
>     r := gin.Default()
> 
>     // GET请求 URL ？后面的是querystring参数
>     // key=value格式，多个key-value用 & 链接
>     // eq : /web?query=xxxxxx&age=xxx
> 
>     r.GET("/web",func(c *gin.Context){
>         //获取浏览器那边发请求携带的 query string 参数
>         name := c.Query("query")  
>         age := c.Query("age")
>         //通过Query获取请求中携带的querystring参数
>         //name := c.DefaultQuery("query","somebody") 
>         /*name,ok := c.GetQuery("query")
>         if !ok {
>             //取不到
>             name ="somebody"
>         }*/
> 
>         c.JSON(http.StatusOK,gin.H{
>             "name": name,
>             "age":age,
>         })
>     })
> 
>     r.Run(":9090")
> ```
> 
> 那么显示：返回出来注意是无序的。
> 
> ```json
> {
>     "age": "23",
>     "name": "里森"
> }
> ```

---

---

## 五、gin获取form参数

```go
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)
func main() {
    r := gin.Default()

    r.LoadHTMLFiles("./web12/login.html")

    r.GET("/login",func(c *gin.Context){
        c.HTML(http.StatusOK,"login.html",nil)
    })

    r.Run(":9090")
}
```

```html
<body>


<form action="/login" method="post">
    <label for="username">username:</label>
    <input type="text" name="username" id="username">

    <label for="password">password:</label>
    <input type="password" name="password" id="password">

    <input type="button" value="登录">
</form>


</body>
</html>
```

显示结果见图片6。

如果想要换行显示：(外面套层div就行

```html
<form action="/login" method="post" novalidate autocomplete="off">
    <div>
    <label for="username">username:</label>
    <input type="text" name="username" id="username">
    </div>

    <div>
    <label for="password">password:</label>
    <input type="password" name="password" id="password">
    </div>

    <div>
    <input type="button" value="登录">
    </div>
</form>
```

form表单通常用post方法发送请求。

如果把type=“button”改成submit，那么action，你提交就会跳转到/login路由。

但是改成submit，<u>由于我们后端还没有编写访问login的post请求</u>，只有前面进入页面的get请求，所以我们如果写了登录和密码会跳转到404页面。

**一次请求对应一个响应**！！

这里没有改两个url，一直在login里，注意区分不同请求！！！

> id属性一般是前端操作这些标签用的标记，form表单提交的键值对的键是name属性。

```go
    r := gin.Default()

    r.LoadHTMLFiles("./web12/login.html","./web12/index.html")

    r.GET("/login", func(c *gin.Context) {
        c.HTML(http.StatusOK, "login.html", nil)
    })

    // 处理这个访问login请求，上面只是进入到页面的get请求

    r.POST("/login", func(c *gin.Context) {
        username := c.PostForm("username")
        password := c.PostForm("password")
        c.HTML(http.StatusOK, "index.html", gin.H{
            "Name":     username,
            "Password": password,
        })
    })
```

```html
    <h1>Helloooooo, {{.Name}} !</h1>
    <p>你的密码是： {{.Password}}</p>
```

```html
<form action="/login" method="post" novalidate autocomplete="off">
    <div>
    <label for="username">username:</label>
    <input type="text" name="username" id="username">
    </div>

    <div>
    <label for="password">password:</label>
    <input type="password" name="password" id="password">
    </div>

    <div>
    <input type="submit" value="登录">
    </div>
</form>
```

> 首先，以get方式我们进入到login页面，进行输入。
> 
> 然后，在我们输入用户名和密码点击登录，submit以后，我们就通过表单的action而且是POST方式进入到login新页面，是新的内容不再是登陆界面，这是这次响应对应的请求。
> 
> 再者，根据标签输入的信息，传输到后端，再传输到页面上，显示可见图片7。

> // 注释  
> Ctrl+K+C  
> 取消//注释  
> Ctrl+K+U  
> /* */ 多行注释  
> Alt + Shift +A  
> 取消多行注释同上

**加入第三种接收方法 会有返回值ok**

```go
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.LoadHTMLFiles("./web12/login.html","./web12/index.html")

    r.GET("/login", func(c *gin.Context) {
        c.HTML(http.StatusOK, "login.html", nil)
    })

    // 处理这个访问login请求，上面只是进入到页面的get请求

    r.POST("/login", func(c *gin.Context) {
        //获取form表达的提交的数据
        /*方法一
        username := c.PostForm("username")
        password := c.PostForm("password")*/

        /*方法二  
        username :=c.DefaultPostForm("username","somebody")
        password :=c.DefaultPostForm("password","***") */

        //方法三
        username ,ok:= c.GetPostForm("username")
        if !ok{
            username = "preciouswxe"
        }
        password ,ok:=c.GetPostForm("password") //ok被重新创建
        if !ok{
            password = "i don't know"
        }

        c.HTML(http.StatusOK, "index.html", gin.H{
            "Name":     username,
            "Password": password,
        })
    })

    r.Run(":9090")
}
```

---

---

## 六、gin获取URL路径参数

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)
//获取请求的path（URI）参数

func main() {
    r := gin.Default()

    r.GET("/:name/:age", func(c *gin.Context) {
        //获取路径参数
        name := c.Param("name")
        age := c.Param("age")

        c.JSON(http.StatusOK,gin.H{
            "name":name,
            "age":age,
        })
    })

    r.Run(":9090")
}
```

主打一个读取输入的url对应的值，并在页面上展示，见如下（返回json格式的内容）或者图片8。也可以试试修改 变成各类字符串

```json
{
    "age": "18",
    "name": "leeson"
}
```

接下来再添加路由，在旧版本golang里，会因为上面的路由指定导致下面冲突报错，保险起见要像下面这么写。

```go
func main() {
    r := gin.Default()

    r.GET("/user/:name/:age", func(c *gin.Context) {
        //获取路径参数
        name := c.Param("name")
        age := c.Param("age")

        c.JSON(http.StatusOK, gin.H{
            "name": name,
            "age":  age,
        })
    })

    r.GET("/blog/:year/:month", func(c *gin.Context) {
        year := c.Param("year")
        month := c.Param("month")
        c.JSON(http.StatusOK, gin.H{
            "year":  year,
            "month": month,
        })
    })

    r.Run(":9090")
}
```

这样输入url为http://127.0.0.1:9090/blog/2024/2    可以得到json显示如下或者看图片9

```json
{
    "month": "2",
    "year": "2024"
}
```

---

---

## 七、gin参数绑定

```go
func main(){
    r:= gin.Default()

    r.GET("/user",func(c *gin.Context){
        username := c.Query("username")
        password := c.Query("password")
        u := UserInfo{
            username: username,
            password: password,
        }
        fmt.Printf("%#v\n",u)
        c.JSON(http.StatusOK,gin.H{
            "message": "ok",
        })
    })
    r.Run(":9090")
}
```

terminal

```
main.UserInfo{username:"", password:""}
[GIN] 2024/02/20 - 14:48:26 | 200 |       997.4µs |       127.0.0.1 | GET      "/user"
```

运行结果中username和password都是空字符串

如果要传入数据，就可以把url写成`http://127.0.0.1:9090/user?username=nidian&password=12345`

这样vscode终端就可以读取到

```
main.UserInfo{username:"nidian", password:"12345"}
[GIN] 2024/02/20 - 14:52:08 | 200 |       291.3µs |       127.0.0.1 | GET      "/user?username=nidian&password=12345"
```

页面仍旧是json格式字符串

```json
{
    "message": "ok"
}
```

但因为每次这样单独命名做成结构体u又用printf打印出来，比较麻烦，就引出了要学的参数绑定。

> %#v表示先打印结构体名字再打印结构体

**关于`c.ShouldBind(u)是什么`**

> 所以c是一个完整的请求 u则是拉取了c里面对应的数据
> 
> 2024/2/20 15:14:37
> 
> 是的，您理解得很正确。
> 
> 在Gin框架中，`c`（也就是Context上下文对象）代表了一个完整的HTTP请求，包含了请求的方法、URL、请求头、请求体等信息。通过`c.ShouldBind(u)`方法，我们可以从这个HTTP请求中提取出对应的数据，并将它们绑定到结构体`u`中。
> 
> 结构体`u`是您预先定义的用于存储请求数据的结构体。`c.ShouldBind(u)`方法会根据结构体的定义，将HTTP请求中的数据解析并赋值给结构体`u`的字段。这样，您就可以方便地使用结构体`u`中的字段来处理请求中的数据。
> 
> 需要注意的是，`c.ShouldBind(u)`方法会自动完成类型转换和数据验证。如果请求数据与结构体定义不匹配，或者无法成功转换为结构体字段所需的类型，将会返回错误信息。

以下是简写上面代码 通过调用ShouldBind函数

```go
package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

type UserInfo struct{
    Username string  `form:"username"`
    Password string  `form:"password"`
}

func main(){
    r:= gin.Default()

    r.GET("/user",func(c *gin.Context){
        /* 
        username := c.Query("username")
        password := c.Query("password")
        u := UserInfo{
            username: username,
            password: password,
        } */
        var u UserInfo //声明一个UserInfo类型的变量u
        err := c.ShouldBind(&u) //顺着指针来修改u的值 本质是拿出c这个http请求文内的数据
                                //一般是用在不知道上传什么数据 要求u的键是大写字母 写上tag
        if err != nil { 
            c.JSON(http.StatusBadRequest,gin.H{
                "error":err.Error(),
            })
        }else{
            c.JSON(http.StatusOK,gin.H{
                "status":"ok",
            })
        }

        r.Run(":9090")
}
```

> `main.UserInfo{Username:"nidian", Password:"12345"}
> [GIN] 2024/02/20 - 15:52:12 | 200 |       709.1µs |       127.0.0.1 | GET      "/user?username=nidian&password=12345"`

在 Go 语言中，结构体字段可以添加标签（tag）来描述该字段的元数据。标签是一种键值对的形式，它可以包含任意的字符串内容，通常被用于描述与该字段相关的附加信息，例如 JSON 序列化、表单数据绑定等。

在您的代码中，`Username` 和 `Password` 两个字段使用了标签 `form:"username"` 和 `form:"password"`，<mark>这些标签告诉 Gin 框架在处理 HTTP 请求时，应该将请求中的 `username` 和 `password` 参数值绑定到 `Username` 和 `Password` 字段上。</mark>

（说白了就是反射，我url输进去虽然和结构体键名不一样，但是我可以通过tag来绑定参数。）

具体来说，`form` 是 Gin 框架内置的一个标签名称，表示绑定表单数据。`"username"` 和 `"password"` 是该标签的参数，表示对应的表单参数名。通过使用标签，<u>Gin 框架能够自动识别 HTTP 请求中包含的参数，并将其与结构体中的字段进行绑定</u>，从而方便地获取和处理请求数据。

> 举个例子，如果某个接口既支持 `username` 参数，也支持 `user_name` 参数作为用户名，可以通过如下方式定义结构体：

> ``type UserInfo struct {     Username string `form:"username" json:"user_name"` }``

```go
type UserInfo struct{
    Username string  `form:"username"`
    Password string  `form:"password"`
}

func main(){
    r:= gin.Default()

    r.GET("/user",func(c *gin.Context){
        /* 
        username := c.Query("username")
        password := c.Query("password")
        u := UserInfo{
            username: username,
            password: password,
        } */
        var u UserInfo //声明一个UserInfo类型的变量u
        err := c.ShouldBind(&u) //顺着指针来修改u的值 本质是拿出c这个http请求文内的数据
                            //一般是用在不知道上传什么数据 要求u的键是大写字母 写上tag
        if err != nil { 
             c.JSON(http.StatusBadRequest,gin.H{
                "error":err.Error(),
            })
        }else{
            c.JSON(http.StatusOK,gin.H{
                "status":"ok",
            })
        }
    })

    r.POST("/form",func(c *gin.Context){
        var u UserInfo //声明一个UserInfo类型的变量u
        err := c.ShouldBind(&u) //顺着指针来修改u的值 本质是拿出c这个http请求文内的数据
                            //一般是用在不知道上传什么数据 要求u的键是大写字母 写上tag
        if err != nil { 
             c.JSON(http.StatusBadRequest,gin.H{
                "error":err.Error(),
            })
        }else{
            c.JSON(http.StatusOK,gin.H{
                "status":"ok",
            })
        }
    })


    r.Run(":9090")
}
```

进阶版 可以通过postman看传输数据，见图片10。

主打一个前端传数据，用postman快速发现是否传到后端，在terminal上面写出来。

可见图片11。（是用postman的body的raw然后选择JSON格式写从前端传入的数据。

```go
    r.POST("/json",func(c *gin.Context){
        var u UserInfo //声明一个UserInfo类型的变量u
        err := c.ShouldBind(&u) //顺着指针来修改u的值 本质是对照着键名拿出c这个http请求文内的数据
                            //一般是用在不知道上传什么数据 要求u的键是大写字母 写上tag
        if err != nil { 
             c.JSON(http.StatusBadRequest,gin.H{
                "error":err.Error(),
            })
        }else{
            c.JSON(http.StatusOK,gin.H{
                "status":"ok",
            })
        }
    })
```

因为函数shouldbind功能很强大，所以我们只需要修改路径就可以了，这个和前面user路径没有区别，只是名字不一样，上传json数据也不需要特别写代码，内容是一样的。

> `ShouldBind`会按照下面的顺序解析请求中的数据完成绑定：
> 
> 1. 如果是 `GET` 请求，只使用 `Form` 绑定引擎（`query`）。
> 2. 如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，然后再使用 `Form`（`form-data`）。

修改json的表单反射名字

```go
type UserInfo struct{
    Username string  `form:"username" json:"user"`
    Password string  `form:"password" json:"pwd"`
}
```

但是哈！！重点来了，如果已经命名好json对应的键名了，postman里不能继续用username了，会接受成空字符串！！要与时俱进改成user和pwd的键名。见图片12.

---

---

## 八、gin文件上传

### （1）单个文件上传

```html
<form action="/upload" method="post" enctype="multipart/form-data">
        <input type="file" name="f1">
        <input type="submit" value="上传">
   </form>
```

> 这里enctype是指定要上传的文件类型，已经填的multipart是复合类型，可以是图片音乐文件等。显示情况如图13

```go
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main(){
    r := gin.Default()

    r.LoadHTMLFiles("./web15/index.html")

    r.GET("/index",func(c *gin.Context){
        c.HTML(http.StatusOK,"index.html",nil)
    })
    r.Run(":9090")
}
```

这个渲染后会形成上传文本框

加入upload路径代码，可以在网页上传文件，根据路径直接存储到后端主文件夹里。见图片14

```go
package main

import (

    "net/http"
    "path"


    "github.com/gin-gonic/gin"
)

func main(){
    r := gin.Default()

    r.LoadHTMLFiles("./web15/index.html")

    r.GET("/index",func(c *gin.Context){
        c.HTML(http.StatusOK,"index.html",nil)
    })

    r.POST("/upload",func(c *gin.Context){
        //从请求中读取文件
        f,err := c.FormFile("f1")   //从请求中获取携带的参数一样的思路
        if err != nil{
            c.JSON(http.StatusBadRequest,gin.H{
                "error" : err.Error(),
            })
        }else{
            //将读取到的文件保存在本地（服务端本地）
            //dst := fmt.Sprintf("./%s",f.Filename)
            dst := path.Join("./",f.Filename)
            c.SaveUploadedFile(f,dst)
            c.JSON(http.StatusOK,gin.H{
                "status":"OK",
            })
        }

    })
    r.Run(":9090")
}
```

### （2）多个文件上传

```go
func main() {
    router := gin.Default()
    // 处理multipart forms提交文件时默认的内存限制是32 MiB
    // 可以通过下面的方式修改
    // router.MaxMultipartMemory = 8 << 20  // 8 MiB
    router.POST("/upload", func(c *gin.Context) {
        // Multipart form
        form, _ := c.MultipartForm()
        files := form.File["file"]

        for index, file := range files {
            log.Println(file.Filename)
            dst := fmt.Sprintf("C:/tmp/%s_%d", file.Filename, index)
            // 上传文件到指定的目录
            c.SaveUploadedFile(file, dst)
        }
        c.JSON(http.StatusOK, gin.H{
            "message": fmt.Sprintf("%d files uploaded!", len(files)),
        })
    })
    router.Run()
}
```

>  多个文件上传，重点在于要循环读取保存到指定的目录下.
> 
> 上传文件也是借助http请求，从客户端到服务端。

---

---

## 九、gin请求重定向

就是把传来的请求由当前服务器转给另外一个服务器，或者在自己的服务器切换路由。

分为永久和不永久。

```go
func main(){
    r := gin.Default()

    r.GET("/index",func(c *gin.Context){
        /* 
        c.JSON(http.StatusOK,gin.H{
            "status":"ok",
        }) */
        c.Redirect(http.StatusMovedPermanently,"http://hdufhq.cn:8888/")
    })

    r.Run(":9090")
}
```

登录这个本机地址加index，就会自动跳转到孵化器地址。见图片15

加入转发操作。

```go
    r.GET("/a",func(c *gin.Context){
        //跳转到/b对应的路由处理函数
        c.Request.URL.Path = "/b"  //把请求的uri修改成去b
        r.HandleContext(c)         //继续后续的处理
    })

    r.GET("/b",func(c *gin.Context){
        c.JSON(http.StatusOK,gin.H{
            "message":"b",
        })
    })
```

所以在登录a路由的时候，会马上把c指向的请求里的路由改成b路径，然后再次处理这个请求，直到跳到正确的路径，这里是b。见图片16

> 重定向一般用在访问网站时需要跳转。比如访问网站的时候跳到登录界面，再跳回来。

>  可以理解为，前端的超链接跳转是一种，后端这个代码跳转是内部重定向，也就是路由的定向发生改变，和超链接略有不同。

---

---

## 十、gin路由和路由组

### （1）普通路由

```go
r.GET("/index", func(c *gin.Context) {...})
r.GET("/login", func(c *gin.Context) {...})
r.POST("/login", func(c *gin.Context) {...})
```

```go
    r:= gin.Default()

    //访问/index的get请求走这条逻辑 会去执行下面这个匿名函数
    //用于访问得到服务器到客户端的数据 比如来获取之前存储的个人信息 个性签名 昵称之类的
    r.GET("/index",func(c *gin.Context){
        c.JSON(http.StatusOK,gin.H{
            "method":"GET",
        })
    })

    //创建操作 比如注册会员等由客户端输入的数据 再从客户端传给服务器存储
    r.POST("/index",func(c *gin.Context){
        c.JSON(http.StatusOK,gin.H{
            "method":"POST",
        })
    })

    //更新个人信息 部分信息 把生日修改为xx
    r.DELETE("/index",func(c *gin.Context){
        c.JSON(http.StatusOK,gin.H{
            "method":"DELETE",
        })
    })

    //删除 比如注销账号
    r.PUT("/index",func(c *gin.Context){
        c.JSON(http.StatusOK,gin.H{
            "method":"PUT",
        })
    })
```

此外，还有一个可以匹配所有请求方法的`Any`方法如下：

```go
r.Any("/test", func(c *gin.Context) {...})
```

```go
    r.Any("/user",func(c *gin.Context){
        switch c.Request.Method{
        case "GET":
            c.JSON(http.StatusOK,gin.H{"method":"GET"})
        case http.MethodPost:
            c.JSON(http.StatusOK,gin.H{"method":"POST"})

        }
```

为没有配置处理函数的路由添加处理程序，默认情况下它返回404代码，下面的代码为没有匹配到路由的请求都返回`views/404.html`页面，或者其他情况如图片18。

```go
r.NoRoute(func(c *gin.Context) {
        c.HTML(http.StatusNotFound, "views/404.html", nil)
    })
```

### （2）路由组

我们可以将拥有<mark>共同URL前缀的路由</mark>划分为一个路由组。习惯性一对`{}`包裹同组的路由，这只是为了看着清晰，你用不用`{}`包裹功能上没什么区别。

```go
func main() {
    r := gin.Default()
    userGroup := r.Group("/user")
    {
        userGroup.GET("/index", func(c *gin.Context) {...})
        //实际是/user/index
        userGroup.GET("/login", func(c *gin.Context) {...})
        userGroup.POST("/login", func(c *gin.Context) {...})

    }
    shopGroup := r.Group("/shop")
    {
        shopGroup.GET("/index", func(c *gin.Context) {...})
        shopGroup.GET("/cart", func(c *gin.Context) {...})
        shopGroup.POST("/checkout", func(c *gin.Context) {...})
    }
    r.Run()
```

```go
videoGroup := r.Group("/video")
    {
        videoGroup.GET("/index",func(c *gin.Context){
            c.JSON(http.StatusOK,gin.H{"msg":"/video/index"})
        })
        videoGroup.GET("/xx",func(c *gin.Context){
            c.JSON(http.StatusOK,gin.H{"msg":"/video/xx"})
        })
        videoGroup.GET("/aaa",func(c *gin.Context){
            c.JSON(http.StatusOK,gin.H{"msg":"/video/aaa"})
        })
    }
```

路由组也是支持<mark>嵌套</mark>的，例如：

```go
shopGroup := r.Group("/shop")
    {
        shopGroup.GET("/index", func(c *gin.Context) {...})
        shopGroup.GET("/cart", func(c *gin.Context) {...})
        shopGroup.POST("/checkout", func(c *gin.Context) {...})
        // 嵌套路由组
        xx := shopGroup.Group("xx")
        xx.GET("/oo", func(c *gin.Context) {...})
    }
```

通常我们将路由分组用在划分业务逻辑或划分API版本时。

---

---

## 十一、gin中间件

Gin框架允许开发者在处理请求的过程中，加入用户自己的钩子（Hook）函数。这个钩子函数就叫中间件，中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、记录日志、耗时统计等。

> c *gin.Context本身就是一个Handlefunc函数，所以可以单独提上来写，上下代码效果等价。
> 
> ```go
> func indexHandler(c *gin.Context){
>     c.JSON(http.StatusOK),gin.H{
>         "msg":"index",
>     }
> }
> 
> func main(){
>     r:=gin.Default()
> 
>     r.GET("/index",func(c *gin.Context){
>         c.JSON(http.StatusOK,gin.H{
>             "msg":"index",
>         })
>     })
> 
>     r.Run(":9090")
> }
> ```

### （1）开始加入中间件

```go
func indexHandler(c *gin.Context){
    fmt.Println("index")
    c.JSON(http.StatusOK,gin.H{
        "msg":"index",
    })
}

//定义一个中间件m1
func m1(c *gin.Context){
    fmt.Println("m1 in ......")
}

func main(){
    r:=gin.Default()

    r.GET("/index",m1,indexHandler)  //m1是中间件，这里注册

    r.Run(":9090")
}
```

终端上就会显示，注意看顺序，先中间件，再访问index，在浏览器页面输出JSON格式数据。

> [GIN-debug] Listening and serving HTTP on :9090
> <mark>m1 in ......
> index</mark>
> [GIN] 2024/02/22 - 23:14:53 | 200 |       736.3µs |       127.0.0.1 | GET      "/index"        

修改一下

```go
// 定义一个中间件m1
func m1(c *gin.Context) {
    fmt.Println("m1 in ......")

    //计时  统计耗时
    strat := time.Now()
    c.Next()  //责任链模式 调用后调用后续的处理函数也就是indexHandler续的处理函数
    //c.Abort() //阻止调用后续的处理函数
    cost := time.Since(start)

    fmt.Printf("cosgt:%v\n",cost)
    fmt.Println("m1 out...")
}
```

进行Next去处理后面的函数indexHandler，然后返回值，再回到中间件做后续，得到运行时间，输出m1结束的字段。

> [GIN-debug] Listening and serving HTTP on :9090
> <mark>m1 in ......
> index
> cosgt:20.2358ms
> m1 out...</mark>
> [GIN] 2024/02/23 - 14:37:49 | 200 |    111.0944ms |       127.0.0.1 | GET      "/index"

### （2）全局中间件

```go
func main() {
    r := gin.Default()

    r.Use(m1)  //注册一个全局中间件 作用在下面所有路由  后续就不用写m1了

    r.GET("/index", indexHandler)
    r.GET("/shop",func(c *gin.Context){
        c.JSON(http.StatusOK,gin.H{
            "msg":"shop",
        })
    })
    r.GET("/user",func(c *gin.Context){
        c.JSON(http.StatusOK,gin.H{
```

当有两个全局中间件时。。。

```go
func indexHandler(c *gin.Context) {
    fmt.Println("index")
    c.JSON(http.StatusOK, gin.H{
        "msg": "index",
    })
}

// 定义一个中间件m1
func m1(c *gin.Context) {
    fmt.Println("m1 in ......")

    //计时  统计耗时
    start := time.Now()
    c.Next() //责任链模式 调用后续的处理函数也就是indexHandler
    //c.Abort() //阻止调用后续的处理函数
    cost := time.Since(start)

    fmt.Printf("cosgt:%v\n", cost)
    fmt.Println("m1 out......")
}

func m2(c *gin.Context) {
    fmt.Println("m2 in ......")
    c.Next()  //调用后续处理函数 打印index
    fmt.Println("m2 out......")
}

func main() {
    r := gin.Default()

    r.Use(m1,m2) //注册全局中间件 作用在下面所有路由  后续就不用写m1了
                 //本质会变成递归 m1调动m2，m2调取index函数，返回index'然后再打印m2结束，最后m1结束，打印m1结束

    r.GET("/index", indexHandler)
    r.GET("/shop", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "msg": "shop",
        })
    })
    r.GET("/user", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "msg": "user",
        })
    })

    r.Run(":9090")
}
```

本质会变成递归 m1调动m2，m2调取index函数，返回index'然后再打印m2结束，最后m1结束，打印耗时然后打印m1 out结束。

> [GIN-debug] Listening and serving HTTP on :9090
> <mark>m1 in ......
> m2 in ......
> index
> m2 out......
> cosgt:670.2µs
> m1 out......</mark>
> [GIN] 2024/02/23 - 15:40:07 | 200 |      3.8972ms |       127.0.0.1 | GET      "/index"

本质是先进后出模型，就是个栈。可见图片20，21。

### （3）阻止响应

```go
func m2(c *gin.Context) {
    fmt.Println("m2 in ......")
    //c.Next()  //调用后续处理函数 打印index
    c.Abort()   //阻止后续的处理函数
    fmt.Println("m2 out......")
}
```

> 执行结果
> 
> <mark>m1 in ......
> m2 in ......
> m2 out......
> cosgt:0s
> m1 out......</mark>
> [GIN] 2024/02/23 - 15:44:27 | 404 |      1.1684ms |       127.0.0.1 | GET      "/favicon.ico"      

如果再加入return ，就会如图片23所示。

```go
func m2(c *gin.Context) {
    fmt.Println("m2 in ......")
    //c.Next()  //调用后续处理函数 打印index
    c.Abort()   //阻止后续的处理函数
    return
    fmt.Println("m2 out......")
}
```

### （3）判断---用于业务

```go
func authMiddleware(doCheck bool)gin.HandlerFunc{
    //链接数据库
    //或者一些其他准备工作
    return func(c *gin.Context){
        if doCheck{
            //存放具体的逻辑
            //是否登录的判断
            //if是登录用户 则c.Next() 否则else c.Abort()
        }else{
            c.Next()
        }

    }

}

func main() {
    r := gin.Default()

    r.Use(m1,m2,authMiddleware(true)) //注册全局中间件 作用在下面所有路由  后续就不用写m1了
                 //本质会变成递归 m1调动m2，m2调取index函数，返回index'然后再打印m2结束，最后m1结束，打印m1结束

    r.GET("/index", indexHandler)
    r.GET("/shop", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "msg": "shop",
        })
    })
    r.GET("/user", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "msg": "user",
        })
    })

    r.Run(":9090")
}
```

### （4）为某个路由单独注册

最开始的写法，再get里添加中间件，非全局注册

```go
// 给/test2路由单独注册中间件（可注册多个）
    r.GET("/test2", StatCost(), func(c *gin.Context) {
        name := c.MustGet("name").(string) // 从上下文取值
        log.Println(name)
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello world!",
        })
    })
```

### （5）为路由组注册中间件

为路由组注册中间件有以下两种写法。

写法1：

```go
shopGroup := r.Group("/shop", StatCost())
{
    shopGroup.GET("/index", func(c *gin.Context) {...})
    ...
}
```

写法2：

```go
shopGroup := r.Group("/shop")
shopGroup.Use(StatCost())
{
    shopGroup.GET("/index", func(c *gin.Context) {...})
    ...
}
```

具体实施：

```go
//路由组注册中间件方法1：
    xxGroup := r.Group("/xx", authMiddleware(true))
    {
        xxGroup.GET("/index", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"msg": "xxGroup"})
        })
    }

    //路由组注册中间件方法2：
    xx2Group := r.Group("/xx2")
    xx2Group.Use(authMiddleware(true))
    {
        xx2Group.GET("/index", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"msg": "xx2Group"})
        })
    }
```

c.Set("name", "Leeson")    //上下文中命名，让后续函数得到name 实现跨中间件

name, ok := c.Get("name") //获取set数据

### （6）  gin默认中间件（注意事项

`gin.Default()`默认使用了`Logger`和`Recovery`中间件，其中：

- `Logger`中间件将日志写入`gin.DefaultWriter`，即使配置了`GIN_MODE=release`。
- `Recovery`中间件会recover任何`panic`。如果有panic的话，会写入500响应码。

.

如果不想使用上面两个默认的中间件，可以使用`gin.New()`新建一个没有任何默认中间件的路由。

### （7）gin中间件中使用goroutine

当在中间件或`handler`中启动新的`goroutine`时，**<u>不能使用**</u>原始的上下文（c *gin.Context），必须使用其只读副本（`c.Copy()`）。

可以避免func修改请求，导致底下运行Next时候并发，不安全。

```go
func m1(c *gin.Context) {
    fmt.Println("m1 in ......")

    //计时  统计耗时
    start := time.Now()

    go funXX(c.Copy()) //goroutine在funcXX中只能使用c的拷贝

    c.Next() //责任链模式 调用后续的处理函数也就是indexHandler
    //c.Abort() //阻止调用后续的处理函数
    cost := time.Since(start)

    fmt.Printf("cosgt:%v\n", cost)
    fmt.Println("m1 out......")

```

### （8）跨中间件传递键值对

主要看c.Set和c.Get！

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	fmt.Println("index")
	name, ok := c.Get("name") //获取set数据
	if !ok {
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

// 定义一个中间件m1
func m1(c *gin.Context) {
	fmt.Println("m1 in ......")

	//计时  统计耗时
	start := time.Now()

	//go funXX(c.Copy()) //goroutine在funcXX中只能使用c的拷贝
	
	c.Next() //责任链模式 调用后续的处理函数也就是indexHandler
	//c.Abort() //阻止调用后续的处理函数
	cost := time.Since(start)

	fmt.Printf("cosgt:%v\n", cost)
	fmt.Println("m1 out......")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in ......")
	c.Set("name", "Leeson")    //上下文中命名，让后续函数得到name 实现跨中间件
	//c.Next()  //调用后续处理函数 打印index
	//c.Abort() //阻止后续的处理函数
	//return
	fmt.Println("m2 out......")
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	//链接数据库
	//或者一些其他准备工作
	return func(c *gin.Context) {
		if doCheck {
			//存放具体的逻辑
			//是否登录的判断
			//if是登录用户 则c.Next() 否则else c.Abort()
			c.Next()
		} else {
			c.Next()
		}

	}

}

func main() {
	r := gin.Default()

	r.Use(m1, m2, authMiddleware(true)) //注册全局中间件 作用在下面所有路由  后续就不用写m1了
	//本质会变成递归 m1调动m2，m2调取index函数，返回index'然后再打印m2结束，最后m1结束，打印m1结束

	r.GET("/index", indexHandler)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})


	/* 
	//路由组注册中间件方法1：
	xxGroup := r.Group("/xx", authMiddleware(true))
	{
		xxGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "xxGroup"})
		})
	}

	//路由组注册中间件方法2：
	xx2Group := r.Group("/xx2")
	xx2Group.Use(authMiddleware(true))
	{
		xx2Group.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "xx2Group"})
		})
	} */

	r.Run(":9090")
}

```

在这段代码中，我们定义了两个中间件`m1`和`m2`，并在路由处理函数中使用了`c.Get("name")`获取了`m2`中设置的值。而在`m2`中，我们使用了`c.Set("name", "Leeson")`将"Leeson"这个值设置到了上下文中。因此，在路由处理函数中执行`c.Get("name")`时，会获取到上下文中名为"name"的键所对应的值"Leeson"。最终的返回结果就是`"msg": "Leeson"`。

需要注意的是，在Gin框架中，可以通过`c.Set(key, value)`方法将键值对设置到上下文中，然后可以通过`c.Get(key)`方法获取该键所对应的值。这样就可以实现跨中间件的值传递。

> [GIN-debug] Listening and serving HTTP on :9090
> <mark>m1 in ......
> m2 in ......
> m2 out......
> index
> cosgt:499.7µs
> m1 out......</mark>
> [GIN] 2024/02/23 - 17:24:30 | 200 |      2.0033ms |       127.0.0.1 | GET      "/index"

> Q:
> 
> 所以实际<mark>是m1到m2到autoMiddleware再到indexHandler再逐级返回的过程</mark>？
> 
> A:
> 
> 2024/2/23 17:39:14
> 
> 是的，中间件的处理过程就是一个责任链模式。在该代码中，请求会先进入`m1`中间件，然后再进入`m2`中间件，接着进入`authMiddleware`中间件，最终进入路由处理函数`indexHandler`。
> 
> 在这个过程中，每个中间件都可以对请求做出一定的处理，比如记录请求日志、统计请求耗时、权限校验等等。而在中间件内部，如果需要将一些数据传递给后续的中间件或者路由处理函数，可以使用`c.Set(key, value)`方法<mark>将键值对设置到上下文中，然后在后续的中间件或路由处理函数中，可以使用`c.Get(key)`方法获取该键所对应的值</mark>。
> 
> 最终，在路由处理函数中，会将响应结果返回给客户端，然后倒序逐级返回，执行完`indexHandler`后，会返回到`authMiddleware`中间件，然后再返回到`m2`中间件，最后返回到`m1`中间件。在逐级返回的过程中，每个中间件都可以对响应结果做出一定的处理，比如添加响应头、修改响应内容等等。

### （9）运行多个服务

我们可以在多个端口启动服务，例如：

```go
package main

import (
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "golang.org/x/sync/errgroup"
)

var (
    g errgroup.Group
)

func router01() http.Handler {
    e := gin.New()
    e.Use(gin.Recovery())
    e.GET("/", func(c *gin.Context) {
        c.JSON(
            http.StatusOK,
            gin.H{
                "code":  http.StatusOK,
                "error": "Welcome server 01",
            },
        )
    })

    return e
}

func router02() http.Handler {
    e := gin.New()
    e.Use(gin.Recovery())
    e.GET("/", func(c *gin.Context) {
        c.JSON(
            http.StatusOK,
            gin.H{
                "code":  http.StatusOK,
                "error": "Welcome server 02",
            },
        )
    })

    return e
}

func main() {
    server01 := &http.Server{
        Addr:         ":8080",
        Handler:      router01(),
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
    }

    server02 := &http.Server{
        Addr:         ":8081",
        Handler:      router02(),
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
    }
   // 借助errgroup.Group或者自行开启两个goroutine分别启动两个服务
    g.Go(func() error {
        return server01.ListenAndServe()
    })

    g.Go(func() error {
        return server02.ListenAndServe()
    })

    if err := g.Wait(); err != nil {
        log.Fatal(err)
    }
}
```
