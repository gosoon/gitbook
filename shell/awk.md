1.计算每列数据占总数的百分比
```
zw	403
yz	263
gz	213
```

```
$ for i in `cat tt | awk '{print $2}' `;do  awk 'BEGIN{printf "%0.2f%%\n",'$i'/1000*100}' ;done
```
