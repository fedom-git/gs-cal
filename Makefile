
export GO111MODULE=on

cal:
	go build -o ./gs-cal main.go

clean:
	rm cal