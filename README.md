# upstash-cli

## Installation

1. Download a binary from

## Using

### Authenticating

Before you can interact with the upstash api you need to authenticate.
Use your email and an api token from [https://console.upstash.com/account/api](https://console.upstash.com/account/api)

Unfortunately right now this is only possible for personal accounts. Teams will be supported at a later date.

```bash
upstash auth --email=YOUR_EMAIL --key=YOUR_API_KEY
```



### Creating a database

```bash
upstash redis create DATABASE_NAME --region="us-easy-1"
```

See `upstash redis create --help` for more options

### Get information about an existing databse

```bash
upstash redis get DATABASE_ID
```

### Delete an existing databse

```bash
upstash redis delete DATABASE_ID
```
