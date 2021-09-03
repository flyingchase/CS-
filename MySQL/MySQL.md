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

