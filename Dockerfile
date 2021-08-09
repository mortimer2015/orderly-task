FROM busybox:1.33.1

COPY cmd/orderlytask/orderlytask /orderlytask

CMD ["/orderlytask", "--master=", "--kubeconfig=/conf/admin.conf"]