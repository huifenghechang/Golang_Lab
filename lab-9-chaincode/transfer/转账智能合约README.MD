# 使用go语言编写转账功能的ChainCode

### 编程心得：

- 在一个函数中，经常用到的值，在函数的最前面声明
- 在考虑可能出错的地方是，要将返回错误的信息写清楚，方便日后调试。


#### 编写transferMarblesBasedOnColor() 遇到的新知识

- 涉及到了复合键的查询
- 读取的参数，用迭代器来接收。更改参数时，采用循环遍历的方式。

````
  stub.GetStateByPartialCompositeKey() // 利用复合键进行查询   
  stub.SplitCompositeKey()  // 将复合键进行分离
````

#### 编写getQueryResultForQueryString（）遇到的新知识

- 在该函数中，使用了富查询。而富查询的数据库是CouchDB，该数据库的查询语法和MongoDB的福查询语法一样。



##### 需要练习的技巧：

- 迭代器的使用
- 富查询语句



#### 