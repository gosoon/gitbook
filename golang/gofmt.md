1.fmt dir all code

```
find ./platforms -name "*.go" | xargs gofmt -w
```

2.use git pre commit hook fmt

.git/hooks/pre-commit

```
#!/bin/sh

# 获取git暂存的所有go代码
# --cached 暂存的
# --name-only 只显示名字
# --diff-filter=ACM 过滤暂存文件，A=Added C=Copied M=Modified, 即筛选出添加/复制/修改的文件
allgofiles=$(git diff --cached --name-only --diff-filter=ACM | grep '.go$')

gofiles=()
godirs=()
for allfile in ${allgofiles[@]}; do
    # 过滤vendor的
    # 过滤prootobuf自动生产的文件
    if [[ $allfile == "vendor"* || $allfile == *".pb.go" ]];then
        continue
    else
        gofiles+=("$allfile")

        # 文件夹去重
        existdir=0
        dir=`echo "$allfile" |xargs -n1 dirname|sort -u`
        for somedir in ${godirs[@]}; do
            if [[ $dir == $somedir ]]; then
                existdir=1
                break
            fi
        done

        if [[ $existdir -eq 0 ]]; then
            godirs+=("$dir")
        fi
    fi
done

[ -z "$gofiles" ] && exit 0


# goimports 自动导包
if [ -f /usr/local/bin/goimports-reviser ]; then  # 检测是否安装
    for file in ${gofiles[@]} ; do
        unimports=$(goimports-reviser -rm-unused -set-alias -format $file)
        if [ -n "$unimports" ]; then
            echo >&2 "goimports FAIL:\nRun following command:"
            for f in ${unimports[@]} ; do
                echo >&2 " goimports -w $PWD/$f"
            done
            echo "\n"
            has_errors=1
        fi
    done
else
	echo 'Error: /usr/local/bin/goimports-reviser not install. Run: "go get -u golang.org/x/tools/cmd/goimports"' >&2
	exit 1
fi
```

ref: [利用 git hook 规范你的代码与 commit message](https://razeen.me/posts/golang-and-git-commit-message-pre-commit/)

3. format import 

[Go语言import分组管理利器: goimports-reviser](https://zhuanlan.zhihu.com/p/411181637)


