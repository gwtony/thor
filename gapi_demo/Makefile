all: main

main:
	@mkdir -p dist/bin dist/conf dist/log
	@echo "Release version"
	@${GOROOT}/bin/go build -v
	@mv gapi_demo dist/bin/gapi_demo
	@cp conf/gapi_demo.conf dist/conf/
	@echo "Build done: binary in dist dir"
debug:
	@mkdir -p dist/bin dist/conf dist/log
	@echo "Debug version"
	@${GOROOT}/bin/go build -o dist/bin/gapi_demo -ldflags '-l -N' main.go
	@cp conf/gapi_demo.conf dist/conf/
	@echo "Build done: binary in dist dir"

#test:
#	@sh -c "'$(CURDIR)/scripts/test.sh'"
#cover:
#	@sh -c "'$(CURDIR)/scripts/test.sh' cover"


clean:
	@rm -rf dist gapi_demo

.PHONY: all main clean debug
