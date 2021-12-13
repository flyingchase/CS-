**CS61B**

# Java基础

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





*变量：*	包括类型、名称、值； declared

![java数据类型](/Users/qlzhou/Desktop/java数据类型.jpg)

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



**类型转换：**	

<img src="/Users/qlzhou/Library/Application Support/typora-user-images/image-20210105144730535.png" alt="image-20210105144730535" style="zoom:33%;" />

​	*自动类型转换* 自动设为高级类别；	char——>int——>long——>long float double

**char byte short之间运算需要转换为int或之上** 

​	*强制转换：*double d1=12.9;  int i=(int)d1; 精度损失的问题

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














**进制：**

- 正数的原码、补码、反码相同；
- 负数反码为除符号为外 其他各位取反；
- 负数的补码元素为反码基础上加1即可；



**运算符：**

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



- **运算符级别：**
  - ![image-20210105150626168](/Users/qlzhou/Library/Application Support/typora-user-images/image-20210105150626168.png)
  - ![image-20210105150647485](/Users/qlzhou/Library/Application Support/typora-user-images/image-20210105150647485.png)





**string——引用数据类型：**

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

    



**输入输出:**

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
    



#### 控制结构

------

- 没用`goto`语句  
- 嵌套块中不可声明同名变量;
- 不同的for循环中可相同的变量,for语句内部的变量无法在循环体之外使用;
- switch 多重选择, 无匹配的case则执行存在的default;

#### 大数

------

`java.math`包的BigInterget和BigDecimal类处理整数和浮点数;

使用类中的`add和multiply`方法计算





**Array数组：**

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

- **增强循环遍历foreach:**

`for(var : collection) statement` 使用var暂存集合collection中的每个元素并执行



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



### 类和引用

------

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

指: 方法名一样,但是参数类型不同

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

    ```java
    
    ```

  - 类类型传参——>引用

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

    





#### 类和对象——package

------

比较接近的类规划在包内

- 同一包下其他类直接使用, 其他包内的类,需要	`import` 相对路径

打包命令 `jar -cvf filename/pathname`







#### 类和对象——访问修饰符、类属性、类方法

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

















































































































































































































































































d


