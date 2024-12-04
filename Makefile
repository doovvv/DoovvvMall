.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo\demo_protobuf && cwgo server -I ../../idl --type RPC --module demo_proto --service demo_proto --idl ../../idl/echo.proto
gen-fronted:
	cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend -module gomall/app/frontend -I ../../idl