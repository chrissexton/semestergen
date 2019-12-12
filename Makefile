run: build *.adoc.tpl c311.toml
	./semestergen c311.toml

build: gendates.go
	go build
