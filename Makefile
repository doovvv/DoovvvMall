.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo\demo_protobuf && cwgo server -I ../../idl --type RPC --module demo_proto --service demo_proto --idl ../../idl/echo.proto
.PHONY: gen-frontend
gen-fronted:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto --service frontend --module gomall/app/frontend -I ../../idl
.PHONY: gen-user
gen-user:
	cwgo client --type RPC --module gomall/rpc_gen --service user -I ../idl --idl ../idl/user.proto
	cwgo server --type RPC --module gomall/app/user --service user -I ../../idl --idl ../../idl/user.proto --pass "-use gomall/rpc_gen/kitex_gen"