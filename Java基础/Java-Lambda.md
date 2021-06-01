# Java-Lambda

## 1 概念

### 1.01 常见 OOP 实现

- 定义函数——>创建对象使用函数

  ``` java
  public class LambdaConcept {
      public void printsth(String waitToPrint) {
          System.out.println(waitToPrint);
      }
      // 创建对象来调用函数
      public static void main(String[] args) {
          LambdaConcept lambdaDemo = new LambdaConcept();
          String waitToPrint = "This is the String waiting to print";
          lambdaDemo.printsth(waitToPrint);
      }
  }
  ```

- 创建功能接口——>对该接口定义抽象方法

  ``` java
  ```

  