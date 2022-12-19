FROM public.ecr.aws/eks-distro-build-tooling/golang:1.19-gcc as builder

WORKDIR /app
RUN go mod init hello
COPY main.go .
RUN go build -o hello main.go

FROM public.ecr.aws/eks-distro-build-tooling/eks-distro-minimal-base-glibc as final

COPY --from=builder /app/hello ./

ENTRYPOINT [ "/hello" ]
