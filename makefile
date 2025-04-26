GITLAB_NAME=ws
GITLAB_PROJECT=41089011
GITWS_BINARY=ws-cmd
GITWS_VERSION=master

build:
	go build -o ${GITWS_BINARY} main.go
 
run:
	go build -o ${GITWS_BINARY} main.go
	./${GITWS_BINARY}
 
clean:
	go clean
	rm ${GITWS_BINARY}

release: build
	@curl --location --header "PRIVATE-TOKEN: ${GITLAB_TOKEN}" \
		--upload-file ${GITWS_BINARY} \
		"https://gitlab.com/api/v4/projects/${GITLAB_PROJECT}/packages/generic/${GITLAB_NAME}/${GITWS_VERSION}/${GITWS_BINARY}"
