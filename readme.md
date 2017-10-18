# Linux 命令行实用程序


## 设计说明
先分析用户的输入，打开相应的文件，再输出指定的内容。

##使用
- sNumber和eNumber为必须有的参数，sNumber代表开始页数，eNumber代表结束页数。
- lNumber为可选参数，代表一页有多少行，若缺少lNumber则默认行数为72行。
- f为可选参数，代表以换页符作为一页的分界，f和lNumber互斥。

##测试结果
####設文檔text內容為 (\f為換頁符)####
>1 1
1 2
1 3\f
2 1
2 2
2 3\f
3 1
3 2
3 3

1. ####selpg -s1 -e2 -l3 text####
![pic_1][1]
***
2. ####selpg -s1 -e2 text####
![pic_2][2]
***
3. ####selpg -s1 -e2 -f text####
![pic_2][3]


  [1]: https://imgsa.baidu.com/forum/w%3D580/sign=17ca3ef38926cffc692abfba89004a7d/e7a029dda3cc7cd92d9f7cb33201213fb80e913d.jpg
  [2]: https://imgsa.baidu.com/forum/w%3D580/sign=593229bd3dd12f2ece05ae687fc3d5ff/e4cc7e310a55b319c91aca1f48a98226cefc1783.jpg
  [3]: https://imgsa.baidu.com/forum/w%3D580/sign=03bc4448b81bb0518f24b320067bda77/1fe831d12f2eb938319b3298de628535e4dd6ff1.jpg