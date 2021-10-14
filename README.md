# Pulumi - Deploy in AWS LocalStack

<h1 align="center">
    <img src="https://www.pulumi.com/logos/brand/avatar-on-white.svg" width=200/>
    <img src="https://avatars.githubusercontent.com/u/28732122?s=280&v=4" width=200/>
</h1>

How to deploy using `pulumi` in an AWS environment running on your machine with `localstack`?

### Dependencies
- [Docker](https://www.docker.com/get-started)
- [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2-linux.html)
- [Pulumi](https://www.pulumi.com/docs/get-started/aws/begin/)

- **Wrappers**
  - [AWS Local CLI](https://github.com/localstack/awscli-local)
  - [Pulumi Local CLI](https://github.com/localstack/pulumi-local)

### Install
```bash
make up
```

### Remove Infra
```bash
make down
```

### Local Stack Logs
```bash
docker logs --tail 1000 -f awslocal
```
