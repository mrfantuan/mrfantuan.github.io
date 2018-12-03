---
layout: post
title: c++设计模式 - 单例模式 (四)
date: 2017-05-26 17:30:00
categories: c++
tag: Design Patterns
---
# 问题描述
现在，不管开发一个多大的系统（至少我现在的部门是这样的），都会带一个日志功能；在实际开发过程中，会专门有一个日志模块，负责写日志，由于在系统的任何地方，我们都有可能要调用日志模块中的函数，进行写日志。那么，如何构造一个日志模块的实例呢？难道，每次new一个日志模块实例，写完日志，再delete，不要告诉我你是这么干的。在C++中，可以构造一个日志模块的全局变量，那么在任何地方就都可以用了，是的，不错。但是，我所在的开发部门的C++编码规范是参照Google的编码规范的。

全局变量在项目中是能不用就不用的，它是一个定时炸弹，是一个不安全隐患，特别是在多线程程序中，会有很多的不可预测性；同时，使用全局变量，也不符合面向对象的封装原则，所以，在纯面向对象的语言Java和C#中，就没有纯粹的全局变量。那么，如何完美的解决这个日志问题，就需要引入设计模式中的单例模式。

# 单例模式

何为单例模式，在GOF的《设计模式：可复用面向对象软件的基础》中是这样说的：保证一个类只有一个实例，并提供一个访问它的全局访问点。首先，需要保证一个类只有一个实例；在类中，要构造一个实例，就必须调用类的构造函数，如此，为了防止在外部调用类的构造函数而构造实例，需要将构造函数的访问权限标记为protected或private；最后，需要提供要给全局访问点，就需要在类中定义一个static函数，返回在类内部唯一构造的实例。意思很明白，使用UML类图表示如下。

# UML类图

对于工厂模式，具体上可以分为三类：

>+ 简单工厂模式
+ 工厂方法模式
+ 抽象工厂模式

对于上面的三种工厂模式，从上到下逐步抽象，并且更具一般性。而这篇博文主要讲的是简单工厂模式，后两种会在之后的博文中接着总结。

![image](/images/cpp-singleton-design-pattern-1-1.png)

ProductA、ProductB和ProductC继承自Product虚拟类，Show方法是不同产品的自描述；Factory依赖于ProductA、ProductB和ProductC，Factory根据不同的条件创建不同的Product对象。

# 代码实现
单例模式，单从UML类图上来说，就一个类，没有错综复杂的关系。但是，在实际项目中，使用代码实现时，还是需要考虑很多方面的。

实现一：
```c++
#include <iostream>
using namespace std;

class Singleton
{
public:
    static Singleton *GetInstance()
    {
        if (m_Instance == NULL )
        {
            m_Instance = new Singleton ();
        }
        return m_Instance;
    }

    static void DestoryInstance()
    {
        if (m_Instance != NULL )
        {
            delete m_Instance;
            m_Instance = NULL ;
        }
    }

    // This is just a operation example
    int GetTest()
    {
        return m_Test;
    }

private:
    Singleton(){ m_Test = 10; }
    static Singleton *m_Instance;
    int m_Test;
};

Singleton *Singleton ::m_Instance = NULL;

int main(int argc , char *argv [])
{
    Singleton *singletonObj = Singleton ::GetInstance();
    cout << singletonObj->GetTest() << endl;

    Singleton ::DestoryInstance();
    return 0;
}
```

这是最简单，也是最普遍的实现方式，也是现在网上各个博客中记述的实现方式，但是，这种实现方式，有很多问题，比如：没有考虑到多线程的问题，在多线程的情况下，就可能创建多个Singleton实例，以下版本是改善的版本。

实现二：
```c++
#include <iostream>
using namespace std;

class Singleton
{
public:
    static Singleton *GetInstance()
    {
        if (m_Instance == NULL )
        {
            Lock(); // C++没有直接的Lock操作，请使用其它库的Lock，比如Boost，此处仅为了说明
            if (m_Instance == NULL )
            {
                m_Instance = new Singleton ();
            }
            UnLock(); // C++没有直接的Lock操作，请使用其它库的Lock，比如Boost，此处仅为了说明
        }
        return m_Instance;
    }

    static void DestoryInstance()
    {
        if (m_Instance != NULL )
        {
            delete m_Instance;
            m_Instance = NULL ;
        }
    }

    int GetTest()
    {
        return m_Test;
    }

private:
    Singleton(){ m_Test = 0; }
    static Singleton *m_Instance;
    int m_Test;
};

Singleton *Singleton ::m_Instance = NULL;

int main(int argc , char *argv [])
{
    Singleton *singletonObj = Singleton ::GetInstance();
    cout<<singletonObj->GetTest()<<endl;
    Singleton ::DestoryInstance();

    return 0;
}
```
此处进行了两次m_Instance == NULL的判断，是借鉴了Java的单例模式实现时，使用的所谓的“双检锁”机制。因为进行一次加锁和解锁是需要付出对应的代价的，而进行两次判断，就可以避免多次加锁与解锁操作，同时也保证了线程安全。但是，这种实现方法在平时的项目开发中用的很好，也没有什么问题？但是，如果进行大数据的操作，加锁操作将成为一个性能的瓶颈；为此，一种新的单例模式的实现也就出现了。

实现三：
```c++
#include <iostream>
using namespace std;

class Singleton
{
public:
    static Singleton *GetInstance()
    {
        return const_cast <Singleton *>(m_Instance);
    }

    static void DestoryInstance()
    {
        if (m_Instance != NULL )
        {
            delete m_Instance;
            m_Instance = NULL ;
        }
    }

    int GetTest()
    {
        return m_Test;
    }

private:
    Singleton(){ m_Test = 10; }
    static const Singleton *m_Instance;
    int m_Test;
};

const Singleton *Singleton ::m_Instance = new Singleton();

int main(int argc , char *argv [])
{
    Singleton *singletonObj = Singleton ::GetInstance();
    cout<<singletonObj->GetTest()<<endl;
    Singleton ::DestoryInstance();
}
```
因为静态初始化在程序开始时，也就是进入主函数之前，由主线程以单线程方式完成了初始化，所以静态初始化实例保证了线程安全性。在性能要求比较高时，就可以使用这种方式，从而避免频繁的加锁和解锁造成的资源浪费。由于上述三种实现，都要考虑到实例的销毁，关于实例的销毁，待会在分析。由此，就出现了第四种实现方式：

实现四：
```c++
#include <iostream>
using namespace std;

class Singleton
{
public:
    static Singleton *GetInstance()
    {
        static Singleton m_Instance;
        return &m_Instance;
    }

    int GetTest()
    {
        return m_Test++;
    }

private:
    Singleton(){ m_Test = 10; };
    int m_Test;
};

int main(int argc , char *argv [])
{
    Singleton *singletonObj = Singleton ::GetInstance();
    cout<<singletonObj->GetTest()<<endl;

    singletonObj = Singleton ::GetInstance();
    cout<<singletonObj->GetTest()<<endl;
}
```
以上就是四种主流的单例模式的实现方式，如果大家还有什么好的实现方式，希望大家能推荐给我。谢谢了。

# 实例销毁
在上述的四种方法中，除了第四种没有使用new操作符实例化对象以外，其余三种都使用了；我们一般的编程观念是，new操作是需要和delete操作进行匹配的；是的，这种观念是正确的。在上述的实现中，是添加了一个DestoryInstance的static函数，这也是最简单，最普通的处理方法了；但是，很多时候，我们是很容易忘记调用DestoryInstance函数，就像你忘记了调用delete操作一样。由于怕忘记delete操作，所以就有了智能指针；那么，在单例模型中，没有“智能单例”，该怎么办？怎么办？

那我先从实际的项目中说起吧，在实际项目中，特别是客户端开发，其实是不在乎这个实例的销毁的。因为，全局就这么一个变量，全局都要用，它的生命周期伴随着软件的生命周期，软件结束了，它也就自然而然的结束了，因为一个程序关闭之后，它会释放它占用的内存资源的，所以，也就没有所谓的内存泄漏了。但是，有以下情况，是必须需要进行实例销毁的：

在类中，有一些文件锁了，文件句柄，数据库连接等等，这些随着程序的关闭而不会立即关闭的资源，必须要在程序关闭前，进行手动释放；
具有强迫症的程序员。
以上，就是我总结的两点。

虽然，在代码实现部分的第四种方法能满足第二个条件，加上析构函数能满足第一个条件。好了，接下来，就介绍另一种方法，这种方法也是我从网上学习而来的，代码实现如下：

```c++
#include <iostream>
using namespace std;

class Singleton
{
public:
    static Singleton *GetInstance()
    {
        return m_Instance;
    }

    int GetTest()
    {
        return m_Test;
    }

private:
    Singleton(){ m_Test = 10; }
    static Singleton *m_Instance;
    int m_Test;

    // This is important
    class GC
    {
    public :
        ~GC()
        {
            // We can destory all the resouce here, eg:db connector, file handle and so on
            if (m_Instance != NULL )
            {
                cout<< "Here is the test" <<endl;
                delete m_Instance;
                m_Instance = NULL ;
            }
        }
    };
    static GC gc;
};

Singleton *Singleton ::m_Instance = new Singleton();
Singleton ::GC Singleton ::gc;

int main(int argc , char *argv [])
{
    Singleton *singletonObj = Singleton ::GetInstance();
    cout<<singletonObj->GetTest()<<endl;

    return 0;
}
```
在程序运行结束时，系统会调用Singleton的静态成员GC的析构函数，该析构函数会进行资源的释放，而这种资源的释放方式是在程序员“不知道”的情况下进行的，而程序员不用特别的去关心，使用单例模式的代码时，不必关心资源的释放。那么这种实现方式的原理是什么呢？我剖析问题时，喜欢剖析到问题的根上去，绝不糊涂的停留在表面。由于程序在结束的时候，系统会自动析构所有的全局变量，实际上，系统也会析构所有类的静态成员变量，就像这些静态变量是全局变量一样。我们知道，静态变量和全局变量在内存中，都是存储在静态存储区的，所以在析构时，是同等对待的。

由于此处使用了一个内部GC类，而该类的作用就是用来释放资源，而这种使用技巧在C++中是广泛存在的。

# 模式扩展
在实际项目中，一个模式不会像我们这里的代码那样简单，只有在熟练了各种设计模式的特点，才能更好的在实际项目中进行运用。单例模式和工厂模式在实际项目中经常见到，两种模式的组合，在项目中也是很常见的。所以，有必要总结一下两种模式的结合使用。

一种产品，在一个工厂中进行生产，这是一个工厂模式的描述；而只需要一个工厂，就可以生产一种产品，这是一个单例模式的描述。所以，在实际中，一种产品，我们只需要一个工厂，此时，就需要工厂模式和单例模式的结合设计。由于单例模式提供对外一个全局的访问点，所以，我们就需要使用简单工厂模式中那样的方法，定义一个标识，用来标识要创建的是哪一个单件。由于模拟代码较多，在文章最后，提供下载链接。

# 总结
为了写这篇文章，自己调查了很多方面的资料，由于网上的资料在各方面都有很多的瑕疵，质量参次不齐，对我也造成了一定的误导。而这篇文章，有我自己的理解，如有错误，请大家指正。

由于该文对设计模式的总结，我认为比网上80%的都全面，希望对大家有用。在实际的开发中，并不会用到单例模式的这么多种，每一种设计模式，都应该在最适合的场合下使用，在日后的项目中，应做到有地放矢，而不能为了使用设计模式而使用设计模式。

# 工程下载
[SingletonPatternEx.zip](/source/projects/SingletonPatternEx.zip)
