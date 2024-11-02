```bash


cwgo server --server_name demo_thrift --type RPC  --idl ../../idl/echo.thrift -module github.com/LXJ0000/gomall/demo/demo_thrift

cwgo server -I ../../idl --server_name demo_proto --type RPC  --idl ../../idl/echo.proto -module github.com/LXJ0000/gomall/demo/demo_proto

```