.PHONY: build

exec-dev:
	sam build; sam local invoke PeriodTweetTrigger \
		--skip-pull-image \
		--env-vars env.json

deploy:
	sam build; sh ./deploy.sh