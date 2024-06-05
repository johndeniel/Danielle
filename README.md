# Danielle Perfume

Danielle Perfume is a web application built with Go, designed to help you discover the perfect scent for any occasion. Explore our range and find your signature scent today. Danielle has been Dockerized for effortless deployment, and its Docker image is readily available on [Docker Hub](https://hub.docker.com/repository/docker/johndeniel/danielle/tags).

To run Danielle using Docker, execute this command in your terminal:

```bash
docker run -d -it --name danielle -p 7860:7860 johndeniel/danielle:build-v1.001
```

## Features

- **Go Integration**: Danielle utilizes Go, a modern programming language known for its speed and efficiency, to power its backend functionality.

- **Automated Dockerization**: Danielle includes a GitHub Action workflow for Dockerizing the Go application, streamlining the deployment process with automated containerization.

- **Hugging Face Platform Integration**: Danielle seamlessly integrates with the Hugging Face Platform, offering access to the live instance on [Hugging Face Spaces](https://huggingface.co/spaces/johndeniel/Danielle).

---

Danielle demonstrates the utilization of cutting-edge technologies like Go, Docker, GitHub Actions, and the Hugging Face Platform, offering developers an efficient development experience.