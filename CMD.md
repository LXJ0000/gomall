```bash


cwgo server --server_name demo_thrift --type RPC  --idl ../../idl/echo.thrift -module github.com/LXJ0000/gomall/demo/demo_thrift

cwgo server -I ../../idl --server_name demo_proto --type RPC  --idl ../../idl/echo.proto -module github.com/LXJ0000/gomall/demo/demo_proto

# 生成用户客户端代码 /root/code/gomall/rpc_gen
cwgo client --type RPC --service user --module github.com/LXJ0000/gomall/rpc_gen --I ../i
dl/ --idl ../idl/user.proto
# 生成用户服务端代码 /root/code/gomall/app/user 直接使用 rpc_gen 下的代码 避免耦合
cwgo server --type RPC --service user --module github.com/LXJ0000/gomall/app/user --I ../../
idl --idl ../../idl/user.proto --pass "use github.com/LXJ0000/gomall/rpc_gen"

```