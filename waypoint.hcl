project = "mark2web-service"

app "mark2web-service" {
    build {
        use "pack" {}
    }

    # Deploy to Docker
    deploy {
        use "docker" {}
    }
}
