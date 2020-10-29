export SERVICE_NAME = $(shell basename `pwd`)
export SERVICE_TYPE = $(shell basename `cd .. && pwd`)
export SERVICE_LANG = go

all:
	$(MAKE) -C ../.. deploy

%:
	$(MAKE) -C ../.. $*
