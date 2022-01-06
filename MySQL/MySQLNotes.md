
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





## 聚合函数





select sum(score) as sum

Count()

Student.name

From student_score

left join student on student.id=student_score.student_id 

order by sum desc limit 0,3









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

`InnoDB`数据按照`数据页`为单位进行读写，默认大小`16KB`

数据页内容包括：

<img src="https://cdn.jsdelivr.net/gh/flyingchase/Private-Img@master/uPic/vGVMIg.png" alt="vGVMIg" style="zoom:45%;" />

其中；

1. `File Header`有两个指针，指向前一个和后一个数据页，形成双向链表；

2. `User Records`：

   数据页中的记录按照主键的顺序形成`单向链表`，`页目录`记录索引作用；

   页目录由多个槽位组成（偏移量），先二分定位到记录分组，再遍历槽内的记录



#### B+树查询

InnoDB 中的B+树中每个节点都是一个数据页

<img src="/Users/qlzhou/Library/Application Support/typora-user-images/image-20220103120143800.png" alt="image-20220103120143800" style="zoom:55%;" />



- 只有叶子结点存放数据，中间结点存放目录项作为索引
- 非叶子结点通过分层达到“矮胖”，降低每层的搜索量；
- 结点按照索引键大小排序，构建双向链表，便于范围查询

查询流程：

在定位记录所在哪一个页时，也是通过二分法快速定位到包含该记录的页。定位到该页后，又会在该页内进行二分法快速定位记录所在的分组（槽号），最后在分组内进行遍历查找



### 聚集索引和非聚集索引（二级索引）

区别在于叶子结点存放的数据：

- 聚集索引的叶子结点存放的是实际数据，非聚集索引叶子结点存放的是主键值，不是实际数据

**InnoDB 创建聚集索引：**

- 默认使用主键作为聚集索引的索引键
- 没有主键，选择第一个不包含`NULL`值的唯一列作为索引键
- 上述不满足，自动生成`隐式自增 id`作为聚簇索引的索引键



**二级索引：**

实现非主键字段的快速搜索，叶子结点为主键值，同样使用 B+树

<img src="/Users/qlzhou/Library/Application Support/typora-user-images/image-20220103120904155.png" alt="image-20220103120904155" style="zoom:45%;" />

**非聚簇索引查询：**

​	查询语句使用二级索引，且查询数据非主键值：

​		在二级索引查找到主键值之后，再去聚簇索引获得数据行，`回表`——>使用两个 B+树查询

​	查询数据为主键值，在二级索引叶子结点就能查询到，`索引覆盖`——>只需一个 B+树





### summary

InnoDB 的数据是按「数据页」为单位来读写的，默认数据页大小为 16 KB。每个数据页之间通过双向链表的形式组织起来，物理上不连续，但是逻辑上连续。

数据页内包含用户记录，每个记录之间用单项链表的方式组织起来，为了加快在数据页内高效查询记录，设计了一个页目录，页目录存储各个槽（分组），且主键值是有序的，于是可以通过二分查找法的方式进行检索从而提高效率。

为了高效查询记录所在的数据页，InnoDB 采用 b+ 树作为索引，每个节点都是一个数据页。

如果叶子节点存储的是实际数据的就是聚簇索引，一个表只能有一个聚簇索引；如果叶子节点存储的不是实际数据，而是主键值则就是二级索引，一个表中可以有多个二级索引。

在使用二级索引进行查找数据时，如果查询的数据能在二级索引找到，那么就是「索引覆盖」操作，如果查询的数据不在二级索引里，就需要先在二级索引找到主键值，需要去聚簇索引中获得数据行，这个过程就叫作「回表」。



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

























































































































