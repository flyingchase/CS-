# Git常见问题

>整理收集于互联网

### 1.修改提交信息

重写`commit`

```shell
git commit --amend -m "new message"
```

### 2.添加文件到最后一次提交

`git push` 后新增文件到最后一次提交

```shell
git add <file_name>git commit --amend HEAD~1
```

### 3.撤消提交

如果要撤消最近一次提交但保留更改，可执行以下操作：

```shell
git reset --soft HEAD~1
```

如果要撤消提交和更改，可执行以下操作：——>丢弃更改

```shell
git reset --hard HEAD~1
```

如果要撤消所有的本地更改，则可以重置为分支的原始版本：

```shell
git reset --hard origin/<branch_name>
```

如果要撤消提交而不修改现有历史记录，则可以使用 git revert，此命令通过创建新的提交来撤消提交。

```shell
git revert HEAD
```

撤消已经推送到远程分支的合并提交的安全方法是使用 git revert 命令：

```shell
git revert -m 1 <commit_id>
```

commit_id 是要还原的合并提交 id。

注意要点：

1. 可以撤消任意数量的提交。例如：`git reset HEAD~3`（返回 HEAD 之前的 3 个提交)；`git reset --hard <commit_id>`（返回特定的提交）。
2. 如果尚未推送提交，并且你不想引入糟糕的提交到远程分支，可以使用 git reset。
3. 使用 git revert 还原已经推送到远程分支的合并提交。
4. 使用 git log 查看提交历史。



```bash
# 删除 untracked files
git clean -f
 
# 连 untracked 的目录也一起删掉
git clean -fd
 
# 连 gitignore 的untrack 文件/目录也一起删掉 （慎用，一般这个是用来删掉编译出来的 .o之类的文件用的）
git clean -xfd
 
# 在用上述 git clean 前，墙裂建议加上 -n 参数来先看看会删掉哪些文件，防止重要文件被误删
git clean -nxfd
git clean -nf
git clean -nfd
```

