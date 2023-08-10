# Basic Go makefile

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# exclude python from std builds
DIRS=`go list ./... | grep -v python`

all: build

build: 
	@echo "GO111MODULE = $(value GO111MODULE)"
	$(GOBUILD) -v $(DIRS)

test: 
	@echo "GO111MODULE = $(value GO111MODULE)"
	$(GOTEST) -v $(DIRS)

clean: 
	@echo "GO111MODULE = $(value GO111MODULE)"
	$(GOCLEAN) ./...

vet:
	@echo "GO111MODULE = $(value GO111MODULE)"
	$(GOCMD) vet $(DIRS)

tidy: export GO111MODULE = on
tidy:
	@echo "GO111MODULE = $(value GO111MODULE)"
	go mod tidy

old:
	@echo "GO111MODULE = $(value GO111MODULE)"
	go list -u -m all | grep '\['
	
mod-update: export GO111MODULE = on
mod-update:
	@echo "GO111MODULE = $(value GO111MODULE)"
	go get -u ./...
	go mod tidy

# gopath-update is for GOPATH to get most things updated.
# need to call it in a target executable directory
gopath-update: export GO111MODULE = off
gopath-update:
	@echo "GO111MODULE = $(value GO111MODULE)"
	cd cmd/pi; go get -u ./...

# NOTE: MUST update version number here prior to running 'make release'
VERS=v1.0.17
PACKAGE=mat32
GIT_COMMIT=`git rev-parse --short HEAD`
VERS_DATE=`date -u +%Y-%m-%d\ %H:%M`
VERS_FILE=version.go

release: export GO111MODULE = on
release:
	/bin/rm -f $(VERS_FILE)
	@echo "// WARNING: auto-generated by Makefile release target -- run 'make release' to update" > $(VERS_FILE)
	@echo "" >> $(VERS_FILE)
	@echo "package $(PACKAGE)" >> $(VERS_FILE)
	@echo "" >> $(VERS_FILE)
	@echo "const (" >> $(VERS_FILE)
	@echo "	Version = \"$(VERS)\"" >> $(VERS_FILE)
	@echo "	GitCommit = \"$(GIT_COMMIT)\" // the commit JUST BEFORE the release" >> $(VERS_FILE)
	@echo "	VersionDate = \"$(VERS_DATE)\" // UTC" >> $(VERS_FILE)
	@echo ")" >> $(VERS_FILE)
	@echo "" >> $(VERS_FILE)
	/bin/cat $(VERS_FILE)
	git commit -am "$(VERS) release -- $(VERS_FILE) updated"
	git tag -a $(VERS) -m "$(VERS) release"
	git push
	git push origin --tags

