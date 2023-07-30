ifdef STEAMPIPE_INSTALL_DIR
STEAMPIPE_DIR := $(STEAMPIPE_INSTALL_DIR)
else
STEAMPIPE_DIR := ~/.steampipe
endif

install:
	go build -o  ${STEAMPIPE_DIR}/plugins/hub.steampipe.io/plugins/francois2metz/gitguardian@latest/steampipe-plugin-gitguardian.plugin *.go
