FROM alpine:3

RUN apk update && apk add --no-cache git

ENV TERRAFORM 0.13.5
RUN wget https://releases.hashicorp.com/terraform/${TERRAFORM}/terraform_${TERRAFORM}_linux_amd64.zip && \
  unzip terraform_${TERRAFORM}_linux_amd64.zip && \
  chmod +x terraform && mv terraform /usr/bin/terraform && rm terraform_${TERRAFORM}_linux_amd64.zip

ENTRYPOINT ["/kh-tf-drift"]

COPY ./build/linux/kh-tf-drift /kh-tf-drift