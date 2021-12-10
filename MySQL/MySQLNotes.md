
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

找到多个表的公共列 `ON`  `JOIN`

``` sql
SELECT *
FROM orders
JOIN customers 
	on customers.customer_id=orders.customer_id
```

































































































