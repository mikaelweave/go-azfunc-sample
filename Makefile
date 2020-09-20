define PROJECT_HELP_MSG
Usage:
    make help               show this message
    make build              build the golang components of the function
    make run                run the solution locally
endef
export PROJECT_HELP_MSG

help:
	@echo "$$PROJECT_HELP_MSG" | less

build:
	bash ./scripts/build.sh

run:
	bash ./scripts/build.sh
	bash ./scripts/run-function.sh