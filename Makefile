name = wiw
img = golang:latest
src = github.com/bornlogic/$(name)
workdir = /go/src/$(src)
run = docker run --net=host -it -v $(PWD):$(workdir) -w $(workdir) \
	-e WORKPLACE_ACCESS_TOKEN=$(WORKPLACE_ACCESS_TOKEN) \
	-e WORKPLACE_GROUP_ID_TEST=$(WORKPLACE_GROUP_ID_TEST) \
	--rm $(img)

shell:
	$(run) bash

serve:
	$(run) go run cmd/server/main.go $(args)

test:
ifeq ($(origin args), undefined)
	$(run) go test ./...
else
	$(run) go test $(args)
endif
