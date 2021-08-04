FROM centos:7

COPY cmd/orderlytask/orderlytask /orderlytask

RUN chmod +x /orderlytask

CMD ["/orderlytask", "--master=''", "--kubeconfig='/config'"]
