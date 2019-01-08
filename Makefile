SHELL := bash

#Environment - start
ifeq ($(origin GOPATH), undefined)
$(error GOPATH environment variable not set!)
endif

GO ?= go

ifndef ENV_MIN_COVERAGE
ENV_MIN_COVERAGE := 85
endif

ifndef ENV_BUILD_LINUX
ENV_BUILD_LINUX := 1
endif

ifndef ENV_BUILD_WINDOWS
ENV_BUILD_WINDOWS := 0
endif

#Environment - end
#----------------------------------

#----------------------------------
#Common variables - start

os := $(shell $(GO) env GOOS)
apps_folder = apps

out_test_path := out/test
out_build_path := out/build

ifdef PATH_TO_MAIN
	main_paths := $(PATH_TO_MAIN)
else
	main_paths := $(shell find ./$(apps_folder) -mindepth 1 -maxdepth 1 -type d 2> /dev/null)
endif

ifdef NAME
	binary_names := $(out_build_path)/$(NAME)
else
	binary_names := $(main_paths:./$(apps_folder)/%=$(out_build_path)/%)
endif

binaries :=

ifeq ($(os), linux)
	os_suffix =
	xargs = xargs -r
	linux_binaries = $(addsuffix -linux,$(binary_names))
	binaries += $(linux_binaries)
	ifeq ($(ENV_BUILD_WINDOWS), 1)
		windows_binaries = $(addsuffix -windows.exe,$(binary_names))
		binaries += $(windows_binaries)
	endif
endif
ifeq ($(os), windows)
	os_suffix = .exe
	xargs = xargs
	windows_binaries = $(addsuffix -windows.exe,$(binary_names))
	binaries += $(windows_binaries)
	ifeq ($(ENV_BUILD_LINUX), 1)
		linux_binaries = $(addsuffix -linux,$(binary_names))
		binaries += $(linux_binaries)
	endif
endif
ifeq ($(os), darwin)
	os_suffix =
	xargs = xargs
	darwin_binaries = $(addsuffix -darwin,$(binary_names))
	binaries += $(darwin_binaries)
	ifeq ($(ENV_BUILD_LINUX), 1)
		linux_binaries = $(addsuffix -linux,$(binary_names))
		binaries += $(linux_binaries)
	endif
	ifeq ($(ENV_BUILD_WINDOWS), 1)
		windows_binaries = $(addsuffix -windows.exe,$(binary_names))
		binaries += $(windows_binaries)
	endif
endif

sources := $(wildcard **/*.go)

top_packages := $(shell glide nv)
top_package_patterns = $(filter-out .,$(top_packages))
top_package_names = $(top_package_patterns:./%/...=%)
all_package_paths = $(shell GOPATH=$(GOPATH) $(GO) list ${top_package_patterns})

top_package_tests = $(top_package_names:%=test.%)
top_package_test_reports = $(top_package_names:%=$(out_test_path)/%.func.txt)
all_test_report = $(out_test_path)/all.func.txt
all_test_report_html = $(out_test_path)/all.func.html
top_root_package_test_reports = $(top_root_package_names:%=$(out_test_path)/%.func.txt)
minimum_coverage_percent := $(ENV_MIN_COVERAGE)

#Common variables - end
#----------------------------------


ifdef PATH_TO_MAIN
$(out_build_path)/%-windows.exe: $(sources)
	@echo "Building windows executable $@"
	@GOOS=windows GOARCH=amd64 $(GO) build -o $@ ./$(PATH_TO_MAIN_GO)/

$(out_build_path)/%-linux: $(sources)
	@echo "Building linux executable $@"
	@GOOS=linux GOARCH=amd64 $(GO) build -o $@ ./$(PATH_TO_MAIN_GO)/

$(out_build_path)/%-darwin: $(sources)
	@echo "Building darwin executable $@"
	@GOOS=darwin GOARCH=amd64 $(GO) build -o $@ ./$(PATH_TO_MAIN_GO)/

else
$(out_build_path)/%-windows.exe: $(sources)
	@echo "Building windows executable $@"
	@GOOS=windows GOARCH=amd64 $(GO) build -o $@ ./$(apps_folder)/$*

$(out_build_path)/%-linux: $(sources)
	@echo "Building linux executable $@"
	@GOOS=linux GOARCH=amd64 $(GO) build -o $@ ./$(apps_folder)/$*

$(out_build_path)/%-darwin: $(sources)
	@echo "Building darwin executable $@"
	@GOOS=darwin GOARCH=amd64 $(GO) build -o $@ ./$(apps_folder)/$*

endif

$(out_test_path)/%.func.txt: $(sources)
	@mkdir -p $(@D)/$*
	@echo "Testing top level package $*";
	@result=$$(find $*  -name '$*.go' -print0 | $(xargs) -0 -n 1 dirname | sort | uniq 2>&1); \
	for package in $$result; do \
		echo "testing package $$package"; \
		profilefile=$(out_test_path)/$$package; mkdir -p `dirname $$profilefile`; \
		$(GO) test -covermode="count" -coverprofile="$(out_test_path)/$$package.func.txt" ./$$package; \
	done;
	@find $(@D)/$* -name '*.func.txt' | $(xargs) cat >> $@
	@rm -rf $(@D)/$*

$(out_test_path)/root.func.txt: $(sources)
	@mkdir -p $(@D)
	@echo "Testing outermost top level package $*"
	@$(GO) test -covermode="count" -coverprofile="$@" ./

$(all_test_report): $(top_root_package_test_reports) $(top_package_test_reports)
	@echo "Creating consolidated coverage file $@"
	@{ echo "mode: count"; \
		for file in $^; do \
			if [[ -f $$file ]]; then \
				cat $$file | \
				sed '/^mode.*count$$/d' | sed '/^$$/d' | sed 's/\r$$/$$/' ; \
			fi \
		done ; } > $@.tmp
	@mv $@.tmp $@

$(all_test_report_html): $(all_test_report)
	@$(GO) tool cover --html $< -o $@


.PHONY: info
info:
	@echo "Build information"
	@$(GO) version
	@echo "OS       :: $(os)"
	@echo "GOPATH   :: $(GOPATH)"

.PHONY: compile
compile: $(binaries)

.PHONY: clean.compile
clean.compile:
	@echo "Removing binaries.. if any"
	@rm -rf $(out_build_path)

.PHONY: clean.test
clean.test:
	@echo "Removing test reports.. if any"
	@rm -rf $(out_test_path)

.PHONY: clean
clean: clean.compile clean.test

.PHONY: fmt
fmt:
	@echo "Checking formatting of go sources"
	@result=$$(gofmt -d -l -e $(top_package_names) 2>&1); \
		if [[ "$$result" ]]; then \
			echo "$$result"; \
			echo 'gofmt failed!'; \
			exit 1; \
		fi

.PHONY: fixfmt
fixfmt:
	@echo "Fixing format of go sources"
	@gofmt -w -l -e $(top_package_names) 2>&1; \
			if [[ "$$?" != 0 ]]; then \
				echo "gofmt failed! (exit-code: '$$?')"; \
				exit 1; \
			fi

.PHONY: vet
vet:
	@echo "Running go vet"
	@$(GO) vet $(all_package_paths)

.PHONY: test.setup

.PHONY: test.teardown

.PHONY: test
test: coverage = $(shell GOPATH=$(GOPATH) $(GO) tool cover --func=out/test/all.func.txt | tail -1 | awk '{ print int($$3) }' | sed 's/%$$//')
test: test.setup $(all_test_report) $(all_test_report) test.teardown
	$(info Total Coverage = $(coverage)%)
	@if [[ $(coverage) -lt $(minimum_coverage_percent) ]]; then \
		echo "Coverage ${coverage} is below $(minimum_coverage_percent)%! Failing build." ;\
		exit 1; \
	fi

.PHONY: check
check: fmt vet test

.PHONY: build
build: info clean compile check copy.resources
