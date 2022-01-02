
#  MySQL
## 00 入门

### 资料

- 小孩子 ``MySQL是怎样运行的:从根上理解MySQL` 

- `MOSH SQL视频`

- 极客时间: `MySQL实战45讲`

   

## 单一表检索数据

``` sql	
USE sql_store;
-- 列选择
SELECT *

FROM customers
-- WHERE customer_id =1
ORDER BY first_name
```



1. `--`则会忽略当前语句；
2. 语句写顺序：`select——>from——>where——>order by`
3. 最好每个子句放在一行



### select 详解

``` sql
SELECT 
		last_name,
		first_name,
		points,
		(points*10-1)%10 AS discount_factor
```

1. 多列选择使用换行，每个子句一行，对应缩进

2. `AS`视为别名`alais`
   - 使用双引号、单引号括起来则可包含空格
3. `select DISTINCT`表明去重，结果不含有重复的值





### Where 详解

条件选择语句，作为列的筛选条件

**比较操作符**

``` sql
<,>,<=,>=,<>,!=,=
```

**筛选**

`AND` `OR` `NOT`

1. `AND` 优先级最高
2. `NOT`表示取反

**多表达式`IN`：**

符合的多个条件内

``` sqlite
SELECT *
FROM customers
WHERE state NOT IN('va','fl','ga')
```

**检索`BETWEEN`**

``` sql	
SELECT *
FROM customers
WHERE points BETWEEN 1000 AND 3000
```

**检索 `LIKE`**

与使用正则表达式`REGEXP`相同效果

``` sql
SELECT *
FROM customers
WHERE last_name LIKE '%b%'
-- %代表任意字符
-- _下划线代表一个字符 '_y' means 以 y 结尾的两个字符
where last_name regexp 'b'
```



**`REGEXP`正则表达式**

`^`起始位置 `|`表示逻辑或 `$`表示末尾位置 `[]`代表多个匹配



### Order by 排序

默认升序，使用`DESC`作为降序

### 限制查询记录的返回条数

`LIMIT *,*` 第一个数字代表跳过的记录数，第二个数字代表现实的内容条数，第一个数字可忽略

**`LIMIT`写在最后**









## 多表检索

找到多个表的公共列 `ON`  `JOIN` 找到公共列

``` sql
SELECT order_id,o.customer_id,first_name,last_name
FROM orders o
JOIN customers c 
	ON o.customer_id=c.customer_id 
```



默认为内连接，多个表合并时，只返回符合条件的记录

左右链接：满足左右两侧的均列出

``` sql
SELECT 
	p.product_id,
	p.`name`,
	oi.quantity
FROM products p
LEFT JOIN  order_items oi 
	on p.product_id=oi.product_id
```



## 复合合并的条件

减少使用`where`作为隐式合并



### using

当不同的表内的字段列名相同时使用`using`而非`on`





















## SQL 语句的更新

redo log binlog

物理日志和逻辑日志





两阶段提交







## 索引

提高数据的查询效率，类似书的目录

### 索引常见模型

#### 哈希表：

​	区间查询速度慢，适合等值查询

#### 有序数组：

​	查询时候二分 logN 即可，但更新数据麻烦，需要挪动后面所有的记录

​	适合静态存储引擎

#### 树

**二叉树：**

​	二叉搜索树查询可以按照中序遍历，logN 复杂，但是保证为平衡二叉树更新操作也为 logN 复杂度

​	

**N 叉树：**

​	N 取决于数据块的大小



#### 其他

跳表、LSM 树（日志结构树）——>将每次操作进行日志的追加，到一定内存或者时间后再进行提交

### InnoDB 的索引模型

索引组织表：按照主键的顺序以索引的形式存放表；使用 B+树索引模型，数据存储在 B+树中



**每个索引在 InnoDB 对应一颗 B+树**





























## 锁

全局锁、表级锁、行锁



### 全局锁

全库逻辑备份时常用，将库内的每个表 select 出来存为文本

`Flush tables with read lock FTWRL` 



### 表级锁

表锁和元数据锁（`meta data lock MDL`）



### 行锁

由引擎层各自实现，`MyISAM`不支持行锁（并发控制只能使用表锁）

针对数据表中行记录的锁

`InnoDB`事务中，行锁在需要时候加上，等待事务完成才释放——>两阶段锁协议





### 死锁和死锁检测

发现死锁后，主动回滚到死锁链中的某个事务，让其他事务继续得以执行。

`innodb_deadlock_detect`设置为`on`

























































































































