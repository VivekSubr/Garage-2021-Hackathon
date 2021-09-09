checkYGOT:
	if test "$(YGOT)" = "" ; then \
        echo "ENV not set"; \
        exit 1; \
    fi

generate: checkYGOT 
	mkdir -p structs
	go run $(YGOT)/generator/generator.go -path=yang -output_file=structs/structs.go -package_name=demo -generate_fakeroot -fakeroot_name=device -compress_paths=true -shorten_enum_leaf_names -typedef_enum_with_defmod yang/*

build: 
	go build

clean:
	go clean