RUN = docker-compose run --rm pulumi_go_localstack

up:
	LAMBDA_HANDLER="../../awslocal/main.zip" $(RUN) pulumilocal up -y

down:
	$(RUN) pulumilocal destroy -s localstack -y

logs:
	docker logs --tail 1000 -f awslocal

console:
	docker-compose up -d
	docker-compose exec pulumi_go_localstack bash
