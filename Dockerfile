# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.8-onbuild
MAINTAINER edinkulovic@gmail.com

EXPOSE 8000