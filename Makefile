VERSION ?= 1.0.0
NAME := shinthj.com

.PHONY: server
server:
	@hugo server -b http://localhost:1313/
.PHONY: build
build:
	@#rm -rf public/*
	@hugo -b https://shinthj.com/
relogin:
	firebase logout
	firebase login
	firebase use xackery
deploy: build
	@firebase deploy
set-version:
	@echo "VERSION=${VERSION}" >> $$GITHUB_ENV

.PHONY: scss
scss:
	@(echo "/* This file is autogenerated via 'make scss'. DO NOT EDIT */" > static/css/tpe-scss.css)
	docker run --rm -v $(PWD):$(PWD) -w $(PWD) jbergknoff/sass scss/app.scss >> static/css/tpe-scss.css

.PHONY: spells
spells:
	go run scripts/spells/main.go

.PHONY: bis
bis:
	go run scripts/bis/main.go 0 > bis/0.txt
	go run scripts/bis/main.go 1 > bis/1.txt
	go run scripts/bis/main.go 2 > bis/2.txt
	go run scripts/bis/main.go 3 > bis/3.txt
	go run scripts/bis/main.go 4 > bis/4.txt
	go run scripts/bis/main.go 5 > bis/5.txt
	go run scripts/bis/main.go 0 1 > bis/0-1.txt
	go run scripts/bis/main.go 0 2 > bis/0-2.txt
	go run scripts/bis/main.go 0 3 > bis/0-3.txt
	go run scripts/bis/main.go 0 4 > bis/0-4.txt
	go run scripts/bis/main.go 0 5 > bis/0-5.txt


.PHONY: proc
proc:
	go run scripts/proc/main.go 0 > proc/0.txt
	go run scripts/proc/main.go 1 > proc/1.txt
	go run scripts/proc/main.go 2 > proc/2.txt
	go run scripts/proc/main.go 3 > proc/3.txt
	go run scripts/proc/main.go 4 > proc/4.txt
	go run scripts/proc/main.go 5 > proc/5.txt
	go run scripts/proc/main.go 5 > proc/6.txt
	go run scripts/proc/main.go 5 > proc/7.txt

aa:
	go run scripts/aas/main.go "all" > content/aa/_index.en.md

sql:
	cd scripts/database && sqlc generate