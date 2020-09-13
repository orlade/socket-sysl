all: build

build:
	sysl pb --mode=pb specs/example.sysl > ../codegen/model.pb \
	&& cd ../codegen \
	&& arrai run --out=dir:../example auto.arrai \
	&& cd ../example \
	&& gofmt -w . \
	&& goimports -w .
