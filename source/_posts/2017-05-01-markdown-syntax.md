---
layout: post
title: markdown 语法手册
date:   2017-05-01 20:00:00
categories: MarkDown
tag: MarkDown
---

# 前言
这篇文章是在学习MarkDown是做的笔记，使用过之后，还是非常推荐文档相关的工作者使用MarkDown来编辑，非常高效且不用拘泥于排版布局之类的问题。

# 1.粗体和斜体
## 1.1 code
```
1. *斜体* 或者 _斜体_
2. **粗体**
3. ***粗体斜体***
4. ~~删除线~~
```
## 1.2 demo
1. *斜体* 或者 _斜体_
2. **粗体**
3. ***粗体斜体***
4. ~~删除线~~

# 2.分级标题
## 2.1 code
```
# 一级标题
## 二级标题
### 三级标题
#### 四级标题
##### 五级标题
###### 六级标题
```
## 2.2 demo
# 一级标题
## 二级标题
### 三级标题
#### 四级标题
##### 五级标题
###### 六级标题

# 3.超链接
MarkDown支持两种形式的连接语法：行内式 和 参考式 两种形式，行内式一般用的会多一些。

## 3.1 行内式

### 3.1.1 syntax
```
[链接文字](链接地址 "链接标题")
```

### 3.1.2 code
```
欢迎来到 [讨人厌的团子蜀GitBlog](http://mrfantuan.github.io)
欢迎来到 [讨人厌的团子蜀GitBlog](http://mrfantuan.github.io "讨人厌的团子蜀")
```
### 3.1.3 demo
欢迎来到 [讨人厌的团子蜀GitBlog](http://mrfantuan.github.io)

欢迎来到 [讨人厌的团子蜀GitBlog](http://mrfantuan.github.io "讨人厌的团子蜀")

## 3.2 参考式
参考式超链接一般用在学术论文上面，或者另一种情况，如果某一个链接在文章中多处使用，那么使用引用 的方式创建链接将非常好，它可以让你对链接进行统一的管理。

## 3.2.1 synatax
```
[链接文字][链接标记]
[链接标记]:链接地址 “链接标题”
```

## 3.2.2 code
```
我经常去的网站 [Google][1] \ [Baidu][2] \ [讨人厌的团子蜀GitBlog][3]

[1]:http://www.google.com "Google"
[2]:http://www.baidu.com "Baidu"
[3]:http://mrfantuan.github.io "讨人厌的团子蜀GitBlog"
```

## 3.2.3 demo
我经常去的网站 [Google][1] \ [Baidu][2] \ [讨人厌的团子蜀GitBlog][3]

[1]:http://www.google.com "Google"
[2]:http://www.baidu.com "Baidu"
[3]:http://mrfantuan.github.io "讨人厌的团子蜀GitBlog"

## 3.3 自动链接
Markdown 支持以比较简短的自动链接形式来处理网址和电子邮件信箱。

### 3.3.1 synatax
```
<http://www.baidu.com>
```

### 3.3.2 code
```
<http://www.baidu.com>
```

### 3.3.3 demo
<http://www.baidu.com>

<mrfantuan@hotmail.com>

# 5.列表
## 5.1无序列表
### 5.1.1 sysnatax
使用 */+/- 表示无序列表

### 5.1.2 code
```
- 无序列表A
- 无序列表B
- 无序列表C
```

### 5.1.3 demo
- 无序列表A
- 无序列表B
- 无序列表C

## 5.2 有序列表
### 5.2.1 sysnatax
有序列表则使用数字接着一个英文句点。

### 5.2.2 code
```
1. 有序列表A
2. 有序列表B
3. 有序列表C
```

### 5.2.3 demo
1. 有序列表A
2. 有序列表B
3. 有序列表C

## 5.3 定义型列表
### 5.3.1 sysnatax
```
header
:   content(左侧有一个冒号和四个空格)
```

### 5.3.2 code
```
Markdown
:   轻量级文本标记语言

code block
:   代码块的定义
        code(左侧8个空格)
```

### 5.3.3 demo
Markdown
:   轻量级文本标记语言

code block
:   代码块的定义

        code(左侧8个空格)

## 5.4 列表缩进
### 5.4.1 sysnatax
列表项目标记通常是放在最左边，但是其实也可以缩进，最多 3 个空格，项目标记后面则一定要接着至少一个空格或制表符。

### 5.4.2 code
```
要让列表看起来更漂亮，你可以把内容用固定的缩进整理好。
*   这里只是需要一段很长的话不要在意细节。这里只是需要一段很长的话不要在意细节。这里只是需要一段很长的话不要在意细节。
*   这里只是需要一段很长的话不要在意细节。这里只是需要一段很长的话不要在意细节。这里只是需要一段很长的话不要在意细节。
```

### 5.4.3 demo
*   这里只是需要一段很长的话不要在意细节。这里只是需要一段很长的话不要在意细节。这里只是需要一段很长的话不要在意细节。
*   这里只是需要一段很长的话不要在意细节。这里只是需要一段很长的话不要在意细节。这里只是需要一段很长的话不要在意细节。

## 5.5  包含段落的列表
### 5.5.1 sysnatax
列表项目可以包含多个段落，每个项目下的段落都必须缩进 4 个空格或是 1 个制表符。

### 5.5.2 code
```
*   这里只是需要一段很长的话不要在意细节。这里只是需要一段很长的话不要在意细节。这里只是需要一段很长的话不要在意细节。

    这里只是需要一段很长的话不要在意细节。这里只是需要一段很长的话不要在意细节。这里只是需要一段很长的话不要在意细节。

*   这里只是需要一段很长的话不要在意细节。这里只是需要一段很长的话不要在意细节。这里只是需要一段很长的话不要在意细节。
```
### 5.5.3 demo
*   这里只是需要一段很长的话不要在意细节。这里只是需要一段很长的话不要在意细节。这里只是需要一段很长的话不要在意细节。

    这里只是需要一段很长的话不要在意细节。这里只是需要一段很长的话不要在意细节。这里只是需要一段很长的话不要在意细节。

*   这里只是需要一段很长的话不要在意细节。这里只是需要一段很长的话不要在意细节。这里只是需要一段很长的话不要在意细节。

## 5.6 包含引用的列表
### 5.6.1 sysnatax
如果要在列表项目内放进引用，那 > 就需要缩进。

### 5.6.2 code
```
*   阅读的方法:
    > 打开书本。
    > 打开电灯。
```
### 5.6.3 demo
*   阅读的方法:
    > 打开书本。
    > 打开电灯。

# 6.引用
## 6.1 引用
### 6.1.1 sysnatax
引用需要在被引用的文本前加上>符号。

## 6.1.2 code
```
> 这是一个有两段文字的引用,
> 无意义的占行文字1.
> 无意义的占行文字2.
>
> 无意义的占行文字3.
> 无意义的占行文字4.
```

## 6.1.3 demo
> 这是一个有两段文字的引用,
> 无意义的占行文字1.
> 无意义的占行文字2.
>
> 无意义的占行文字3.
> 无意义的占行文字4.


## 6.2 引用的多层嵌套
### 6.2.1 sysnatax
区块引用可以嵌套(例如：引用内的引用)，只要根据层次加上不同数量的 >

### 6.2.2 code
```
>>> 请问 Markdwon 怎么用？ - 路人A
>> 自己看教程！ - 路人B
> 教程在哪？ - 路人C
```

### 6.2.3 demo
>>> 请问 Markdwon 怎么用？ - 路人A
>> 自己看教程！ - 路人B
> 教程在哪？ - 路人C

## 6.3 引用其他要素
### 6.3.1 sysnatax
引用的区块内也可以使用其他的 Markdown 语法，包括标题、列表、代码区块等。

### 6.3.2 code
```
> 1.   这是第一行列表项。
> 2.   这是第二行列表项。
>
> 给出一些例子代码：
>
>     return shell_exec("echo $input | $markdown_script");
```

### 6.3.3 demo
> 1.   这是第一行列表项。
> 2.   这是第二行列表项。
>
> 给出一些例子代码：
>
>     return shell_exec("echo $input | $markdown_script");


# 7.插入图像
图片的创建方式与超链接相似，而且和超链接一样也有两种写法，行内式和参考式写法。
语法中图片Alt的意思是如果图片因为某些原因不能显示，就用定义的图片Alt文字来代替图片。 图片Title则和链接中的Title一样，表示鼠标悬停与图片上时出现的文字。 Alt 和 Title 都不是必须的，可以省略，但建议写上。

## 7.1 行内式
### 7.1.1 sysnatax
```
![图片Alt](图片地址 "图片Title")
```
### 7.1.2 code
```
![image](/images/markdown-syntax-7.1.1.gif "吃草的猫")
```
### 7.1.3 demo
![image](/images/markdown-syntax-7.1.1.gif "吃草的猫")

## 7.2 参考式
### 7.2.1 sysnatax
在文档要插入图片的地方写![图片Alt][标记]

在文档的最后写上[标记]:图片地址 “Title”
### 7.2.2 code
```
![吃草的猫][grasscat]

[grasscat]: /images/markdown-syntax-7.1.1.gif "吃草的猫"
```
### 7.2.3 demo
![吃草的猫][grasscat]

[grasscat]: /images/markdown-syntax-7.1.1.gif "吃草的猫"

```c++
@requires_authorization
def somefunc(param1='', param2=0):
    '''A docstring'''
    if param1 > param2: # interesting
        print 'Greater'
    return (param2 - param1 + 1) or None

class SomeClass:
    pass

>>> message = '''interpreter
... prompt'''
```

``` javascript
/**
* nth element in the fibonacci series.
* @param n >= 0
* @return the nth element, >= 0.
*/
function fib(n) {
  var a = 1, b = 1;
  var tmp;
  while (--n >= 0) {
    tmp = a;
    a += b;
    b = tmp;
  }
  return a;
}

document.write(fib(10));
```

```flow
st=>start: Start:>https://www.zybuluo.com
io=>inputoutput: verification
op=>operation: Your Operation
cond=>condition: Yes or No?
sub=>subroutine: Your Subroutine
e=>end

st->io->op->cond
cond(yes)->e
cond(no)->sub->io
```

```seq
Alice->Bob: Hello Bob, how are you?
Note right of Bob: Bob thinks
Bob-->Alice: I am good thanks!
```

```seq
Title: Here is a title
A->B: Normal line
B-->C: Dashed line
C->>D: Open arrow
D-->>A: Dashed open arrow
```


```flow
st=>start: Start
e=>end: End
op1=>operation: My Operation
sub1=>subroutine: My Subroutine
cond=>condition: Yes or No?
io=>inputoutput: catch something...
st->op1->cond
cond(yes)->io->e
cond(no)->sub1(right)->op1
```

Title: Here is a title
A->B: Normal line
B-->C: Dashed line
C->>D: Open arrow
D-->>A: Dashed open arrow



```flow
st=>start: Start|past:>http://www.google.com[blank]
e=>end: End:>http://www.google.com
op1=>operation: get_hotel_ids|past
op2=>operation: get_proxy|current
sub1=>subroutine: get_proxy|current
op3=>operation: save_comment|current
op4=>operation: set_sentiment|current
op5=>operation: set_record|current

cond1=>condition: ids_remain空?
cond2=>condition: proxy_list空?
cond3=>condition: ids_got空?
cond4=>condition: 爬取成功??
cond5=>condition: ids_remain空?

io1=>inputoutput: ids-remain
io2=>inputoutput: proxy_list
io3=>inputoutput: ids-got

st->op1(right)->io1->cond1
cond1(yes)->sub1->io2->cond2
cond2(no)->op3
cond2(yes)->sub1
cond1(no)->op3->cond4
cond4(yes)->io3->cond3
cond4(no)->io1
cond3(no)->op4
cond3(yes, right)->cond5
cond5(yes)->op5
cond5(no)->cond3
op5->e
```

```c++
#include <stdio.h>

int main()
{
    printf("Hello world\n" );
}

```
