# ChainCode之资产管理编写笔记

#### 对资产Marble进行存储和查询时，可以用到复合键来查询


``
stub.createCompositeKey
``

#### 编写transferMarblesBasedOnColor() 遇到的新知识

- 涉及到了复合键的查询
- 读取的参数，用迭代器来接收。更改参数时，采用循环遍历的方式。

````
  stub.GetStateByPartialCompositeKey() // 利用复合键进行查询   
  stub.SplitCompositeKey()  // 将复合键进行分离
````

#### 编写getQueryResultForQueryString（）遇到的新知识

- 在该函数中，使用了富查询。而富查询的数据库是CouchDB，该数据库的查询语法和MongoDB的富查询语法一样。


#### 编写getMarbleByRange() 调用的接口

```` 
    stub.getStateByRange()
````

#### 编写getMarbleHistory() 调用的接口

````
    stub.GetHistoryForKey()
````
#### 需要练习的技巧：

- 迭代器的使用
- 富查询语句