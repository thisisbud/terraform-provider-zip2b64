resource "http-b64" "provisioningzipfile" {
  provider = http2b64
  url      = "http://localhost/secure-connect-bud-josh-develop-default-ade.zip"
}


resource "zip2b64" "extractedfile" {
  provider   = zip2b64
  base64file = http-b64.provisioningzipfile.response_body_base64
  filename   = "/ca.crt"
}
