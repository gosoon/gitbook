保存密码到硬盘一条命令就可以

$ git config credential.helper store

添加远端

git remote add upstream git@git.xiaojukeji.com:devops/odin-router.git 

git pull --rebase  upstream master 

git pull upstream master

———————————— 冲突处理

先切到远端：

git branch -u upstream/release

查看冲突：

git pull -r

解决冲突：

git rebase —continue

git add monitor_init.sh

git rebase —continue

git branch -u origin/release

git push origin release:release -f

提交 pr

覆盖分支：

git push origin feature_dev_new:feature_dev_gaofeilei_1 -f

---------------------删除文件
当我们需要删除暂存区或分支上的文件, 同时工作区也不需要这个文件了, 可以使用

git rm file_path 
当我们需要删除暂存区或分支上的文件, 但本地又需要使用, 只是不希望这个文件被版本控制, 可以使用

git rm -r --cached file_path

———————-----cherry-pick 合并部分 commit 到
git checkout -b new_release

git cherry-pick 4d10ef33a3b5869cdecca16018e4a9a395a65a2e

git checkout  feature_preview_gaofeilei_3

git merge new_release

git push  origin feature_preview_gaofeilei_3 -f

—————————  合并分支
合并多个 commit 
git rebase -i  $commit_id
----------git 删除已经提交的文件

git rm --cached logs/xx.log 

然后更新 .gitignore 忽略掉目标文件，最后 

git commit -m "We really don't want Git to track this anymore!”


中转：

 git cherry-pick 4d10ef33a3b5869cdecca16018e4a9a395a65a2e
