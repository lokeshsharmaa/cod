build:
	cd characterservice && $(MAKE) build && cd ..
	cd externalservice && $(MAKE) build && cd ..
	cd inventoryservice && $(MAKE) build && cd ..
	cd itemservice && $(MAKE) build && cd ..