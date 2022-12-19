FROM public.ecr.aws/eks-distro-build-tooling/golang:1.19-gcc as builder

RUN yum install -y zlib zlib-devel

WORKDIR /app
RUN go mod init hello
RUN go get github.com/4kills/go-libdeflate/v2
COPY main.go .
RUN go build -o hello main.go
RUN objdump -x hello | grep NEEDED
FROM public.ecr.aws/eks-distro-build-tooling/eks-distro-minimal-base-glibc-builder as package-installs

RUN clean_install zlib

FROM public.ecr.aws/eks-distro-build-tooling/eks-distro-minimal-base-glibc as final

COPY --from=package-installs /newroot /
COPY --from=builder /app/hello ./

ENTRYPOINT [ "/hello" ]
