.PHONY: models generate

# ==============================================================================
# Swagger Models
models:
	$(call print-target)
	find ./pkg/api/model -type f -not -name '*_test.go' -delete
	swagger generate model -f pkg/docs/e-commerce.yml -m pkg/api/model

generate: models