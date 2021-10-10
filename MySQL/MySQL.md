# MySQL基础笔记

## 数据库基础操作



### 启动MySQL

``` shell
mysql.server start
```





### 导入 sqlName.sql

``` mysql
# 若不存在则创建
	## way1
CREATE DATABASE IF NOT EXISTS sqlName DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
	## way2
DROP DATABASE IF EXISTS tmall;
CREATE DATABASE tmall DEFAULT CHARACTER SET utf8;
# 使用该数据库
use sqlName;

```







| 1. 了解数据库和表 |                                                              |
| ----------------- | ------------------------------------------------------------ |
|                   | show databases;                                              |
|                   | show tables;                                                 |
|                   | show columns from customers;                                 |
|                   | desc customers;                                              |
|                   | show status;   -- 显示服务器状态                             |
|                   | show grants;   -- 显示用户权限                               |
|                   | show errors;    -- 显示服务器错误                            |
|                   | show warnings;  -- 显示服务器警告                            |
|                   |                                                              |
|                   | 2. 通配符                                                    |
|                   | select * from seat where student like 'a%';   -- 注意%不区分大小写 |
|                   | select * from seat where student like '_b%';  -- 下划线表示一个字符 |
|                   |                                                              |
|                   | 3. 正则表达式                                                |
|                   | select * from seat where student regexp '^[a-zA-Z]{5}$';     |
|                   |                                                              |
|                   | 4. 计算字段                                                  |
|                   | 1) concat                                                    |
|                   | select concat(student, '-', class) from courses ;            |
|                   | 2) rtrim                                                     |
|                   | select length(rtrim(class)) from courses;                    |
|                   |                                                              |
|                   | 5. 函数                                                      |
|                   | 1) 文本处理                                                  |
|                   | left	返回左边字符                                         |
|                   | length	返回串长度                                         |
|                   | locate	找到子串位置                                       |
|                   | lower	将串变成小写                                        |
|                   | ltrim	去除串左边的空格                                    |
|                   | right	返回右边                                            |
|                   | rtrim 	去除串右边的空格                                   |
|                   | soundex	返回串的soundex值                                 |
|                   | substring	返回子串                                        |
|                   | upper	将串转换成小写                                      |
|                   | 2) 时间处理                                                  |
|                   | adddate	增加日期                                          |
|                   | addtime	增加时间                                          |
|                   | curdate	当前日期                                          |
|                   | curtime	当前时间                                          |
|                   | date		返回时间的日期部分                               |
|                   | datediff	计算两个日期之差                                 |
|                   | date_format	返回格式化的日期字符串                        |
|                   | day		返回日期的天数部分                                |
|                   | dayofweek	返回星期几                                      |
|                   | hour	返回时间的小时部分                                   |
|                   | minute	返回时间的分钟部分                                 |
|                   | month	返回时间的月份部分                                  |
|                   | now		返回当前日期加时间                                |
|                   | second	返回当前时间的秒部分                               |
|                   | time		返回日期时间的时间部分                           |
|                   | year		返回日期的年份部分                               |
|                   | 3) 数值处理                                                  |
|                   | abs                                                          |
|                   | cos                                                          |
|                   | exp                                                          |
|                   | pi                                                           |
|                   | rand                                                         |
|                   | sin                                                          |
|                   | sqrt                                                         |
|                   | tan                                                          |
|                   | 4) 聚合                                                      |
|                   | avg                                                          |
|                   | count   注意，count(*)计算表的行数，count(字段)计算该字段不为NULL的行数 |
|                   | min                                                          |
|                   | max                                                          |
|                   | sum                                                          |
|                   |                                                              |
|                   | 6. 分组                                                      |
|                   | 7. 视图                                                      |
|                   | create view XXX as 查询sql;                                  |
|                   | 8. 存储过程                                                  |
|                   | 9. 游标                                                      |
|                   | 10. 触发器                                                   |





# [MySQL快速入门](https://www.cnblogs.com/Finley/p/5284865.html)

MySQL是一款流行的关系型数据库, 结构化查询语言SQL(Structured Query Language)是关系型数据库通用的操作语言, 但是不同数据库系统的行为还是有着细微的不同, 本文将以MySQL为例进行介绍.

一个MySql服务实例下可以维护多个database, 其中部分数据库用于维护用户和权限信息. Mysql提供了一些数据库管理语句:

- `SHOW DATABASES;` 显示当前实例所拥有的数据库
- `DROP DATABASES db_name;` 删除指定数据库
- `CREATE DATABASE [IF NOT EXISTS] db_name;` 创建数据库
- `use db_name;` 进入数据库

MySQL将数据存储在数据表中, 数据表由数据库来管理. 此外, 有一些命令可以查看数据表的元信息:

- `SHOW TABLES` 显示当前数据库所有数据表
- `DESC <table_name>` 或 `DESCRIBE <table_name>` 显示表结构

# DDL[#](https://www.cnblogs.com/Finley/p/5284865.html#ddl)

数据定义语言DDL(Data Definition Language)用于操作数据表等数据库对象，对数据表、索引等进行定义。

数据表table是行列的逻辑结构, 每行是一条记录, 每列是记录的一个字段. 同一张表的所有记录拥有相同的结构, 创建表时最重要的信息就是表的结构信息.

```sql
CREATE TABLE `user` (
	`id` INT PRIMARY KEY AUTO_INCREMENT,
	`username` VARCHAR(20) DEFAULT "",
	`password` CHAR(20) DEFAULT "",
	`gender` CHAR(1),
	`age` INT
);
```

SQL语句不区分大小写, 通常我们将SQL关键字大写, 自定义标识符小写. 用反点号包括自定义标识符可以避免它们被认为是SQL关键字.

SQL标识符以字母开头, 可以使用字母, 数字或三个特殊符号`#`, `_`, `$`.标识符用于表示表, 列和其它对象的名称.

所有DDL语句中的`DEFAULT`项都是可以省略的.

MySQL中常用的数据类型有:

- `INT`: 32位有符号整数
- `DOUBLE` 双精度浮点数
- `DECIMAL` 精确小数
- `CHAR(limit)`, 定长字符串
- `VARCHAR(limit)` 变长字符串
- `TEXT` 长文本数据
- `DATE` 日期
- `DATETIME` 日期和时间
- `TIMESTAMP` UNIX时间戳

使用`drop`语句来删除数据表:

```sql
DROP TABLE `user`;
```

表在创建完成后仍可以修改其结构, 不过要注意对其中数据的影响:

- 删除列

```sql
ALTER TABLE `user` DROP `age`;
```

- 增加列:

```sql
ALTER TABLE `user` ADD `age` INT DEFAULT 0;
```

- 修改列的类型或名称

```sql
ALTER TABLE `user` CHANGE `gender` `sex` CHAR(1) DEFAULT "M";
ALTER TABLE `user` MODIFY `gender` INT DEFAULT 0;
```

# DML[#](https://www.cnblogs.com/Finley/p/5284865.html#dml)

数据操作语言DML(Data Manipulation Language), 用于操作表中的数据。

插入一条记录:

```sql
INSERT INTO `user` (`username`, `password`) VALUES ("abcd", "1234");
INSERT INTO `user` VALUES ("abcd", "1234");
```

要向所有没有默认值且不允许为空的列插入数据.

更新已存在的记录:

```sql
UPTATE `user` SET `username`="abc" WHERE `id`=1; 
```

`update`可以更新所有符合`WHERE`子句的记录, 当没有`WHERE`子句时更新该表所有记录.

当键重复时更新记录, 否则插入记录:

```sql
INSERT INTO `user` VALUES ("a", "2") ON DUPLICATE KEY UPDATE `username`="a", "password"=2; 
```

删除记录:

```sql
DELETE FROM `user` WHERE `id`=1;
```

删除所有符合`WHERE`子句的记录, 当没有`WHERE`子句时删除该表所有记录.

# SELECT[#](https://www.cnblogs.com/Finley/p/5284865.html#select)

查询语句SELECT是最灵活最复杂也是最重要的SQL语句.

我们创建用户`user`和文章`post`两张表做为示例:

```sql
CREATE TABLE `user` (
	`id` INT PRIMARY KEY AUTO_INCREMENT,
	`username` VARCHAR(20) DEFAULT "",
	`password` CHAR(20) DEFAULT "",
	`gender` CHAR(1),
	`age` INT
);
CREATE TABLE `post` (
	id INT PRIMARY KEY AUTO_INCREMENT,
	uid INT,
	title TEXT,
	CONTENT TEXT
);
```

查询`user`表中所有记录的所有字段.

```sql
SELECT * FROM `user`;
```

SELECT可以可以查询某个表中符合条件记录的某些字段.如所有男性用户的username:

```sql
SELECT `username`
FROM `user`
WHERE `gender`='M';
```

我们可以使用LIMIT限制返回结果的数量:

```sql
SELECT * FROM `user` LIMIT 10 OFFSET 10;
```

上述查询跳过前10个(OFFSET 10)结果， 并返回最多10条记录(LIMIT 10)。

在多表查询时可能需要指定列的所在的表:

```sql
SELECT `user`.`username`
FROM `user`
WHERE `user`.`gender`='M';
```

为了简化表达式可以为列和表指定别名:

```sql
SELECT u.`username` AS name
FROM `user` u
WHERE u.`gender`='M';
```

表中可能出现用户名相同的情况, 使用`DISTINCT`函数去除重复的用户名:

```sql
SELECT DISTINCT(`username`)
FROM `user`;
```

## 聚集[#](https://www.cnblogs.com/Finley/p/5284865.html#聚集)

MySQL可以进行取平均值, 求和等聚集操作.

```sql
SELECT AVG(`age`) FROM `user`;
```

上述语句查询`user`表中所有记录`age`字段的平均值. 可以使用`GROUP BY`子句指定分组聚集方式.

```sql
SELECT AVG(`age`) FROM `user`
GROUP BY `gender`;
```

上述语句将`gender`字段相同的记录视作一组, 求出`age`字段的平均值. 我们可以附加where子句查询男性用户的平均年龄:

```sql
SELECT gender, AVG(`age`) FROM `user`
WHERE gender=`M`
GROUP BY gender;
```

`WHERE`子句中无法使用聚集函数, 可以使用`HAVING`子句完成该功能, 比如查询发表文章超过3篇的作者:

```sql
SELECT uid FROM `post`
HAVING count(id) > 3
GROUP BY uid;
```

常用的聚集函数有:

- `SUM` 求和
- `AVG` 求平均值
- `MIN` 求最小值
- `MAX` 求最大值
- `COUNT` 求记录数

此外还有一些工具函数也顺便介绍:

- `LENGETH(txt)`, `LEN(txt)`: 求文本长度
- `REPLACE(txt, from, to)` 替换文本
- `UCASE(txt)`, `LCASE(txt)`: 转为大写 / 小写
- `mid(txt, start, len)`: 提取子串

## 排序[#](https://www.cnblogs.com/Finley/p/5284865.html#排序)

ORDER BY 语句可以根据结果中某一列进行排序:

```sql
SELECT * FROM user ORDER BY age DESC LIMIT 10;
```

把user表中的记录根据age字段进行降序排列， 并返回前10条记录。

```sql
SELECT * FROM user ORDER BY age ASC LIMIT 10;
```

把user表中的记录根据age字段进行升序排列， 并返回前10条记录。

ORDER BY 语句可以进行多列排序:

```sql
SELECT * FROM user ORDER BY age DESC, username ASC;
```

把user表中的记录根据age字段进行升序排列，age相同的根据username字典序升序排列。

ORDER BY 语句可以根据聚集函数进行排序:

```sql
SELECT post_id, count(*) as comment_count 
FROM comment
GROUP BY post_id
ORDER BY count(*) DESC;
```

根据文章的评论数进行降序排列。

## 联合查询[#](https://www.cnblogs.com/Finley/p/5284865.html#联合查询)

现在查询所有女性用户发表的文章:

```sql
SELECT p.`id` AS post_id, `title`, `content`, u.`username` AS author
FROM `user` u, `post` p
WHERE
	p.`uid`=u.`id`
	AND u.`gender`='F'; 
```

> 字符串可以使用单引号`'`或双引号`"`表示.

上述查询将搜索`user`和`post`表中记录所有组合, 满足条件`p.uid=u.id`的`user`记录和`post`记录将组合为同一条记录. 然后查询组合记录中符合条件`u.`gender`='F'`的记录.

使用等价的`join`操作来代替上述查询:

```sql
SELECT p.`id` AS post_id, title, content, u.`username` AS author 
FROM `post` p
JOIN `user` u 
ON p.`uid`=u.`id`
WHERE u.`gender`='F'; 
```

`join`操作有3种:

- `INNER JOIN`, `JOIN`: 查询出的记录必须满足`ON`条件，若左表或右表中没有对应的记录，查询结果中均不会包含相应记录。
- `LEFT JOIN`: 查询左表`post`中所有记录, 即使右表中没有对应的记录. 当没有右表记录时, 查询结果中右表的字段为空值.
- `REIGHT JOIN` 查询右表`user`中所有记录, 即使左表中没有对应记录. 当没有左表记录时, 查询结果中左表的字段为空值.
- `OUTER JOIN`: `LEFT JOIN` 和 `RIGHT JOIN` 结果的并集，即左右表只有一个表中有记录即可。

注意若`user`表中有m条记录, `post`表中有n条记录则连接查询将会扫描`m*n`条记录. 这在表中数据量较大时会非常耗时.

## IN 和 EXISTS 查询[#](https://www.cnblogs.com/Finley/p/5284865.html#in-和-exists-查询)

使用`IN`和`EXISTS`子查询可以完成连接查询的任务.

```sql
SELECT *
FROM `post`
WHERE `uid` IN (
	SELECT `user`.`id` FROM `user` WHERE `gender`='F'
);
```

`IN`运算符将匹配`uid`和子查询的结果, 若记录的`uid`在结果集中则`IN`条件为真.

`IN`运算符后可以直接书写列表:

```sql
SELECT * FROM user
WHERE username in ('a', 'b');
```

查询username为'a'或'b'的用户。

`EXISTS`运算符可以起到类似的作用:

```sql
SELECT *
FROM `post`
WHERE EXISTS (
	SELECT `user`.`id` FROM `user` WHERE `gender`='F' 
	AND `user`.`id`=`post`.`uid`
); 
```

在EXISTS子查询中可以访问外表，如上例exists子查询中访问了`post`表。

上述SQL语句查询使EXISTS子查询`SELECT user.id FROM user WHERE gender='F' AND user.id=post.uid` 结果不为空集的`post`记录。即查询女性作者发表的所有文章，与本节开头的IN查询作用相同。

`IN`和`EXISTS`都有对应的否定形式:

```sql
SELECT *
FROM `post`
WHERE `uid` NOT IN (
	SELECT `user`.`id` FROM `user` WHERE `gender`='F'
);
SELECT *
FROM `post`
WHERE NOT EXISTS (
	SELECT `user`.`id` FROM `user` WHERE `gender`='F' 
	AND `user`.`id`=`post`.`uid`
); 
```

## 子查询[#](https://www.cnblogs.com/Finley/p/5284865.html#子查询)

`SELECT * FROM`语句可以把另外一个查询结果作为数据源。

添加两张表:

```sql
CREATE TABLE `likes` (
	`uid` INT,
	`post_id` INT,
	PRIMARY KEY (`uid`, `post_id`) 
);

CREATE TABLE `comment` (
	`id` INT PRIMARY KEY AUTO_INCREMENT,
	`uid` INT,
	`post_id` INT,
	`content` TEXT,
);
```

我们要搜索所有文章的点赞`like`数和评论`comment`数:

```sql
SELECT 
	lc.post_id,
	lc.like_count,
	cc.comment_count
FROM (
	SELECT `post_id`, count(*) AS like_count FROM `likes` GROUP BY `post_id`
) lc
OUTER JOIN (
	SELECT `post_id`, count(*) AS comment_count FROM `comment` GROUP BY `post_id`
) cc
ON lc.post_id = cc.post_id;
```

上述语句中我们将查询语句`SELECT post_id, count(*) AS like_count FROM likes GROUP BY post_id`的结果作为一个数据表， 并给它一个别名`lc`。

类似地，我们将另一个查询的结果作为数据表`cc`, 然后将两个数据表进行JOIN。

## 视图[#](https://www.cnblogs.com/Finley/p/5284865.html#视图)

MySQL可以根据查询创建视图`view`对象, 可以用访问数据表的方式访问视图. 视图只保存查询操作不保存数据, 视图在被访问时从数据表中取出数据.

```sql
CREATE VIEW `post_detail` AS
SELECT p.`id` AS post_id, title, content, 
	p.`uid` AS author_id, u.`username` AS author_name 
FROM `post` p
JOIN `user` u 
ON p.`uid`=u.`id`
WHERE u.`gender`='F'; 
SELECT * FROM `post_detail`;
```

# 约束[#](https://www.cnblogs.com/Finley/p/5284865.html#约束)

约束是在表上强制执行的校验规则, 用于保证数据的完整性. 约束分为对单列的约束和对多个列集合的约束.

## 非空约束[#](https://www.cnblogs.com/Finley/p/5284865.html#非空约束)

MySQL中所有数据类型都可以使用null, 通常使用`field IS NULL`来判断某个字段是否为空.

非空约束限制某一列不允许出现null值, 在建表时添加非空约束:

```sql
CREATE TABLE `user` {
  `username` VARCHAR(20) NOT NULL
}
```

也可以随时添加或删除非空约束:

```sql
ALTER TABLE `user` modify `username` VARCHAR(20) NULL;
ALTER TABLE `user` modify `username` VARCHAR(20) NOT NULL;
```

## 唯一约束[#](https://www.cnblogs.com/Finley/p/5284865.html#唯一约束)

唯一约束要求单列或列的集合不允许出现多个相同的非NULL值, 不允许`username`列出现重复值:

```sql
CREATE TABLE `user` {
  `username` VARCHAR(20) UNIQUE
}
```

允许`username`或`email`重复, 但不允许任意两条记录在`username`相同时`email`也相同.

```sql
CREATE TABLE `user` {
  `username` VARCHAR(20),
  `email` VARCHAR(20),
  UNIQUE(`username`, `email`)
}
```

为了便于后续操作, 最好使用命名约束:

```sql
CREATE TABLE `user` {
  username VARCHAR(20),
  email VARCHAR(20),
  CONSTRAINT uc_user UNIQUE(username, email)
}
```

使用`ALTER`添加唯一约束:

```sql
ALTER TABLE `user` ADD UNIQUE(`username`);
ALTER TABLE `user` ADD CONSTRAINT uc_user UNIQUE(`username`, `email`);
```

撤销唯一约束:

```sql
ALTER TABLE `user` DROP CONSTRAINT uc_user;
```

## 主键约束[#](https://www.cnblogs.com/Finley/p/5284865.html#主键约束)

主键是表上的一列或几列的集合, 主键列必须非空且唯一. 主键必须可以唯一确定表中的记录, 即不存在主键列相同的两条记录.

每个表最多包含一个主键约束:

```sql
CREATE TABLE `user` (
	`id` INT PRIMARY KEY AUTO_INCREMENT,
	`username` VARCHAR(20) DEFAULT "",
);
CREATE TABLE `user` {
  username VARCHAR(20),
  email VARCHAR(20),
  CONSTRAINT user_key PRIMARY KEY(username, email)
}
```

使用`ALTER`修改主键约束:

```sql
ALTER TABLE `user` ADD PRIMARY KEY(`id`);
ALTER TABLE `user` ADD CONSTRAINT user_key PRIMARY KEY(`username`, `email`);
```

因为最多有一个主键约束, 所以删除时不用指定约束对象:

```sql
ALTER TABLE `user` DROP PRIMARY KEY;
```

## 外键约束[#](https://www.cnblogs.com/Finley/p/5284865.html#外键约束)

外键约束将数据表B上的某列或列集合与数据表A的主键关联, 外键列的结构必须与被参照主键的结构一致, 但允许外键列的值重复. 数据表B上每条记录的外键列必须与被参照表A的某条记录的主键相同.

```sql
CREATE TABLE `user` (
	`id` INT PRIMARY KEY AUTO_INCREMENT,
	`username` VARCHAR(20) DEFAULT "",
	`password` CHAR(20) DEFAULT "",
	`gender` CHAR(1),
	`age` INT
);
CREATE TABLE `post` (
	id INT PRIMARY KEY AUTO_INCREMENT,
	uid INT,
	title TEXT,
	CONTENT TEXT,
	FOREIGN KEY(uid) REFERENCES user(`id`) 
);
```

建立在`post`表上的外键可以保证`post`表中每条记录的uid都指向`user`表中一条记录. 避免存在找不到作者的文章.

当删除被参照表上的某条记录时, 必须删除所有参照它的记录. 即删除用户前必须前必须删除他发表的所有文章, 以保证外键约束不被破坏.

在拥有外键约束的表`post`上添加数据时必须扫描一遍被参照的`user`, 这可能消耗很多时间, 在使用外键时必须要考虑到这一点.