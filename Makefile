DIST_DIR=dist
COMPILED_PACKAGE=raindrop-alfred-workflow
DIST_FILE=RaindropWorkflow.alfredworkflow

.DEFAULT_GOAL := package

package: clean build zip

build:
	go build -o ./dist/$(COMPILED_PACKAGE)  ./cmd/$(COMPILED_PACKAGE)/main.go

zip:
	zip -j $(DIST_FILE) $(CURDIR)/$(DIST_DIR)/*

.PHONY: clean
clean:
	rm -f $(DIST_FILE)
	rm -f $(CURDIR)/$(DIST_DIR)/$(COMPILED_PACKAGE)
