# Go Project for Hosting WA Websuite

The project is primarily for learning purposes. I restarted from the beginning to incorporate Clean Code advice from the book *Clean Code* by Robert C. Martin into it.

The goal is to quickly set up a working website and then expand on it over time. This means I plan to create the code quickly, then perform multiple refactorings. The approach is to avoid letting the project die in perfection.

Currently, the website is built with Docker and then deployed to an Ubuntu VM.

## Used Commands

1. **Build Docker Image**
    ```bash
    docker build --tag lukasloetscher/lulo_tryout .
    ```
    -> Then use Docker Desktop to push it.

2. **Open CMD on Ubuntu VM**
    - Stop the container:
        ```bash
        docker ps
        docker stop {name}
        ```

3. **Update the Image**
    - Pull the updated image:
        ```bash
        docker pull lukasloetscher/lulo_tryout
        ```
    - Check with:
        ```bash
        docker images
        ```

4. **Start the Docker Container**
    - Start the container (you may need to change ports):
        ```bash
        docker run -d -p 80:8080 lukasloetscher/lulo_tryout
        ```
    - Check with:
        ```bash
        docker ps
        ```
    - Check by visiting the webpage.

## My Notes
- [DigitalOcean: How to Install and Use Docker on Ubuntu 20.04](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-20-04)
