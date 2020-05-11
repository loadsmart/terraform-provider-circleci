mod-download:
	go mod download

build:
	go build

install:
	mv terraform-provider-circleci ~/.terraform.d/plugins
