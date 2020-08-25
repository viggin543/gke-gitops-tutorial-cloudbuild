FROM golang
WORKDIR /app
COPY gke-gitops-tutorial-cloudbuild /app/app
CMD ["/app/app"]
