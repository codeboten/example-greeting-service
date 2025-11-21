
.PHONY: all
all:
	cd golang/year-service && go build ./...

.PHONY: run
run:
	cd golang/year-service && \
		OTEL_EXPORTER_OTLP_ENDPOINT=localhost:4317 HONEYCOMB_API_KEY=none OTEL_EXPERIMENTAL_CONFIG_FILE=../../otelconf.yaml go test ./...
