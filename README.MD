# gokeeper

## 作用

一个简单的命令行版本密码记录工具，用于帮助你记录和管理众多密码。且可以使用文件作为密钥为保存的密码进行加密存储。

## 使用

此工具包括的子命令如下，更多选项可以通过对应子命令的 `--help` 进行查看。

```txt
COMMANDS:
  add      add new password
  get      get password
  list     list all password
  update   update password
  remove   remove password
  import   import passwords from csv
  export   export to csv
  decrypt  decrypt all
  encrypt  encrypt all
  gen      generate password
  help, h  Shows a list of commands or help for one command
```

都通过 `--key` 指定密钥文件，在添加密码的时候对密码，登录名，绑定信息进行加密存储，也可以不指定使用明文存储。允许每一次添加使用不同的密钥文件，但在查看时必须使用正确的密钥文件才能查看到正确的信息。

## 示例

- 新增密码

```bash
gokeeper add --key linus.jpg --name nvidia --loginid yuanxinhai@outlook.com --password so_nvidia_fuxx_u
```

新增密码时 `--name`, `--loginid`, `--password` 为必选参数，其中 `--name` 为你为该条密码记录起的名字，`--loginid` 为可用于登录的用户名，`--password` 为密码。

- 查询密码
查询所有名字，别名，注释中包含 nvi 的记录

```bash
gokeeper get --key linus.jpg nvi
```

更多使用方法可通过 `--help` 进行查看
