# makefile

test_srv:
	go run ./cli/main.go

run_consul:
	consul agent -dev

run_srv:
	go run ./srv/main.go

run_api:
	go run ./api/api.go

list_services:
	micro --registry_address=127.0.0.1:8500 list services

run_api_gw:
	micro api --handler=api

test_api:
	curl http://localhost:8080/puffin/ipdb/IpDb/GetIpInfo?Ip=8.8.8.8





