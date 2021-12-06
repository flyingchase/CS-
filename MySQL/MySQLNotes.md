
#  MySQL
## 00 入门

### 资料

- 小孩子 ``MySQL是怎样运行的:从根上理解MySQL` 

- `MOSH SQL视频`

- 极客时间: `MySQL实战45讲`

   

## 检索数据

### 单一表检索数据

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
<,>,<=,>=,<>,!=
```

