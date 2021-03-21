**CS61B**

# Java基础

------

## 00 前言



**JDK&JRE&JVM：**

JDK= JRE + Java的开发工具（javac.exe java.exe java.doc.exe）

JRE= JVM +Java核心类库

**编译与运行**

.java文件（源文件）——>java.exe来编译——>.class文件（字节码文件）——>java.exe来解释运行——>结果

其中字节码中的文件名称为class的类名 多个类则有多个.class文件  与源文件名不一定重复; 注释掉的语句不参与编译，即*.class中不存在这些字节码

运行java

终端操作javac *.java  java *(无须后缀.class) 









**mac查看java安装位置：**

/usr/libexec/java_home -V

/Library/Java/JavaVirtualMachines/jdk-15.0.1.jdk/Contents/Home

```shell
➜  / /usr/libexec/java_home -V
Matching Java Virtual Machines (2):
    15.0.1, x86_64:	"Java SE 15.0.1"	/Library/Java/JavaVirtualMachines/jdk-15.0.1.jdk/Contents/Home
    11.0.9, x86_64:	"Amazon Corretto 11"	/Users/qlzhou/Library/Java/JavaVirtualMachines/corretto-11.0.9/Contents/Home

/Library/Java/JavaVirtualMachines/jdk-15.0.1.jdk/Contents/Home
➜  /
```

**文档注释（java特有）：**

/**  */（可以被JDK的工具javadoc解析，生成网页文件来说明）

cd进入*java后 使用javadoc -d 形成的网页文件夹名称 -author -version *.java 在新建的文件内有 *.html



**Java API文档：**

应用程序编程接口（Application Programming Interface）

**Notes:**

- 一个java源文件可以有多个class但只有最多一个为public class，要求public的类名必须与源文件名相同；
- 程序入口main()方法:"public static void main(String[] args)" arg==arguments 参数 
- 输出语句：Stystem.out.println("");// 先输出后换行  或者 Stystem.out.print();
- 执行语句以;结束



**java语言特性：**

- 面向对象（类和对象、特征：封装 继承 多态）
- 健壮性（去除指针 自动垃圾回收机制（仍存在内存泄漏 溢出））
- 跨平台：JVM 

**Java常见关键字：** 所有字母均小写

*数据类型*  class interface enum byte short int long float double char boolean void

*流程控制*  if else switch case default while do for break continue return

*访问权限修饰符*  private protected public

*类 函数 变量*	abstract final static synchronized

*类与类之间关系*	extends implement

*异常处理*	try catch finally throw throws

*包* 	package import

*数据类型值* ture false null

java没有无符号类型 unsigned

**java标志符**	

对各类变量 方法 类命名的名字 （自己命名的）

- 不可包含空格 区分大小写 内部可包含关键字 保留字 

- 数字不可开头 字母 0-9 , _ $

- 不可包含空格

  *命名规则：*包名均小写；类名首字母大写；常量均大写；变量第一个字母小写第二个单词首字母大写；



## 01 变量

​	包括类型、名称、值； declared

| 类型      | 大小   | 默认值     | 值范围                                      |
| :-------- | :----- | :--------- | :------------------------------------------ |
| `boolean` | 不适用 | false      | true 或 false                               |
| `byte`    | 8 位   | 0          | -128 到 127                                 |
| `char`    | 16 位  | （无符号） | \u0000′ \u0000′ 到 \uffff’ 或 0 到 65535    |
| `short`   | 16 位  | 0          | -32768 到 32767                             |
| `int`     | 32 位  | 0          | -2147483648 到 2147483647                   |
| `long`    | 64 位  | 0          | -9223372036854775808 到 9223372036854775807 |
| `float`   | 32 位  | 0.0        | 1.17549435e-38 到 3.4028235e+38             |
| `double`  | 64 位  | 0.0        | 4.9e-324 到 1.7976931348623157e+308         |

**字符串属于class类别** **数据类型不受OS影响，可移植性**

​	整型数据类型：**整型默认为int**

​		byte/short/int/long (1字节 2字节 4 8字节) long x= 98231821L; 后面需要*加L*(默认为int 选long必须加L) 

​		1字节=8byte 2^8=128  (-128～127 一共256数字)

​	浮点类型：**浮点默认为double**

​		float/double 4字节 8字节 将整数部分与10次方分开存储 故表达范围广 使用float y=12.21F（float必须使用F 默认为double  不加f发生自动类型转换，报错;

​	字符型：

​		char 1字符=2字节=16byte; 

​		char c='a'; 有且只能一个字符，不可空，字符串必须class；转义字符 char c='\n';

​	布尔类型：

​		ture,false  不同于C中非0为真 只用ture/false

​		

```java
if (x=0) 	//编译错误 整数表达式x=0无法转换成boolean 但CPP中可以
```



### 01.1 类型转换：	

<img src="/Users/qlzhou/Library/Application Support/typora-user-images/image-20210105144730535.png" alt="image-20210105144730535" style="zoom:33%;" />

​	*自动类型转换* 自动设为高级类别；	char——>int——>long——>long float double

**char byte short之间运算需要转换为int或之上** 

​	*强制转换：*double d1=12.9;  int i=(int)d1; 精度损失的问题

父类到子类的类型转换必须强制类型转换实现 子类到父类可以自动进行 无继承关系之间不可转换



**字符串String：** 

- 属于引用数据类型 class;
- 双引号"" 关键字String
- 可空 String z=""；
- 可以和基本数据类型做连接运算，+，运算结果为String，包括boolean，+boolearn为bool



**变量初始化：**

- Java可以将声明放在代码任何地方

  ```java
  double salary=6000;
  System.out.println(salary);
  int vacationDays=12;	//It is okay.
  ```



**常量：**

- 关键字`final` 定义常量 

  ```java
  final double PI=3.14159;	//变量仅被赋值一次且之后无法更改，常量名通常全大写
  //类常量定义位于main方法外部 
  public static final double PITWO=3.14159;//类常量，使得其在类的多个方法中均可使用
  ```



**数学函数：**

- Math类中包括各类函数

  - Math.sqrt(x)对x开方 返回double；Math.pow(x,a) 返回x^a^ double类型；Math.floorMod(position+adjustment,x) 前者对x取余 
  - Math.sin cos tan atan atan2 exp log log10 *Math.PI Math.E* （数学常量）

- 静态导入：

  - 不必在方法和常量添加前缀Math 在源文件顶加入

    ```java
    import static java.lang.Math
    System.out.println("The square root of \u03c0 is"+ sqrt(PI));
    ```







### 01.2 装箱和拆箱

Java 语言中的每种原语类型都有一个对应的 JDK 类，如表 4 所示。

#####  原语和对应的 JDK 类

| 原语      | 对应的 JDK 类         |
| :-------- | :-------------------- |
| `boolean` | `java.lang.Boolean`   |
| `byte`    | `java.lang.Byte`      |
| `char`    | `java.lang.Character` |
| `short`   | `java.lang.Short`     |
| `int`     | `java.lang.Integer`   |
| `long`    | `java.lang.Long`      |
| `float`   | `java.lang.Float`     |
| `double`  | `java.lang.Double`    |

每个 JDK 类都提供了相应方法来解析它的内部表示，并将其转换为相应的原语类型。例如，下面这段代码将十进制值 238 转换为一个 `Integer`：

```
int value = 238;
Integer boxedValue = Integer.valueOf(value);
```



此技术称为 *装箱* ，因为您将原语放在一个包装器或箱子中。

类似地，要将 `Integer` 表示转换为对应的 `int` 类，可以对它进行 *拆箱* ：

```
Integer boxedValue = Integer.valueOf(238);
int intValue = boxedValue.intValue();
```





#### 自动装箱和自动拆箱

严格来讲，不需要显式对原语进行装箱和拆箱。可以使用 Java 语言的自动装箱和自动拆箱特性：

```
int intValue = 238;

Integer boxedValue = intValue;
//
intValue = boxedValue;
```



但是，建议避免使用自动装箱和自动拆箱，因为它们可能导致代码可读性问题。装箱和拆箱代码段中的代码比自动装箱的代码更一目了然，因此更容易阅读；为此投入额外的工作是值得的。

#### 解析和转换装箱的类型

您已经了解了如何获取一个装箱的类型，但如何才能将一个您怀疑拥有装箱类型的数字 `String` 解析到它的正确箱子中呢？JDK 包装器类也拥有实现此目标的方法：

```
String characterNumeric = "238";
Integer convertedValue = Integer.parseInt(characterNumeric);
```

显示更多

还可以将 JDK 包装器类型的内容转换为 `String`：

```
Integer boxedValue = Integer.valueOf(238);
String characterNumeric = boxedValue.toString();
```

显示更多

请注意，在 `String` 表达式中使用串联运算符时（您已在对 `Logger` 的调用中看到过），原语类型已自动装箱，而且包装器类型会自动在它们之上调用 `toString()`。非常方便。






**进制：**

- 正数的原码、补码、反码相同；
- 负数反码为除符号为外 其他各位取反；
- 负数的补码元素为反码基础上加1即可；



## 02 运算符

算术、赋值、比较（关系）、 逻辑、位运算、三元运算符

- 算术：+ - * / %(取模) ++ -- +(字符串连接)

  - 取模运算符号与被模数符号相同 -12%5= -2 

  - ```java
    short s1=10;
    s1=s1+1;	//报错，1为int出现转换进制错误
    s1=(short)(s1+1);
    s1+=1;
    s1++;	//均不会报错
    ```

    

- 三元操作符：
  
  - Condition?expression1:expression2; 条件为true则为express1值
- 位运算：
  - &（and)、|（or）、^（xor）、~（not)
  - &和｜*不使用短路*方式求值 两表达式均会计算
  - \>> 和<<运算符进行左移和右移操作



### 02.1 运算符级别

- ![image-20210105150626168](/Users/qlzhou/Library/Application Support/typora-user-images/image-20210105150626168.png)
- ![image-20210105150647485](/Users/qlzhou/Library/Application Support/typora-user-images/image-20210105150647485.png)





## 03 字符串

每个用双引号""括起来的字符串都是String类的实例

String类没有提供直接修改字符串的方法 String对象又称为*不可变字符串* 

- 子串：

  - String类的substring方法 类似切片

    ```java
    String gretting="hello";
    String s1=gretting.substring(0,3);//不包括第三位数
    ```

- 拼接**+**

  - 字符串与非字符串的值拼接时，后者被转换为字符串；
  - 任何一个Java对象均可以转换成字符串

- *equals方法*检测字符串相等 

  - 忽略大小写比较字符串使用equalsIgnoreCsae

  - s.equals(t);	//检测s与t是否相等 返回boolean类型

  - s/t 可以是变量也可以是常量

    ```java
    "hello".equals(t);
    ```

  - 不可使用==判断相等 这里比较的是内存中的位置 这里将Java中String类的对象当作指针处理

    - CPP可以直接==判断相等 C中使用strcmp函数 
    - compareTo可以类似strcmp函数使用

    ```java
    if(s.compareTo("hello")==0)
    ```

- 空串和Null串

  - 空串为长度为0的字符串 仍为Java对象 具有长度和内容

  ```java
  s.equals("");
  if(s.compareTo("")==0);
  ```

  - Null为String变量存储的特殊的值 名为null 表示没有任何对象与该变量关联

  ```java
  if (s=null)
  if(s!=null&&s.length()!=0)
  ```

- 常用的String类的方法**API**

  API文档阅读 API文档是JDK中的一部分 是HTML格式 JDK/docs/api/index.html可以查看

  - bool s.equals(s1)、bool s.startwith(s1)（以s1开头则返回true）、bool s.endwith(s1)、
  - int返回值方法有：
    - indexOf(s)/indexOf(s,int fromeIndex)（从0/未忽略的from index开始搜索第一个匹配的子串位置
    - lastIndexOf(s,fromIndex int)
  - String类返回
    - s.replace(charSequence oldString, charSequence newString)  将原始字符串中的oldString换为newString
    - Substring(int begin, int end) 取出子串 类似切片 end不包括
    - toLowerCase() toUpCase() 大小写转换
    - s.trim()返回去除首尾空格的字符串

- 字符串构造: **StrinBuidler类**

  - ```java
    StringBuilder s=new StringBuilder();	//构建空字符串构建器 默认为长度为0的string
    s.append(str);	//使用append方法追加字符
    s.append("test");
    //将StringBuilder类转为String类 ——>使用toString方法
    String s1=s.toString;
    
    ```

  - StringBuilder类的常用方法:

    | StringBuidler()                     | 构建空字符串构造器           | int length()                 | 返回代码单元数量                 |
    | ----------------------------------- | :--------------------------- | ---------------------------- | -------------------------------- |
    | StringBuilder append(str/char)      | 追加字符串/字符              | void setCharAt(int i,char c) | 将i处代码单元设置为c             |
    | StringBuilder inset(int i,str/char) | 在i处插入字符串/字符         | StringBuilder delete(i,j)    | 删除从i到j-1的代码单元并返回初始 |
    | String toString()                   | 构建与构建器相同内容的字符串 |                              |                                  |

    



## 04 输入输出

- 输入流:

  - ```java
    import java.util.*;	//Scanner类定义在java.until包内 需要import进入(非java.lang包)
    //构造Scanner对象 再与输入流关联
    Scanner in=new Scanner(System.in);
    //nextLine方法读取一行(包括空格)
    String s=in.nextLine();
    //next方法以空白符作为分隔 读取单个单词
    String s1=in.next();
    //nextInt方法 读取单个整数 nextDouble方法读取单个浮点数
    
    ```

  - Scanner(inputStream in) 给定输入流创建Scanner对象 `String nextLine()/next()/ int nextInt()/double nextDouble()` 读取 `boolean hasNext()/hasNextInt()/hasNextDouble()`检测输入中是否还有其他

- 输出:

  - 格式化输出 `System.out.print(x)`将x对应的类型所允许的最大非0数字位数打印输出x

  - `System.out.printf("test%* test%*",xi...)` 每一个按照%字符开始的格式说明符将其与对应参数替换 

  - <img src="/Users/qlzhou/Library/Application Support/typora-user-images/image-20210106154602643.png" alt="image-20210106154602643" style="zoom:40%;" />

  - 常见转换符 *s转换符可以格式化任意对象*

    `String m=String.format("tets this is %s you age is %d",s,a);` 使用静态String.format方法创建格式化字符串

  - 时间输出: `System.out.printf("%tc",new Date());` %t+* 即可输出格式化后的Date()时间

  - 

<img src="/Users/qlzhou/Library/Application Support/typora-user-images/image-20210106154724239.png" alt="image-20210106154724239" style="zoom:40%;" />

- 文件输入输出:

  - 读取:

    ```java
    //in即为文件中的数据流
    Scanner in=new Scanner(Paths.get("file.txt"),"UTF-8"); //文件地址\\
    e.g.:
          Scanner in=new Scanner(Paths.get("/Users/qlzhou/Movies/test.txt"),"UTF-8");
          String show=in.nextLine();
          System.out.println(show);
          int sInt=in.nextInt();
          System.out.println(sInt);
          System.out.println(in.hasNextInt());
    
    ```

  - 输出:

    构造PrintWriter对象 `PrintWriter out = new PrintWriter("file.txt","UTF-8");` 再如同system.out使用println等 

    ```java
    			Scanner in=new Scanner(Paths.get("/Users/qlzhou/Movies/test.txt"),"UTF-8");
          String show=in.nextLine();
      System.out.println(show);
          int sInt=in.nextInt();
          System.out.println(sInt);
          System.out.println(in.hasNextInt());
    
          PrintWriter out1=new PrintWriter("test1.txt","UTF-8");
          String test="hello\n test\n test\t";
          out1.println(test);
          String name= new String();
          name="test";
          out1.println(name);
          out1.print(name); 
    ```
    



## 05 控制结构

------

- 没用`goto`语句  
- 嵌套块中不可声明同名变量;
- 不同的for循环中可相同的变量,for语句内部的变量无法在循环体之外使用;
- switch 多重选择, 无匹配的case则执行存在的default;

##### 大数

------

`java.math`包的BigInterget和BigDecimal类处理整数和浮点数;

使用类中的`add和multiply`方法计算





## 06 Array数组

- 三部曲：声明数组名称和类型；创建数组；初始化元素

```java
//完整三部曲
double[] a;
a=new double[N];
for(int i=0;i<N;++i){
  a[i]=0;
}
//简易 double类型变量默认为0.0
double[] a=new double[N];

//声明、创建、初始化
int[] a={1,2,3,4,5};
```

- 数字数组初始化自动为0, boolean类型为false, 对象数组为nil;
- 允许**二维数组**每行长度不同 
  - 数组名.length——>行数
- **增强循环遍历foreach:**

`for(var : collection) statement` 使用var暂存集合collection中的每个元素并执行

- 常用`API`

  `System.arrayCopy(source,startIndexInSource,dest,startIndexInDest,n);` 从source数组的第startIndexInSource的n个元素到dest数组的startIndexInDest

  `Array.sort(数组名)`   需要导入java.util包





- **Arrays.toString():**

返回[*,*,*,]中间逗号隔开的字符串

- 数组拷贝

  Arrays.copyOf方法

- 数组排序

  对数值类数组 Arrays类中sort方法 将原数组进行递增排序 无返回值



- Be declared before java variables can be used;
- java variables must hava a specific type;
- Java variables types can't chage;
- Types are verified before running; //运行前检查变量格式（Python不先检查）



- Functions must be part of a class in java;

- 使用pulic static define 函数；

- parameters of a functions must have a declared type, and the return value of the function must have a declared type; //函数参数以及函数返回值必须被声明；

- Functions in java return only one value; //函数返回值仅一个

  ```java
  public class HelloWordl {
      public static int larger(int x,int y){
          if(x>y){
              return x;
          }
          return y;
      }
      public static void main(String[] args) {
          System.out.println(larger(-5,10));
            }
  }
  ```

  

  

  

  在该文件中需使用 Add 为类名方可编译通过。

  ```java

  输入与输出
import java.util.Scanner;
  public class PrintHello {

  	public static void main(String[] args) {
		// TODO Auto-generated method stub
  		System.out.println("hello world,");
		Scanner in = new Scanner(System.in);
  		System.out.println("echo: "+in.next());
		
  	}

  }
//注意在读入用户的输入的时候，并不是在Scanner in = new Scanner(System.in);这句话进行读入的，这里只是定义了一个对象
  //真正读入变量的是in.next()，用户的输入是在这句话执行的
  
  //in.next()表示读入一个字符串，如果是in.nextInt()就表示读入一个数字
  //格式化输出
  System.out.printf("%.2f",3.14159);
  //字符串链接
  import java.util.Scanner;
  public class PrintHello {
  
  	public static void main(String[] args) {
  		// TODO Auto-generated method stub
		System.out.println("2+3="+(2+3));
  	}

  }
  //当字符串与数字相加的时候，数字会自动转换成字符串
  
  定义变量
import java.util.Scanner;
  public class PrintHello {

  	public static void main(String[] args) {
		// TODO Auto-generated method stub
  		System.out.println("please input two integers:");
  		Scanner in = new Scanner(System.in);//一个用于读入的对象
  		int a = in.nextInt();
		int b = in.nextInt();
  		System.out.println("answer: "+a+"+"+b+"="+(a+b));
		in.close();
  	}
  }
  

  一行定义多个变量
int a = 1,b = 2;
  常量的定义
final int money = 100; //定义之后money就一直是100，不能被修改
  
  运算符优先级
  
  ```





- 文件名必须与public类名匹配且public仅有一个 `javac 编译产生多个.class文件 `可以将包含main方法的类名提供给字节码解释器`java *`来启动
- 构造器
  - 与类同名 
  - 构造器无返回值
  - 伴随`new`操作符一起调用
- `main`方法
  - 静态方法不对对象进行操作 
  - 每一个类均可有一个main方法 



## 07 类和对象

------

`instanceof`关键字判断对象属于哪个类

引用的概念，如果一个变量的类型是 类类型，而非基本类型，那么该变量又叫做引用。

`Hero h=new Hero();` h这个引用指向这个对象Hero

多个引用指向同一个对象

```java
public class Hero{
    String name;
    int moveSpeed;
    public static void main(String[] args){
        //h这个引用指向对象Hero
        Hero h = new Hero();
        //h1这个引用同样指向h所指向的对象
        Hero h1 = h;
    }
}
```



#### 继承

------

`extends` 继承别的类使用其内的方法、属性



#### 方法的重载

------

指: 方法名一样,但是参数个数、类型、顺序不同 

自动依据传递参数的类型 数量,自动调用相同名称中的方法

```java
import java.util.*;

import javax.lang.model.element.Name;
public class base_learning {
    public static void main(String[] args) {
        class  Hero {
            String name;
            int moveSpeed;            
        }
        class ADHero extends Hero{
            public void attack(){
                System.out.println(name+"攻击不知道对象");
            }
            public void attack(Hero h1){
                System.out.println(name+"攻击一次"+h1.name);
            }
            public void attack(Hero h1,Hero h2) {
                System.out.println(name+"攻击两次"+h1.name+" "+h2.name);
            }
        }
        ADHero b = new ADHero();
        b.name="wang";
        
        Hero h1=new Hero();
        Hero h2=new Hero();
        h1.name="zhou";
        h2.name="xu";
        //自动依据传递参数在相同的方法中选择
        b.attack();
        b.attack(h1);
        b.attack(h1,h2);
    }
}

```



#### 类和对象——构造方法

------

实例化对象时,必然调用构造器 无

参的构造方法时不显式出来则默认提供无参的构造方法; 若提供有参的构造方法则默认的无参构造器不存在

构造器也可以重载

子类的构造方法调用父类的多个构造方法中的指定使用`super`完成

```java
public class Hero{
    String name;
    float hp;
    float armor;
    int moveSpeed;
    //一个参数的构造器
    public Hero(String heroname){
        name=heroname;
    }
    //两个参数的构造器——>重载
    public Hero(String heroname, float herohp){
        name=heroname;
        hp=herohp;
    }
}
```



#### 类和对象——this

------

**this关键字:**代表 **当前对象**

- **this**访问属性			

  ```java
  //只能访问到属性,而无法访问类
  //参数名称与属性名称一致, 从而需要将参数设置不同的变量名
  public void setName1(String name){
      name=name;		//无法赋值
  }
  
  //this.name 代表属性name
  public void setName2(String name){
      this.name=name;
  }
  
  ```

- **this**调用其他构造器 

  - 必须放在方法的首行 super也必须在首行 溢出无法共存
  
  ```java
  	 public Hero(String name){
          System.out.println("此处调用一个参数的构造器");
          this.name=name;
      }
      public Hero(String name, float hp){
  //这里this(name)会调用第一个一参数的构造器——>相当于new Hero(name)使用
          this(name);
          System.out.println("此处调用两个参数的构造器");
          this.hp=hp;
      }
  ```



#### 类和对象——传参

------

- 八大基本类型的传递

  - 在方法内无法修改方法外的基本类型的参数——仅为值传递

    
    
  - 类的类型传参——>引用

    ```java
    	public void attack(Hero hero, float damage) {
            hero.hp-=damage;			//可以修改hero.hp值
        }
    
    //此处为新的同名Hero
    	public void revive(Hero h) {
            h= new Hero("wang",886);
            System.out.println("复活后的h的血量值: "+h.hp);
        }
    
    	h.attack(h, 900);
        System.out.println("受到攻击之后的血量值:"+h.hp);
    
        h.revive(h);
    //仍然为收到伤害后的负值 而非revie(h)中的h的hp
        System.out.println("复活之后的血量值:"+h.hp);
    
    
    ```

  - 可变参数 (String... args) 





#### 类和对象——package

------

比较接近的类规划在包内

- 同一包下其他类直接使用, 其他包内的类,需要	`import` 相对路径

打包命令 `jar -cvf filename/pathname`

- package 内可以嵌套package  使用`import`导入







#### 类和对象——访问修饰符、类属性/静态属性、类方法/静态方法 (用 static修饰的属性和方法 )

------

成员变量的修饰符

- `private` 私有 	`public` `protected`   `package/friendly/default/` (无修饰符)

  

当属性被`static`修饰——>静态属性/类属性

- 所有对象共享一个值

- 访问方法: 对象.类属性  类.类属性

  ```java
  public class Hero {
      String name;
      float hp;
      float armor;
      int moveSpeed;
      static String copyRight;
      public Hero(String name){
          System.out.println("此处调用一个参数的构造器");
          this.name=name;
      }
      public Hero(String name, float hp){
          this.name=name;
          // this(name);
          System.out.println("此处调用两个参数的构造器");
          this.hp=hp;
      }
      public Hero(){
  
      }
      public static void main(String[] args) {
          Hero h = new Hero("wang",886);
  
          Hero.copyRight="zhou";
          Hero g=new Hero("hao",100);
          System.out.println(g.copyRight);
          System.out.println(h.copyRight);
  		System.out.println(Hero.copyRight);
  
      }
  }
  
  ```

  

静态方法/类方法: 访问类方法无需建立在存在对象的基础上



​				访问同样是两种方式 对象访问和类访问调用 

```java
		//静态方法/类方法——可以直接通过Hero.battleWin()访问
   		public static void battleWin() {
            System.out.println("这是类方法");        
        }
		//必须在对象存在 Hero h = new Hero(); 的基础上访问h.die();
        public void die() {
            hp=0;
        }
		h.die();
        System.out.println(h.hp);

        Hero.battleWin();
```



**枚举enmu:**

使得枚举的switch中的case限定在`enmu`规定的values内

```java
public class Hero {
    public enum Season {
        a,b,c,d
    }

    public static void main(String[] args) {
        Season s=Season.a;
        switch (s) {
            case a:
                System.out.println("aaa");
                break;
        
            default:System.out.println("hhh");
                break;
        }
    }
    
}
```

####  `main`方法

------

- JVM调用类的main方法——>`public`
- 执行main方法无需创建对象——>`static`
- `String[] args` 表示接受运行时参数

## 08 接口与继承

------

#### 08.1 接口

相当于约定, 实现接口必须提供接口声明总的方法



使用关键字`implements` 

```java
	public interface AD {
        public void physicAttack();
        
    }

    public class ADHero extends Hero implements AD {
        public void physicAttack(){
            System.out.println("攻击");
    }
```



#### 08.2 override 方法的重写/覆盖---隐藏

------

子类继承父类后对对象方法的重写 调用override的方法则会执行重写的方法

`suoer`显式调用父类的构造器/属性/方法——避免重写/隐藏



#### 08.3 多态

------

0. 操作符的多态——>`+`可以作为算术运算/字符串连接
1. **类的多态——>** 父类引用指向子类对象

#### 08.4 Object类

------

- 是所有类的父类,子类默认继承`Object`类

- `Object`提供`toString()`方法 返回当前对象的字符串表达  

  ```java
  public class Hero {
      String name;
      float hp;
      float armor;
      int moveSpeed;
      public String toString(){
          return name;
      }        
  
      public static void main(String[] args) {
          Hero h = new Hero();
          h.name="zhou";
          System.out.println(h.toString());
      //直接打印对象则默认为Object.toString()返回
          System.out.println(h);
      }
  }
  ```

  

- 对象没有任何引用时, JVM自动调用`finalize()`方法进行**垃圾回收**

- `equals()`判断对象内容是否相同 返回boolean 指的是引用对象

  ```java
  public class Hero {
  
      public String name;
      protected float hp;
  
  
      public boolean equals(Object o) {
          if(o instanceof Hero) {
              Hero h = (Hero) o;
              return this.hp==h.hp;
          }
          return false;
      }
  
      public static void main(String[] args) {
          Hero h1=new Hero();
          h1.hp=300;
          Hero h2=new Hero();
          h2.hp=400;
          Hero h3=new Hero();
          h3.hp=300;
          
          System.out.println(h1.equals(h2));	//FALSE
          System.out.println(h1.equals(h3));	//TRUE
      }
  }
  ```

- `==` 判断两个引用指向同一个对象否

- `hashCode()`

  - 获取Hash码

- `wait() notify() notifyAll() `线程同步方法

- `getClass()` 返回对象的类对象

  ```java
  
  ```

- `instanceOf` 返回boolean 是否是类的类型 





#### 08.5 final

------

- **修饰类:**
  - 类被修饰成final时表示无法被继承
- **修饰方法:**
  - 无法被重写
- **修饰基本变量:**
  - 只能一次赋值
- **修饰引用:**
  - 只能一次指向对象机会





#### 08.6 抽象类 abstract

------

​	类中的方法没有实现体 是一个空方法 

​	类中存在抽象方法时 必须被声明为抽象类

`public abstract class Hero{ public abstract void attack();}	` 

​	不存在抽象方法时候 也可声明为抽象类——>无法被实例化

 `Hero h = new Hero() `

​	抽象类不可直接产生对象实例化

**抽象类和接口的区别:**

接口数据成员必须初始化,接口内方法必须声明为abstract

区别1：

子类只能继承一个抽象类，不能继承多个
子类可以实现**多个**接口
区别2：
抽象类可以定义
public,protected,package,private
静态和非静态属性
final和非final属性
但是接口中声明的属性，只能是
public
静态
final的
即便没有显式的声明



#### 08.7 内部类

------

非静态内部类 静态内部类匿名类 本地类 



- **非静态内部类:**
  
  - 实例化语法 `new 外部类().内部类()` 建立在外部类存在的基础上 可以直接访问外部类的`private`实例属性
  
  - 非静态内部类可以直接调用外部类的属性
  
  - ```java
    public class Hero{
        private String name;
        float hp;
        int moveSpeed;
        //非静态内部类——>只有外部类Hero存在时才有意义
        class battleScore{
            int kill;
        }
        public static void main(String){
            Hero g = new Hero();
            //实例化对象
    		h.name="wo";
            //实例化内部类
            battleScore bb =g.battleScore();
            
            
        }
    }
    ```
  
    
  
- **静态内部类:** 

  - 直接实例化 `外部类.内部类 * = new 外部类.内部类`

  - ```java
    public class Hero{
        ...
        static class EnemyCrystal{
            int hp=500;
            public void chechVictory(){
                if(hp==0) Hero.battlewin();
                //静态内部类不可访问外部Hero的属性name
            }
        }
        public static void main(String[] args){
            //实例化静态类可以直接不存在Hero的对象基础
            Hero.EnemyCrystal c= new Hero.Enemyrystal();
        }
    }
    ```

    

- **匿名类:**

- **本地类:**





## 09 包和访问权限

### 00 package & import

------

jar命令可以打包

import导入不在同一包内其他类





## 10 多线程

### 00 进程和线程

------



- 进程是程序的动态执行过程 由于CPU时间片使得多进程同时运行

- 线程是进程内部单一的顺序控制流

- 区别:

  - 线程和进程都是并发的基本单位
  - 线程是比进程更小的执行单元
  - 进程有专用的内存区域 线程共享内存单元 

  

  

### 01 激活多线程:

------



- 线程需要扩展自`Thread`类 成为其子类
- 线程处理需要被编写在`run()`方法内部 是定义在Thread类内的方法 新代码编写在run()内就是覆盖的操作
- ![2tor1J](https://cdn.jsdelivr.net/gh/flyingchase/Private-Img@master/uPic/2tor1J.png)
- 当子类已经有其他父类时 无法再`extends Thread` 使用 `Runnable接口`

![0gB4J5](https://cdn.jsdelivr.net/gh/flyingchase/Private-Img@master/uPic/0gB4J5.png)

- 实现了Runable接口还需要调用Tread()中的start()方法才可以 

- **Thread类和Runable接口关系:**
  - Thread类实现Ruanble接口——>是Runable接口的子类
  - Runable接口声明了抽象的run()方法——>必须在实现Runable接口的类内定义run() 方法
  - ![JbYt7E](https://cdn.jsdelivr.net/gh/flyingchase/Private-Img@master/uPic/JbYt7E.png)





### 02 线程操作方法

------

**状态**

- 新建  `Thread thread = new Thread()` 具有相应的内存空间但不可运行
- 就绪 调用`start()`方法启动线程 队列排队
- 运行 就绪状态被调用获得cpu资源 自动调用`run()`方法 
- 堵塞 `sleep() suspend() wait()` 线程之间的`join()`
- 死亡 `stop()` 或者`run()`执行完

**操作:**

<img src="https://cdn.jsdelivr.net/gh/flyingchase/Private-Img@master/uPic/C3Zgl4.png" alt="C3Zgl4" style="zoom:50%;" />



### 03 线程安全/同步

------

保证当前运行代码的**原子性**:

 即cpu不去执行其他线程中、可能影响当前线程中的下一句代码的执行结果的代码块

等下一句执行完成后再去执行其他线程中的有关代码块

```java
//对代码块进行同步
synchronized(对象){
    '''
}

//对函数同步
访问控制符 synchronized 返回值类型 方法名称 (参数) {
    '''
}
```



**死锁:**

一组线程或者进程 其中每个都在等待一个只有其他线程/进程才可以执行的操作

















## 11 IO操作







## 12 Java常用类库/集合

### 12.01 String/StringBuilder/StringBuffer

------

位于`java.lang`包内 无需导入import即可使用

- String类的对象内容一旦被初始化即不可被改变
- 

使用`dash`查看常用的方法

**String、StringBuffer、StringBuilder 有什么区别？**

​	1、String 一旦创建不可变，如果修改即创建新的对象，StringBuffer 和 StringBuilder 可变，修改之后引用不变。

​	2、String 对象直接拼接效率高，但是如果执行的是间接拼接，效率很低，而 StringBuffer 和 StringBuilder 的效率更高，同时 StringBuilder 的效率高于 StringBuffer。

​	3、StringBuffer 的方法是线程安全的，StringBuilder 是线程不安全的，在考虑线程安全的情况下，应该使用 StringBuffer。

`StringBuffer`和`StringBuilder`有相同的父类`AbstractStringBuilder`

```java
StringBuffer stringBuffer = new StringBuffer();
System.out.println("StringBuffer:"+stringBuffer);
System.out.println("StringBuffer的长度:"+stringBuffer.length());

stringBuffer = new StringBuffer("Hello World");
System.out.println("StringBuffer:"+stringBuffer);
System.out.println("下标为2的字符是："+stringBuffer.charAt(2));

stringBuffer = stringBuffer.append("Java");
System.out.println("append之后的StringBuffer："+stringBuffer);

stringBuffer = stringBuffer.delete(3, 6);
System.out.println("delete之后的StringBuffer："+stringBuffer);

stringBuffer = stringBuffer.deleteCharAt(3);
System.out.println("deleteCharAt之后的StringBuffer："+stringBuffer);

stringBuffer = stringBuffer.replace(2,3,"StringBuffer");
System.out.println("replace之后的StringBuffer："+stringBuffer);

String str = stringBuffer.substring(2);
System.out.println("substring之后的String："+str);

str = stringBuffer.substring(2,8);
System.out.println("substring之后的String："+str);

stringBuffer = stringBuffer.insert(6,"six");
System.out.println("insert之后的StringBuffer："+stringBuffer);
System.out.println("e的下标是："+stringBuffer.indexOf("e"));
System.out.println("下标6之后的e的下标是："+stringBuffer.indexOf("e",6));

stringBuffer = stringBuffer.reverse();
System.out.println("reverse之后的StringBuffer："+stringBuffer);

str = stringBuffer.toString();
System.out.println("StringBuffer对应的String："+str);
```



### 12.02 基本数据类型的包装类 拆装箱

------

![GCmxNT](https://cdn.jsdelivr.net/gh/flyingchase/Private-Img@master/uPic/GCmxNT.png)

包装类与基本数据类型之间的转换:

```java
String a="123";
int i = Integer.parseInt(a);
System.out.println(i);


```



### 12.03 类集合框架

------

Collection定义的方法:

![NpsW79](https://cdn.jsdelivr.net/gh/flyingchase/Private-Img@master/uPic/NpsW79.png)



#### 12.03.01 List接口







































































































































































d

