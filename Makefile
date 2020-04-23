build: gendates.go
	packr2 build

run: build *.adoc.tpl c311.toml
	./semestergen c311.toml

clean:
	@rm -rf packrd main-packr.go semestergen

.PHONY: clean
