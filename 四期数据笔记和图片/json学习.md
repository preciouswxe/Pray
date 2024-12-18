# $\color{#FFA}{||json学习||}$

学习资料来自csdn

---

## 一、JSON 是什么？

JSON，全称是 JavaScript Object Notation，即 JavaScript对象标记法。
JSON是一种轻量级（Light-Meight)、基于文本的(Text-Based)、可读的(Human-Readable)格式。
JSON 的名称中虽然带有JavaScript，但这是指其语法规则是参考JavaScript对象的，而不是指只能用于JavaScript 语言。
JSON无论对于人，还是对于机器来说，都是十分便于阅读和书写的，而且相比 XML(另一种常见的数据交换格式)，文件更小，因此迅速成为网络上十分流行的交换格式。
近年来JavaScript已经成为浏览器上事实上的标准语言，JavaScript 的风靡，与JSON 的流行也有密切的关系。
因为JSON本身就是参考JavaScript 对象的规则定义的，其语法与JavaScript定义对象的语法几乎完全相同。
JSON格式的创始人声称此格式永远不升级，这就表示这种格式具有长时间的稳定性，10 年前写的文件，10年后也能用,没有任何兼容性问题。

## 二、JSON 的语法规则是怎样的？

JSON 的语法规则十分简单，可称得上“优雅完美”，总结起来有：

数组（Array）用方括号(“[]”)表示。
对象（0bject）用大括号(“{}”)表示。
名称/值对(name/value）组合成数组和对象。
名称(name）置于双引号中，值（value）有字符串、数值、布尔值、null、对象和数组。
并列的数据之间用逗号(“,”）分隔
{
    "name": "xdr630",
    "favorite": "programming"
}

---

JSON常被拿来与XML做比较，因为JSON 的诞生本来就多多少少要有取代XNL的意思。相比 XML，JSON的优势如下:

没有结束标签,长度更短,读写更快
能够直接被JavaScript解释器解析
可以使用数组
两者比较

JSON：

```json
{
 "name":"兮动人",
 "age":22,
 "fruits":["apple","pear","grape"]
}
```

XML：

```xml
<root> <name>兮动人</name>
 <age>22</age>
 <fruits>apple</fruits>
 <fruits>pear</fruits>
 <fruits>grape</fruits></root>
```

## 三、JSON的解析和生成（JSON 和 JS 对象互转）

在JavaScript中，有两个方法与此相关: JSON.parse和 JSON.stringify 。

首先是 JSON 和 JS 对象互转

要实现从JSON字符串转换为JS对象，使用 JSON.parse() 方法：

```js
<script>
    var str = '{"name": "兮动人","age":22}';
    var obj = JSON.parse(str);
    console.log(obj);
</script>
```

要实现从JS对象转换为JSON字符串，使用 JSON.stringify() 方法：

```js
<script>
    var str = '{"name": "兮动人","age":22}';
    var obj = JSON.parse(str);
    console.log(obj);

    var jsonstr = JSON.stringify(obj);
    console.log(jsonstr);
</script>
```

---

## 四、JSON格式规定

对象（0bject)

- 对象用大括号(“`{}`”）括起来，大括号里是一系列的“`名称/值对`”。

- 两个并列的数据之间用逗号(“`,`”）隔开，注意两点:
1. 使用英文的逗号（“`,`”)，不要用中文的逗号(“`，`”)
2. 最后一个“`名称/值对`“之后<u>不要加</u>逗号。

**JSON中Key/Value不能使用单引号**

名称/值对（Name/Value）

名称（Name）是一个字符串，要用双引号括起来，不能用单引号，也不能没有引号，这一点与JavaScript不同。
值的类型只有七种:字符串（string)、数值(number)、对象（object)、数组（array), true、false、null。不能有这之外的类型，例如undefined、函数等。

字符串（`string`）的规则如下:

1. 英文双引号括起来,不能用单引号，也不能没有。
2. 字符串中不能单独出现双引号（`”`）和右斜杠（“`\`")。
3. 如果要打双引号或右斜杠，需要使用“`右斜杠+字符`”的形式，例如`\”`和`\\`，其它的转义字符也是如此字符串的概念图。
- 转义字符

```html
{
"string":"\\ \" "
}
```

- 数值类型，可以使用科学计数法表示  

```html
{
"number":1e3,
"n1":1e2,
"n2":-100
}
```

---

## 五、字符串转化成js对象

方法

1. 使用`eval()`
2. 使用`JSON.parse()`
3. 使用第三方库,例如JQuery等

1、 `eval()`

- `eval(）`函数的参数是一个字符串，其作用是直接执行其中的 JavaScript代码。
- 案例：eval()解析字符串

```js
<script>
    var str = "console.log('hello')";
    eval(str);
</script>
```

- eval(）能够解析JSON字符串。从这里也可以看得出，JSON 和JavaScript是高度嵌合的。
- 案例：eval()解析JSON字符串

```js
<script>
	var str = '{"name":"兮动人","age":22}';
	var obj = eval("("+str+")");
	console.log(obj)
</script>

```

但是,现在已经很少直接使用eval()来解析了，如果您的浏览器版本真的是很旧，可能才需要这个方法。此外，eval()是一个相对危险的函数，因为字符串中可能含有未知因素。在这里，作为学习，还是要知道这也是一种方法。
请注意 eval()的参数，在字符串两旁加了括号，这是必须的，否则会报错。
因为JSON字符串是被大括号（“{}”）包围的，直接放到 eval()会被当成语句块来执行，因此要在两旁加上括号，使其变成表达式。


2. `JSON. parse()`
- 现在绝大多数浏览器都以支持`JSON.parse()`，**是推荐使用的方式**。
- 如果输入了不符合规范的字符串,会报错。

案例：JSON字符串转换为JS对象

```js
<script>
	var str = '{"name":"兮动人","age":22}';
	var obj = JSON.parse(str)
	console.log(obj)
</script>

```

JSON.parse()可以有第二个参数，是一个函数。此函数有两个参数:name和value，分别代表名称和值。

当传入一个JSON字符串后，JSON的每一组名称/值对都要调用此函数。该函数有返回值，返回值将赋值给当前的名称(name)。
利用第二个参数，可以在解析JSON字符串的同时对数据进行一些处理。

```js
<script>
    var str = '{"name":"兮动人","age":22}';
    var obj = JSON.parse(str,fun);
    function fun(name,value){
        console.log(name+":"+value);
        return value
    }
    console.log(obj)
</script>
```

- 可以做判断处理，当JSON字符串的 name=age 时，设置age的value=14

```js
<script>
	var str = '{"name":"兮动人","age":22}';
	var obj = JSON.parse(str,fun);
	function fun(name,value){
	if (name == "age")
		value = 14;
		return value
	}
 	console.log(obj)
</script>

```

---

## 六、js格式转换成字符串

- 序列化，指将 JavaScript 值转化为JSON字符串的过程。
- `JSON.stringify(）`能够将JavaScript值转换成JSON字符串。`JSON.stringify()`生成的字符串可以用`JSON.parse()`再还原成JavaScript值。

参数的含义

`JSON.stringify(value[, replacer[, space]])`
value:必选参数。被变换的JavaScript值，一般是对象或数组。
replace:可以省略。有两种选择:函数或数组。
如果是函数，则每一组名称/值对都会调用此函数，该函数返回一个值，作为名称的值变换到结果字符串中,如果返回undefined，则该成员被忽略。
案例：

```js
<script>
        var obj = {
            name: "兮动人",
            age: 22
        };
        console.log(obj);

        var jsonstr = JSON.stringify(obj,fun);

        function fun(name,value) {
            if (name=="age")
                value = 18;
                return value;
         }

        console.log(jsonstr)

</script>

```

- 如果是数组，则只有数组中存在名称才能够被转换，且转换后顺序与数组中的值保持一致。
- ```js
  <script>
          var obj = {
              a: 1,
              b: 2,
              c: 3,
              d: 4
          };
          console.log(obj);
  
          var jsonstr = JSON.stringify(obj,["a","b","c"]);
          
          console.log(jsonstr)
  </script>
  ```
  
  
- 把顺序改下，对应转换的JSON字符串的数值不变

```js
var jsonstr = JSON.stringify(obj,["c","a","b"]);

```

3. `space`:可以省略。这是为了排版、方便阅读而存在的。可以在JSON字符串中添加空白或制表符等。

```js
<script>
var obj = {
            name: "xxx",
            age: 22
        }
	console.log(obj);
	var jsonstr = JSON.stringify(obj);
	console.log(jsonstr)
</script>


```

- 当有不符合JSON语法规则时，就不会被转换成JSON字符串。
- 数组中有函数时会被转换成 `null`

```js
<script>
	var obj = {
            name: "xxx",
            age: 22,
            a: undefined,
            f: function () {

            },
            b:[function () {}]
        }
	console.log(obj);
	var jsonstr = JSON.stringify(obj);
	console.log(jsonstr)        
</script>


```

space的用法

- 案例：在上面的基础上添加

```js
<script>
        var obj = {
            a: 1,
            b: 2,
            c: 3,
            d: 4
        };
        console.log(obj);

        var jsonstr = JSON.stringify(obj,["c","a","b"],"one");

        console.log(jsonstr)
</script>

```

改成制表符：\t

```js
<script>
        var obj = {
            a: 1,
            b: 2,
            c: 3,
            d: 4
        };
        console.log(obj);

        var jsonstr = JSON.stringify(obj,["c","a","b"],"\t");

        console.log(jsonstr)
</script>
```


