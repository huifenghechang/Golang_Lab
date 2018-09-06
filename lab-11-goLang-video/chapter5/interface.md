# 类型断言

在Go语言中，对接口进行类型断言，一共有两种方法


````

var k = 8

// 方法1: comma-ok
if value , ok := k.(Int) ok{
    fmt.printf("the type of k is Int")
}

// 方法2: Switch 测试
switch value = k.(type){
    case int: //
    case string: //
}


````
