export CMD_ROOT=github.com/LXJ0000/gomall
.PHONY: consul
consul:
	@consul agent -dev -client=0.0.0.0

.PHONY: gen-gateway
gen-gateway:
	@cd app/gateway && cwgo server --type HTTP -I ../../idl/ --idl  ../../idl/gateway/common.proto --service gateway --module ${CMD_ROOT}/app/gateway
	@cd app/gateway && cwgo server --type HTTP -I ../../idl/ --idl  ../../idl/gateway/ping.proto --service gateway --module ${CMD_ROOT}/app/gateway
	@cd app/gateway && cwgo server --type HTTP -I ../../idl/ --idl  ../../idl/gateway/auth.proto --service gateway --module ${CMD_ROOT}/app/gateway

.PHONY: gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC --service user --module ${CMD_ROOT}/rpc_gen --I ../idl --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC --service user --module ${CMD_ROOT}/app/user --I ../../idl --idl ../../idl/user.proto --pass "use ${CMD_ROOT}/rpc_gen"

.PHONY: gen-product
gen-product:
	@cd rpc_gen && cwgo client --type RPC --service product --module ${CMD_ROOT}/rpc_gen --I ../idl --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --service product --module ${CMD_ROOT}/app/product --I ../../idl --idl ../../idl/product.proto --pass "use ${CMD_ROOT}/rpc_gen"

.PHONY: gen-cart
gen-cart:
	@cd rpc_gen && cwgo client --type RPC --service cart --module ${CMD_ROOT}/rpc_gen --I ../idl --idl ../idl/cart.proto
	@cd app/cart && cwgo server --type RPC --service cart --module ${CMD_ROOT}/app/cart --I ../../idl --idl ../../idl/cart.proto --pass "use ${CMD_ROOT}/rpc_gen"

.PHONY: gen-order
gen-order:
	@cd rpc_gen && cwgo client --type RPC --service order --module ${CMD_ROOT}/rpc_gen --I ../idl --idl ../idl/order.proto
	@cd app/order && cwgo server --type RPC --service order --module ${CMD_ROOT}/app/order --I ../../idl --idl ../../idl/order.proto --pass "use ${CMD_ROOT}/rpc_gen"

.PHONY: gen-payment
gen-payment:
	@cd rpc_gen && cwgo client --type RPC --service payment --module ${CMD_ROOT}/rpc_gen --I ../idl --idl ../idl/payment.proto
	@cd app/payment && cwgo server --type RPC --service payment --module ${CMD_ROOT}/app/payment --I ../../idl --idl ../../idl/payment.proto --pass "use ${CMD_ROOT}/rpc_gen"

.PHONY: gen-checkout
gen-checkout:
	@cd rpc_gen && cwgo client --type RPC --service checkout --module ${CMD_ROOT}/rpc_gen --I ../idl --idl ../idl/checkout.proto
	@cd app/checkout && cwgo server --type RPC --service checkout --module ${CMD_ROOT}/app/checkout --I ../../idl --idl ../../idl/checkout.proto --pass "use ${CMD_ROOT}/rpc_gen"

.PHONY: gen-email
gen-email:
	@cd rpc_gen && cwgo client --type RPC --service email --module ${CMD_ROOT}/rpc_gen --I ../idl --idl ../idl/email.proto
	@cd app/email && cwgo server --type RPC --service email --module ${CMD_ROOT}/app/email --I ../../idl --idl ../../idl/email.proto --pass "use ${CMD_ROOT}/rpc_gen"

.PHONY: gen-code
gen-code:
	@cd rpc_gen && cwgo client --type RPC --service code --module ${CMD_ROOT}/rpc_gen --I ../idl --idl ../idl/code.proto
	@cd app/code && cwgo server --type RPC --service code --module ${CMD_ROOT}/app/code --I ../../idl --idl ../../idl/code.proto --pass "use ${CMD_ROOT}/rpc_gen"
