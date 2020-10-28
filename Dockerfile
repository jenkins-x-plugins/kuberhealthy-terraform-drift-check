FROM scratch

ENTRYPOINT ["/jx-terraform-drift-check"]

COPY ./build/linux/jx-terraform-drift-check /jx-terraform-drift-check