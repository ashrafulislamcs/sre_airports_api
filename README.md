# Airport API

<!-- My thought process and decisions goes here -->
---
Solution Overview
--
Steps Taken:
-
Provisioning Cloud Storage: Created an AWS S3 bucket to store airport images using Terraform.

Make Endpoint: the /update_airport_image endpoint in Go to handle image uploads.
Manual Testing: Initially ran the application manually and resolved errors to ensure it was functioning correctly.
Containerization: Dockerized the Go application for easier deployment.
Kubernetes Deployment: Created a simplest deployment and service manifest files to run the application in a Kubernetes cluster.
CI/CD Pipeline: Configured a simplest CI/CD workflow using GitHub Actions to automate the build and deployment process but due to the unavailability of a Kubernetes cluster, I was unable to test the CI/CD workflow.

---
_For tasks, checkout [tasks.md](tasks.md)_
