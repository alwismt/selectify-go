ADMIN_APP_NAME = selectifyAdmin
CUSTOMER_APP_NAME = selectifyCustomer
.DEFAULT_GOAL := help
APP ?= adminApp

domain:
ifndef d
	$(error DOMAIN is not set)
endif
	chmod +x ./scripts/createDomain.sh
	./scripts/createDomain.sh $(APP) $(d)

debug-admin: 
	export APP=admin && $(MAKE) rice
	chmod +x ./scripts/compileDaemon.sh
	./scripts/compileDaemon.sh $(ADMIN_APP_NAME)


run-admin: 
	export APP=admin && $(MAKE) rice
	env GOOS=linux CGO_ENABLED=0 go build -o ./bin/${ADMIN_APP_NAME} ./cmd/admin
	./bin/${ADMIN_APP_NAME}


debug-cus: 
	export APP=customer && $(MAKE) rice
	CompileDaemon -exclude=".env,.gitignore,*_test.go" -include="*.gotpl" -build="env GOOS=linux CGO_ENABLED=0 go build -o ./bin/${CUSTOMER_APP_NAME} ./cmd/customer" -command="./bin/${CUSTOMER_APP_NAME}"

rice:
	chmod +x ./scripts/rice.sh
	./scripts/rice.sh $(APP)

help:
	@echo "Usage:"
	@echo "  make domain APP=adminApp d=example"
	@echo "  make debug-admin	- For development"
	@echo ""
	@echo "Options:"
	@echo "  APP         App name (default: admin-app)"
	@echo "  d      Domain name (required)"