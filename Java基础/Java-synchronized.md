# Java-synchronized

## 概念

​	是 Java 内建的同步机制 线程对`共享资源`加锁后其他想要获取的线程必须等待，具有互斥和排他性



## 实例

- `synchronized`修饰实例方法

  对类的实例进行加锁 进入同步代码前需要获得当前实例的锁

  ``` java
  // public synchronized void method() {...}
  
  package Synchronized;
  
  public class TSynochronized implements  Runnable{
  
      @Override
      public void run() {
          for (int i = 0; i < 1000; i++) {
              increase();
          }
      }
  
      static int i=0;
      public synchronized void increase() {
          i++;
          System.out.println(Thread.currentThread().getName());
      }
  
      public static void main(String[] args) throws InterruptedException{
          TSynochronized tSynochronized = new TSynochronized();
          Thread aThread = new Thread(tSynochronized);
          Thread bThread = new Thread(tSynochronized);
  
          aThread.start();
          bThread.start();
          aThread.join();
          bThread.join();
          System.out.println("i = " + i);
      }
  }
  
  ```

  

- 修饰静态方法

  给类的对象加锁

- 修饰代码开

  给对象加锁 

