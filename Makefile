SHELL := /bin/sh

.PHONY: classifier-test classifier-run dashboard-install dashboard-build sdr-configure

classifier-test:
	cd signal_classifier && go test ./...

classifier-run:
	cd signal_classifier && go run ./cmd/classifier

dashboard-install:
	cd tactical_dashboard && npm install

dashboard-build:
	cd tactical_dashboard && npm run build

sdr-configure:
	cmake -S sdr_ingestion -B build/sdr_ingestion
