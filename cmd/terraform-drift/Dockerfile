FROM alpine AS builder

RUN apk update && apk add --no-cache git

ENV TERRAFORM 0.13.5
RUN wget https://releases.hashicorp.com/terraform/${TERRAFORM}/terraform_${TERRAFORM}_linux_amd64.zip && \
  unzip terraform_${TERRAFORM}_linux_amd64.zip && \
  chmod +x terraform && mv terraform /usr/bin/terraform && rm terraform_${TERRAFORM}_linux_amd64.zip

COPY build/linux/terraform-drift /
ENTRYPOINT ["/terraform-drift"]
