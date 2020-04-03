name = wiw
img = golang:latest
src = github.com/bornlogic/$(name)
workdir = /go/src/$(src)
run = docker run --net=host -it -v $(PWD):$(workdir) -w $(workdir) \
	-e WORKPLACE_ACCESS_TOKEN=$(WORKPLACE_ACCESS_TOKEN) \
	-e WORKPLACE_GROUP_ID_TEST=$(WORKPLACE_GROUP_ID_TEST) \
	--rm $(img)
bin ?= /usr/local/bin

shell:
	$(run) bash

clean-%:
	rm -f ${*}

clean: clean-wiw clean-wSendToGroup

uninstall-%:
	rm -f $(bin)/${*}

uninstall: uninstall-wiw uninstall-wSendToGroup

build-%:
	$(run) go build -o ${*} cmd/${*}/main.go

build: clean build-wiw build-wSendToGroup

install-%:
	cp ${*} $(bin)

install: build uninstall install-wiw install-wSendToGroup

serve:
	$(run) go run cmd/wiw/main.go $(args)

test:
ifeq ($(origin args), undefined)
	$(run) go test ./...
else
	$(run) go test $(args)
endif
