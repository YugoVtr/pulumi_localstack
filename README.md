# Pulumi - Deploy in AWS LocalStack

<h1 align="center">
    <img src="https://www.pulumi.com/logos/brand/avatar-on-white.svg" width=200/>
    <img src="https://avatars.githubusercontent.com/u/28732122?s=280&v=4" width=200/>
</h1>

How to deploy using `pulumi` in an AWS environment running on your machine with `localstack`?

### Dependencies
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install)

## Running the container Pulumi

1. First we need to configure the env `Pulumi_Access_Token` in the` .env` file. For this execute:

```shell
$ cp .env.sample .env
```

2. Open the file, fill in the value of the env `pulumi_access_token` with its respective value to generate your token, follow this [document](https://www.pulumi.com/docs/intro/console/accounts-and-organizations/accounts/#access-tokens)

3. To build and run the container that will rotate the pulumi, run:

```shell
$ docker-compose up -d
$ make up
```

4. To access the container execute:

```shell
$ make console
```

5. To accompany the logs:

```shell
$ make logs
```
