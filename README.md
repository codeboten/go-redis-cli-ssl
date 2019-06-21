# go-redis-cli

Need a simple client to run `redis info` when connecting to redis over TLS? Then this repo's for you. No support for any additional commands at this point. Happy to implement more functionality if ever needed.

### Usage

```bash
git clone https://github.com/codeboten/go-redis-cli && go-redis-cli
make
./go-redis-cli --help
Usage of ./go-redis-cli:
  -h string
        Server hostname (default: 127.0.0.1) (default "127.0.0.1")
  -n int
        Database number (default: 0)
  -p int
        Server port (default: 6379) (default 6379)
```

