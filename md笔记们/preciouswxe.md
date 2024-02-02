王欣儿 学号23040905 组长:邵哥

# 最终结果展示（全屏情况）纯净版本

![](C:\Users\13561\AppData\Roaming\marktext\images\2024-01-31-23-34-16-image.png)

## 如果给fieldset添加颜色

![](C:\Users\13561\AppData\Roaming\marktext\images\2024-01-31-23-34-45-image.png)

## 如果在之上去掉背景

![](C:\Users\13561\AppData\Roaming\marktext\images\2024-01-31-23-35-14-image.png)

## 咳咳 其实还有蓝色版本

![](C:\Users\13561\AppData\Roaming\marktext\images\2024-01-31-17-13-34-image.png)

按键可以按动

 ![](C:\Users\13561\AppData\Roaming\marktext\images\2024-01-30-10-47-46-image.png) 

## 跳转部分

顶部的导航条也可以按动 会变成浅粉色 都设置了跳转链接

左上角logo是从实验室那边嫖下载的，没有设置跳转实验室的链接，底部的跳转链接可以跳转到实验室

随机输入账号密码 没有专门设置对错  会跳转到到我的vscode代码文件夹

<img title="" src="笔记图片/2024-01-30-10-54-08-image.png" alt="">

用户条款和忘记密码的跳转被我设置到一些相关的页面，知乎啊csdn啥的了

## 补充

时间不太够 没设计后端 可以后面写

所有图片都是被我放在我html学习文件夹一起的 所以要用代码试试登陆界面的话得下载好图片（我贴在最后面）

---

# 代码 第一个版本对应的

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>


    <style>
        p
        {
            text-align:center;
        } 
        body{
        top:0;
        width: 100%;
        height: 100%;
        object-fit: cover;
        background: url("下载 (8).png")no-repeat center center;
        /*background-repeat: no-repeat; /* 不重复 */
        /*background-position:  center center; /* 居中 */
        /*background-size: 100% 100%; /* 覆盖整个元素 */
        background-attachment: fixed;
        }
    </style>

       <!--设置孵化器登录大标题-->
    <style>
    .colortxt{
        font-size: 30px;
        background: linear-gradient(to right, #e4bbda, rgb(129, 36, 104));
        -webkit-background-clip: text;
        color: transparent;
        }
    </style>

        <!--设置整体表单-->
    <style type="text/css"> 
        fieldset
        {
            /*background-color: #ffff;*/
            opacity: 0.9;
            margin:30px auto;
            border-style:outset;
            width:600px;
            height:400px;
            padding:20px auto;
            color:rgb(0, 0, 0);
            text-align:center;
            box-shadow: 2px 2px 6px;
        }

        legend {
            color: rgb(202, 107, 164); /* 设置<legend>元素就是外边框的文本颜色为红色  但是没有采用已经删掉*/
                    }


        /*label部分CSS才是重点*/
        label
        {
            display: inline-block;
            width: 120px;
            text-align: left;
            text-align-last: justify;
            margin-right: 10px;
        }
    </style>

    <!--单独设置按钮-->
    <style>
    .button {
        display: inline-block;
        border-radius: 20px;
        background-color: #ad7296;
        border: none;
        color: #ffff;
        text-align: center;
        font-size: 18px;
        font-weight: 400;
        padding: 10px;
        width: 140px;
        transition: all 0.5s;
        cursor: pointer;
        margin: 5px;
        vertical-align: middle;
      }
      .button span {
        cursor: pointer;
        display: inline-block;
        position: relative;
        transition: 0.5s;
      }
      .button span::after {
        content: ">";
        position: absolute;
        opacity: 0;
        top: 0;
        right: -20px;
        transition: 0.5s;
      }
      .button:hover span {
        padding-right: 30px;
      }
      .button:hover span::after {
        opacity: 1;
        right: 0;
      }
    </style>

    <!--设置头部图片位置、色块和阴影以及点击效果-->
    <style>   
        .img{
            display: inline;
            float: left;
        }
        .topnav {
        /*overflow: hidden;*/
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        text-align: center;
        height: 45px;
        background-image: linear-gradient(270deg, #e38fcf,rgb(245, 242, 242));
        box-shadow:1px 2px 5px #867979;
        }

        /* 导航链接 */
        .topnav a {
        float: right;
        display: block;
        color: #f2f2f2;
        text-align: center;
        padding: 10px 16px;
        text-decoration: none;
    }

      /* 链接 - 修改颜色 */
      .topnav a:hover {
        background-color: #efd1e7;
        color: black;
    }
    </style>

</head>
<br>
    <!--插入的图片logo就在当前文件夹下-->    
    <body></body>

    <div class="topnav">
        <div class="img">
            <img src="logoIcon.png" style="height: 60px;position: relative; top: px; left: 250px;" alt="无法显示">
        </div>
        <a href="#">排名</a>
        <a href="#">评测记录</a>
        <a href="#">讨论</a>
        <a href="#">作业</a>
        <a href="#">比赛</a>
        <a href="#">训练</a>
        <a href="#">题库</a>
        <a href="#">首页</a>
        <a href="#"><p  style="color:#521154;position:fixed;top:0px;left:310px;">孵化器实验室之菜比冒险记</p></a>
      </div>


    <p class="container">
    <span class="colortxt"><b>孵化器三轮一期<br>登录界面制作</b></br></span></p>
    <form action="/" method="post">
        <fieldset>
            <!-- 文本输入框 -->

            <br><br>
            <!--<legend><b>用户登录</b></legend><br>-->
            <label for="name" style="position:fixed;left:670px;width:110px;" >账号:</label>
            <input type="text" id="name" name="name" placeholder="用户名、手机号、邮箱" 
            style="position:fixed;left:800px;width: 200px;height:22px;border-radius: 10px;" required>

            <br><br>

            <!-- 密码输入框 -->
            <label for="password" style="position:fixed;left:670px;width:110px;" >密码:</label>
            <input type="password" id="password" name="password" placeholder="密码" 
            style="position:fixed;left:800px;width:200px;height:22px;border-radius: 10px;" required>




            <br><br>

            <!-- 单选按钮 -->
            <label style="position: fixed;left:670px;width:90px;">登录方式 ：</label>

            <input type="radio" id="dl" name="fs" value="dl" checked>
            <label for="dl" style="width:40px;">登录</label>
            <input type="radio" id="dl" name="fs" value="zc">
            <label for="zc" style="width:80px;">自动注册</label>


        <br><br><br>
            <input type="checkbox" id="subscribe" name="subscribe" checked>
            <label for="subscribe" style="width:70px;">记住密码</label>

            <br><br>


            <br><br>

                <a href="https://blog.csdn.net/YouYou_GO/article/details/105574407"><i>忘记密码？</i></a>

            <!-- 提交按钮 -->

            <button class="button"><span>LOGIN IT</span></button>

            <br>

        <!-- 复选框 -->
            <input type="checkbox" id="subscribe" name="subscribe" style="margin-left: -300px;margin-top: 50px" checked>
            <label for="subscribe" style="width: 220px;">是否同意
            <a href="https://zhuanlan.zhihu.com/p/540432837"><i>用户协议和隐私政策</i></a></label>

            <p style="position:fixed;left:900px;top:497px;">其他方式登录：
                <a href="#"><img src="R-C (2).png" style="width: 30px;"></a>
                <a href="#"><img src="R-C (1).png" style="width: 30px;"></a>

            </p>


        </form>
    </fieldset>

        <a href="#"><img src="erweima.jpg" style="width: 320px;height:400px;position:fixed;left:200px;top:160px;"></a>
        <p style="width: 300px;height:400px;position:fixed;left:210px;top:500px;">请使用<a href="http://hdufhq.cn:8888">孵化器客户端</a>扫描二维码登录</p>

        <p style="position:fixed; bottom:30px; left: 780px;"><i>Designed by wxe</i></p>
        <p style="position:fixed; bottom:0px; left:750px;"><a href="http://hdufhq.cn:8888">！戳这里访问泥电孵化器地址</a></p>
</body>

</html>
```

---

# 使用图片

<img src="file:///D:/代码！/html学习/图片/erweima.jpg" title="" alt="erweima.jpg" width="616">

![logoIcon.png](D:\代码！\html学习\图片\logoIcon.png)

![R-C (1).png](D:\代码！\html学习\图片\R-C%20(1).png)

![R-C (2).png](D:\代码！\html学习\图片\R-C%20(2).png)

![下载 (8).png](D:\代码！\html学习\图片\下载%20(8).png)
