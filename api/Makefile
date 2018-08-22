##### Convenient command ######

REPO:=github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example
GAE_PROJECT ?= projectName

SERVER_BASE=$(REPO)/api

LOCAL_PORT ?= 8080

dep_init:
	@dep init

dep_ensure:
	@dep ensure

goa_bootstrap:
	@mv vendor vendor.bak
	@goagen bootstrap -d $(SERVER_BASE)/design ; mv vendor.bak vendor

init: install bootstrap import
gen: clean generate import

# goagen douments
# https://goa.design/implement/goagen/
# https://goa.design/ja/implement/goagen/

# Rename vendor during executing goagen
#	https://github.com/goadesign/goa/issues/923#issuecomment-290424097
bootstrap: generate main

main: controller appengine/main.go

appengine/main.go:
	@mv vendor vendor.bak
	@goagen main -d $(SERVER_BASE)/design >/dev/null ; mv vendor.bak vendor
	@mkdir -p appengine
	@mv main.go appengine
	@rm *.go
	@echo 'appengine/main.go'
	@echo '1. Change package from "main" to "appengine"'
	@echo '2. Add "net/http" to import section'
	@echo '3. Add "$(SERVER_BASE)/controller" to import section'
	@echo '4. Change "func main()" to "func init()"'
	@echo '5. Add "controller." before each "NewXxxxController"'
	@echo '6. Comment out the lines below the comment "Start service"'
	@echo '7. Add http.HandleFunc("/", service.Mux.ServeHTTP) at the end of init func'

app:
	@mv vendor vendor.bak
	@goagen app -d $(SERVER_BASE)/design ; mv vendor.bak vendor

controller: goa_controller converter

goa_controller:
	@mv vendor vendor.bak
	@mkdir -p controller
	@goagen controller  -d $(SERVER_BASE)/design --pkg controller --out controller --app-pkg ../app ; mv vendor.bak vendor

converter: converter_gen
converter_gen:
	@goa_model_gen converter design/*.yaml

model: model_gen
model_gen:
	@mkdir -p ./model
	@goa_model_gen model design/*.yaml

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf tool
	@rm -rf swagger

goa_gen:
	@mv vendor vendor.bak
	@goagen app     -d $(SERVER_BASE)/design && \
	 goagen swagger -d $(SERVER_BASE)/design && \
	 goagen client  -d $(SERVER_BASE)/design ; mv vendor.bak vendor

generate: goa_gen model_gen converter_gen

swaggerui:
	rm -rf ./tmp/swaggerui && \
	mkdir -p ./tmp/swaggerui && \
	cd ./tmp/swaggerui && \
		git init && \
		git config core.sparsecheckout true && \
		git remote add origin git@github.com:swagger-api/swagger-ui.git && \
		echo "dist" > .git/info/sparse-checkout && \
		git pull origin master && \
		cd - && \
	mkdir -p swaggerui && \
	cp -R tmp/swaggerui/dist swaggerui/

install:
	@which dep || go get -u github.com/golang/dep/cmd/dep
	@dep ensure

import:
	@which gorep || go get -v github.com/novalagung/gorep
	@gorep -path="./" \
          -from="../app" \
          -to="$(REPO)/app"
	@gorep -path="./" \
          -from="../client" \
          -to="$(REPO)/client"
	@gorep -path="./" \
          -from="../tool/cli" \
          -to="$(REPO)/tool/cli"

fmt:
	@gofmt -w controller/*.go model/*.go

test:
	goapp test $(SERVER_BASE)/model

build:
	goapp build $(SERVER_BASE)/appengine

deploy:
	goapp deploy -application $(GAE_PROJECT) ./app

rollback:
	appcfg.py rollback ./app -A $(GAE_PROJECT)

local: swaggerui
	dev_appserver.py --port=$(LOCAL_PORT) --enable_console --skip_sdk_update_check=yes appengine/app.yaml
# goapp serve ./appengine

local_https: swaggerui
	dev_appserver.py --port=$(LOCAL_PORT) --enable_console --skip_sdk_update_check=yes appengine/app.yaml \
		--ssl_certificate_path $(SSL_CERTIFICATE_PATH) \
		--ssl_certificate_key_path $(SSL_CERTIFICATE_KEY_PATH)
# goapp serve ./appengine
