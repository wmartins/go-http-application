# Go HTTP Application

This is a simple Go HTTP Application, build just for test purposes.

## Building

There's a simple `Makefile` that abstracts all the build process. Just type:

```bash
make
```

After that, a Docker image (`wmartins/go-http-application:${VERSION}`) will be
built, and you can run that.

## Running using Docker

You can run the application using Docker just by typing:

```bash
docker run -p 80:80 wmartins/go-http-application:${VERSION}
```

Now, you can `curl` to check the application version:

```bash
curl http://localhost/version
```
