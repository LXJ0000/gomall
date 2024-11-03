export CMD_ROOT=github.com/LXJ0000/gomall
.PHONY: gen-gateway
gen-gateway:
	@cd app/gateway && cwgo server --type HTTP -I ../../idl/ --idl  ../../idl/gateway/ping.proto --service gateway --module ${CMD_ROOT}/app/gateway
	@cd app/gateway && cwgo server --type HTTP -I ../../idl/ --idl  ../../idl/gateway/auth.proto --service gateway --module ${CMD_ROOT}/app/gateway

.PHONY: gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC --service user --module ${CMD_ROOT}/rpc_gen --I ../idl --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC --service user --module ${CMD_ROOT}/app/user --I ../../idl --idl ../../idl/user.proto --pass "use ${CMD_ROOT}/rpc_gen"