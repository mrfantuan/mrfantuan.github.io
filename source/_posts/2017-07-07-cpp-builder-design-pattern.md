---
layout: post
title: c++设计模式 - 建造者模式 (五)
date: 2017-07-07 14:30:00
categories: c++
tag: Design Patterns
---
# 问题描述
在GOF的《设计模式 可复用面向对象软件的基础》中是这样说的：将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示。

这句话，似懂非懂的。一个复杂对象的创建，其通常是由很多的子对象构成；如果一个对象能够直接就创建好了，那么也不会称之为复杂对象。由于项目中需求的变化，这个复杂对象的各个部分经常会发生剧烈的变化，但是，不管怎么变化，将它们组合在一起，组成一个复杂的对象的事实是不会变的。建造者模式就提供了一种“封装机制”来将各个对象的变化隔离开，最终，组合成复杂对象的过程是不会变的。

在《大话设计模式》一书中，例举了一个很好的例子————建造小人。建造一个小人，要分为六步：头部、身体、左手、右手、左脚和右脚。与抽象工厂模式不同的是，建造者模式是在Director的控制下一步一步的构造出来的，在建造的过程中，建造者模式可以进行更精细的控制。不管人的头部、身体、左手、右手、左脚或者右脚如何变化，但是最终还是由这几部分组合在一起形成一个人，虽然是同一个建造过程，但是这个人就会有不同的表示，比如，胖子，瘦子，个高的，个低的等等。

# UML图

## 类图
![ClassDiagrams](/images/cpp-builder-design-pattern-1-1.png)

## 时序图
![ClassDiagrams](/images/cpp-builder-design-pattern-1-2.png)




# 代码实现
```c++
#include <iostream>
using namespace std;

typedef enum MANTYPETag
{
	kFatMan,
	kThinMan,
	kNormal
}MANTYPE;

class Man
{
public:
	void SetHead(MANTYPE type){ m_Type = type; }
	void SetBody(MANTYPE type){ m_Type = type; }
	void SetLeftHand(MANTYPE type){ m_Type = type; }
	void SetRightHand(MANTYPE type){ m_Type = type; }
	void SetLeftFoot(MANTYPE type){ m_Type = type; }
	void SetRightFoot(MANTYPE type){ m_Type = type; }
	void ShowMan()
	{
		switch (m_Type)
		{
		case kFatMan:
			cout<<"I'm a fat man"<<endl;
			return;

		case kThinMan:
			cout<<"I'm a thin man"<<endl;
			return;

		default:
			cout<<"I'm a normal man"<<endl;
			return;
		}
	}

private:
	MANTYPE m_Type;
};

// Builder
class Builder
{
public:
	virtual void BuildHead(){}
	virtual void BuildBody(){}
	virtual void BuildLeftHand(){}
	virtual void BuildRightHand(){}
	virtual void BuildLeftFoot(){}
	virtual void BuildRightFoot(){}
	virtual Man *GetMan(){ return NULL; }
};

// FatManBuilder
class FatManBuilder : public Builder
{
public:
	FatManBuilder(){ m_FatMan = new Man(); }
	void BuildHead(){ m_FatMan->SetHead(kFatMan); }
	void BuildBody(){ m_FatMan->SetBody(kFatMan); }
	void BuildLeftHand(){ m_FatMan->SetLeftHand(kFatMan); }
	void BuildRightHand(){ m_FatMan->SetRightHand(kFatMan); }
	void BuildLeftFoot(){ m_FatMan->SetLeftFoot(kFatMan); }
	void BuildRightFoot(){ m_FatMan->SetRightFoot(kFatMan); }
	Man *GetMan(){ return m_FatMan; }

private:
	Man *m_FatMan;
};

// ThisManBuilder
class ThinManBuilder : public Builder
{
public:
	ThinManBuilder(){ m_ThinMan = new Man(); }
	void BuildHead(){ m_ThinMan->SetHead(kThinMan); }
	void BuildBody(){ m_ThinMan->SetBody(kThinMan); }
	void BuildLeftHand(){ m_ThinMan->SetLeftHand(kThinMan); }
	void BuildRightHand(){ m_ThinMan->SetRightHand(kThinMan); }
	void BuildLeftFoot(){ m_ThinMan->SetLeftFoot(kThinMan); }
	void BuildRightFoot(){ m_ThinMan->SetRightFoot(kThinMan); }
	Man *GetMan(){ return m_ThinMan; }

private:
	Man *m_ThinMan;
};

// Director
class Director
{
public:
	Director(Builder *builder) { m_Builder = builder; }
	void CreateMan();

private:
	Builder *m_Builder;
};

void Director::CreateMan()
{
	m_Builder->BuildHead();
	m_Builder->BuildBody();
	m_Builder->BuildLeftHand();
	m_Builder->BuildRightHand();
	m_Builder->BuildLeftHand();
	m_Builder->BuildRightHand();
}

int main(int argc, char *argv[])
{
	Builder *builderObj = new FatManBuilder();
	Director directorObj(builderObj);
	directorObj.CreateMan();
	Man *manObj = builderObj->GetMan();
	if (manObj == NULL)
		return 0;

	manObj->ShowMan();

	delete manObj; // 感谢张小张同学的review
	manObj = NULL;

	delete builderObj;
	builderObj = NULL;

	return 0;
}
```

---
上面这个例子比较杂，但是也是建造者模式的应用。下面这个例子是建造者最一般，最简单的实现方法：

```c++
#include <iostream>
#include <vector>
using namespace std;

class Builder;

// Product
class Product
{
public:
	void AddPart(const char *info) { m_PartInfoVec.push_back(info); }
	void ShowProduct()
	{
		for (std::vector<const char *>::iterator item = m_PartInfoVec.begin();
			item != m_PartInfoVec.end(); ++item)
		{
			cout<<*item<<endl;
		}
	}

private:
	std::vector<const char *> m_PartInfoVec;
};

// Builder
class Builder
{
public:
	virtual void BuildPartA() {}
	virtual void BuildPartB() {}
	virtual Product *GetProduct() { return NULL; }
};

// ConcreteBuilder
class ConcreteBuilder : public Builder
{
public:
	ConcreteBuilder() { m_Product = new Product(); }
	void BuildPartA()
	{
		m_Product->AddPart("PartA completed");
	}

	void BuildPartB()
	{
		m_Product->AddPart("PartB completed");
	}

	Product *GetProduct() { return m_Product; }

private:
	Product *m_Product;
};

// Director
class Director
{
public:
	Director(Builder *builder) { m_Builder = builder; }
	void CreateProduct()
	{
		m_Builder->BuildPartA();
		m_Builder->BuildPartB();
	}

private:
	Builder *m_Builder;
};

// main
int main()
{
	Builder *builderObj = new ConcreteBuilder();
	Director directorObj(builderObj);
	directorObj.CreateProduct();
	Product *productObj = builderObj->GetProduct();
	if (productObj == NULL)
	{
		return 0;
	}
	productObj->ShowProduct();

        delete productObj;
        productObj = NULL; // 谢谢宾零同学的review
	delete builderObj;
	builderObj = NULL;
}
```
通过比较上面的两个例子，可以很容易的把建造者模式的骨架抽象出来。


# 使用要点
1. 建造者模式生成的对象有复杂的内部结构，将分步骤的去构建一个复杂的对象，分多少步是确定的，而每一步的实现是不同的，可能经常发生变化；

2. 在上面的例子中，我们都看到了最终生成的Man和Product都没有抽象类，这又导出建造者适用的一种情况，当需要创建复杂对象的过程中，复杂对象没有多少共同的特点，很难抽象出来时，而复杂对象的组装又有一定的相似点时，建造者模式就可以发挥出作用。简单的说，可能使用了建造者模式，最终建造的对象可能没有多大的关系，关于这一点，阅读《设计模式 可复用面向对象软件的基础》中的建造者模式时是最有体会的。

# 总结

一个复杂对象是由多个部件组成的，建造者模式是把复杂对象的创建和部件的创建分别开来，分别用Builder类和Director类来表示。用Director构建最后的复杂对象，而在上面Builder接口中封装的是如何创建一个个部件（复杂对象是由这些部件组成的），也就是说，Director负责如何将部件最后组装成产品。这样建造者模式就让设计和实现解耦了。

刚开始接触建造者模式的时候，最容易把建造者和抽象工厂模式混淆了。由于而这都属于创建型的设计模式，所以二者之间是有公共点的，但是建造者模式注重于对象组合，即不同的小对象组成一个整体的复杂大对象，而抽象工厂模式针对于接口编程，只是对外提供创建对象的工厂接口，不负责对象之后的处理。

建造者模式，是一个比较复杂，不容易权衡的设计模式。大家应该更多的阅读开源代码，理解他人是如何使用该模式的。从实际的应用中学习设计模式。
