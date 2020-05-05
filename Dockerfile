FROM golang:1.14-buster as build
ENV GOPROXY=https://goproxy.cn
WORKDIR /opt
ADD . .
RUN go build -o app ./cmd/serv
RUN go get github.com/grpc-ecosystem/grpc-health-probe

FROM gcr.io/distroless/base-debian10
#FROM longkai/dockerfiles:distroless-base-debian10 #mirror gcr registry
COPY --from=build /opt/app /
COPY --from=build /go/bin/grpc-health-probe /bin
CMD ["/app"]

#this only works within a shell
#COPY --from=busybox /bin/busybox /busybox/busybox
#RUN ["/busybox/busybox", "--install", "/bin"]
#ENV HW=lhw
#CMD ["/app", "-hw=123", "-hw2=$HW"] # not work!
#CMD /app -hw=123 -hw2=$HW
