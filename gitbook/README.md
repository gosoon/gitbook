hexo：

```
$ hexo g
$ hexo s 
$ hexo d
```

source-code-reading-notes 发布流程：

```
$ gh-md-toc xxx.md   // 为文章生成目录

需要在 SUMMARY.md 中添加对应的文件到章节中

$ gitbook build   // 有可能会失败，多试几次

$ cp _book/index.html .

$ cp _book/kubernetes/* kubernetes

$ ...
```


使用 gitalk 评论



教程参考：

https://blog.cugxuan.cn/2018/12/03/Markdown/How-to-use-gitbook-elegantly/

https://github.com/zhangjikai/gitbook-use

