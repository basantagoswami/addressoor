<a href="https://nbd.wtf"><img align="right" height="50" src="https://i.imgur.com/EXcpaWc.png" /></a>

# addressoor

Lightning & Nostr Address Server

## How to run

1. Compile with `go build`
2. Set the following environment variables somehow (for ex: using `export VAR=VALUE`):

```
PORT=5555
DOMAIN=example.com
SECRET=askdbasjdhvakjvsdjasd
SITE_OWNER_URL=https://twitter.com/basanta_goswami
SITE_OWNER_NAME=@basantagowami
SITE_NAME=Addressoor
```

3. Start the app with `./satdress`
4. Serve the app on your own domain