# AWS S3 - POC

When adding the `otelaws.AppendMiddlewares(&cfg.APIOptions)`, all pre-signed URLs for downloading S3 files start to break with code `SignatureDoesNotMatch` and the message:

> "The request signature we calculated does not match the signature you provided. Check your key and signing method."

![image](https://user-images.githubusercontent.com/33238105/217125791-14a45626-7937-4fae-a1a3-f6f4997bc7f7.png)

When removing `otelaws.AppendMiddlewares(&cfg.APIOptions)` they immediately return to work.

I'm attaching an image from a diff window with the generated URLs. The only difference that is not related to the AWS account nor the link expiration time is the `X-Amz-SignedHeaders`. The URL generated from code with **otelaws middlewares** has two values `host` and `traceparent`, while the one without **otelaws middlewares** has only `host`.

![image](https://user-images.githubusercontent.com/33238105/217124120-f12c5afa-1625-49af-bf42-03bc38be119e.png)

## Usage

- Configure **default** credentials at the  `~/.aws/credentials` file
- Start the **OTEL collector** and **Jaeger** containers

```bash
  docker compose up -d
```

- Export the environment variable `BUCKET_NAME` with an existent AWS S3 Bucket name:

```bash
  export BUCKET_NAME=<your_bucket>
```

- Start the application:

```bash
  go run main.go
```

- Open this URL in the browser:

> http://localhost:8080/swagger/index.html

- Call the `/s3/files/{name}` endpoint with a valid S3 file within the your bucket.

- Try to download the file with the generated URL, **this URL will work**.

- Call the `/s3/files/otel/{name}` endpoint with a valid S3 file within the your bucket.

- Try to download the file with the generated URL, **this URL will not work**.

- Open Jaeger UI:

> http://localhost:16686/

- The trace for path `/s3/files/{name}` should have one span, while the trace for `/s3/files/otel/{name}` should have two spans.
