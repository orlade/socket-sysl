all: build

build:
	cd ../codegen \
	&& arrai run --out=dir:../example auto.arrai \
	&& cd ../example \
	&& gofmt -w . \
	&& goimports -w .
