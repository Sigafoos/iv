compile: main gener8r

main:
	GOOS=linux GOARCH=arm GOARM=5 go build

gener8r:
	cd generator && GOOS=linux GOARCH=arm GOARM=5 go build
