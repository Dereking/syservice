一个 系统服务模版，主要功能:

1. 系统服务注册，支持linux mac和windows
   1. xxxx install    # 安装服务,/etc/systemd/system/xxxx.service
   2. xxxx uninstall   # 卸载删除服务
   3. xxxx stop   # 停止服务
   4. xxxx start  # 启动服务
   5. xxxx restart  # 重启服务
2. 简单http服务器，基于gin
3. yaml配置文件，config.yaml
4. redis操作客户端
5. 邮件发送客户端



# 使用方法

git clone https://github.com/Dereking/syservice.git

修改go.mod 文件里面的 module syservice 名称为你自己要的
执行

go mod tidy

修改main.go,把 服务相关参数修改为你要的，再写业务代码。

go build

执行命令安装：

xxxx install


启动 服务
systemctl start xxxx
