makefile_dir		:= $(abspath $(shell pwd))

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

build:
	go build

watch:
	#CompileDaemon -exclude-dir=vendor -command='./devfile brewfile read --help'
	CompileDaemon -exclude-dir=vendor -command='./devfile brewfile write --help'
	#CompileDaemon -exclude-dir=vendor -command='./devfile brewfile -p test-output/Brewfile -f json -d'
	#CompileDaemon -exclude-dir=vendor -command='./devfile brewfile -p test-output/Brewfile -f json'
	#CompileDaemon -exclude-dir=vendor -command='./devfile brewfile -p test-output/Brewfile -f yaml'
