.PHONY: build

build:
	sam build

local: build
	sam local start-api --skip-pull-image -p 9090 --debug
