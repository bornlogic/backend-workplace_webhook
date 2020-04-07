name = wiw
img = golang:1.14.1
src = github.com/bornlogic/$(name)
workdir = /go/src/$(src)
run = docker run --net=host -it -v $(PWD):$(workdir) -w $(workdir) \
	-e WORKPLACE_ACCESS_TOKEN \
	-e WORKPLACE_GROUP_ID_TEST \
	--rm $(img)
prefix ?= /usr/local/
binpath ?= $(prefix)/bin

shell:
	$(run) bash

clean-%:
	rm -f ${*}

clean: clean-wiw clean-wSendToGroup

uninstall-%:
	rm -f $(binpath)/${*}

uninstall: uninstall-wiw uninstall-wSendToGroup

build-%:
	$(run) go build -o ${*} cmd/${*}/main.go

build: clean build-wiw build-wSendToGroup

install-%:
	mkdir -p $(binpath)
	mv ${*} $(binpath)

install: build uninstall install-wiw install-wSendToGroup

serve:
	$(run) go run cmd/wiw/main.go $(args)

test:
ifeq ($(origin args), undefined)
	$(run) go test ./...
else
	$(run) go test $(args)
endif
