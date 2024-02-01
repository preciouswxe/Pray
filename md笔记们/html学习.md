# $\color{lime}{||HTML学习||}$

---

学习内容：菜鸟教程的文档[HTML 教程 | 菜鸟教程 (runoob.com)](https://www.runoob.com/html/html-tutorial.html)

---

## 一、引入初始框架

```html
<!DOCTYPE>
<html>
<head>
<meta charset="utf-8">
<title>泥电</title> 
</head>
<body>
    <h1>第三轮孵化器考核</h1>
    <p>我的第一个任务。</p>
</body>
</html>
```

--- 

## 二、简介

- **<!DOCTYPE html>** 声明为   HTML5 文档

- **<html>** 元素是 HTML 页面的  根元素

- **<head>** 元素包含了文档的元（meta）数据，如 <meta charset="utf-8"> 定义网页编码格式为 **utf-8**。

- **<title>** 元素描述了文档的标题

- **<body>** 元素包含了可见的页面内容

- **<h1>** 元素定义一个大标题

- **<p>** 元素定义一个段落
  
  <mark><img title="" src="https://www.runoob.com/wp-content/uploads/2013/06/02A7DD95-22B4-4FB9-B994-DDB5393F7F03.jpg" alt="" width="716"></mark>

### （1）什么是HTML？

HTML 是用来描述网页的一种语言。

- HTML 指的是超文本标记语言: **H**yper**T**ext **M**arkup **L**anguage
- HTML 不是一种编程语言，而是一种**标记**语言
- 标记语言是一套**标记标签** (markup tag)
- HTML 使用标记标签来**描述**网页
- HTML 文档包含了HTML **标签**及**文本**内容
- HTML文档也叫做 **web 页面**

---

## （2） HTML 标签

HTML 标记标签通常被称为 HTML 标签 (HTML tag)。

- HTML 标签是由*尖括号*包围的关键词，比如 <html>
- HTML 标签通常是*成对出现*的
- 标签对中的第一个标签是*开始标签*，第二个标签是*结束标签*
- 开始和结束标签也被称为*开放标签*和*闭合标签*

---

## (3)    vs配置和实时预览

<img src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-01-28-11-00-48-image.png" title="" alt="" width="860">

---

# 三、基础

## ① HTML 标题

HTML 标题（Heading）是通过<h1> - <h6> 标签来定义的。

```html
<h1>这是一个标题</h1>
<h2>这是一个标题</h2>
<h3>这是一个标题</h3>
```

## ② HTML 段落

HTML 段落是通过标签 <p> 来定义的。

```html
<p>这是一个段落。</p>
<p>这是另外一个段落。</p>
```

## ③ HTML 链接

HTML 链接是通过标签 <a> 来定义的。

```html
<a href="https://www.runoob.com">这是一个链接</a>
```

## ④ HTML 图像

```html
HTML 图像是通过标签<img>来定义的.
<img src="/images/logo.png" width="258" height="39" />
```

**注意：** 图像的名称和尺寸是以属性的形式提供的。

---

# 

# 四、HTML元素

| 开始标签 *                 | 元素内容   | 结束标签 * |
| ---------------------- | ------ | ------ |
| <p>                    | 这是一个段落 | </p>   |
| <a href="default.htm"> | 这是一个链接 | </a>   |
| <br>                   | 换行     |        |

**开始标签常被称为**起始标签（opening tag）**，结束标签常称为**闭合标签（closing tag）。

## ① HTML 元素语法

- HTML 元素以**开始标签**起始
- HTML 元素以**结束标签**终止
- **元素的内容**是开始标签与结束标签之间的内容
- 某些 HTML 元素具有**空内容（empty content）**
- 空元素**在开始标签中进行关闭**（以开始标签的结束而结束）
- 大多数 HTML 元素可拥有**属性**

## ② 嵌套的 HTML 元素

大多数 HTML 元素可以嵌套（HTML 元素可以包含其他 HTML 元素）。

HTML 文档由相互嵌套的 HTML 元素构成。

## ④ HTML 文档实例

```html
<!DOCTYPE html>
<html>

<body>
<p>这是第一个段落。</p>
</body>

</html>
以上实例包含了三个 HTML 元素。
```

## ⑤ HTML 实例解析

**<p> 元素:**

```html
<p>这是第一个段落。</p>

这个 <p> 元素定义了 HTML 文档中的一个段落。  
这个元素拥有一个开始标签 <p> 以及一个结束标签 </p>.  
元素内容是: 这是第一个段落。
```

**<body> 元素:**

```html
<body>  
<p>这是第一个段落。</p>  
</body>
```

<body> 元素定义了 HTML 文档的主体。  
这个元素拥有一个开始标签 <body> 以及一个结束标签 </body>。  
元素内容是另一个 HTML 元素（p 元素）。

**<html> 元素：**

<html>

```html
<html>

<body>
<p>这是第一个段落。</p>
</body>

</html>
```

</html>

<html> 元素定义了整个 HTML 文档。  
这个元素拥有一个开始标签 <html> ，以及一个结束标签 </html>.  
元素内容是另一个 HTML 元素（body 元素）。

## ⑥ HTML 空元素

没有内容的 HTML 元素被称为空元素。空元素是在开始标签中关闭的。

<br> 就是没有关闭标签的空元素（<br> 标签定义换行）。

在 XHTML、XML 以及未来版本的 HTML 中，所有元素都必须被关闭。

在开始标签中添加斜杠，比如 <br />，是关闭空元素的正确方法，HTML、XHTML 和 XML 都接受这种方式。

即使 <br> 在所有浏览器中都是有效的，但使用 <br /> 其实是更长远的保障。

## ⑦ HTML 提示：使用小写标签

HTML 标签对大小写不敏感：<P> 等同于 <p>。许多网站都使用大写的 HTML 标签。

菜鸟教程使用的是小写标签，因为万维网联盟（W3C）在 HTML 4 中**推荐**使用小写，而在未来 (X)HTML 版本中**强制**使用小写。

---

# 五、HTML属性

- HTML 元素可以设置**属性**
- 属性可以在元素中添加**附加信息**
- <mark>属性一般描述于**开始标签**</mark>
- 属性总是以名称/值对的形式出现，**比如：name="value"**。

---

## ① 属性实例

HTML 链接由 <a> 标签定义。链接的地址在 **href 属性**中指定：

```html
<a href="http://www.runoob.com">这是一个链接</a>
```

## ② HTML 属性常用引用属性值

<u>属性值应该始终被包括在引号内。</u>

双引号是最常用的，不过使用单引号也没有问题。

> ![Remark](https://www.runoob.com/images/lamp.gif)**提示:** 在某些个别的情况下，比如属性值本身就含有双引号，那么您必须使用单引号，例如：
> 
> name='John "ShotGun" Nelson'

| 属性    | 描述                                     |
| ----- | -------------------------------------- |
| class | 为html元素定义一个或多个类名（classname）(类名从样式文件引入) |
| id    | 定义元素的唯一id                              |
| style | 规定元素的行内样式（inline style）                |
| title | 描述了元素的额外信息 (作为工具条使用)                   |

更多标准属性说明： [HTML 标准属性参考手册](https://www.runoob.com/tags/ref-standardattributes.html).

例如，以下是一个使用属性的示例：

```html
<a href="https://www.example.com" target="_blank">这是一个链接</a>`
```

在上面的例子中，`href`是链接元素(`<a>`)的属性之一，它指定了链接的目标URL。`

target`是另一个属性，它指定了链接在何处打开，`_blank`表示在新的浏览器窗口中打开链接。

HTML中的属性可以用于各种元素，包括图片、表单元素、样式等。每个元素都具有一组特定的属性，用于定义其行为和外观。可以通过在开始标签中使用属性来设置元素的属性值。

---

# 六、HTML标题

标题（Heading）是通过 <h1> - <h6> 标签进行定义的。

```html
<h1> 定义最大的标题。 <h6> 定义最小的标题。

<h1>这是一个标题。</h1>
<h2>这是一个标题。</h2>
<h3>这是一个标题。</h3>

注释: 浏览器会自动地在标题的前后添加空行。
```

请确保将 HTML 标题 标签只用于标题。不要仅仅是为了生成**粗体**或**大号**的文本而使用标题。

搜索引擎使用标题为您的网页的结构和内容编制索引。

因为用户可以通过标题来快速浏览您的网页，所以用标题来呈现文档结构是很重要的。

应该将 h1 用作主标题（最重要的），其后是 h2（次重要的），再其次是 h3，以此类推。

## 

## HTML 水平线

hr 元素可用于分隔内容。

```html
<p>这是一个段落。</p>
<hr>
<p>这是一个段落。</p>
<hr>
<p>这是一个段落。</p>
```

## HTML 注释

可以将注释插入 HTML 代码中，这样可以提高其可读性，使代码更易被人理解。浏览器会忽略注释，也不会显示它们。

注释写法如下:

```html
<!-- 这是一个注释 -->
```

**注释:** 开始括号之后（左边的括号）需要紧跟一个叹号 ! (英文标点符号)，结束括号之前（右边的括号）不需要，合理地使用注释可以对未来的代码编辑工作产生帮助。

## HTML 折行

如果您希望在不产生一个新段落的情况下进行换行（新行），请使用 `<br>`标签：

```html
<p>这个<br>段落<br>演示了分行的效果</p>
```

`<br />` 元素是一个空的 HTML 元素。由于关闭标签没有任何意义，因此它没有结束标签。

## 实例

```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>菜鸟教程(runoob.com)</title>
</head>

<body>

<h1>春晓</h1>

<p>
    春眠不觉晓，
      处处闻啼鸟。
        夜来风雨声，
          花落知多少。
</p>

<p>注意，浏览器忽略了源代码中的排版（省略了多余的空格和换行）。</p>
</body>
</html>
```

输出结果：

```html
春晓

春眠不觉晓， 处处闻啼鸟。 夜来风雨声， 花落知多少。

注意，浏览器忽略了源代码中的排版（省略了多余的空格和换行）。
```

  ---

# 

# 七、HTML 文本格式化

<!DOCTYPE html>

```html
<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>菜鸟教程(runoob.com)</title> 
</head> 
<body>

<b>加粗文本</b><br><br>
<i>斜体文本</i><br><br>
<code>电脑自动输出</code><br><br>
这是 <sub> 下标</sub> 和 <sup> 上标</sup>

</body>
</html>
```

<img title="" src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-01-28-14-58-30-image.png" alt="" width="264" data-align="left">

```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>菜鸟教程(runoob.com)</title>
</head>

<body>

<b>这个文本是加粗的</b>

<br />

<strong>这个文本是加粗的</strong>

<br />

<big>这个文本字体放大</big>

<br />

<em>这个文本是斜体的</em>

<br />

<i>这个文本是斜体的</i>

<br />

<small>这个文本是缩小的</small>

<br />

这个文本包含
<sub>下标</sub>

<br />

这个文本包含
<sup>上标</sup>

</body>
</html>
```

<img src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-01-28-15-01-48-image.png" title="" alt="" width="259">

```html
<pre>
此例演示如何使用 pre 标签
对空行和    空格
进行控制
</pre>
```

<img title="" src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-01-28-15-04-53-image.png" alt="" data-align="inline" width="257">

```html
<code>计算机输出</code>
<br />
<kbd>键盘输入</kbd>
<br />
<tt>打字机文本</tt>
<br />
<samp>计算机代码样本</samp>
<br />
<var>计算机变量</var>
<br />

<p>
<b>注释：</b>这些标签常用于显示计算机/编程代码。
</p>
```

<img title="" src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-01-28-15-05-47-image.png" alt="" width="317">

### 写地址

```html
<address>
Written by <a href="mailto:webmaster@example.com">Jon Doe</a>.<br> 
Visit us at:<br>
Example.com<br>
Box 564, Disneyland<br>
USA
</address>
```

<img src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-01-28-15-07-10-image.png" title="" alt="" width="278">

### 缩写和首字母缩写

```html
<abbr title="etcetera">etc.</abbr>
<br />
<acronym title="World Wide Web">WWW</acronym>

<p>在某些浏览器中，当您把鼠标移至缩略词语上时，title 可用于展示表达的完整版本。</p>

<p>仅对于 IE 5 中的 acronym 元素有效。</p>

<p>对于 Netscape 6.2 中的 abbr 和 acronym 元素都有效。</p>
```

<img src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-01-28-15-08-35-image.png" title="" alt="" width="478">

### 左右换向

```html
<p>该段落文字从左到右显示。</p>  
<p><bdo dir="rtl">该段落文字从右到左显示。</bdo></p> 
```

<img src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-01-28-15-10-46-image.png" title="" alt="" width="235">

### 引用

```html
<p>WWF's goal is to: 
<q>Build a future where people live in harmony with nature.</q>
We hope they succeed.</p>
```

<img src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-01-28-15-12-11-image.png" title="" alt="" width="672">

### 标记删除文字

```html
<p>My favorite color is <del>blue</del> <ins>red</ins>!</p>
```

<img src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-01-28-15-13-41-image.png" title="" alt="" width="327">

## 总结 HTML 文本格式化标签

| 标签       | 描述     |
| -------- | ------ |
| <b>      | 定义粗体文本 |
| <em>     | 定义着重文字 |
| <i>      | 定义斜体字  |
| <small>  | 定义小号字  |
| <strong> | 定义加重语气 |
| <sub>    | 定义下标字  |
| <sup>    | 定义上标字  |
| <ins>    | 定义插入字  |
| <del>    | 定义删除字  |

## HTML "计算机输出" 标签

| 标签     | 描述        |
| ------ | --------- |
| <code> | 定义计算机代码   |
| <kde>  | 定义键盘码     |
| <samp> | 定义计算机代码样本 |
| <var>  | 定义变量      |
| <pre>  | 定义预格式文本   |

## HTML 引文, 引用, 及标签定义

| 标签           | 描述        |
| ------------ | --------- |
| <abbr>       | 定义缩写      |
| <address>    | 定义地址      |
| <bdo>        | 定义文字方向    |
| <blockquote> | 定义长的引用    |
| <q>          | 定义短的引用语   |
| <cite>       | 定义引用、引证   |
| <dfn>        | 定义一个定义项目。 |

---

# 

# 八、链接

```html
<p>
<a href="/index.html">本文本</a> 是一个指向本网站中的一个页面的链接。</p>

<p><a href="http://www.microsoft.com/">本文本</a> 是一个指向万维网上的页面的链接。</p>

</body>
```

输出结果：

<p>
<a href="/index.html">本文本</a> 是一个指向本网站中的一个页面的链接。</p>

<p><a href="http://www.microsoft.com/">本文本</a> 是一个指向万维网上的页面的链接。</p>

</body>

## ① HTML 超链接

HTML使用标签 <a> 来设置超文本链接。

超链接可以是一个字，一个词，或者一组词，也可以是一幅图像，您可以点击这些内容来跳转到新的文档或者当前文档中的某个部分。

当您把鼠标指针移动到网页中的某个链接上时，箭头会变为一只小手。

在标签 <a> 中使用了 href 属性来描述链接的地址。

默认情况下，链接将以以下形式出现在浏览器中：

- 一个未访问过的链接显示为蓝色字体并带有下划线。
- 访问过的链接显示为紫色并带有下划线。
- 点击链接时，链接显示为红色并带有下划线。

> 注意：如果为这些超链接设置了 CSS 样式，展示样式会根据 CSS 的设定而显示。

## ② HTML 链接语法

以下是 HTML 中创建链接的基本语法和属性：$\color{lime}{<a>}$ 元素：创建链接的主要HTML元素是`<a>`（锚）元素。`<a>`元素具有以下属性：

- `href`：指定链接目标的URL，这是链接的最重要属性。可以是另一个网页的URL、文件的URL或其他资源的URL。
- `target`（可选）：指定链接如何在浏览器中打开。常见的值包括 `_blank`（在新标签或窗口中打开链接）和 `_self`（在当前标签或窗口中打开链接）。
- `title`（可选）：提供链接的额外信息，通常在鼠标悬停在链接上时显示为工具提示。
- `rel`（可选）：指定与链接目标的关系，如 nofollow、noopener 等。

链接的 HTML 代码很简单，它类似这样：

`<a href="url">链接文本</a>`

href 属性描述了链接的目标。

**提示:** *"链接文本"* 不必一定是文本。图片或其他 HTML 元素都可以成为链接。

**文本链接**：最常见的链接类型是文本链接，它使用 <a> 元素将一段文本转化为可点击的链接，例如：

```html
<a href="https://www.example.com">访问示例网站</a>
```

**图像链接**：您还可以使用图像作为链接。在这种情况下，`<a>` 元素包围着 `<img> `元素。例如：

```html
<a href="https://www.example.com">
  <img src="example.jpg" alt="示例图片">
</a>
```

**锚点链接**：除了链接到其他网页外，您还可以在<u>同一页面内创建内部链接</u>，这称为锚点链接。要创建锚点链接，需要在目标位置使用 元素定义一个标记，并使用#符号引用该标记。例如：

```html
<a href="#section2">跳转到第二部分</a>
<!-- 在页面中的某个位置 -->
<a name="section2"></a>
```

**下载链接**：如果您希望链接用于下载文件而不是导航到另一个网页，可以使用 download 属性。例如：

```html
<a href="document.pdf" download>下载文档</a>
```

## ③ HTML 链接- id 属性

id 属性可用于创建一个 HTML 文档书签。

**提示:** 书签不会以任何特殊方式显示，即在 HTML 页面中是不显示的，所以对于读者来说是<mark>隐藏</mark>的。

### 实例

在HTML文档中**插入ID**:

```html
<a id="tips">有用的提示部分</a>
```

在HTML文档中创建一个链接到"有用的提示部分(id="tips"）"：

```html
<a href="#tips">访问有用的提示部分</a>
```

或者，从另一个页面创建一个链接到"有用的提示部分(id="tips"）"：

```html
<a href="https://www.runoob.com/html/html-links.html#tips">
访问有用的提示部分</a>
```

**注意事项：**

 $\color{yellow}{请始终将正斜杠添加到子文件夹}$。假如这样书写链接：href="https://www.runoob.com/html"，就会向服务器产生两次 HTTP 请求。这是因为服务器会添加正斜杠到这个地址，然后创建一个新的请求，就像这样：href="https://www.runoob.com/html/"。

## ④ 实例

**图片链接**

```html
<p>创建图片链接:
<a href="http://www.runoob.com/html/html-tutorial.html">
<img  border="10" src="smiley.gif" alt="HTML 教程" width="32" height="32"></a></p>

<p>无边框的图片链接:
<a href="http://www.runoob.com/html/html-tutorial.html">
<img border="0" src="smiley.gif" alt="HTML 教程" width="32" height="32"></a></p>
```

<img src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-01-28-21-54-16-image.png" title="" alt="" width="255">

**定位当前位置**

```html
<p>
<a href="#C4">查看章节 4</a>
</p>

<h2>章节 1</h2>
<p>这边显示该章节的内容……</p>

<h2>章节 2</h2>
<p>这边显示该章节的内容……</p>

<h2>章节 3</h2>
<p>这边显示该章节的内容……</p>

<h2><a id="C4">章节 4</a></h2>
<p>这边显示该章节的内容……</p>

<h2>章节 5</h2>
<p>这边显示该章节的内容……</p>
```

先定义跳转到C4的id的超文本链接，然后给章节四定义了C4的id，这样就可以生成了。见下图首

<img src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-01-28-22-01-45-image.png" title="" alt="" width="200">

**跳出框架**

```html
<p>跳出框架?</p> 
<a href="http://www.runoob.com/" target="_top">点击这里!</a> 
```

- `_top`是一个特殊的目标名称，代表顶层窗口或顶层框架。当链接的目标设置为`_top`时，浏览器会忽略当前页面的所有框架和嵌套页面，直接在整个窗口中加载链接的目标页面。
- 这意味着，不论当前页面是否包含在框架中，点击链接后都会跳出框架，将整个窗口显示链接的目标页面。

<img src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-01-28-22-13-30-image.png" title="" alt="" width="199">

**电子邮件链接**

```html
<p>
这是一个电子邮件链接：
<a href="mailto:someone@example.com?Subject=Hello%20again" target="_top">
发送邮件</a>
</p>

<p>
<b>注意:</b> 单词之间空格使用 %20 代替，以确保浏览器可以正常显示文本。
</p>
```

<img src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-01-28-22-14-47-image.png" title="" alt="" width="629">

```html
<p>
这是另一个电子邮件链接：
<a href="mailto:someone@example.com?cc=someoneelse@example.com&bcc=andsomeoneelse@example.com&subject=Summer%20Party&body=You%20are%20invited%20to%20a%20big%20summer%20party!" target="_top">发送邮件!</a>
</p>

<p>
<b>注意:</b> 单词之间的空格使用 %20 代替，以确保浏览器可以正常显示文本。
</p>
```

<img src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-01-28-22-15-55-image.png" title="" alt="" width="649">

---

# 九、头部

## <head>

<head> 元素包含了所有的头部标签元素。在 <head>元素中你可以插入脚本（scripts）, 样式文件（CSS），及各种meta信息。

可以添加在头部区域的元素标签为: <title>, <style>, <meta>, <link>, <script>, <noscript> 和 <base>。

## **<title>**

```html
<title> 标签定义了不同文档的标题。

<title> 在 HTML/XHTML 文档中是必需的。

<title> 元素:
定义了浏览器工具栏的标题
当网页添加到收藏夹时，显示在收藏夹中的标题
显示在搜索引擎结果页面的标题
```

```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8"> 
<title>我的 HTML 的第一页</title>
</head>

<body>
<p>浏览器中包含body元素的内容。</p>
<p>浏览器的标题包含title元素的内容</p>
</body>

</html>
```

<img src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-01-28-22-20-19-image.png" title="" alt="" width="331">

## **<base>**

描述基本的链接地址/链接目标，该标签作为HTML文档中所有的链接标签的默认链接:

```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8"> 
<title>菜鸟教程(runoob.com)</title> 
<base href="https://www.runoob.com/images/" target="_blank">
</head>

<body>
<img src="logo.png"> - 注意这里我们设置了图片的相对地址。能正常显示是因为我们在 head 部分设置了 base 标签，该标签指定了页面上所有链接的默认 URL，所以该图片的访问地址为 "https://www.runoob.com/images/logo.png"
<br><br>
<a href="https://www.runoob.com">菜鸟教程</a> - 注意这个链接会在新窗口打开，即便它没有 target="_blank" 属性。因为在 base 标签里我们已经设置了 target 属性的值为 "_blank"。

</body>
</html>
```

注意：target为blank的意思前面有提到，其实就是在新窗口打开，self是在当前窗口，top则会跳出。我对base的理解就是基础，body所有的基础，在此之上的宏定义。

<img title="" src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-01-28-22-27-53-image.png" alt="" width="559">

## <link>

定义了文档与外部资源之间的关系。

通常用于链接到样式表。

```html
<head>
<link rel="stylesheet" type="text/css" href="mystyle.css">
</head>
```

- `rel="stylesheet"`指定了被链接的资源是一个**样式表**。
- `type="text/css"`表示被链接的资源是CSS文件。
- `href="mystyle.css"`指定了CSS文件的路径。在这个例子中，CSS文件的路径为<u>相对路径</u>，即相对于当前HTML文件的位置。

通过将上述代码添加到HTML文档的`<head>`标签内，浏览器会根据给定的路径加载并应用名为`mystyle.css`的CSS文件，从而改变HTML文档中相关元素的样式。这样可以将样式与HTML内容分离，使得样式的管理和维护更加方便。

## <style>

```html
<style> 标签定义了HTML文档的样式文件引用地址.
在<style> 元素中你也可以直接添加样式来渲染 HTML 文档:
```

```html
<head>
<style type="text/css">
body {
    background-color:yellow;
}
p {
    color:blue
}
</style>
</head>
```

## <meta>

meta标签描述了一些基本的元数据。

<meta> 标签提供了元数据.元数据也不显示在页面上，但会被浏览器解析。

META 元素通常用于指定网页的描述，关键词，文件的最后修改时间，作者，和其他元数据。

元数据可以使用于浏览器（如何显示内容或重新加载页面），搜索引擎（关键词），或其他Web服务。

```html
<meta> 一般放置于 <head> 区域

为搜索引擎定义关键词:
<meta name="keywords" content="HTML, CSS, XML, XHTML, JavaScript">
为网页定义描述内容:
<meta name="description" content="免费 Web & 编程 教程">提供网页的简短描述，通常用于搜索引擎结果页面的显示。
定义网页作者:
<meta name="author" content="Runoob">
每30秒钟刷新当前页面:
<meta http-equiv="refresh" content="30">指定页面的自动刷新或重定向行为
```

是的，`<meta>`元素中包含的元数据通常不会直接在页面上显示出来，而是提供给浏览器、搜索引擎和其他Web工具使用的信息。这些信息可以影响页面的呈现方式、搜索引擎索引行为以及用户体验，但<mark>不会直接呈现为可见内容</mark>。

例如，设置字符编码的`<meta>`元素确保了页面中的文本能够以正确的编码方式显示，但这个元素本身并不在页面上显示出来。同样，描述网页内容的`<meta name="description" content="描述内容">`元素提供给搜索引擎<u>用于生成搜索结果摘要</u>，但它本身也不在页面上呈现为可见文本。

---

# 

# 十、HTML CSS

CSS (Cascading Style Sheets) 用于渲染HTML元素标签的样式。

## 如何使用CSS

> CSS 是在 HTML 4 开始使用的,是为了更好的渲染HTML元素而引入的.
> 
> CSS 可以通过以下方式添加到HTML中:
> 
> - 内联样式- 在HTML元素中使用"style" **属性**
> - 内部样式表 -在HTML文档头部 区域使用

## HTML样式实例 - 背景颜色

> 背景色属性（background-color）定义一个元素的背景颜色：

```html
<body style="background-color:yellow;">
 <h2 style="background-color:red;">这是一个标题</h2>
 <p style="background-color:green;">这是一个段落。
</p>
</body>
```

## HTML 样式实例 - 字体, 字体颜色 ，字体大小

> 我们可以使用font-family（字体），color（颜色），和font-size（字体大小）属性来定义字体的样式:

```html
<h1 style="font-family:verdana;">一个标题</h1>
<p style="font-family:arial;color:red;font-size:20px;">一个段落。</p>
```

> 现在通常使用font-family（字体），color（颜色），和font-size（字体大小）属性来定义文本样式，而不是使用<font>标签。

---

## HTML 样式实例 - 文本对齐方式

> 使用 text-align（文字对齐）属性指定文本的水平与垂直对齐方式：
> 
> ```html
> <h1 style="text-align:center;">居中对齐的标题</h1> <p>这是一个段落。</p>
> ```

## 内部样式表

> 当单个文件需要特别样式时，就可以使用内部样式表。你可以在<head> 部分通过 <style>标签定义内部样式表:
> 
> ```html
> <head>
> <style type="text/css">
> body {background-color:yellow;}
> p {color:blue;}
> </style>
> </head>
> ```

## !!!!!!!!!!!!外部样式表!!!

当样式需要被应用到很多页面的时候，外部样式表将是理想的选择。使用外部样式表，你就可以通过更改一个文件来改变整个站点的外观。

> ```html
> <head>
> <link rel="stylesheet" type="text/css" href="mystyle.css">
> </head>
> ```

<img title="" src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-01-31-22-39-45-image.png" alt="" width="512">先写css文档 然后在html的head部分链接进来 作为格式 可以重复使用<img title="" src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-01-31-22-38-14-image.png" alt="" width="511">

## HTML 样式标签

| 标签                                                    | 描述       |
| ----------------------------------------------------- | -------- |
| [<style>](https://www.runoob.com/tags/tag-style.html) | 定义文本样式   |
| [<link>](https://www.runoob.com/tags/tag-link.html)   | 定义资源引用地址 |

---

## 十一、图像

`<img>` 是空标签，意思是说，它只包含属性，并且没有闭合标签。

要在页面上显示图像，你需要使用源属性（src）。src 指 "source"。源属性的值是图像的 URL 地址。

**定义图像的语法是：**

`<img src="*url*" alt="*some_text*">`

URL 指存储图像的位置。如果名为 "pulpit.jpg" 的图像位于 www.runoob.com 的 images 目录中，那么其 URL 为 [http://www.runoob.com/images/pulpit.jpg](https://www.runoob.com/images/pulpit.jpg)。

浏览器将图像显示在文档中图像标签出现的地方。如果你将图像标签置于两个段落之间，那么浏览器会首先显示第一个段落，然后显示图片，最后显示第二段。

alt 属性用来为图像定义一串预备的可替换的文本。

替换文本属性的值是用户定义的。

`<img src="boat.gif" alt="Big Boat">`

在浏览器无法载入图像时，替换文本属性告诉读者失去的信息。

height（高度） 与 width（宽度）属性用于设置图像的高度与宽度。

属性值默认单位为像素:

```html
<img src="pulpit.jpg" alt="Pulpit rock" width="304" height="228">
```

**浮动实例**

```html
<p>
<img src="smiley.gif" alt="Smiley face" style="float:left" width="32" height="32"> 一个带图片的段落，图片浮动在这个文本的左边。
</p>

<p>
<img src="smiley.gif" alt="Smiley face" style="float:right" width="32" height="32"> 一个带图片的段落，图片浮动在这个文本的右边。
</p>
```

<img src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-01-31-22-51-25-image.png" title="" alt="" width="738">

**图片链接实例**

> ```html
> <p>创建图片链接:
> <a href="http://www.runoob.com/html/html-tutorial.html">
> <img  border="10" src="smiley.gif" alt="HTML 教程" width="32" height="32"></a></p>
> 
> <p>无边框的图片链接:
> <a href="http://www.runoob.com/html/html-tutorial.html">
> <img border="0" src="smiley.gif" alt="HTML 教程" width="32" height="32"></
> ```

> ![](C:\Users\13561\AppData\Roaming\marktext\images\2024-01-31-22-54-20-image.png)

**图片映射！！！！感觉会用到！！点击图片不同位置可以得到不同图片跳转**

> ```html
> <p>点击太阳或其他行星，注意变化：</p>
> 
> <img src="planets.gif" width="145" height="126" alt="Planets" usemap="#planetmap">
> 
> <map name="planetmap">
>   <area shape="rect" coords="0,0,82,126" alt="Sun" href="sun.htm">
>   <area shape="circle" coords="90,58,3" alt="Mercury" href="mercur.htm">
>   <area shape="circle" coords="124,58,8" alt="Venus" href="venus.htm">
> </map>
> ```

> <img src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-01-31-22-56-08-image.png" title="" alt="" width="255">

## HTML 图像标签

| 标签       | 描述            |
| -------- | ------------- |
| `<img>`  | 定义图像          |
| `<map>`  | 定义图像地图        |
| `<area>` | 定义图像地图中的可点击区域 |

---

## 

## 十二、表格

# HTML 表格

HTML 表格由 <table> 标签来定义。

HTML 表格是一种用于展示结构化数据的标记语言元素。

每个表格均有若干行（由 <tr> 标签定义），每行被分割为若干单元格（由 <td> 标签定义），表格可以包含标题行（<th>）用于定义列的标题。

- **tr**：tr 是 table row 的缩写，表示表格的<mark>一行</mark>。
- **td**：td 是 table data 的缩写，表示表格的<mark>数据</mark>单元格。
- **th**：th 是 table header的缩写，表示表格的<mark>表头</mark>单元格。

数据单元格可以包含文本、图片、列表、段落、表单、水平线、表格等等。

```html
<table>
  <thead>
    <tr>
      <th>列标题1</th>
      <th>列标题2</th>
      <th>列标题3</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>行1，列1</td>
      <td>行1，列2</td>
      <td>行1，列3</td>
    </tr>
    <tr>
      <td>行2，列1</td>
      <td>行2，列2</td>
      <td>行2，列3</td>
    </tr>
  </tbody>
</table>
```

- **<thead > 用于定义表格的标题部分:** 在 <thead > 中，使用 <th > 元素定义列的标题，以上实例中列标题分别为"列标题1"，"列标题2"和"列标题3"。

- **<tbody > 用于定义表格的主体部分:** 在 <tbody > 中，使用 <tr > 元素定义行，并在每行中使用 <td > 元素定义单元格数据，以上实例中有两行数据，每行包含三个单元格。

通过使用 <th > 元素定义列标题，可以使其在表格中以粗体显示，与普通单元格区分开来。

HTML 表格还可以具有其他部分，如 <tfoot > （表格页脚）和 <caption > （表格标题），<tfoot > 可用于在表格的底部定义摘要、统计信息等内容。 <caption > 可用于为整个表格定义标题。

HTML 表格还支持合并单元格和跨行/跨列的操作，以及其他样式和属性的应用，以满足各种需求。

我们也可以使用 CSS 来进一步自定义表格的样式和外观。



>  `<table border=" 数字 ">`可以设置表格边框宽度,数字为0则没有边框。
> 
> ```html
> <h4>水平标题:</h4>
> <table border="1">
> <tr>
>   <th>Name</th>
>   <th>Telephone</th>
>   <th>Telephone</th>
> </tr>
> <tr>
>   <td>Bill Gates</td>
>   <td>555 77 854</td>
>   <td>555 77 855</td>
> </tr>
> </table>
> 
> <h4>垂直标题:</h4>
> <table border="1">
> <tr>
>   <th>First Name:</th>
>   <td>Bill Gates</td>
> </tr>
> <tr>
>   <th>Telephone:</th>
>   <td>555 77 854</td>
> </tr>
> <tr>
>   <th>Telephone:</th>
>   <td>555 77 855</td>
> </tr>
> </table>
> ```

> 表格想要有标题，用`<caption></caption>`

> 表格想单独跨两行，可以用rouspan=“ 数字 ”
> 
> ```html
> <h4>单元格跨两行:</h4>
> <table border="1">
> <tr>
>   <th>First Name:</th>
>   <td>Bill Gates</td>
> </tr>
> <tr>
>   <th rowspan="2">Telephone:</th>
>   <td>555 77 854</td>
> </tr>
> <tr>
>   <td>555 77 855</td>
> </tr>
> </table>
> ```

> 单元格边距可以这样写：
> 
> ```html
> <h4>有单元格边距:</h4>
> <table border="1" 
> cellpadding="10">
> <tr>
>   <td>First</td>
>   <td>Row</td>
> </tr>   
> <tr>
>   <td>Second</td>
>   <td>Row</td>
> </tr>
> </table>
> ```



---



## 十三、列表

> <ul>是无序列表，<li>是每个列表始于的标签。
> 
> <ol>是有序列表标签
> 
> ```html
> <ol>
> <li>Coffee</li>
> <li>Milk</li>
> </ol>
> 浏览器中显示如下：
> 
> 1.Coffee
> 2.Milk
> ```

> 自定义列表从<dl>开始，每个列表项从<dt>开始，每个列表项的定义从<dd>开始。
> 
> > ```html
> > <dl>
> > <dt>Coffee</dt>
> > <dd>- black hot drink</dd>
> > <dt>Milk</dt>
> > <dd>- white cold drink</dd>
> > </dl>
> > ```

> <ol>标签里也可以写入type来定义排序符号。
> 
> > ```html
> > <h4>大写字母列表：</h4>
> > <ol type="A">
> >  <li>Apples</li>
> >  <li>Bananas</li>
> >  <li>Lemons</li>
> >  <li>Oranges</li>
> > </ol>
> > 得到
> > 大写字母列表：
> > A.Apples
> > B.Bananas
> > C.Lemons
> > D.Oranges
> > ```

> 嵌套列表
> 
> ```html
> <h4>嵌套列表：</h4>
> <ul>
>   <li>Coffee</li>
>   <li>Tea
>     <ul>
>       <li>Black tea</li>
>       <li>Green tea</li>
>     </ul>
>   </li>
>   <li>Milk</li>
> </ul>
> ```

---



## 十四、区块

```html
HTML 可以通过 <div> 和 <span>将元素组合起来。

HTML 区块元素
大多数 HTML 元素被定义为块级元素或内联元素。

块级元素在浏览器显示时，通常会以新行来开始（和结束）。

实例: <h1>, <p>, <ul>, <table>

HTML内联元素在显示时通常不会以新行开始。

实例: <b>, <td>, <a>, <img>


HTML <div> 元素是块级元素，它可用于组合其他 HTML 元素的容器。

<div> 元素没有特定的含义。除此之外，由于它属于块级元素，浏览器会在其前后显示折行。

如果与 CSS 一同使用，<div> 元素可用于对大的内容块设置样式属性。

<div> 元素的另一个常见的用途是文档布局。它取代了使用表格定义布局的老式方法。使用 <table> 元素进行文档布局不是表格的正确用法。<table> 元素的作用是显示表格化的数据。

HTML <span> 元素
HTML <span> 元素是内联元素，可用作文本的容器

<span> 元素也没有特定的含义。

当与 CSS 一同使用时，<span> 元素可用于为部分文本设置样式属性。


<div>	定义了文档的区域，块级 (block-level)
<span>	用来组合文档中的行内元素， 内联元素(inline)

```



> 用div把页面分成一块块，它本身只是用于标记，个人感觉，前面先预先定义（这个时候不用结束/div）某块风格然后名一个id，下文调用它。
> 
> ```html
> <div id="container" style="width:500px">
> 
> <div id="header" style="background-color:#FFA500;">
> 	
> <h1 style="margin-bottom:0;">主要的网页标题</h1></div>
> 
> <div id="menu" style="background-color:#FFD700;height:200px;width:100px;float:left;">
> <b>菜单</b><br>
> HTML<br>
> CSS<br>
> JavaScript</div>
> 
> <div id="content" style="background-color:#EEEEEE;height:200px;width:400px;float:left;">
> 内容在这里</div>
> 
> <div id="footer" style="background-color:#FFA500;clear:both;text-align:center;">
> 版权 © runoob.com</div>
> 
> </div>
> ```



---

## 十五、表单

```html
HTML 表单用于收集用户的输入信息。

HTML 表单表示文档中的一个区域，此区域包含交互控件，将用户收集到的信息发送到 Web 服务器。

HTML 表单通常包含各种输入字段、复选框、单选按钮、下拉列表等元素。

以下是一个简单的HTML表单的例子：

<form> 元素用于创建表单，action 属性定义了表单数据提交的目标 URL，method 属性定义了提交数据的 HTTP 方法（这里使用的是 "post"）。
<label> 元素用于为表单元素添加标签，提高可访问性。
<input> 元素是最常用的表单元素之一，它可以创建文本输入框、密码框、单选按钮、复选框等。type 属性定义了输入框的类型，id 属性用于关联 <label> 元素，name 属性用于标识表单字段。
<select> 元素用于创建下拉列表，而 <option> 元素用于定义下拉列表中的选项.
表单元素是允许用户在表单中输入内容，比如：文本域（textarea）、下拉列表（select）、单选框（radio-buttons）、复选框（checkbox） 等等。

```

```html
接下来我们介绍几种常用的输入类型。

一、
文本域通过 <input type="text"> 标签来设定，当用户要在表单中键入字母、数字等内容时，就会用到文本域。

实例
<form>
First name: <input type="text" name="firstname"><br>
Last name: <input type="text" name="lastname">
</form>

二、
密码字段通过标签 <input type="password"> 来定义:

实例
<form>
Password: <input type="password" name="pwd">
</form>

三、
单选按钮（Radio Buttons）
<input type="radio"> 标签定义了表单的单选框选项:

实例
<form action="">
<input type="radio" name="sex" value="male">男<br>
<input type="radio" name="sex" value="female">女
</form>

四、
复选框（Checkboxes）
<input type="checkbox"> 定义了复选框。

复选框可以选取一个或多个选项：

实例
<form>
<input type="checkbox" name="vehicle" value="Bike">我喜欢自行车<br>
<input type="checkbox" name="vehicle" value="Car">我喜欢小汽车
</form>

五、
提交按钮(Submit)
<input type="submit"> 定义了提交按钮。

当用户单击确认按钮时，表单的内容会被传送到服务器。表单的动作属性 action 定义了服务端的文件名。

action 属性会对接收到的用户输入数据进行相关的处理:

实例
<form name="input" action="html_form_action.php" method="get">
Username: <input type="text" name="user">
<input type="submit" value="Submit">
</form>

以上实例中有一个 method 属性，它用于定义表单数据的提交方式，可以是以下值：

post：指的是 HTTP POST 方法，表单数据会包含在表单体内然后发送给服务器，用于提交敏感数据，如用户名与密码等。

get：默认值，指的是 HTTP GET 方法，表单数据会附加在 action 属性的 URL 中，并以 ?作为分隔符，一般用于不敏感信息，如分页等。例如：https://www.runoob.com/?page=1，这里的 page=1 就是 get 方法提交的数据。
```

---



## 十六、框架

通过使用框架，你可以在同一个浏览器窗口中显示不止一个页面。

```html
iframe语法:

<iframe src="URL"></iframe>
该URL指向不同的网页。

iframe - 设置高度与宽度
height 和 width 属性用来定义iframe标签的高度与宽度。

属性默认以像素为单位, 但是你可以指定其按比例显示 (如："80%")。

实例
<iframe src="demo_iframe.htm" width="200" height="200"></iframe>


iframe - 移除边框
frameborder 属性用于定义iframe表示是否显示边框。

设置属性值为 "0" 移除iframe的边框:

实例
<iframe src="demo_iframe.htm" frameborder="0"></iframe>
```

```html
使用 iframe 来显示目标链接页面
iframe 可以显示一个目标链接的页面

目标链接的属性必须使用 iframe 的属性，如下实例:

实例
<iframe src="demo_iframe.htm" name="iframe_a"></iframe>
<p><a href="https://www.runoob.com" target="iframe_a" rel="noopener">RUNOOB.COM</a></p>
```

<img src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-02-01-15-58-21-image.png" title="" alt="" width="278">

<img src="file:///C:/Users/13561/AppData/Roaming/marktext/images/2024-02-01-15-56-24-image.png" title="" alt="" width="458">

> target就是指向对应name的声明，然后点击一下那个图片，会显示链接的页面。

---
















