.PHONY: build

build:
	GOOS=linux GOARCH=amd64 go build -o hello-world/hello-world ./PeriodTweetTrigger

exec-dev:
	sam build; sam local invoke PeriodTweetTrigger