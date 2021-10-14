up:
	LAMBDA_HANDLER="../../awslocal/main.zip" pulumilocal up -y

down:
	pulumilocal destroy -s localstack -y

logs:
	docker logs --tail 1000 -f awslocal
