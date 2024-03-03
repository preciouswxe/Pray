# $\color{#AAF}{||MySQL数据库学习||}$

---

学习视频：尚硅谷[MySQL数据库入门到大牛，mysql安装到优化，百科全书级，全网天花板_哔哩哔哩_bilibili](https://www.bilibili.com/video/BV1iq4y1u7vj/?spm_id_from=333.337.search-card.all.click&vd_source=16d869afecffa8da2167060c3de0b7c6)

---

## 一、初步认识

### 1.1为什么要使用数据库

持久化(persistence)：把数据保存到可掉电式存储设备中以供之后使用。大多数情况下，特别是企业级应用，数据持久化意味着将内存中的数据保存到硬盘上加以”固化”，而持久化的实现过程大多通过各种关系数据库来完成。
持久化的主要作用是将内存中的数据存储在关系型数据库中，当然也可以存储在磁盘文件、XML数据文件中。

### 1.2名称

DB：数据库（Database）
即存储数据的“仓库”，其本质是一个文件系统。它保存了一系列有组织的数据。
DBMS：数据库管理系统（Database Management System）
是一种操纵和管理数据库的大型软件，用于建立、使用和维护数据库，对数据库进行统一管理和控制。用户通过数据库管理系统访问数据库中表内的数据。
SQL：结构化查询语言（Structured Query Language）
专门用来与数据库通信的语言。

> 常见的数据库管理软件有Oracle、MySQL、MS SQL Server、DB2、PostgreSQL、Access、
> Sybase、Informix这几种。

### 1.3 介绍

MySQL6.x 版本之后分为 社区版 和 商业版 。
MySQL是一种关联数据库管理系统，将数据保存在不同的表中，而不是将所有数据放在一个大仓库
内，这样就增加了速度并提高了灵活性。
MySQL是开源的，所以你不需要支付额外的费用。
MySQL是可以定制的，采用了 GPL（GNU General Public License） 协议，你可以修改源码来
开发自己的MySQL系统。
MySQL支持大型的数据库。可以处理拥有上千万条记录的大型数据库。
MySQL支持大型数据库，支持5000万条记录的数据仓库，32位系统表文件最大可支持 4GB ，64位系
统支持最大的表文件为 8TB 。
MySQL使用 标准的SQL数据语言 形式。
MySQL可以允许运行于多个系统上，并且支持多种语言。这些编程语言包括C、C++、Python、
Java、Perl、PHP和Ruby等。

MySQL从5.7版本直接跳跃发布了8.0版本。

> 为什么如此多的厂商要选用MySQL？大概总结的原因主要有以下几点：
> 
> 1. 开放源代码，使用成本低。
> 2. 性能卓越，服务稳定。
> 3. 软件体积小，使用简单，并且易于维护。
> 4. 历史悠久，社区用户非常活跃，遇到问题可以寻求帮助。
> 5. 许多互联网公司在用，经过了时间的验证。

> Oracle 更适合大型跨国企业的使用，因为他们对费用不敏感，但是对性能要求以及安全性有更高的要
> 求。
> MySQL 由于其体积小、速度快、总体拥有成本低，可处理上千万条记录的大型数据库，尤其是开放源码
> 这一特点，使得很多互联网公司、中小型网站选择了MySQL作为网站数据库（Facebook，Twitter，
> YouTube，阿里巴巴/蚂蚁金服，去哪儿，美团外卖，腾讯）。

### 1.4 关系型数据库RDBMS

实质

> 这种类型的数据库是最古老的数据库类型，关系型数据库模型是把复杂的数据结构归结为简单的二元关系 （即二维表格形式）。
> 关系型数据库以 行(row) 和 列(column) 的形式存储数据，以便于用户理解。
> 
> 行列组成表，表组成库

优势

> 复杂查询 可以用SQL语句方便的在一个表以及多个表之间做非常复杂的数据查询。
> 事务支持 使得对于安全性能很高的数据访问要求得以实现。

### 1.5 非RDBMS

非关系型数据库

> 可看成传统关系型数据库的功能 阉割版本 ，基于键值对存储数据，不需要经过SQL层
> 的解析， 性能非常高 。同时，通过减少不常用的功能，进一步提高性能。

有哪些非关系型数据库

> 相比于 SQL，NoSQL 泛指非关系型数据库，包括了榜单上的键值型数据库、文档型数据库、搜索引擎和
> 列存储等，除此以外还包括图形数据库。也只有用 NoSQL 一词才能将这些技术囊括进来。

键值型数据库

> 键值型数据库通过 Key-Value 键值的方式来存储数据，其中 Key 和 Value 可以是简单的对象，也可以是复
> 杂的对象。Key 作为唯一的标识符，优点是查找速度快，在这方面明显优于关系型数据库，缺点是无法
> 像关系型数据库一样使用条件过滤（比如 WHERE），如果你不知道去哪里找数据，就要遍历所有的键，
> 这就会消耗大量的计算。
> 键值型数据库典型的使用场景是作为 内存缓存 。 Redis 是最流行的键值型数据库。

文档型数据库

> 此类数据库可存放并获取文档，可以是XML、JSON等格式。在数据库中文档作为处理信息的基本单位，
> 一个文档就相当于一条记录。文档数据库所存放的文档，就相当于键值数据库所存放的“值”。MongoDB
> 是最流行的文档型数据库。此外，还有CouchDB等。

搜索引擎数据库

> 虽然关系型数据库采用了索引提升检索效率，但是针对全文索引效率却较低。搜索引擎数据库是应用在
> 搜索引擎领域的数据存储形式，由于搜索引擎会爬取大量的数据，并以特定的格式进行存储，这样在检
> 索的时候才能保证性能最优。核心原理是“倒排索引”。
> 
> 百度、淘宝
> 
> 典型产品：Solr、Elasticsearch、Splunk 等。

列式数据库

> 列式数据库是<u>相对于行式存储</u>的数据库，<u>Oracle、MySQL、SQL Server 等数据库都是采用的行式存储</u>
> （Row-based），而列式数据库是将数据按照列存储到数据库中，这样做的好处是<u>可以大量降低系统的
> I/O（减少数据查询），适合于分布式文件系统</u>，不足在于功能相对有限。典型产品：HBase等。

图形数据库

> 图形数据库顾名思义，就是一种存储图形关系的数据库。它利用了图这种数据结构存储了实体（对象）
> 之间的关系。关系型数据用于存储明确关系的数据，但对于复杂关系的数据存储却有些力不从心。如社
> 交网络中人物之间的关系，如果用关系型数据库则非常复杂，用图形数据库将非常简单。典型产品：
> Neo4J、InfoGrid等。

由于 SQL 一直称霸 DBMS，因此许多人在思考是否有一种数据库技术能远离 SQL，于是 NoSQL 诞生了，
但是随着发展却发现越来越离不开 SQL。到目前为止 NoSQL 阵营中的 DBMS 都会有实现类似 SQL 的功
能。

NoSQL 对 SQL 做出了很好的补充，比如实际开发中，有很多业务需求，其实并不需要完整的关系型数据
库功能，非关系型数据库的功能就足够使用了。这种情况下，使用 性能更高 、 成本更低 的非关系型数据
库当然是更明智的选择。比如：日志收集、排行榜、定时器等。

### 1.6 关系型数据库设计

> 一个数据库中可以有多个表，每个表都有一个名字，用来标识自己。表名具有唯一性。
> 表具有一些特性，这些特性定义了数据在表中如何存储，类似Java和Python中 “类”的设计。

ORM思想（object relational mapping）

> 数据库中的一个表   <----> java或python中的一个表
> 
> 表中一条数据 <----> 类中的一个对象（或实体）
> 
> 表中一个列  <----> 类中的一个字段、水属性

表、记录、字段

> E-R（entity-relationship，实体-联系）模型中有三个主要概念是： 实体集 、 属性 、 联系集 。
> 一个实体集（class）对应于数据库中的一个表（table），一个实体（instance）则对应于数据库表中的一行（row），也称为一条记录（record）。一个属性（attribute）对应于数据库表中的一列（column），也称为一个字段（field）。

表的关联关系

> 表与表之间的数据记录有关系(relationship)。现实世界中的各种实体以及实体之间的各种联系均用关系模型来表示。
> 四种：一对一关联（学生信息表填写）、一对多关联（一门课对应很多学生成绩）、多对多关联（必须有第三个表叫联接表，举例如下）、自我引用

多对多举例：产品-订单

> “订单”表和“产品”表有一种多对多的关系，这种关系是通过与“订单明细”表建立两个一对多关系来
> 定义的。一个订单可以有多个产品，每个产品可以出现在多个订单中。
> 产品表 ：“产品”表中的每条记录表示一个产品。
> 订单表 ：“订单”表中的每条记录表示一个订单。
> 订单明细表 ：每个产品可以与“订单”表中的多条记录对应，即出现在多个订单中。一个订单
> 可以与“产品”表中的多条记录对应，即包含多个产品。

---

## 二、mysql安装

装社区版就行，开源免费，自由下载。集群版开源免费，用于社区版基础上。

企业版和高级集群版付费。

注意！

> 把data路径和server路径放在不同地址
> 
> server：D:\MySQL\MySQL8.0server  （其实是8.3server）
> 
> data位置：C:\ProgramData\MySQL\8.3data   （programdata是隐藏的文件夹）

> 8.3和5.7密码都是wxe138.。。。，端口号占用的是3306（8.3版本）和13306（5.7.29版本）

> mysql8.3版本服务名叫mysql83，另一个是mysql57

MySQL的登录
服务的启动与停止

> MySQL安装完毕之后，需要启动服务器进程，不然客户端无法连接数据库。
> 在前面的配置过程中，已经将MySQL安装为Windows服务，并且勾选当Windows启动、停止时，MySQL也
> 自动启动、停止。
> 方式1：使用图形界面工具
> 步骤1：打开windows服务
> 方式1：计算机（点击鼠标右键）→ 管理（点击）→ 服务和应用程序（点击）→ 服务（点
> 击）
> 方式2：控制面板（点击）→ 系统和安全（点击）→ 管理工具（点击）→ 服务（点击）
> 方式3：任务栏（点击鼠标右键）→ 启动任务管理器（点击）→ 服务（点击）
> 方式4：单击【开始】菜单，在搜索框中输入“services.msc”，按Enter键确认
> 步骤2：找到MySQL80（点击鼠标右键）→ 启动或停止

```
# 启动 MySQL 服务命令：
net start MySQL服务名
# 停止 MySQL 服务命令：
net stop MySQL服务名
```

> C:\Windows\System32>net stop mysql83
> MySQL83 服务正在停止.
> MySQL83 服务已成功停止。
> 
> C:\Windows\System32>mysql -uroot -pwxe13867187633
> mysql: [Warning] Using a password on the command line interface can be insecure.
> ERROR 2003 (HY000): Can't connect to MySQL server on 'localhost:3306' (10061)
> 
> C:\Windows\System32>net start mysql83
> MySQL83 服务正在启动 .
> MySQL83 服务已经启动成功。
> 
> C:\Windows\System32>mysql -uroot -pwxe13867187633
> mysql: [Warning] Using a password on the command line interface can be insecure.
> Welcome to the MySQL monitor.  Commands end with ; or \g.
> Your MySQL connection id is 8
> Server version: 8.3.0 MySQL Community Server - GPL
> 
> Copyright (c) 2000, 2024, Oracle and/or its affiliates.
> 
> Oracle is a registered trademark of Oracle Corporation and/or its
> affiliates. Other names may be trademarks of their respective
> owners.
> 
> Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.
> 
> mysql>

等价登录方式（进入8.3版本）

> <mark>C:\Windows\System32>mysql -u root -pwxe13867187633</mark>
> mysql: [Warning] Using a password on the command line interface can be insecure.
> Welcome to the MySQL monitor.  Commands end with ; or \g.
> Your MySQL connection id is 10
> Server version: 8.3.0 MySQL Community Server - GPL
> 
> Copyright (c) 2000, 2024, Oracle and/or its affiliates.
> 
> Oracle is a registered trademark of Oracle Corporation and/or its
> affiliates. Other names may be trademarks of their respective
> owners.
> 
> Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.
> 
> mysql> quit
> Bye
> 
> <mark>C:\Windows\System32>mysql -u root -p
> Enter password: ********</mark>******
> Welcome to the MySQL monitor.  Commands end with ; or \g.
> Your MySQL connection id is 11
> Server version: 8.3.0 MySQL Community Server - GPL
> 
> Copyright (c) 2000, 2024, Oracle and/or its affiliates.
> 
> Oracle is a registered trademark of Oracle Corporation and/or its
> affiliates. Other names may be trademarks of their respective
> owners.
> 
> Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.
> 
> mysql>

访问5.7的端口13306可以进入（5.7版本）

> <mark>C:\Windows\System32>mysql -u root -P13306 -p
> Enter password: ********</mark>******
> Welcome to the MySQL monitor.  Commands end with ; or \g.
> Your MySQL connection id is 7
> Server version: 5.7.29-log MySQL Community Server (GPL)

详细版——通过端口加主机（localhost或者127.0.0.1，h和localhost中间可以空格）<u>可以进入8.3版本</u>

> <mark>C:\Windows\System32>mysql -u root -hlocalhost -P3306 -p
> Enter password: ********</mark>******
> Welcome to the MySQL monitor.  Commands end with ; or \g.
> Your MySQL connection id is 12
> Server version: 8.3.0 MySQL Community Server - GPL

这里-P是端口port，先写h还是先写P是一样的，因为环境变量里面设置了8.3版本，所以真的便捷cmd登陆的时候可以不写端口和主机

---

## 三、使用演示

演示数据库cmd

> mysql> <mark>show databases;</mark>
> +--------------------+
> | Database           |
> +--------------------+
> | information_schema |
> | mysql              |
> | performance_schema |
> | sys                |
> +--------------------+
> 4 rows in set (0.00 sec)
> 
> mysql>

创建数据库 记得分号！！

> mysql> <mark>create database dbtest1;</mark>
> Query OK, 1 row affected (0.00 sec)
> 
> mysql> show databases;
> +--------------------+
> | Database           |
> +--------------------+
> | information_schema |
> | dbtest1            |
> | mysql              |
> | performance_schema |
> | sys                |
> +--------------------+
> 5 rows in set (0.00 sec)
> 
> mysql>

直接创建表，进行use，发现里面没有表，在新库dbtest1里面创建表employee，然后插入两条数据

记得分号！！

> <u>mysql> create table employee(id int,name varchar(15));</u>
> ERROR 1046 (3D000): No database selected
> <mark>mysql> use dbtest1;</mark>
> Database changed
> <mark>mysql> show tables;</mark>
> Empty set (0.00 sec)
> 
> <mark>mysql> create table employee(id int,name varchar(15));</mark>
> Query OK, 0 rows affected (0.02 sec)
> 
> <mark>mysql> show tables;</mark>
> +-------------------+
> | Tables_in_dbtest1 |
> +-------------------+
> | employee          |
> +-------------------+
> 1 row in set (0.00 sec)
> 
> <mark>mysql> select * from employee;</mark>       //查找是否有数据
> Empty set (0.00 sec)
> 
> <mark>mysql> insert into employee values(1001,'TOM');</mark>          //插入数据
> Query OK, 1 row affected (0.00 sec)
> 
> mysql> insert into employee values(1002),'afwefg');       //这里输入多了个括号所以报错
> ERROR 1064 (42000): You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near ''afwefg')' at line 1
> <mark>mysql>  insert into employee values(1002,'afwefg');</mark>
> Query OK, 1 row affected (0.01 sec)
> 
> <mark>mysql> select * from employee;</mark>
> +------+--------+
> | id   | name   |
> +------+--------+
> | 1001 | TOM    |
> | 1002 | afwefg |
> +------+--------+
> 2 rows in set (0.00 sec)     //展示了表单employee中的数据
> 
> mysql>

但是不能直接让名字是中文，会报错

> <mark>mysql> show create table employee;</mark>     //查看表单
> +----------+---------------------------------------------------------------------------------------------------------------------------------+
> | Table    | Create Table
>                      |
> +----------+---------------------------------------------------------------------------------------------------------------------------------+
> | employee | CREATE TABLE `employee` (
>   `id` int(11) DEFAULT NULL,
>   `name` varchar(15) DEFAULT NULL
> ) ENGINE=InnoDB DEFAULT <mark>CHARSET=latin1</mark> |
> +----------+---------------------------------------------------------------------------------------------------------------------------------+
> 1 row in set (0.00 sec)

上面的拉丁，因为起初是瑞士开发，所以不包含汉字，所以需要引入UTF-8（主流文字集合）才能

> <mark>mysql> show create database dbtest1;</mark>      //查看库
> +----------+--------------------------------------------------------------------+
> | Database | Create Database                                                    |
> +----------+--------------------------------------------------------------------+
> | dbtest1  | CREATE DATABASE `dbtest1` /*!40100 DEFAULT CHARACTER SET latin1 */ |
> +----------+--------------------------------------------------------------------+
> 1 row in set (0.00 sec)

编码命令

```
show variables like 'character_%';
show variables like 'collation_%';
```

查看属性

> <mark>mysql> show variables like 'character_%';</mark>
> +--------------------------+-----------------------------------------+
> | Variable_name            | Value                                   |
> +--------------------------+-----------------------------------------+
> | character_set_client     | gbk                                     |
> | character_set_connection | gbk                                     |
> | character_set_database   | <mark>latin1      </mark>                            |
> | character_set_filesystem | binary                                  |
> | character_set_results    | gbk                                     |
> | character_set_server     | <mark>latin1   </mark>                               |
> | character_set_system     | utf8                                    |
> | character_sets_dir       | D:\MySQL\MySQL5.7server\share\charsets\ |
> +--------------------------+-----------------------------------------+
> 8 rows in set, 3 warnings (0.00 sec)

说明一般默认使用的字符集是latin1

> <mark>mysql> show variables like 'collation_%';</mark>
> +----------------------+-------------------+
> | Variable_name        | Value             |
> +----------------------+-------------------+
> | collation_connection | gbk_chinese_ci    |
> | collation_database   |<mark> latin1_swedish_ci </mark>|
> | collation_server     | <mark>latin1_swedish_ci </mark>|
> +----------------------+-------------------+
> 3 rows in set, 3 warnings (0.00 sec)

对promdata文件夹里给mysql57的文件夹ini文件，对[mysql]和[mysqld]都加入utf8语句（两句有点不一样），成功把设置改成utf8字符集。

> <mark>mysql>  show variables like 'character_%';</mark>
> +--------------------------+-----------------------------------------+
> | Variable_name            | Value                                   |
> +--------------------------+-----------------------------------------+
> | character_set_client     | gbk                                     |
> | character_set_connection | gbk                                     |
> | character_set_database   | utf8mb4                                 |
> | character_set_filesystem | binary                                  |
> | character_set_results    | gbk                                     |
> | character_set_server     | <u>utf8</u>mb4                                 |
> | character_set_system     | utf8mb3                                 |
> | character_sets_dir       | D:\MySQL\MySQL8.0server\share\charsets\ |
> +--------------------------+-----------------------------------------+
> 8 rows in set, 6 warnings (0.02 sec)
> 
> <mark>mysql> show variables like 'collation_%';</mark>
> +----------------------+--------------------+
> | Variable_name        | Value              |
> +----------------------+--------------------+
> | collation_connection | gbk_chinese_ci     |
> | collation_database   | <u>utf8</u>mb4_0900_ai_ci |
> | collation_server     | utf8mb4_0900_ai_ci |
> +----------------------+--------------------+
> 3 rows in set, 6 warnings (0.00 sec)

出了小问题，因为是用记事本打开进行修改的ini文件，导致出现错误了，服务无法重启！！！nnd全网找不到解决办法

误打误撞把修改部分删掉，可以重新用5.7不用卸载了！！记得重启服务，这个时候是成功的。

所以专门下了notepad++，修改以后可以成功了。

> C:\Users\13561>mysql -u root -P13306 -hlocalhost -p
> Enter password: **************
> Welcome to the MySQL monitor.  Commands end with ; or \g.
> Your MySQL connection id is 2
> Server version: 5.7.29-log MySQL Community Server (GPL)
> 
> Copyright (c) 2000, 2024, Oracle and/or its affiliates.
> 
> Oracle is a registered trademark of Oracle Corporation and/or its
> affiliates. Other names may be trademarks of their respective
> owners.
> 
> Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.
> 
> <mark>mysql> show variables like 'character_%';</mark>
> +--------------------------+-----------------------------------------+
> | Variable_name            | Value                                   |
> +--------------------------+-----------------------------------------+
> | character_set_client     | gbk                                     |
> | character_set_connection | gbk                                     |
> | character_set_database   | utf8                                    |
> | character_set_filesystem | binary                                  |
> | character_set_results    | gbk                                     |
> | character_set_server     | utf8                                    |
> | character_set_system     | utf8                                    |
> | character_sets_dir       | D:\MySQL\MySQL5.7server\share\charsets\ |
> +--------------------------+-----------------------------------------+
> 8 rows in set, 3 warnings (0.00 sec)
> 
> <mark>mysql> show variables like 'collation_%'</mark>
>     -> ;
> +----------------------+-----------------+
> | Variable_name        | Value           |
> +----------------------+-----------------+
> | collation_connection | gbk_chinese_ci  |
> | collation_database   | utf8_general_ci |
> | collation_server     | utf8_general_ci |
> +----------------------+-----------------+
> 3 rows in set, 3 warnings (0.00 sec)

修改完以后就如以上属性。

这个时候原来的employee表还是latin的，要先drop删除掉database，才可以新建进行新的，新的库和表会变成utf8，能成功插入中文字符杰瑞。

> mysql> show create table employee;
> +----------+---------------------------------------------------------------------------------------------------------------------------------+
> | Table    | Create Table
>                      |
> +----------+---------------------------------------------------------------------------------------------------------------------------------+
> | employee | CREATE TABLE `employee` (
>   `id` int(11) DEFAULT NULL,
>   `name` varchar(15) DEFAULT NULL
> ) ENGINE=InnoDB DEFAULT <mark>CHARSET=latin1 </mark>| 之前演示的还留存着，要删除
> +----------+---------------------------------------------------------------------------------------------------------------------------------+
> 1 row in set (0.01 sec)
> 
> mysql><mark> drop database dbtest1;</mark>
> Query OK, 1 row affected (0.01 sec)
> 
> mysql> show databases;
> +--------------------+
> | Database           |
> +--------------------+
> | information_schema |
> | mysql              |
> | performance_schema |
> | sys                |
> +--------------------+
> 4 rows in set (0.00 sec)
> 
> mysql> <mark>create database dbtest1;</mark>
> Query OK, 1 row affected (0.00 sec)
> 
> mysql><mark> show create database dbtest1;</mark>
> +----------+------------------------------------------------------------------+
> | Database | Create Database                                                  |
> +----------+------------------------------------------------------------------+
> | dbtest1  | CREATE DATABASE `dbtest1` /*!40100 DEFAULT CHARACTER SET <mark>utf8</mark> */ |
> +----------+------------------------------------------------------------------+
> 1 row in set (0.00 sec)
> 
> mysql> <mark>use dbtest1;</mark>
> Database changed
> mysql> create table employees(id int,name varchar(15));
> Query OK, 0 rows affected (0.04 sec)
> 
> mysql> show create table employees;
> +-----------+--------------------------------------------------------------------------------------------------------------------------------+
> | Table     | Create Table
>                      |
> +-----------+--------------------------------------------------------------------------------------------------------------------------------+
> | employees | CREATE TABLE `employees` (
>   `id` int(11) DEFAULT NULL,
>   `name` varchar(15) DEFAULT NULL
> ) ENGINE=InnoDB DEFAULT <mark>CHARSET=utf8 |</mark>

和5.7不同的是，8.0和8.3是设置本就为utf8，如下

> mysql><mark> show variables like 'character_%';</mark>
> +--------------------------+-----------------------------------------+
> | Variable_name            | Value                                   |
> +--------------------------+-----------------------------------------+
> | character_set_client     | gbk                                     |
> | character_set_connection | gbk                                     |
> | character_set_database   | utf8mb4                                 |
> | character_set_filesystem | binary                                  |
> | character_set_results    | gbk                                     |
> | character_set_server     | utf8mb4                                 |
> | character_set_system     | utf8mb3                                 |
> | character_sets_dir       | D:\MySQL\MySQL8.0server\share\charsets\ |
> +--------------------------+-----------------------------------------+
> 8 rows in set, 6 warnings (0.03 sec)
> 
> mysql><mark> show variables like 'collation_%';</mark>
> +----------------------+--------------------+
> | Variable_name        | Value              |
> +----------------------+--------------------+
> | collation_connection | gbk_chinese_ci     |
> | collation_database   | utf8mb4_0900_ai_ci |
> | collation_server     | utf8mb4_0900_ai_ci |
> +----------------------+--------------------+
> 3 rows in set, 6 warnings (0.00 sec)
> 
> mysql><mark> show databases;</mark>
> +--------------------+
> | Database           |
> +--------------------+
> | information_schema |
> | mysql              |
> | performance_schema |
> | sys                |
> +--------------------+
> 4 rows in set (0.01 sec)

本身就是utf8，所以我们可以在新建的表里直接输入中文汤姆，不会报错

> <mark>mysql> create database dbtest2;</mark>
> Query OK, 1 row affected (0.01 sec)
> 
> mysql> show create database dbtest2;
> +----------+-----------------------------------------------------------------------------------------------------------------------------------+
> | Database | Create Database
>                        |
> +----------+-----------------------------------------------------------------------------------------------------------------------------------+
> | dbtest2  | CREATE DATABASE `dbtest2` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */ |
> +----------+-----------------------------------------------------------------------------------------------------------------------------------+
> 1 row in set (0.00 sec)
> 
> <mark>mysql> use dbtest2;</mark>
> Database changed
> <mark>mysql> create table emp(id int,name varchar(15));</mark>
> Query OK, 0 rows affected (0.03 sec)
> 
> <mark>mysql> show create table emp;</mark>
> +-------+----------------------------------------------------------------------------------------------------------------------------------------------------+
> | Table | Create Table
>                                      |
> +-------+----------------------------------------------------------------------------------------------------------------------------------------------------+
> | emp   | CREATE TABLE `emp` (
>   `id` int DEFAULT NULL,
>   `name` varchar(15) DEFAULT NULL
> ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci |
> +-------+----------------------------------------------------------------------------------------------------------------------------------------------------+
> 1 row in set (0.01 sec)
> 
> <mark>mysql> insert into emp values(1001,'汤姆');</mark>
> Query OK, 1 row affected (0.00 sec)
> 
> <mark>mysql> select * from emp;</mark>
> +------+------+
> | id   | name |
> +------+------+
> | 1001 | <mark>汤姆</mark> |
> +------+------+
> 1 row in set (0.01 sec)

---

## 四、语句

SQL语言在功能上主要分为如下3大类：

> <mark>DDL（Data Definition Languages、数据定义语言）</mark>，这些语句定义了不同的数据库、表、视图、索
> 引等数据库对象，还可以用来创建、删除、修改数据库和数据表的结构。
> 主要的语句关键字包括 CREATE 、 DROP 、 ALTER 等。
> <mark>DML（Data Manipulation Language、数据操作语言）</mark>，用于添加、删除、更新和查询数据库记录，并检查数据完整性。
> 主要的语句关键字包括 INSERT 、 DELETE 、 UPDATE 、 SELECT 等。
> SELECT是SQL语言的基础，最为重要。
> <mark>DCL（Data Control Language、数据控制语言）</mark>，用于定义数据库、表、字段、用户的访问权限和
> 安全级别。
> 主要的语句关键字包括 GRANT 、 REVOKE 、 COMMIT 、 ROLLBACK 、 SAVEPOINT 等。
> 因为查询语句使用的非常的频繁，所以很多人把查询语句单拎出来一类：DQL（数据查询语言）。
> 还有单独将 COMMIT 、 ROLLBACK 取出来称为TCL （Transaction Control Language，事务控制语
> 言）。

---

## 五、go连接mysql

做任务思路

1. 后端（Go）
   
   - 使用 `database/sql` 包和适当的数据库驱动程序连接到 MySQL 数据库。
   - 创建路由和处理程序，处理来自前端的请求并与数据库进行交互。
   - 将处理程序暴露为 RESTful API，以便前端可以通过 HTTP 请求进行访问。

2. 前端
   
   - 使用 HTML、CSS 和 JavaScript 构建用户界面。
   - 使用 AJAX 或 Fetch API 与后端的 API 进行通信，发送请求并接收响应。
   - 处理从后端接收的数据，并将其呈现在界面上。

大致思路：
1.写好go怎么连接mysql，写好增删改查操作。
2.编写前端界面，用action操作把前端提交的表单数据，由gin框架接收然后传到后端，用变量存储起来。
3.调用增删改查函数，里面变量就用上传的数据，如果单纯删除和查询应该是不用变量的，操作表。

---

## 六、基本的select语句

### SQL

（Structured Query Language，结构化查询语言）是使用关系模型的数据库应用语言， 与数据直
接打交道 ，由 IBM 上世纪70年代开发出来。后由美国国家标准局（ANSI）开始着手制定SQL标准，
先后有 SQL-86 ， SQL-89 ， SQL-92 ， SQL-99 等标准。
SQL 有两个重要的标准，分别是 SQL92 和 SQL99，它们分别代表了 92 年和 99 年颁布的 SQL 标
准，我们今天使用的 SQL 语言依然遵循这些标准。

----

SQL语言在功能上主要分为如下3大类：

> <mark>DDL（Data Definition Languages、数据定义语言）</mark>，这些语句定义了不同的数据库、表、视图、索
> 引等数据库对象，还可以用来<u>创建、删除、修改数据库和数据表</u>的结构。
> 主要的语句关键字包括<u> CREATE 、 DROP 、 ALTER </u>等。
> 
> ---
> 
> <mark>DML（Data Manipulation Language、数据操作语言）</mark>，用于添加、删除、更新和查询数据库记
> 录，并检查数据完整性。
> 主要的语句关键字包括 <u>INSERT 、 DELETE 、 UPDATE 、 SELECT</u> 等。
> SELECT是SQL语言的基础，最为重要。
> 
> ---
> 
> <mark>DCL（Data Control Language、数据控制语言）</mark>，用于定义数据库、表、字段、用户的访问权限和
> 安全级别。
> 主要的语句关键字包括 <u>GRANT 、 REVOKE 、 COMMIT 、 ROLLBACK 、 SAVEPOINT </u>等。
> 
> > 因为查询语句使用的非常的频繁，所以很多人把查询语句单拎出来一类：DQL（数据查询语言）。
> > 还有单独将 COMMIT 、 ROLLBACK 取出来称为TCL （Transaction Control Language，事务控制语
> > 言）。

---

基本规则
SQL 可以写在一行或者多行。为了提高可读性，各子句分行写，必要时使用缩进
每条命令以 ; 或 \g 或 \G 结束
关键字不能被缩写也不能分行
关于标点符号
必须保证所有的()、单引号、双引号是成对结束的
必须使用英文状态下的半角输入方式
字符串型和日期时间类型的数据可以使用单引号（' '）表示
列的别名，尽量使用双引号（" "），而且不建议省略as

---

SQL大小写规范 （建议遵守）
MySQL 在 Windows 环境下是大小写不敏感的
<u>MySQL 在 Linux 环境下是大小写敏感的</u>
数据库名、表名、表的别名、变量名是严格区分大小写的
关键字、函数名、列名(或字段名)、列的别名(字段的别名) 是忽略大小写的。
推荐采用统一的书写规范：
数据库名、表名、表别名、字段名、字段别名等都小写
SQL 关键字、函数名、绑定变量等都大写

---

注释

```
单行注释：#注释文字(MySQL特有的方式)
单行注释：-- 注释文字(--后面必须包含一个空格。)
多行注释：/* 注释文字 */
```

---

### 命名规则

数据库、表名不得超过30个字符，变量名限制为29个
必须只能包含 A–Z, a–z, 0–9, _共63个字符
数据库名、表名、字段名等对象名中间不要包含空格
同一个MySQL软件中，数据库不能同名；同一个库中，表不能重名；同一个表中，字段不能重名
必须保证你的字段没有和保留字、数据库系统或常用方法冲突。如果坚持使用，请在SQL语句中使
用`（着重号）引起来
保持字段名和类型的一致性，在命名字段并为其指定数据类型的时候一定要保证一致性。假如数据
类型在一个表里是整数，那在另一个表里可就别变成字符型了

```
#以下两句是一样的，不区分大小写
show databases;
SHOW DATABASES;
#创建表格
#create table student info(...); #表名错误，因为表名有空格
create table student_info(...);
#其中order使用``飘号，因为order和系统关键字或系统函数名等预定义标识符重名了
CREATE TABLE `order`(
id INT,
lname VARCHAR(20)
);
select id as "编号", `name` as "姓名" from t_stu; #起别名时，as都可以省略
select id as 编号, `name` as 姓名 from t_stu; #如果字段别名中没有空格，那么可以省略""
select id as 编 号, `name` as 姓 名 from t_stu; #错误，如果字段别名中有空格，那么不能省略""

```

数据导入指令
在命令行客户端登录mysql，使用source指令导入

`mysql> source d:\mysqldb.sql`

---

### 介绍select格式

```
SELECT 1; #没有任何子句
SELECT 9/2; #没有任何子句
```



select...from

```mysql
SELECT 标识选择哪些列
FROM 标识从哪个表中选择
```

如：一个表里选择全部列

```
SELECT *
FROM departments;
```

> 在生产环境下，不推荐你直接使用 SELECT * 进行查询。

```sql
SELECT department_id, location_id
FROM departments;
```

---

### 列的别名

重命名一个列
便于计算
紧跟列名，也可以在列名和别名之间加入关键字AS，别名使用双引号，以便在别名中包含空格或特
殊的字符并区分大小写。
AS 可以省略
建议别名简短，见名知意

```sql
SELECT last_name AS name, commission_pct comm
FROM employees;
```

去除重复行
默认情况下，查询会返回全部行，包括重复行。

```sql
SELECT department_id
FROM employees;
```

在SELECT语句中使用关键字DISTINCT去除重复行

```sql
SELECT DISTINCT department_id
FROM employees;
```


针对于：

```sql
SELECT DISTINCT department_id,salary
FROM employees;
```


这里有两点需要注意：

DISTINCT 需要放到所有列名的前面，如果写成 SELECT salary, DISTINCT department_id
FROM employees 会报错。

DISTINCT 其实是对后面所有列名的组合进行去重，你能看到最后的结果是 74 条，因为这 74 个部
门id不同，都有 salary 这个属性值。如果你想要看都有哪些不同的部门（department_id），只需
要写 DISTINCT department_id 即可，后面不需要再加其他的列名了。

---

### 查询常数

SELECT 查询还可以对常数进行查询。对的，就是在 SELECT 查询结果中增加一列固定的常数列。这列的
取值是我们指定的，而不是从数据表中动态取出的。
你可能会问为什么我们还要对常数进行查询呢？
SQL 中的 SELECT 语法的确提供了这个功能，一般来说我们只从一个表中查询数据，通常不需要增加一个
固定的常数列，但如果我们想整合不同的数据源，用常数列作为这个表的标记，就需要查询常数。
比如说，我们想对 employees 数据表中的员工姓名进行查询，同时增加一列字段 corporation ，这个
字段固定值为“尚硅谷”，可以这样写：

```sql
SELECT '尚硅谷' as corporation, last_name FROM employees;
```

---

### 显示表结构

使用DESCRIBE 或 DESC 命令，表示表结构。

```sql
DESCRIBE employees;
或
DESC employees;
```

> mysql> desc employees;
> +----------------+-------------+------+-----+---------+-------+
> | Field | Type | Null | Key | Default | Extra |
> +----------------+-------------+------+-----+---------+-------+
> | employee_id | int(6) | NO | PRI | 0 | |
> | first_name | varchar(20) | YES | | NULL | |
> | last_name | varchar(25) | NO | | NULL | |
> | email | varchar(25) | NO | UNI | NULL | |
> | phone_number | varchar(20) | YES | | NULL | |
> | hire_date | date | NO | | NULL | |
> | job_id | varchar(10) | NO | MUL | NULL | |
> | salary | double(8,2) | YES | | NULL | |
> | commission_pct | double(2,2) | YES | | NULL | |
> | manager_id | int(6) | YES | MUL | NULL | |
> | department_id | int(4) | YES | MUL | NULL | |
> +----------------+-------------+------+-----+---------+-------+
> 11 rows in set (0.00 sec)
> 
> 其中，各个字段的含义分别解释如下：
> Field：表示字段名称。
> Type：表示字段类型，这里 barcode、goodsname 是文本型的，price 是整数类型的。
> Null：表示该列是否可以存储NULL值。
> Key：表示该列是否已编制索引。PRI表示该列是表主键的一部分；UNI表示该列是UNIQUE索引的一
> 部分；MUL表示在列中某个给定值允许出现多次。
> Default：表示该列是否有默认值，如果有，那么值是多少。
> Extra：表示可以获取的与给定列有关的附加信息，例如AUTO_INCREMENT等。

---

### 过滤数据

语法：

```sql
SELECT 字段1,字段2
FROM 表名
WHERE 过滤条件
```

使用WHERE 子句，将不满足条件的行过滤掉
WHERE子句紧随 FROM子句
举例

```sql
SELECT employee_id, last_name, job_id, department_id
FROM employees
WHERE department_id = 90 
```

---

## 七、列类型

（1）整数类型
tinyint 迷你整型 一个字节 = 8位 范围 0-255；
smallint 小整形 2个字节 0-65535 ；
mediumint 中整形，3字节
int 整形 4字节；
bigint 大整形 8字节；
注意：因为整形有负数，所以是正负各一半。例如tinyint，实际上是-128-127；
无符号标识：类型后加上 unsigned；
填充长度：类型后加上 zerofill；左侧补零，使用后自动设置上unsigned。
//有符号整数
create table mytable (
int1 tinyint;
int2 smallint;
int3 mediumint;
int4 int;
int5 bigint;
);
desc mytable;
insert into mytable values(1,65534,255,255,255);
select * from mytable;

//无符号整数
create table mytable (
int1 tinyint unsigned,
int2 smallint unsigned,
int3 mediumint unsigned;,
int4 int unsigned,
int5 bigint unsigned,
);
desc mytable;
insert into mytable values(1,65534,255,255,255);
select * from mytable;
（2）小数类型
//浮点数 4字节
float：不指定小数位数
float(M,D):M个有效数字，小数占D位；

//定点数 整数部分一定准确，小数部分超出长度会四舍五入
decimal(M,D):M是总长度，不能超过65，D是小数长度，不能超过30

（3）时间
date:yyyy-mm-dd 3字节 10000-01-01到9999-12-12 初始值位0000-00-00；
time：hh:mm:ss 3字节,能表示-838：59：59~838：59：59，用于描述时间段；
datetime：yyyy-mm-dd  hh:mm:ss 8字节，能表示1000-01-01 00：00：00~9999-12-12 23：59：59，可以为0，表示0000-00-00 00：00：00；
timestamp：时间戳，当数据被修改，自动更新，不能为空
year：1900-2155

（4）字符串
//定长字符串 
char(L):L为0到255，中英文一样；
//变长字符串
varchar(L):L为0到65535，中英文一样,varchar要记录长度，所以会多使用1-2字节，数据小于127，则增加1字节，大于127就增加2字节。

//text 四种text，不用刻意选择 ，系统会自动选择合适的长度。
tinytext:1字节  长度为2^8+1;
text:2字节  长度为2^16+2;
mediumtext:3字节  长度为2^24+3;
longtext:4字节  长度为2^32+4;

//枚举类型 enum(值1，值2，...) 值为固定的几个选项，存储时是存的元素的下标，从1开始
create table mytb(
sex enum('male‘，’female‘）
）；

//set集合 将多个数据同时保存，1表示选中，0表示为为选中，自动计算存储空间大小
//1个字节对应8个选项，存储例如为选中（1，0，1，1，1，1，1，0），但最后在内存中会前后颠倒存放，变为（0，1，1，1，1，1，0，1）再转为10进制存储
create table mytb(
sex set('1‘，’2‘，'3'，'4'）
）；
注：mysql中会进行自动类型转换，遇到+ - * / 会自动将数据转换为数据，普通字符串转化为0.

### 列属性

> //列属性即字段属性有6个：null，默认值，列描述，主键，唯一键，自动增长
> 
> //null，允许为空；但尽量避免为空，如果有字段允许为空，则系统需要保留1字节储存null。
> //not null 不允许为空；
> create table mytable(
> name varchar(10) not null
> );
> 
> //default  val 默认值为val；
> //显示告知 insert into mytable values（default）；
> create table mytable(
> name varchar(10) default 'zs',
> );
> insert into mytable values(default);
> 
> //comment '注释说明'
> create table mytable(
> name varchar(10) comment '这是名字'，
> );
> //show create table mytable 查看说明
> 
> //主键 ，在一张表内有且只有一个字段做主键，且值惟一；
> //在列表字段后面加上 primary key
> //方式一 创建中的时候选定
> create table mytable(
> name varchar(10) comment '这是名字'，
> id int primary key,
> );
> //方式二 创建后选定
> create table mytable(
> name varchar(10) comment '这是名字'，
> id int,
> );
> //创建后增加主键
> alter table mytable add primary key(id);
> //查看主键--查看表结构
> desc mytable；
> //查看主键--查看创建语法
> show create table mytable;
> //删除逐渐
> alter table mytable drop primary key;
> 
> //复合主键
> create table mytable(
> primary key(std_id,course_id),//复合主键
> );
> 
> //主键约束：主键对应的字段不能为空，且不能重复
> //主键分类为：业务主键，具有业务意义；逻辑主键，自然增长的整形，应用广泛。
> 
> //自动增长：auto_increment,某字段给定该属性后，在没有提供数据时，会根据已存在的数据自动增长并填充，仅使用于数值类型，一张表最多一个自增长；
> 
> create table my_at(
> id int primary key auto_increment,
> name varchar(10)
> );
> //修改自增长的值
> alter table my_at auto_increment = 10;
> //删除自增长，字段后不再保留自增长属性
> alter table my_at modify id int;
> //系统中维护自增长的初始值和步长
> show variables like 'auto_increment%';
> //修改自增长
> set auto_increment_increment  = 10;

> //唯一键
> // unique key，保证字段中数据唯一
> //一张表内可以有多个唯一键，但只能有一个主键
> //唯一键允许为null，null可以有多个，null不参与比较
> //字段后后添加
> create table my_uq(
> id int primary key auto_increment,
> username varchar(10) unique
> );
> //表后添加
> create table my_uq(
> id int primary key auto_increment,
> username varchar(10) ，
> unique key(username)
> ) charset utf8;
> //创建后增加
> alter table my_uq add unique key(username);
> //查看唯一键
> desc my_uq;
> //删除唯一键
> alter table my_uq drop index username;
> //修改唯一键，先删除后增加

复合唯一键，使用多个字段共同保证唯一性。

---

## 八、增删改查

### 数据增删改查

（1）一次性插入多条
insert into mytable values('zs'),('ls'),('ww');

（2）主键冲突
在插入数据时，不确定是否已经存在主键，会导致错误

//主键冲突更新
insert into my_table values('id0001','zs') on duplicate key update name='zs';
//主键冲突替换
replace into my_table values('id0001','zs');

（3）蠕虫复制
成倍增长，从已有的数据中获取数据，并插入到数据表中,可能有主见冲突。

insert into my_table(name) select name from my_table;

（4）更新数据
注意加上where，否则是全部更新

update mytable set name = ’new_name' where name = 'zs';
//限制数量
update mytable set name = ’new_name' where name = 'zs' limit 4;

（5）删除数据
注意加上where，否则是全部删除，删除时无法充值auto_increment.

delete from mytable where name = 'zs';
//限制数量
delete from mytable where name = 'zs' limit 4;

//重置自增长
truncate mytable;

（5）查询数据
选项包括：all 全部，默认值；distinct 去重。所有字段都相同表示重复。
限制条件：

where 条件
group by 分组
having 条件
order by 排序
limit 限制


```sql
//指令为： select 选项 字段列表 from 数据表 where 条件 group by 分组 having 条件 order by 排序 limit 限制数量
select all name from mytable where name='zs' ;

//取别名。避免多张表命名冲突
select distinct name as name1 from mytable;

//从多张表同时取数据,得到记录数相乘，字段拼接的，即笛卡尔积。
select * from mytable1,mytable2;

//动态数据，from后面不是一个实体表，而是从一个表中查询的返回结果，再从里面查询
select * from （select name from mytable）as my_table2;

//where 子句，全表查询，保留符合条件的数据
select * from mytable where name ='zs';

//group by 字句，根据指定的字段进行分组统计，如果只想看数据显示，group by没特别的含义，分组后只会显示第一条数据
//count() ;统计每组的数量，如果统计目标字段则不统计null字段，如果为*，代表统计记录
//avg();求平均
//sum();求和
//max();求最大
//min();求最小
select  class_id,count(*),max(age),min(age),avg(age) from myclass group by class_id;

//group_concat(name) 将name 字段合并在一起
select  class_id,count(*)，group_concat(name),max(age),min(age),avg(age) from myclass group by class_id;

//多分组，按照某个字段分组之后再按照另一个字段分组，group by class_id,gender;
select  class_id,count(*)，group_concat(name),max(age),min(age),avg(age) from myclass group by class_id,gender;

//分组排序 group by class_id asc|desc,gender asc|desc
select  class_id,count(*)，group_concat(name),max(age),min(age),avg(age) from myclass group by class_id asc;

//回溯统计，多分组后往上统计需要层层上报，每次回溯将产生一个新的统计数据
//group by with rollup;
select  class_id,count(*)，group_concat(name),max(age),min(age),avg(age) from myclass group by class_id asc with rollup;


//having子句，和where一样，根据条件筛选
//having在group by 之后，可以针对分组数据进行筛选，where不行
//where 不能使用聚合函数
select class_id ,count(*) as number from mytable group by calss_id having count(*) >=3;


//order by 子句，根据校对规则对数据进行排序
//order by 字段 asc|desc，默认为asc
select  * from mytable order by age asc;
//order by 也可以多字段排序，按字段顺序依次排序
select  * from mytable order by age asc,class_id desc;

//limit 字句；限制数量，从第一条到指定的数量
select * from mytable limit 2;
//limit offset,length;
select * from mytable limit 1,2;

```



### 查询中的运算符

```sql
//+ - * / % 基础运算，通常不在在查询条件中使用，而是用于结果，其中除法的结果是浮点数
select age+age from mytable;

// > 、>=、 <、 <=、 =、 <> 比较运算符用于查询中做限定。
//<=> 相等比较
select * from mytable where age>=20;
//区间查找 between num1 and num2；
select * from mytable where age between 8 and 30;

//and 、or 、not 逻辑运算。
select * from mytable where age between 8 and 30 and class_id = 1;

//in 运算，表示在集合中出现过
select * from mytable where age in (10,20);

//is 运算，专门判断字段是否是null，is null  ； is not null.
select * from mytable where age is null;

//like 运算，模糊匹配，匹配字符串  
// _ :匹配单个字符
//%：匹配多个字符
select * from mytable where name like ’z%';

```



### 联合查询

联合查询是可以合并多个相似的查询结果，等同于一个表追加到另一个表，使用union 或者union all
是纵向合并，字段数不变
常见应用场景是：数据量大，会分表存储，从多张表取出数据，再合并显示
union理论上只需要拿到的字段数一致，并不需要字段一样，且只保留第一个select 语句的字段名。




```sql
//基本语法
select 语句1
union [选项]
select 语句2；
//选项包括 all 和distinct；
select name from mytable
union 
select age  from mytable;

//联合查询中使用order by子句，对应的select 语句需要用括号括起来,且 必须使用limit 数量 才会生效，通常取较大的数值。
（select name from mytable order by age asc limit 100000)
union 
select age  from mytable;


```



### 连接查询

将多张表连在一起查询，会导致字段和行数改变。  
原因是关系型数据库有一对一。一对多。多对多的关系，需要利用这个关系保证数据的完整性。

```sql
//交叉连接 记录数=table1*table2，字段数 = table1+table2，即笛卡尔积,没有实际应用
//table1 cross join table2
select * from able1 cross join table2;


//内连接 从一张表中取出所有的记录去另一张表进行匹配，利用匹配条件进行匹配，成功则保留，如果没有匹配条件就是交叉连接
//内连接通常用在对数据有精准要求的地方，必须保证两个表中能进行数据匹配。
//语法 select * from table1 inner join table2 on 条件
//无条件
 select * from table1 inner join table2 ；
 //有条件,避免同名字段干扰，通常用  表名.字段名  作为条件。
  select * from table1 inner join table2 on table1.class_id = table2.id;
//条件也可以使用where来限制，但一般使用on
  select * from table1 inner join table2 where table1.class_id = table2.id;
  //表名称较长的时候也可以使用别名
  select * from table1  as t1 inner join table2 as t2 where t1.class_id = t2.id;


//外连接，把一张表作为主表，他的所有数据都会保留，根据条件去另一张表获取目标数据，从表中未匹配到，则从表的字段全部设置为null
//分为 左外连接，即左表作为主表；右外连接，即右表做为主表
//非常常用的方法。
select * from table1 as t1 left join table2 as t2 on t1.class_id =t2.id;
select * from table1 as t1 right join table2 as t2 on t1.class_id =t2.id;

//using关键字，在使用on的地方使用using代替，必须保证字段同名，链接后的同名字段只会保留一个
//select * from table1 as t1 right join table2 as t2 using (同名字段列表)；


```



### 子查询

主查询：第一条select语句，确定所获取的数据源。

子查询： sub query,是指当一个查询时另一个查询的条件时，称为子查询。在一条select语句中嵌入另一条select语句，被嵌入的select 语句就是子查询语句。
子查询要么作为条件，要么作为数据源，子查询是独立存在的，是完整的select语句。

分类：
按功能分为：

标量子查询：子查询返回的结果是一个数据，即只有一行一列
列子查询：返回一列多行
行子查询：返回一行多列
表子查询：返回多行多列
exists子查询：返回结果为1或0
按位置分为：
where 子查询：子查询位置出现在where条件中
from子查询：子查询位置出现在from数据源中，作为数据源


```sql
//标量子查询，需求决定主查询，条件决定子查询。
//基本语法 select * from table1 where 条件判断 = (select 字段 from 数据源 where 条件判断)
select * from student where class_id = (select class_id from student where name='zs');

//列子查询
//语法：主查询 where 条件 in （列子查询）
select * from student where id in (select id from student where id between 1 and 3);

//行子查询
//语法：主查询 where 条件 (构造行元素) =（行子查询）,构造的行元素和子查询的数量必须一致。
select * from student where (stu_id,age) = (select max(stu_id),max(age) from student);

//表子查询,使用的是from
//语法：select 字段  from （表子查询）as 别名 [where ,group by having ...]
select * from (select * from student where class_id=1) as tb1 where age between 10 and 20;

//exists子查询
//语法 where exists(查询语句)，exists根据查询结果进行判断，如果存在就返回1，1的就保留
select name from student as  s1 where exists(select stu_id from student as s2 where s1.class_id = s2.stu_id);


```



子查询中的特定关键字：

```sql
in：主查询 where 条件 in （列子查寻）
select * from student where stu_id in (select stu_id from student where stu_id >=2);

any：任意一个
=any(子查询)，条件在查询结果中有任意一个匹配，等价于in
<>any(子查询),条件在查询结果中与任意一个都不匹配
select * from student where stu_id =any (select stu_id from student where stu_id >=2);
select * from student where stu_id <>any(select stu_id from student where stu_id >=2);

some：与any 一样。

all：
=all(子查询)：等于里面所有
<>=all(子查询)：不等于里面所有

数据
```





### 数据库备份和还原

```
整库备份，也叫做sql数据备份，备份的结果都是sql指令。
mysql中提供专门用于备份的客户端：mysqldump.exe
sql备份是常见备份方式，不止备份数据，还会备份sql指令，即使数据库完全毁坏也可以实现数据还原，但是被挖坟文件特别大，不适合大型数据备份。

```
