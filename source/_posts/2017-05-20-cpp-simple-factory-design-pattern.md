---
layout: post
title: c++设计模式 - 简单工厂模式 (一)
date: 2017-05-20 10:30:00
categories: c++
tag: Design Patterns
---

# 问题描述
之前在公司做了一个阅读器。首先，需要将电子书中的内容渲染到屏幕上，而电子书每一页都包含各种各样的内容，比如：图形、图像和文字等等；不同的内容，就是不同的对象；在将不同的内容渲染到屏幕上之前，就需要new操作，建立不同的对象，然后再在屏幕上进行描绘。这个时候，就需要进行很多new操作，new操作分布在代码的不同地方，管理起来很麻烦，而且也很乱，到后期扩展和维护的时候，有的时候，对象多的让开发人员不知道这个对象是干什么的，这就增加了难度；同时，new操作，都会有对应的异常处理，最后，就会发现，在代码中，new了一个对象，然后，就跟着一段异常处理代码，这时编码变的极其混乱和臃肿。那么怎么办？怎么办？此时，我们需要一个新的类，专门从事对象的建立和释放，之后，对象的各种操作，和这个类没有任何关系。这个专门建立对象的类，向外暴漏创建对象的接口，供外部调用。

工厂模式有一种非常形象的描述，建立对象的类就如一个工厂，而需要被建立的对象就是一个个产品；在工厂中加工产品，使用产品的人，不用在乎产品是如何生产出来的。从软件开发的角度来说，这样就有效的降低了模块之间的耦合。

# UML类图

对于工厂模式，具体上可以分为三类：

>+ 简单工厂模式
+ 工厂方法模式
+ 抽象工厂模式

对于上面的三种工厂模式，从上到下逐步抽象，并且更具一般性。而这篇博文主要讲的是简单工厂模式，后两种会在之后的博文中接着总结。

![image](/images/cpp-simple-factory-design-pattern-1-1.png)

ProductA、ProductB和ProductC继承自Product虚拟类，Show方法是不同产品的自描述；Factory依赖于ProductA、ProductB和ProductC，Factory根据不同的条件创建不同的Product对象。

# 适用场合

1. 在程序中，需要创建的对象很多，导致对象的new操作多且杂时，需要使用简单工厂模式；
2. 由于对象的创建过程是我们不需要去关心的，而我们注重的是对象的实际操作，所以，我们需要分离对象的创建和操作两部分，如此，方便后期的程序扩展和维护。

# 代码实现
```c++
#include <iostream>
#include <vector>

using namespace std;

typedef enum ProductTypeTag
{
    TypeA,
    TypeB,
    TypeC
}PRODUCTTYPE;

// Here is the product class
class Product
{
public:
    virtual void Show() = 0;
};

class ProductA : public Product
{
public:
    void Show()
    {
        cout<<"I'm ProductA"<<endl;
    }
};

class ProductB : public Product
{
public:
    void Show()
    {
        cout<<"I'm ProductB"<<endl;
    }
};

class ProductC : public Product
{
public:
    void Show()
    {
        cout<<"I'm ProductC"<<endl;
    }
};

// Here is the Factory class
class Factory
{
public:
    Product* CreateProduct(PRODUCTTYPE type)
    {
        switch (type)
        {
        case TypeA:
            return new ProductA();

        case TypeB:
            return new ProductB();

        case TypeC:
            return new ProductC();

        default:
            return NULL;
        }
    }
};

int main(int argc, char *argv[])
{
    // First, create a factory object
    Factory *ProductFactory = new Factory();
    Product *productObjA = ProductFactory->CreateProduct(TypeA);
    if (productObjA != NULL)
        productObjA->Show();

    Product *productObjB = ProductFactory->CreateProduct(TypeB);
    if (productObjB != NULL)
        productObjB->Show();

    Product *productObjC = ProductFactory->CreateProduct(TypeC);
    if (productObjC != NULL)
        productObjC->Show();

    delete ProductFactory;
    ProductFactory = NULL;

    delete productObjA;
    productObjA = NULL;

    delete productObjB;
    productObjB = NULL;

    delete productObjC;
    productObjC = NULL;

    return 0;
}
```
