# Tech Curriculum BGP

## Local Development

### 1. Install Docker

To install docker, simply just install the docker desktop app. It will install everything you need. You can download from [here](https://www.docker.com/products/docker-desktop/)

### 2. Run Docker Containers

We've already create a makefile command to compose up the docker container and also migrate the db migration using soda.Run `make composeup` to compose the docker and use it for your development.

### 3. Run Migration and Seeding

We've already create a migrate command to run db migration and seeding data using `make gomigrate` and use it for your development.

### 4. Run Server

We also created makefile commands to run the server using `make gorun` and open `localhost:8000` on your browser
