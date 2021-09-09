package main

import (
	rest "hack.com/rest"
)

//go:generate go run /home/vivek/hack/yang/ygot/generator/generator.go -path=yang -output_file=structs/structs.go -package_name=demo -generate_fakeroot -fakeroot_name=device -compress_paths=true -shorten_enum_leaf_names -typedef_enum_with_defmod yangSimple/*

func main() {
	rest.ServeYang()
}
