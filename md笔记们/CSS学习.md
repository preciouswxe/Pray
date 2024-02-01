# $\color{#FAA}{||CSS学习||}$

---

## 一、 什么是 CSS?

- CSS 指层叠样式表 (**C**ascading **S**tyle **S**heets)
- 样式定义**如何显示** HTML 元素
- 样式通常存储在**样式表**中
- 把样式添加到 HTML 4.0 中，是为了**解决内容与表现分离的问题**
- **外部样式表**可以极大提高工作效率
- 外部样式表通常存储在 **CSS 文件**中
- 多个样式定义可**层叠**为一个

> ## CSS 实例
> 
> CSS声明总是以分号 ; 结束，声明总以大括号 {} 括起来:
> 
> p {color:red;text-align:center;}
> 
> 为了让CSS可读性更强，你可以每行只描述一个属性:
> 
> ## 实例
> 
> p { color:red; text-align:center; }
> 
> ## CSS 注释
> 
> 注释是用来解释你的代码，并且可以随意编辑它，浏览器会忽略它。
> 
> CSS注释以 /* 开始, 以 */ 结束

---

## 二、id 和 class选择器

如果你要在HTML元素中设置CSS样式，你需要在元素中设置"id" 和 "class"选择器。

## id 选择器

id 选择器可以为标有特定 id 的 HTML 元素指定特定的样式。

HTML元素以id属性来设置id选择器,CSS 中 <mark>id 选择器以 "#" 来定义</mark>。

以下的样式规则应用于元素属性 id="para1":

## 实例

> ```html
> #para1
> {
>     text-align:center;
>     color:red;
> }
>  ID属性不要以数字开头，数字开头的ID在 Mozilla/Firefox 浏览器中不起作用。
> ```

## class 选择器

class 选择器用于描述一组元素的样式，class 选择器有别于id选择器，class可以在多个元素中使用。

class 选择器在 HTML 中以 class 属性表示, 在 CSS 中，类选择器以一个点 . 号显示：

> ```html
> 在以下的例子中，所有拥有 center 类的 HTML 元素均为居中。
> .center {text-align:center;}
> 
> 你也可以指定特定的 HTML 元素使用 class。
> 在以下实例中, 所有的 p 元素使用 class="center" 让该元素的文本居中:
> p.center {text-align:center;}
> 
> 
> 多个 class 选择器可以使用空格分开：
> .center { text-align:center; }
> .color { color:#ff0000; }
> 
> 类名的第一个字符不能使用数字！它无法在 Mozilla 或 Firefox 中起作用。
> ```

---

## 三、创建

外部样式表

> ```html
> 当样式需要应用于很多页面时，外部样式表将是理想的选择。在使用外部样式表的情况下，你可以通过改变一个文件来改变整个站点的外观。每个页面使用 <link> 标签链接到样式表。 <link> 标签在（文档的）头部：
> 
> <head>
> <link rel="stylesheet" type="text/css" href="mystyle.css">
> </head>
> ```

内部样式表

> ```html
> 当单个文档需要特殊的样式时，就应该使用内部样式表。你可以使用 <style> 标签在文档头部定义内部样式表，就像这样:
> 
> <head>
> <style>
> hr {color:sienna;}
> p {margin-left:20px;}
> body {background-image:url("images/back40.gif");}
> </style>
> </head>
> ```

内联样式

> ```html
> 将表现和内容混杂在一起，内联样式会损失掉样式表的许多优势。请慎用这种方法，例如当样式仅需要在一个元素上应用一次时。
> 
> 要使用内联样式，你需要在相关的标签内使用样式（style）属性。Style 属性可以包含任何 CSS 属性。本例展示如何改变段落的颜色和左外边距：
> 
> <p style="color:sienna;margin-left:20px">这是一个段落。</p>
> ```

多重样式

> ```html
> 如果某些属性在不同的样式表中被同样的选择器定义，那么属性值将从更具体的样式表中被继承过来。 
> 
> 例如，外部样式表拥有针对 h3 选择器的三个属性：
> h3
> {
>     color:red;
>     text-align:left;
>     font-size:8pt;
> }
> 
> 而内部样式表拥有针对 h3 选择器的两个属性：
> h3
> {
>     text-align:right;
>     font-size:20pt;
> }
> 假如拥有内部样式表的这个页面同时与外部样式表链接，那么 h3 得到的样式是：
> 
> color:red;
> text-align:right;
> font-size:20pt;
> 即颜色属性将被继承于外部样式表，而文字排列（text-alignment）和字体尺寸（font-size）会被内部样式表中的规则取代。
> ```

多重样式优先级

> 样式表允许以多种方式规定样式信息。样式可以规定在单个的 HTML 元素中，在 HTML 页的头元素中，或在一个外部的 CSS 文件中。甚至可以在同一个 HTML 文档内部引用多个外部样式表。
> 
> 一般情况下，优先级如下：
> 
> **（内联样式）Inline style > （内部样式）Internal style sheet >（外部样式）External style sheet > 浏览器默认样式**

---

## 四、背景+文本+字体

CSS 属性定义背景效果:

- background-color
- background-image 如`body {background-image:url('paper.gif');}`
- background-repeat 如水平铺开`background-repeat:repeat-x;`不想铺开就用`background-repeat:no-repeat;`
- background-attachment
- background-position 如`background-position:right top;`或者fixed

> 如何简写？
> 
> 当使用简写属性时，属性值的顺序为：:
> 
> - background-color
> - background-image
> - background-repeat
> - background-attachment
> - background-position
> 
> 以上属性无需全部使用，你可以按照页面的实际需要使用.
> 
> 如：
> 
> ```html
> body {background:#ffffff url('img_tree.png') no-repeat right top;}
> ```

## CSS字型

在CSS中，有两种类型的字体系列名称：

- **通用字体系列** - 拥有相似外观的字体系统组合（如 "Serif" 或 "Monospace"）
- **特定字体系列** - 一个特定的字体系列（如 "Times" 或 "Courier"）

| Generic family | 字体系列                          | 说明                         |
| -------------- | ----------------------------- | -------------------------- |
| Serif          | Times New Roman<br>Georgia    | Serif字体中字符在行的末端拥有额外的装饰     |
| Sans-serif     | Arial<br>Verdana              | "Sans"是指无 - 这些字体在末端没有额外的装饰 |
| Monospace      | Courier New<br>Lucida Console | 所有的等宽字符具有相同的宽度             |

## 字体系列

font-family 属性设置文本的字体系列。

font-family 属性应该设置几个字体名称作为一种"后备"机制，如果浏览器不支持第一种字体，他将尝试下一种字体。

**注意**: 如果字体系列的名称超过一个字，它必须用引号，如Font Family："宋体"。

多个字体系列是用一个逗号分隔指明：

> p{font-family:"Times New Roman", Times, serif;}

## 字体样式

主要是用于指定斜体文字的字体样式属性。

这个属性有三个值：

- 正常 - 正常显示文本
- 斜体 - 以斜体字显示的文字
- 倾斜的文字 - 文字向一边倾斜（和斜体非常类似，但不太支持）

> p.normal {font-style:normal;}  
> p.italic {font-style:italic;}  
> p.oblique {font-style:oblique;}

## 字体大小

font-size 属性设置文本的大小。

能否管理文字的大小，在网页设计中是非常重要的。但是，你不能通过调整字体大小使段落看上去像标题，或者使标题看上去像段落。

请务必使用正确的HTML标签，就<h1> - <h6>表示标题和<p>表示段落：

字体大小的值可以是绝对或相对的大小。

绝对大小：

- 设置一个指定大小的文本
- 不允许用户在所有浏览器中改变文本大小
- 确定了输出的物理尺寸时绝对大小很有用

相对大小：

- 相对于周围的元素来设置大小
- 允许用户在浏览器中改变文字大小

![Remark](https://www.runoob.com/images/lamp.gif) 如果你不指定一个字体的大小，默认大小和普通文本段落一样，是16像素（16px=1em）。

## 设置字体大小像素

```html
设置字体大小像素
设置文字的大小与像素，让您完全控制文字大小：

实例
h1 {font-size:40px;}
h2 {font-size:30px;}
p {font-size:14px;}
尝试一下 »
上面的例子可以在 Internet Explorer 9, Firefox, Chrome, Opera, 和 Safari 中通过缩放浏览器调整文本大小。

虽然可以通过浏览器的缩放工具调整文本大小，但是，这种调整是整个页面，而不仅仅是文本
```

## 用em设置字体大小

```html
为了避免Internet Explorer 中无法调整文本的问题，许多开发者使用 em 单位代替像素。

em的尺寸单位由W3C建议。

1em和当前字体大小相等。在浏览器中默认的文字大小是16px。

因此，1em的默认大小是16px。可以通过下面这个公式将像素转换为em：px/16=em

实例
h1 {font-size:2.5em;} /* 40px/16=2.5em */
h2 {font-size:1.875em;} /* 30px/16=1.875em */
p {font-size:0.875em;} /* 14px/16=0.875em */

在上面的例子，em的文字大小是与前面的例子中像素一样。不过，如果使用 em 单位，则可以在所有浏览器中调整文本大小。

不幸的是，仍然是IE浏览器的问题。调整文本的大小时，会比正常的尺寸更大或更小。
```

## 使用百分比和EM组合

```html
使用百分比和EM组合
在所有浏览器的解决方案中，设置 <body>元素的默认字体大小的是百分比：

实例
body {font-size:100%;}
h1 {font-size:2.5em;}
h2 {font-size:1.875em;}
p {font-size:0.875em;}

我们的代码非常有效。在所有浏览器中，可以显示相同的文本大小，并允许所有浏览器缩放文本的大小。
```

## 加粗字体

```html
<style>
p.normal {font-weight:normal;}
p.light {font-weight:lighter;}
p.thick {font-weight:bold;}
p.thicker {font-weight:900;}
</style>
</head>

<body>
<p class="normal">This is a paragraph.</p>
<p class="light">This is a paragraph.</p>
<p class="thick">This is a paragraph.</p>
<p class="thicker">This is a paragraph.</p>
</body>
```

## 设置字体转变

```html
<style>
p.normal {font-variant:normal;}
p.small {font-variant:small-caps;}
</style>
</head>

<body>
<p class="normal">My name is Hege Refsnes.</p>
<p class="small">My name is Hege Refsnes.</p>
</body>
```

## 简写字体属性在一个声明之内

```html
<style>
p.ex1
{
    font:15px arial,sans-serif;
}

p.ex2
{
    font:italic bold 12px/30px Georgia,serif;
}
</style>
```

## 所有CSS字体属性

| Property                                                                | 描述                 |
| ----------------------------------------------------------------------- | ------------------ |
| [font](https://www.runoob.com/cssref/pr-font-font.html)                 | 在一个声明中设置所有的字体属性    |
| [font-family](https://www.runoob.com/cssref/pr-font-font-family.html)   | 指定文本的字体系列          |
| [font-size](https://www.runoob.com/cssref/pr-font-font-size.html)       | 指定文本的字体大小          |
| [font-style](https://www.runoob.com/cssref/pr-font-font-style.html)     | 指定文本的字体样式          |
| [font-variant](https://www.runoob.com/cssref/pr-font-font-variant.html) | 以小型大写字体或者正常字体显示文本。 |
| [font-weight](https://www.runoob.com/cssref/pr-font-weight.html)        | 指定字体的粗细。           |

## 文本的对齐方式

```html
文本排列属性是用来设置文本的水平对齐方式。

文本可居中或对齐到左或右,两端对齐.

当text-align设置为"justify"，每一行被展开为宽度相等，左，右外边距是对齐（如杂志和报纸）。

实例
h1 {text-align:center;}
p.date {text-align:right;}
p.main {text-align:justify;}
```

## 文本修饰

```html
文本修饰
text-decoration 属性用来设置或删除文本的装饰。

从设计的角度看 text-decoration属性主要是用来删除链接的下划线：

实例
a {text-decoration:none;}

也可以这样装饰文字：

实例
h1 {text-decoration:overline;}
h2 {text-decoration:line-through;}
h3 {text-decoration:underline;}
```

## 文本转换

```html
文本转换
文本转换属性是用来指定在一个文本中的大写和小写字母。

可用于所有字句变成大写或小写字母，或每个单词的首字母大写。

实例
p.uppercase {text-transform:uppercase;} 全大写
p.lowercase {text-transform:lowercase;} 全小写
p.capitalize {text-transform:capitalize;} 每个首字母大写
```

## 文本缩进

```html
文本缩进属性是用来指定文本的第一行的缩进。

实例
p {text-indent:50px;}
```

## 所有CSS文本属性

| 属性                                                                            | 描述           |
| ----------------------------------------------------------------------------- | ------------ |
| [color](https://www.runoob.com/cssref/pr-text-color.html)                     | 设置文本颜色       |
| [direction](https://www.runoob.com/cssref/pr-text-direction.html)             | 设置文本方向。      |
| [letter-spacing](https://www.runoob.com/cssref/pr-text-letter-spacing.html)   | 设置字符间距       |
| [line-height](https://www.runoob.com/cssref/pr-dim-line-height.html)          | 设置行高         |
| [text-align](https://www.runoob.com/cssref/pr-text-text-align.html)           | 对齐元素中的文本     |
| [text-decoration](https://www.runoob.com/cssref/pr-text-text-decoration.html) | 向文本添加修饰      |
| [text-indent](https://www.runoob.com/cssref/pr-text-text-indent.html)         | 缩进元素中文本的首行   |
| [text-shadow](https://www.runoob.com/cssref/css3-pr-text-shadow.html)         | 设置文本阴影       |
| [text-transform](https://www.runoob.com/cssref/pr-text-text-transform.html)   | 控制元素中的字母     |
| [unicode-bidi](https://www.runoob.com/cssref/pr-text-unicode-bidi.html)       | 设置或返回文本是否被重写 |
| [vertical-align](https://www.runoob.com/cssref/pr-pos-vertical-align.html)    | 设置元素的垂直对齐    |
| [white-space](https://www.runoob.com/cssref/pr-text-white-space.html)         | 设置元素中空白的处理方式 |
| [word-spacing](https://www.runoob.com/cssref/pr-text-word-spacing.html)       | 设置字间距        |

### 垂直对齐图像

> ```html
> <style>
> img.top {vertical-align:text-top;}
> img.bottom {vertical-align:text-bottom;}
> </style>
> </head>
> 
> <body>
> <p>一个<img src="logo.png" alt="w3cschool" width="270" height="50" />默认对齐的图像。</p> 
> <p>一个<img class="top" src="logo.png" alt="w3cschool" width="270" height="50" />  text-top 对齐的图像。</p> 
> <p>一个<img class="bottom" src="logo.png" alt="w3cschool" width="270" height="50" /> text-bottom 对齐的图像。</p>
> </body>
> ```

### 指定字符间空间

> ```html
> <style>
> h1 {letter-spacing:2px;}
> h2 {letter-spacing:-3px;}
> </style>
> </head>
> 
> <body>
> <h1>This is heading 1</h1>
> <h2>This is heading 2</h2>
> </body>
> ```

### 指定书写方向

> ```html
> <style type="text/css">
> div.ex1 {direction:rtl;}
> </style>
> </head>
> <body>
> 
> <div>一些文本。 默认书写方向</div>
> <div class="ex1">一些文本。从右到左的书写方向。</div>
> ```

### 添加文本阴影

> ```html
> <style>
> h1 {text-shadow:2px 2px #FF0000;}
> </style>
> </head>
> <body>
> 
> <h1>Text-shadow 效果</h1>
> 
> <p><b>注意：</b> Internet Explorer 9 以及更早的浏览器不支持 text-shadow 属性。</p>
> ```

---

## 五、链接

> ## 链接样式
> 
> 链接的样式，可以用任何CSS属性（如颜色，字体，背景等）。
> 
> 特别的链接，可以有不同的样式，这取决于他们是什么状态。
> 
> 这四个链接状态是：
> 
> - a:link - 正常，未访问过的链接
> 
> - a:visited - 用户已访问过的链接
> 
> - a:hover - 当用户鼠标放在链接上时
> 
> - a:active - 链接被点击的那一刻
>   
>   > ```html
>   > a:link {color:#000000;}      /* 未访问链接*/
>   > a:visited {color:#00FF00;}  /* 已访问链接 */
>   > a:hover {color:#FF00FF;}  /* 鼠标移动到链接上 */
>   > a:active {color:#0000FF;}  /* 鼠标点击时 */
>   > 
>   > 
>   > 当设置为若干链路状态的样式，也有一些顺序规则：
>   > 
>   > a:hover 必须跟在 a:link 和 a:visited后面
>   > a:active 必须跟在 a:hover后面
>   > ```

改变的除了color，还可以是text-decoration，background-color，background，font-size

> ### 创建链接框
> 
> > ```html
> > <style>
> > a:link,a:visited
> > {
> >     display:block;
> >     font-weight:bold;
> >     color:#FFFFFF;
> >     background-color:#98bf21;
> >     width:120px;
> >     text-align:center;
> >     padding:4px;
> >     text-decoration:none;
> > }
> > a:hover,a:active
> > {
> >     background-color:#7A991A;
> > }
> > </style>
> > </head>
> > 
> > <body>
> > <a href="/css/" target="_blank">这是一个链接</a>
> > </body>
> > ```

---

## 六、列表

CSS 列表属性作用如下：

- 设置不同的列表项标记为有序列表
- 设置不同的列表项标记为无序列表
- 设置列表项标记为图像

---

## 列表

在 HTML中，有两种类型的列表：

- 无序列表 ul - 列表项标记用特殊图形（如小黑点、小方框等）
- 有序列表 ol - 列表项的标记有数字或字母

使用 CSS，可以列出进一步的样式，并可用图像作列表项标记。

### 无序列表如下所示:

- Coffee

- Tea

- Coca Cola

- Coffee

- Tea

- Coca Cola

### 有序列表如下所示:

1. Coffee

2. Tea

3. Coca Cola

4. Coffee

5. Tea

6. Coca Cola

# 

[所有不同的列表项标记](https://www.runoob.com/try/try.php?filename=trycss_list-style-type_all)  

list-style-type属性指定列表项标记的类型是：

## 实例

```html
ul.a {list-style-type: circle;}
ul.b {list-style-type: square;}

ol.c {list-style-type: upper-roman;}
ol.d {list-style-type: lower-alpha;}
```

一些值是无序列表，以及有些是有序列表。

## 作为列表项标记的图像

要指定列表项标记的图像，使用列表样式图像属性：

## 实例

ul { list-style-image: url('sqpurple.gif'); }

## 浏览器兼容性解决方案

同样在所有的浏览器，下面的例子会显示的图像标记：

## 实例

```html
ul
{
    list-style-type: none;
    padding: 0px;
    margin: 0px;
}
ul li
{
    background-image: url(sqpurple.gif);
    background-repeat: no-repeat;
    background-position: 0px 5px; 
    padding-left: 14px; 
}
```

例子解释：

- ul:
  - 设置列表类型为没有列表项标记
  - 设置填充和边距 0px（浏览器兼容性）
- ul 中所有 li:
  - 设置图像的 URL，并设置它只显示一次（无重复）
  - 您需要的定位图像位置（左 0px 和上下 5px）
  - 用 padding-left 属性把文本置于列表中

---

## 列表 - 简写属性

在单个属性中可以指定所有的列表属性。这就是所谓的简写属性。

为列表使用简写属性，列表样式属性设置如下：

## 实例

ul { list-style: square url("sqpurple.gif"); }

可以按顺序设置如下属性：

- list-style-type
- list-style-position (有关说明，请参见下面的CSS属性表)
- list-style-image

---

## 七、盒子模型

## CSS 盒子模型(Box Model)

所有HTML元素可以看作盒子，在CSS中，"box model"这一术语是用来设计和布局时使用。

CSS盒模型本质上是一个盒子，封装周围的HTML元素，它包括：边距，边框，填充，和实际内容。

盒模型允许我们在其它元素和周围元素边框之间的空间放置元素。

下面的图片说明了盒子模型(Box Model)：

![CSS box-model](https://www.runoob.com/images/box-model.gif)

不同部分的说明：

- **Margin(外边距)** - 清除边框外的区域，外边距是透明的。
- **Border(边框)** - 围绕在内边距和内容外的边框。
- **Padding(内边距)** - 清除内容周围的区域，内边距是透明的。
- **Content(内容)** - 盒子的内容，显示文本和图像。

## 元素的宽度和高度

![Remark](https://www.runoob.com/images/lamp.gif)**重要:** 当您指定一个 CSS 元素的宽度和高度属性时，你只是设置内容区域的宽度和高度。要知道，完整大小的元素，你还必须添加内边距，边框和外边距。

下面的例子中的元素的总宽度为 450px：

## 实例

div { width: 300px; border: 25px solid green; padding: 25px; margin: 25px; }

让我们自己算算：  
300px (宽)  

+ 50px (左 + 右填充)  
+ 50px (左 + 右边框)  
+ 50px (左 + 右边距)  
  = 450px

试想一下，你只有 250 像素的空间。让我们设置总宽度为 250 像素的元素:

## 实例

div { width: 220px; padding: 10px; border: 5px solid gray; margin: 0; }

最终元素的总宽度计算公式是这样的：

总元素的宽度=宽度+左填充+右填充+左边框+右边框+左边距+右边距

元素的总高度最终计算公式是这样的：

总元素的高度=高度+顶部填充+底部填充+上边框+下边框+上边距+下边距

---
