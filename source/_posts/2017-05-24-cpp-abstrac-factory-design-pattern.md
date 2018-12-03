---
layout: post
title: c++设计模式 - 抽象工厂模式 (三)
date: 2017-05-24 15:30:00
categories: c++
tag: Design Patterns
---
# 问题描述
之前讲到了C++设计模式——工厂方法模式，我们可能会想到，后期产品会越来越多了，建立的工厂也会越来越多，工厂进行了增长，工厂变的凌乱而难于管理；由于工厂方法模式创建的对象都是继承于Product的，所以工厂方法模式中，每个工厂只能创建单一种类的产品，当需要生产一种全新的产品（不继承自Product）时，发现工厂方法是心有余而力不足。

举个例子来说：一个显示器电路板厂商，旗下的显示器电路板种类有非液晶的和液晶的；这个时候，厂商建造两个工厂，工厂A负责生产非液晶显示器电路板，工厂B负责生产液晶显示器电路板；工厂一直就这样运行着。有一天，总经理发现，直接生产显示器的其余部分也挺挣钱，所以，总经理决定，再建立两个工厂C和D；C负责生产非液晶显示器的其余部件，D负责生产液晶显示器的其余部件。此时，旁边参谋的人就说了，经理，这样做不好，我们可以直接在工厂A中添加一条负责生产非液晶显示器的其余部件的生产线，在工厂B中添加一条生产液晶显示器的其余部件的生产线，这样就可以不用增加厂房，只用将现有厂房进行扩大一下，同时也方便工厂的管理，而且生产非液晶显示器电路板的技术人员对非液晶显示的其余部件的生产具有指导的作用，生产液晶显示器电路板也是同理。总经理发现这是一个不错的主意。

再回到软件开发的过程中来，工厂A和B就是之前所说的C++设计模式——工厂方法模式；总经理再次建立工厂C和D，就是重复C++设计模式——工厂方法模式，只是生产的产品不同罢了。这样做的弊端就如参谋所说的那样，增加了管理成本和人力成本。在面向对象开发的过程中，是很注重对象管理和维护的，对象越多，就越难进行管理和维护；如果工厂数量过多，那么管理和维护的成本将大大增加；虽然生产的是不同的产品，但是可以二者之间是有微妙的关系的，如参谋所说，技术人员的一些技术经验是可以借鉴的，这就相当于同一个类中的不同对象，之间是可以公用某些资源的。那么，增加一条流水线，扩大厂房，当然是最好的主意了。

实际问题已经得到了解决，那么如何使用设计模式模拟这个实际的问题呢？那就是接下来所说的抽象工厂模式。

# UML类图

现在要讲的抽象工厂模式，就是工厂方法模式的扩展和延伸，但是抽象工厂模式，更有一般性和代表性；它具有工厂方法具有的优点，也增加了解决实际问题的能力。

![image](/images/cpp-abstrac-factory-design-pattern-1-1.png)

如图所示，抽象工厂模式，就好比是两个工厂方法模式的叠加。抽象工厂创建的是一系列相关的对象，其中创建的实现其实就是采用的工厂方法模式。在工厂Factory中的每一个方法，就好比是一条生产线，而生产线实际需要生产什么样的产品，这是由Factory1和Factory2去决定的，这样便延迟了具体子类的实例化；同时集中化了生产线的管理，节省了资源的浪费。

# 适用场合

工厂方法模式适用于产品种类结构单一的场合，为一类产品提供创建的接口；而抽象工厂方法适用于产品种类结构多的场合，主要用于创建一组（有多个种类）相关的产品，为它们提供创建的接口；就是当具有多个抽象角色时，抽象工厂便可以派上用场。

# 代码实现
```c++
#include <iostream>
using namespace std;

// Product A
class ProductA
{
public:
    virtual void Show() = 0;
};

class ProductA1 : public ProductA
{
public:
    void Show()
    {
        cout<<"I'm ProductA1"<<endl;
    }
};

class ProductA2 : public ProductA
{
public:
    void Show()
    {
        cout<<"I'm ProductA2"<<endl;
    }
};

// Product B
class ProductB
{
public:
    virtual void Show() = 0;
};

class ProductB1 : public ProductB
{
public:
    void Show()
    {
        cout<<"I'm ProductB1"<<endl;
    }
};

class ProductB2 : public ProductB
{
public:
    void Show()
    {
        cout<<"I'm ProductB2"<<endl;
    }
};

// Factory
class Factory
{
public:
    virtual ProductA *CreateProductA() = 0;
    virtual ProductB *CreateProductB() = 0;
};

class Factory1 : public Factory
{
public:
    ProductA *CreateProductA()
    {
        return new ProductA1();
    }

    ProductB *CreateProductB()
    {
        return new ProductB1();
    }
};

class Factory2 : public Factory
{
    ProductA *CreateProductA()
    {
        return new ProductA2();
    }

    ProductB *CreateProductB()
    {
        return new ProductB2();
    }
};

int main(int argc, char *argv[])
{
    Factory *factoryObj1 = new Factory1();
    ProductA *productObjA1 = factoryObj1->CreateProductA();
    ProductB *productObjB1 = factoryObj1->CreateProductB();

    productObjA1->Show();
    productObjB1->Show();

    Factory *factoryObj2 = new Factory2();
    ProductA *productObjA2 = factoryObj2->CreateProductA();
    ProductB *productObjB2 = factoryObj2->CreateProductB();

    productObjA2->Show();
    productObjB2->Show();

    if (factoryObj1 != NULL)
    {
        delete factoryObj1;
        factoryObj1 = NULL;
    }

    if (productObjA1 != NULL)
    {
        delete productObjA1;
        productObjA1= NULL;
    }

    if (productObjB1 != NULL)
    {
        delete productObjB1;
        productObjB1 = NULL;
    }

    if (factoryObj2 != NULL)
    {
        delete factoryObj2;
        factoryObj2 = NULL;
    }

    if (productObjA2 != NULL)
    {
        delete productObjA2;
        productObjA2 = NULL;
    }

    if (productObjB2 != NULL)
    {
        delete productObjB2;
        productObjB2 = NULL;
    }
}
```
