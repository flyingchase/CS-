#  **TO DO LIST**

- top命令
- ps命令
- netstat命令
- awk命令
- find命令
- grep命令
- wc命令
- sed命令
- head和tail命令
- 正则表达式
- 如何查找出现频率最高的100个IP地址
- linux如何统计文件中某个字符串出现的频率
- linux启动的第一个进程
- linux查看端口占用
- linux查看CPU和内存使用
- Linux查看系统负载命令
- Linux调试程序
- Linux硬链接和软连接
- core dump
- cmake和makefile
- Shell脚本基本语法和使用



# Linux网络编程

- 孤儿进程、僵尸进程和守护进程
- 进程间通信方式signal、file、pipe、shm、sem、msg、socket
- 线程同步机制线程：互斥量、锁机制、条件变量、信号量、读写锁
- fork返回值
- 五大IO模型：阻塞I/O、非阻塞I/O、I/O复用、信号驱动I/O、异步I/O
- IO复用机制
- epoll与select/poll
- LT水平触发和ET边缘触发
- Reactor和Proactor模式
- 反向代理、负载均衡

 

# Linux操作指令-笔记

## 1 文件权限

### 1.1 ls命令

帮助 `man ls` `info ls`

`ls -al` ——>list all include those hided

e.g:

`-rwxr-xr-x   1 qlzhou  staff   5797544 12 21 15:59 impl`

第一列10个字符 文件类型和权限

​	1 \- 代表文件 d代表文件夹

​	234拥有者权限 rwx 可读可写可执行

​	567 同用户组权限 r-x 可读可执行

​	8910 其他用户权限 r-x 可读可执行

 	其中r w x顺序不变 不存在该权限则为- 

第二列 1代表多少文件名连接到此节点i-node

第三列 qlzhou代表所有者账号

第四列 表示文件所属用户组 

第五列 文件大小 默认单位B

第六列 修改日期

​	完整时间显示 `ls -l --full-time`

第七列 文件名 隐藏文件前含.



### 1.2 文件权限修改

使用cp复制时会保留原有的用户权限 从而他人无法修改

- chgrp修改所属用户组

  change group简写

  -R递归 recursive 持续修改 包括子目录下文件 目录





- Chown修改文件所有者

  `chown username filename/dirname` 可加-R





- Chmod























## **man - 获得帮助**

```text
man ls        # 许多Linux自带命令可以通过man查看使用帮助
ls --help     # 有些程序可以通过-h, --help查看使用帮助
```

## **ls - 显示目录内容**

```bash
ls                      # 显示目录内容
ls -l                   # 以列表显示形式显示目录内容，通常在~/.bashrc文件中增加一行：alias ll='ls -l'
                        # 以后就可以直接使用别名ll了，更方便
ll -h                   # 以人类可读的方式显示文件大小
ll -t                   # 以文件的修改时间排序，最新修改的在最前面
ll -tr                  # 以文件的修改时间排序，最新修改的在最后面
watch -n 3 -dc ls -l    # 追踪目录内容的变化，每3秒刷新一次
```

## **pwd - 显示当前目录**

```bash
pwd                 # 显示当前目录的绝对路径
ls `pwd`/file       # 显示文件的绝对路径
```

## **cd - 切换目录**

```bash
cd dir    # 切换到目录dir
cd        # 切换到用户的HOME目录
cd ~      # 同cd，~表示HOME目录
cd ..     # 切换到上一级目录；一个点.表示当前目录，两个点..表示上一级目录
cd -      # 切换到进入当前目录之前所在的目录
```

## **mkdir - 创建目录**

```bash
mkdir dir           # 创建dir目录
mkdir -p dir1/dir2  # 递归创建目录，如dir1不存在，会先创建dir1
```

## **cat - 合并文件（按行）**

```bash
cat file              # 合并一个或多个文件至标准输出，当只有一个文件时，相当于显示所有文件内容
cat file1 file2       # 合并file1和file2的内容，并在屏幕上输出
cat R1.fq.gz R2.fq.gz # 可以合并gzip压缩文件，如测序数据原始reads的合并
```

## **paste - 合并文件（按列）**

```bash
paste -d ' ' file1 file2    # 按列对列的方式一行一行合并文件。默认列中间加TAB键， -d参数可以改变列之间的分隔符
```

## **split - 分割文件**

```bash
split -d -l 10000 file chunk_   # 按行数分割文件，每个文件最多10000行，分割成的文件名为chunk_01, chunk_02。。。
split -d -b 100m file chunk_    # 按大小分割文件，每个文件最多100m，分割成的文件名为chunk_01, chunk_02。。。
```

## **cut - 剪切文件**

```bash
cut -f 1 file                   # 剪切文件的第1列
cut -f 1,2                      # 剪切文件的第1，2列
cut -f 3-                       # 剪切第3列及之后的所有列
cut -d ' ' -f 1 file            # 剪切第1列，但以空格作为列与列之间的分隔符。默认以TAB作为分隔符
grep '^>' test.fa | cut -c 2-   # 得到fasta文件中的序列名称（去掉了>符号）
```

## **less, head, tail - 显示文件内容**

```bash
less file       # 分屏显示文件内容，按空格键显示下一页，按下/后可以搜索内容
less -SN file   # 显示文件的行号，并且截断太长的行 

head file       # 默认显示文件前10行
head -n 20 file # 显示文件前20行

tail file       # 默认显示文件后10行
tail -n 20 file # 显示文件后20行
tail -n +2 file # 跳过第1行，显示从第2行开始的所有行，可用于跳过文件的标题行
tail -f file    # 当文件的内容还在增加时，实时显示末尾增加的内容，常用于查看日志文件的更新情况
```

## **wc - 统计文件内容**

```bash
wc -l file      # 统计文件行数
```

## **touch - 创建文件**

```bash
touch file                  # 创建一个空文件
touch {file1,file2,file3}   # 同时创建3个文件
```

## **cp, mv, rm- 文件/目录的复制，移动，删除**

```bash
scp file1 file2     # 将file1复制一份，命名为file2，复制目录要加-r参数：scp -r
mv file1 dir1/      # 将file1移动到dir1/目录下
mv file1 file2      # 重命名：即将file1移动成为file2
rm file             # 删除文件，删除目录要加-r参数：rm -r
rm -f file          # 文件若不存在，删除时会报错，加-f参数就不会报错
```

## **tar - 文件打包/压缩**

```bash
# 平时tar基本上就能完成打包、压缩、解压的任务了
tar czvf file.tar.gz files  # 打包并压缩
tar xvf file.tar.gz         # 解包，解压缩

gzip file                   # 压缩
gunzip file.gz              # 解压
```

## **chmod - 改变文件/目录权限**

```bash
chmod +x file   # 增加[本人]可执行权限
chmod -x file   # 取消[本人]可执行权限
chmod a+x file  # 增加[所有人]可执行权限
chmod a-x file  # 取消[所有人]可执行权限
```

## **chown - 改变文件/目录归属**

```bash
chown jianzuoyi:jianzuoyi file      # 将文件的所有权给jianzuoyi
chown -R jianzuoyi:jianzuoyi dirname    # 将目录以及目录内的文件的所有权给jianzuoyi
```

## **sort, uniq - 排序，去重**

```bash
sort file				# 默认按字典序对文件进行排序
sort -k2,2 -k3,3 file	# 先按第2列排序，第2列相同，再按第3列排序
sort -k2,2n file		# 按第2列排序，且第2列是数字，升序
sort -k2,2nr file		# 按第2列排序，且第2列是数字，降序
sort -u file			# 先排序文件，然后去除相邻的重复行，只保留一条记录
sort file | uniq		# 去除相信的重复行，只保留一条记录，相当于： sort -u file

# 利用sort, uniq取两个文件的交、并、补集
sort a b | uniq			# 并集
sort a b | uniq -d > c	# 交集
sort a c | uniq -u 		# 补集
```

## **wget - 下载文件**

```bash
wget https://repo.anaconda.com/archive/Anaconda3-2020.07-Linux-x86_64.sh	# 下载文件到当前目录，文件名保持不变
```

## **ssh - 远程登录**

```bash
ssh username@host					 # ssh 远程连接至服务器
```

## **scp - 远程文件传输**

```bash
scp username@host:/path/to/file .	 # 将远程服务器上的文件传输到当前目录，文件名保持不变，复制目录加参数-r
scp file username@host:/path/to/dir/ # 将本地文件复制到远程服务器，文件名保持不变，复制目录加参数-r
```

## **rsync - 远程文件拷贝**

rsync与scp不同，它只是做增量更新且支持断点续传，也就是要复制的文件存在于目标文件夹且内容与当前要复制的相同，则不会复制。

```bash
rsync -azvP dir1 dir2					# 将dir1的内容同步至dir2
rsync -azvP --delete dir1 dir2			# 同步dir2与dir1，dir1中删除的文件，dir2中也要跟着删除
rsync -azvP --exclude 'file' dir1 dir2  # 同步dir2与dir2，且将file排除在外
```

## **df, du, free - 查看磁盘/内存使用情况**

```bash
df -h		# 查看磁盘使用情况，-h表示以人类可读的方式显示容量大小
du -sh		# 查看当前目录使用了多少磁盘空间
du -sh *	# 查看当前目录下各文件或文件夹使用的磁盘空间
free -h		# 查看内存使用情况
```

## **top, htop, ps, kill - 任务管理**

```bash
top -c		# 查看CPU，内存的使用情况
htop		# top的完美替代品，Linux系统不自带，需要安装， ubuntu系统：apt install htop
ps aut		# 查看后台任务运行情况，第2列是任务的PID号
kill -9 PID # 删除编号为PID的任务
```

## **nohup，disown - 远程任务管理**

```bash
nohup ./run.sh &> run.sh.o &	# 远程SSH登录服务器，在后台运行任务，断开远程连接后任务仍然在后台跑
```

- 如果运行任务时没有加nohup命令，但任务运行时间长，但又必须断开（比如快下班了），若不想让任务因为断开远程连接而中断，可以用disown命令补救

```bash
./run.sh	# 假如任务是直接这样开始跑的
ctrl + z	# 按ctrl + z，将任务放到后台
jobs		# 输入jobs命令，回车，可以看到任务是暂停的： [1]+  Stopped(SIGTSTP)        bash run.sh
bg			# 让后台暂停的任务开始运行
jobs		# 再次运行jobs，可以看到任务已经跑起来了：   [1]+  Running                 bash run.sh &
disown -r 	# 从当前shell中移除运行中的作业，至此，可以关掉终端回家了
```

## **| - 管道**

管道，将前一个命令的输出作为后一个命令的输入

```bash
command1 | command2
```

## **>, >> - 输入输出重定向**

Linux中常用重定向操作符有：

1. 标准输入（/dev/stdin）：代码为0， 使用<或<<
2. 标准输出（/dev/stdout）：代码为1，使用>（覆盖）或>>（追加）
3. 标准错误输出（/dev/stderr）：代码为2，使用2>或2>>
4. &> 标准输出和错误输出同时重定向
5. /dev/null 代表垃圾箱，不想要保存的东西都可以重定向到这里

- 输出重定向就是将命令的结果重定向到文件，而不是输出到屏幕，通常用于保存命令的结果

```bash
./run.sh > run.sh.o		# 标准输出到run.sh.o日志文件
./run.sh 2> run.sh.e	# 标准错误输出到run.sh.e错误日志文件
./run.sh &> run.sh.log	# 标准输出和标准错误都输出到定一个文件
./run.sh &> /dev/null	# 丢弃标准输出和标准错误信息
```

- 输入重定向是将文件作为输入的来源，而不是键盘

```bash
command < file			# 将file的内容作为command的输入 
command << END			# 从标准输入（键盘）中读取数据，直到遇到分界符END时停止（分界符用户可以自定义）
command <file1 > file2	# 将file1作为command的输入，并将处理结果输出到file2
```

- 综合运用

```bash
#!/bin/bash

while read line
do
    do something
done < file.txt > result.txt
```

逐行读入file.txt的内容，处理之后，将结果保存到result.txt文件中。

## **find, locate, which - 文件查找**

```bash
find -name file					# 在当前目录查找名为file的文件
find dir/ -name file			# 在dir/目录下查找名为file的文件
find dir/ -name '*file*'		# 在dir/目录下查找包含file关键词的文件，-name参数支持正则表达式
find dir/ -name file -delete	# 查找文件并删除

locate file						# 查找文件
which command					# 显示命令的绝对路径
```

## **xargs - 命令组合工具**

```bash
cat file | xargs		# 将file的内容显示成一行
cat file | xargs -n3	# 将file的内容每3列一行进行输出
find /ifs/result -name '*.fq.gz' | xargs -n1 -I{} cp {} /ifs/data/	# 查找fq.gz文件并复制到/ifs/data目录下
find /ifs/result -name '*.fq.gz' | xargs tar czvf all.fq.gz			# 查找fq.gz文件并打包在一起
find . -type f -name '*.log' -print0 | xargs -0 rm -f				# 当rm文件过多时，可以这样删除
find . -type f -name '*.py' -print0 | xargs -0 wc -l				# 统计一个目录中所有python文件的行数
```

## **parallel - 并行工具**

parallel是增强版的xargs。假如一个脚本文件中有4条命令：

```bash
# cat run.sh
echo a
echo b
echo c
echo d

# 同时执行4个任务，生信中常通过这种方式并行执行多个任务
cat run.sh | parallel -j 4	
find *.fq | parallel -j 12 "fastqc {} --outdir ."	# 同时执行12个Fastqc任务
find *.bam | parallel --dry-run 'samtools index {}' # 同时执行samtools index任务，--dry-run显示任务命令但不实际执行，用于命令检查
```

## **useradd - 添加用户**

```bash
useradd -m username	# 创建用户并为其在/home下创建一个以其名称命名的目录
```

## **passwd - 更改密码**

```bash
passwd	    	   # 更改当前用户的密码
passwd username	   # 更改指定用户的密码
```

## **dos2unix - 文件格式转换**

Linux很多工具都是针对纯文本文件的，并且需要是Unix-like格式的文本文件。但是很多时候文件是从Windows或Mac系统上传到Linux服务器上的，这可能导致文件格式不兼容，原因是不同平台生成的文本文件的换行符不一样。

| 操作系统 | 符号 | 正则表达式 |
| -------- | ---- | ---------- |
| Mac      | ^M   | \r         |
| Linux    | $    | \n         |
| Windows  | ^M$  | \r\n       |

```bash
cat -A file			# 查看文件换行符情况
dos2unix file			# Windows格式转换成Unix-like格式
```

## **grep**

**用于查找文件里符合条件的字符串。**

```bash
grep [-abcEFGhHilLnqrsvVwxy][-A<显示列数>][-B<显示列数>][-C<显示列数>][-d<进行动作>][-e<范本样式>][-f<范本文件>][--help][范本样式][文件或目录...]
grep pattern files			 # 搜索文件中包含pattern的行
grep -v pattern files		 # 搜索文件中不包含pattern的行

grep -f pattern.txt files	 # 搜索的pattern来自于文件中
grep -i pattern files		 # 不区分大小写。默认搜索是区分大小写的
grep -i pattern files		 # 只匹配整个单词，而不是字符串的一部分（如搜索hello，不会匹配到helloworld）
grep -n pattern files		 # 显示行号信息
grep -c pattern files		 # 显示匹配的行数
grep -l pattern files		 # 只显示匹配的文件名
grep -L pattern files		 # 显示不匹配的文件名
grep -C number pattern files     # 额外显示匹配行的上下[number]行
grep pattern1 | grep pattern2 files     # 显示既匹配pattern1，又匹配pattern2的行
grep -E "pattern1|pattern2" files	# 显示匹配pattern1或者pattern2的行, grep -E相当于egrep

# 用于搜索的特殊字符
^: 表示行前
$: 表示行尾

grep '^#' result.vcf		# 显示VCF文件的表头信息
grep '^hello$' files		# 显示只包含hello的行
grep -v '^\s*$' file		# 删除空白行
```

## **sed**

sed是stream editor的缩写，中文称之为“流编辑器”。

```bash
sed command file
```

- command部分，针对每行要进行的处理
- file，要处理的文件

### **Actions**

- d：删除该行
- p：打印该行
- i：在行的前面插入新行
- a：在行的后面插入新行
- r：读取指定文件的内容。
- w：写入指定文件。

```bash
sed -n '10p' file		# 显示第10行
sed -n '10,20p' file	# 显示第10到20之间的行
sed -n '/pattern/p' file# 显示含有pattern的行
sed -n '/pattern1/,/pattern2/p' file # 显示patter1与pattern2之间的行

sed '10d' file			# 删除第10行
sed '10,20d' file		# 删除第10到20之间的行
sed '/pattern/d'		# 删除匹配pattern的行
sed '/^\s*$/d' file		# 删除空白行
sed 's/^\s*//' file		# 删除行前的空白：空格，制表符
sed 's/\s*$//' file		# 删除行尾的空白：空格，制表符
sed 's/^\s*//;s/\s*$//' file # 删除行首和行尾的空白：空格，制表符

sed 's/AA/BB/' file		# 将文件中的AA替换成BB，只替换一行中第一次出现的AA，替换后的结果输出到屏幕
sed 's/AA/BB/g' file	# 将文件中的所有AA都替换成BB，替换后的结果输出到屏幕
sed -i 's/AA/BB/g' file # 将文件中的所有AA都替换成BB，直接更改文件的内容
sed '/CC/s/AA/BB/g' file# 只替换那些含有CC的行
sed 's/pattern/&XXXX/' file	# 在pattern之后加上XXXX。&表示之前被匹配的内容
sed 's/pattern.*/&XXXX' file# 在匹配pattern的行尾加上XXXX。pattern.*表示包含pattern的整行内容

sed -n '1~4s/^@/>/p;2~4p' file.fq > file.fa	# Fastq文件转Fasta文件
sed -n '2~4p' file.fq		# 提取Fastq文件的序列

sed 'y/ABC/XYZ/' file	# 将ABC逐字替换成XYZ

sed '1i\hello' file		# 在第1行前面插入一行，内容为hello，通常用来为文件增加标题
sed '1a\hello' file		# 在第1行后面插入一行，内容为hello
sed '1r file2' file1	# 在第1行后面读入file2的内容
sed '/pattern/w file2' file1 # 将匹配的行写入file2中
```

## **awk**

Awk是一个强大的文本分析工具，它每次读入一条记录，并把每条记录切分成字段后进行分析。Awk官方文档是非常好的学习材料，通过`man awk`查看。

```bash
awk 'BEGIN { action } pattern { action } END { action }'
```

**Awk程序通常是一系列 pattern {action}对：**

`pattern`，表示模式匹配，只处理匹配的行。pattern可以省略，表示匹配所有行

`action`，表示对匹配行所做的动作。{actions}可以省略，表示{ print }。`BEGIN`和`END`的{action}不能省略



**pattern可能是：**

`BEGIN`， 执行初始化操作，程序开始时执行一次

`END`，执行收尾工作，程序结束时执行一次

`expression`，一个表达式，既可以是判断语句，也可以是正则表达式

### **常用参数**

- `-F value` 设置域分隔符，相当于给FS内置变量赋值
- `-v var=value` 将变量value的值赋给程序变量var，-v可以多次使用

### **记录与字段**

记录是一次读入的内容，通常是文件的一行，保存在字段变量$0中，记录可以被分割成字段，保存在变量$1，$2，...，$NF中。

### **表达式与操作符**

Awk表达式的符号与C语言的类似，基本的表达式有数字，字符串，变量，字段，数组以及函数调用。变量无需声明，它们在首次使用时被初始化为`null`。

```bash
assignment          =  +=  -=  *=  /=  %=  ^=
conditional         ?  :
logical and         &&
logical or          ||
logical not         !
array membership    in
matching       		~   !~
relational          <  >   <=  >=  ==  !=
concatenation       (no explicit operator)
add ops             +  -
mul ops             *  /  %
unary               +  -
exponentiation      ^
inc and dec         ++ -- (both post and pre)
field               $
```

### **正则表达式**

在Awk中语言中，通常测试一个记录、字段或字符串是否与一个正则表达式匹配，匹配返回1，不匹配返回0。正则表达式用两个反斜杠`/`包围。

```bash
expr ~ /r/							 # 评估expr是否与r匹配。匹配的意思是expr的一个子串是否在正则表达式r定义的字符串集中。

/r/ { action }, $0 ~ /r/ { action }	 # 两者相同， /r/ 等于 $0 ~ /r/
```

任何表达式都可以放到`~`和`!~`右边或者内建的需要正则表达式的地方。在必要的时候，该表达式会被转变成字符串，然后作为一个正则表达式来解释。以下三行awk命令完成同样的功能：输出第5列为10的的行。

```bash
seq 20 | xargs -n5 > file
# cat file
1 2 3 4 5
6 7 8 9 10
11 12 13 14 15
16 17 18 19 20

awk '$5 ~ /10/' file
awk '$5 ~ "10"' file
awk '$5 ~ 10' file
```

### **数组**

Awk支持一维数组。其表示方法为`array[expr]`，`expr`在内部被统一转换成字符串类型，因此A[1]，与A["1"]相同，事实上索引都是“1”。索引为字符串的数组被称为关联数组。`expr in array`用于判断数组元素array[expr]是否存在。

```bash
for ( var in array ) statement
```

### **控制语句**

```bash
if ( expr ) statement
if ( expr ) statement else statement
while ( expr ) statement
do statement while ( expr )
for ( opt_expr ; opt_expr ; opt_expr ) statement
for ( var in array ) statement
continue
break
```

### **内置变量**

- `NR` - 当前行数
- `NF` - 当前行的列数
- `RS`，行分隔符，默认是换行符
- `FS`，列分隔符，默认是空格和制表符
- `ORS`，输出行分隔符，默认为换行符
- `OFS`，输出列分隔符，默认为空格
- `FILENAME`，当前文件名

### **内置函数**

### **字符串函数**

sub()、substr()、gsub()，sprintf()，index()，length()， match()，split()，tolower(), toupper()

### **数学函数**

sin()，cos(), ...

### **输入输出**

有两个输出语句，`print`和`printf`

```python
print							# 打印整条记录到标准输出，相当于print $0
print expr1, expr2, ..., exprn	# 打印指定字段到标准输出
printf format, expr-list		# C语言printf函数的重用
```

输入函数getline有以下几种形式：

```bash
getline							# 读取下一条记录到$0，更新NF，NR和FNR
getline var						# 读取下一条记录到var，更新NR和FNR
getline < file					# 从文件读取记录到$0，更新NF
getline var < file				# 从文件读取记录到var
command | getline				# 通过管道传递command的结果到$0，更新NF
command | getline var			# 通过管道传递command的结果到var
seq 10 | awk '{print $0;getline}'					  # 显示奇数行
seq 10 | awk '{getline; print $0}'					  # 显示偶数行
seq 10 | awk '{getline tmp; print tmp; print $0}'	  # 奇偶行对调

awk 'BEGIN {"date" | getline;close("date");print $0}' # 得到系统当前时间

# fastq转换成fasta
awk '{getline seq; getline comment; getline quality; sub("@", ">", $0); print $0"\n"seq}' file
```

### **示例**

```bash
awk '{print $0}' file	# 打印整行
awk '{print $1}' file	# 打印第一列
awk '{print $2}' file	# 打印第二列
awk '{print $NF}' file	# 打印最后一列
awk '{print $(NF-1)}' file#打印倒数第二列
awk -F ';' -v OFS='\t' '{print $1,$2,$NF}' file	# 读入的文件以逗号;分隔列，打印第1列，第2列和最后一列，并且打印时以制表符作为列的分隔符
number=10;awk -v n=$number '{print n}' file	# number的值被传给了程序变量n
awk '$2 > 100' file		# 打印第2列大于100的行
awk 'NR>1 && NR<4' file # 打印第2~3行

awk '/EGFR/' file		# 打印含有EGFR的行，相当于grep EGFR file
awk '$1 ~ /EGFR/' file	# 打印第1列含有EGFR的列

# 按指定列去除重复行
# cat file
1 2 3 4 5
6 2 8 9 10
11 12 13 14 15
16 17 18 19 20
awk '!a[$2]++' file		# 第二列出现两次2，只保留第一次出现的那一行，结果如下：
1 2 3 4 5
11 12 13 14 15
16 17 18 19 20

awk '{sum+=$1} END {print sum}' file	# 累加文件的第一列
awk '{sum+=$1} END {print sum/NR}' file	# 求第一列的平均数

# 从含有多条fasta序列的文件中提取指定序列
 awk -v RS=">" '/chr1/ {print $0}' hg19.fa	# 提取chr1的序列
 awk -v RS=">" '/chr1|chr2/ {print $0}' hg19.fa	# 提取chr1和chr2的序列
```

## **Bash脚本模板**

```bash
#!/bin/bash

command1

command2

...
```

`chmod +x run.sh` 给run.sh脚本增加可执行权限

执行脚本，以下三种方式都可以：

```bash
# 脚本在前台执行，标准输出和标准错误输出到屏幕
./run.sh
bash run.sh
sh run.sh		# 前提sh链接到了bash，如果没有，需要root权限执行命令：ln -sf /bin/bash /bin/sh

# 脚本在前台执行，标准输出和标准错误保存到文件
./run.sh &> run.sh.o

# 脚本在后台执行，在最后加上一个&符号
./run.sh &> run.sh.o &

# 脚本在后台执行，并且防断线（长时间运行任务时使用）
nohup ./run.sh &> run.sh.o &
```

## **其他命令**

```bash
echo $PATH		# 显示环境变量
time command	# 显示命令执行时间
date			# 显示日期和时间
history			# 显示历史命令
export PATH=$PATH:/path/to/bin	# 将路径加入环境变量中
ln -s file file2# 为file文件创建软链接，名称为file2
exit			# 退出登录
Tab键自动补全	 # Tab键可以补全命令或文件路径，输入部分命令或路径时，尝试按Tab键补全
Ctrl + c		# 中止当前命令的执行
seq 10			# 产生1到10的整数
md5sum			# 生成，或验证文件的MD5值
```









------

## 1. 帮助命令

1.1 `man 命令名称`，man 的意思是 manual，即手册，对于大多数命令都可以使用这个命令来查看其使用的方法。

1.2 `help 命令名称` 或者 `命令名称 --help`，可以查看命令的使用帮助。

## 2. 目录操作

2.1 `pwd`，查看当前所在的目录路径。

2.2 `ls 路径`，查看指定路径下的文件列表，可以加上多个路径，例如：`ls /usr /etc`，分别显示根目录下的 user 目录和 etc 目录。

ls 命令的常用参数有：

- `-a`，显示隐藏的目录或文件
- `-l`，显示文件的详细信息，`ls -l` 等价于 `ll`
- `-h`，格式化显示文件的大小，如 1K，3M，1G，方便阅读
- `-t`，以修改时间排序文件列表
- `-r`，反序排序列表
- `-R`，递归显示所有子文件夹的内容
- `-S`，大写的 S，以文件的大小排序

2.3 `cd 路径`，cd 的意思是 change directory，改变路径。

2.4 `mkdir 目录名称`，创建目录，可以加多个参数创建多个目录，例如 `mkdir /tmp/a /tmp/b`。

常用参数：`-p`，表示创建多级不存在的目录。

2.5 `rmdir 目录`，删除目录，注意只能删除空目录。

2.6 `cp`，复制文件或者目录的命令，例如将一个文件复制到另一个目录中：`cp test.txt /usr/local/`

常用参数：

- `-r`，用于复制目录，递归目录中的所有内容
- `-v`，显示复制的详细信息

2.7 `mv`，这个命令有两个功能，一是重命名文件或文件夹，例如 `mv a.txt b.txt`，二是移动文件，例如将文件 a.txt 移动至 root 目录下 `mv a.txt /root/`。

## 3. 查看文件

3.1 `cat 文件名`，查看文件的全部内容，常用参数 `-n`，可显示文件的行号。

3.2 `head 文件名`，查看文件头部的内容，默认显示前 10 行，可加参数 `-行数` 查看前 n 行，例如 `head -30 文件名`。

3.3 `tail 文件名`，查看文件尾部的内容，默认显示最后 10 行，常用参数 `-f`，可查看追加的文件内容。

3.4 `wc 文件名`，查看文件的统计信息，常用参数：

- `-l`，显示文件的行数
- `-w`，显示文件内的总单词数
- `-c`，文件的字节数大小

## 4. 压缩解压

4.1 打包命令：`tar cf 打包后存放路径和文件名 源文件`，例如打包根目录下的 /etc，命令：`tar cf /temp/etc-backup.tar /etc`，使用这个参数打包，不用压缩源文件。

4.1 压缩打包：

- gz 格式：`tar czf /temp/etc-backup.tar.gz /etc`
- bz2 格式：`tar cjf /temp/etc-backup.tar.bz2 /etc`

两种格式的区别：gz 格式压缩速度更快，bz2 个格式压缩后的文件更小。

.tar.gz 可以缩写为 **.tgz**，.tar.bz2 可以缩写为 **.tbz2**。

4.2 解压命令：

- 解压 **.tar** 格式的文件：`tar xf 文件名`，可加参数 -C，指定解压后的存储路径
- 解压 **.tar.gz** 格式的文件：`tar zxf 文件名`
- 解压 **.tar.bz2** 格式的文件：`tar jxf 文件名`

## 5. 权限操作

5.1 添加用户：`useradd 用户名`。

5.2 修改用户密码：`passwd 用户名`，不加用户名的话，则默认修改当前用户的密码。

5.3 删除用户：`userdel 用户名`，可加参数 -f 强制删除，-r 删除用户主目录。

5.4 新建用户组：`groupadd 用户组名称`。

5.5 将某个用户加入到某个用户组中：`usermod -g 用户组名称 用户名`

5.6 修改文件权限，文件权限的表示为 r（可读），w（可写），x（可执行），数字表示分别是 4， 2，1。

`chmod u+x a.txt`，表示对文件 a.txt 添加可执行权限，选项 u 表示所属用户，类似的有 g（属组）、a（全部）；+ 表示添加权限，- 表示删除权限

5.7 修改文件的属主和属组权限：`chown root:root a.txt`，其中 : 前面的表示所属用户，后面的表示所属用户组。

也可以单独改变属主和属组的权限：

- `chown root a.txt`，改变所属用户。
- `chown :root a.txt`，改变所属用户组。

## 6. 软件安装

6.1 rpm 包管理，文件后缀一般是以 .rpm 结尾的。

- 安装的命令：`rpm -i xxx.rpm`，如果安装的包需要有其他的依赖，则会提示先手动安装依赖。
- 查看安装的软件包：`rpm -qa`，可加管道过滤想查询的软件包，`rpm -qa | grep mysql` 。
- 卸载软件包命令：`rpm -e 软件包名称`，注意这里必须输入软件包的名字全程。常加参数 `--nodeps` 表示卸载时不检查依赖。

6.2 yum 是一个基于 rpm 的包管理器，可以自己检查依赖，维护包的升级移除，更加的自动化，在实际中的使用也更多。

- 安装软件包：`yum install 软件包名称`
- 卸载：`yum remove 软件包名称`
- 查看安装的软件：`yum list`

6.3 Linux 上安装软件的第三种方式是源代码编译安装，在下载好的源代码中一般有一个 configure 文件，安装的步骤为：

- 进入软件包目录，执行 configure 文件，`./configure --prefix=/软件包安装路径`
- 然后执行命令 `make`，或者按照指示执行 `gmake`，两者是类似的命令
- 然后再执行 `make install`命令

## 7. Vim命令

Vim 是在 Linux 上经常使用的一个文本编辑器，熟练使用 Vim 的常用命令也是有必要的。

7.1 输入 `vim 文件名` 则进入了 vim 的正常模式，常用命令如下：

- 插入：
  - i（光标处前一个字符编辑），I（大写的 i，光标所处行首编辑）
  - a（光标处后一个字符编辑），A（光标所处行尾编辑）
  - o（光标处向下新建一行编辑），O（大写的 o，光标处向上新建一行编辑）
- 复制：
  - 复制单行：光标处所处行按 YY，然后在目标行按 P
  - 复制多行：按数字 + YY，例如复制 3 行，在光标处按 3 YY，然后在目标行按 P
- 剪切：
  - 剪切单行：光标所在行按 DD
  - 剪切多行：数字 + DD
- 撤销：按 U
- 删除光标所处的字符：X
- 替换光标所处的字符：按 R 后输入新的字符
- 光标移动至某行：行数 + GG
- 显示文件行数：按 ：后输入命令 `set nu`
- 移动至文件的第一行：GG
- 移动至文件最后一行：shift + G
- 移动至一行的开头：shift + 6
- 移动至一行的末尾：shift + 4

7.2 vim 的命令模式：

- `:w /usr/local/a.txt`，w 表示保存，后面可以跟上保存的路径和文件名称

- `:wq` 表示保存并退出，`:q!` 表示强制退出，不保存修改

- `:! 其他Linux命令`，输入 ! 可以进入临时的命令模式，在编辑文件的时候执行其他的命令，按 Enter 键回到正在编辑的文件中

- ```
  / 字符
  ```

  ，可以搜索文件中的内容，搜索后，匹配的文本会高亮显示，按 N 移动至下一个匹配的结果，按 shift + N 移动至上一个。

  - 搜索后的文件会高亮显示，如果需要取消高亮显示，输入命令：`:set nohlsearch`，如果需要重新高亮显示，去掉命令中的 no 即可：`:set hlsearch`

- 替换搜索到的内容：`:s/原文本/新的内容`，默认只会替换光标所在行的第一处搜索匹配的内容

- 全局替换搜索到的内容：`:%s/原文本/新的内容`

- 如果匹配的内容有连续的，则全局替换时需要加参数 g：`:%s/原文本/新文件/g`

- 替换某行的内容：`m,ns/原文本/新的内容`， m、n 分别表示起始行和结束行，同理也可加参数 /g 替换有连续匹配的内容。







# 考题

------

## 00 平常用什么 linux 命令比较多？如何打开文件并进行查找某个单词？怎么在某个目录下找到包含 txt 的文件？
pwd：显示当前所在位置

sudo + 其他命令：以系统管理者的身份执行指令，也就是说，经由 sudo 所执行的指令就好像是 root 亲自执行。

grep：要搜索的字符串 要搜索的文件 --color ： 搜索命令，--color 代表高亮显示

ps - ef/ps aux： 这两个命令都是查看当前系统正在运行进程，两者的区别是展示格式不同。如果想要查看特定的进程可以使用这样的格式：

```shell
 ps aux|grep redis
```


（查看包括redis的进程），也可使用

```shell
pgrep redis -a
```


注意：如果直接用ps（（Process Status））命令，会显示所有进程的状态，通常结合 grep 命令查看某进程的状态。

kill -9 进程的 pid ： 杀死进程（-9 表示强制终止），先用 ps 查找进程，然后用 kill 杀掉。

find 目录 参数 ： 寻找目录（查）。在/home目录下查找以 .txt 结尾的文件名:

```shell
find /home -name "*.txt"
```


  ls 或者 ll :（ll 是 ls -l 的别名，ll 命令可以看到该目录下的所有目录和文件的详细信息）： 查看目录信息。

free : 显示系统内存的使用情况，包括物理内存、交换内存(swap)和内核缓冲区内存。

tar -zcvf 打包压缩后的文件名 要打包压缩的文件 : 打包并压缩文件，一般情况下打包和压缩是一起进行的，打包并压缩后的文件的后缀名一般 .tar.gz。c：压缩。

tar -xvf 压缩文件 - C 解压的位置 : 解压压缩包。x: 解压。

wget : 是从远程下载的工具。

vmstat : 虚拟内存性能监控、CPU 监控。

top : 常用来监控Linux的系统状况，比如CPU、内存的使用，显示系统上正在运行的进程。load average：系统负载，就是进程队列的长度。当这个值>cpu核心数的时候就说明有进程在等待处理了，是负载过重。

2.18 用过 ping 命令么？简单介绍一下。TTL 是什么意思？
ping : 查看与某台机器的连接情况。TTL：生存时间。数据报被路由器丢弃之前允许通过的网段数量。

2.19 怎么判断一个主机是不是开放某个端口？
telnet IP 地址 端口

```shell
telnet  127.0.0.1 3389 
```









## 快捷键

接下来是大家很熟悉的一个环节了，许多软件都内置了快捷键供用户使用，Linux/Shell 也不例外，熟悉这些快捷键可以极大提高我们的工作效率，我尤其推荐那些用来“移动光标”的命令。有时候我们输入的命令很长，会遇到一些“卡在中间、进退两难”的情况，这时候它们可能会发挥大的作用哦：

| 快捷键名 | 作用                          |
| -------- | ----------------------------- |
| CTRL+A   | 把光标移到命令行开头          |
| CTRL+E   | 把光标移到命令行结尾          |
| CTRL+C   | 强制终止当前命令              |
| CTRL+L   | 清屏，相当于clear命令         |
| CTRL+U   | 删除/剪切当前行、光标前的内容 |
| CTRL+K   | 删除/剪切当前行、光标后的内容 |
| CTRL+Y   | 粘贴CTRL+U/K中剪切的内容      |
| CTRL+R   | 在历史命令中搜索              |
| CTRL+D   | 退出当前终端                  |
| CTRL+Z   | 刮起当前进程                  |
| CTRL+S   | 暂停屏幕输出                  |
| CTRL+Q   | 恢复屏幕输出                  |