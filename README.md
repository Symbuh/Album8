# Foundant Coding Challenge

Fullstack image sharing service (golang/docker/react/typescript/postgreSQL)

---

### Running the App

#### Through docker

1. Run:

```
sudo docker run -p 3000:8080 nsabadicci/foundant-coding-challenge
```

in your terminal after installing docker and logging in with a dockerhub account

2. visit http://localhost:3000 in your browser to view the client

##### Locally

1. open two terminals
2. CD into the server directory on one terminal then run the command:
```
$ go mod download
```
3. cd into client on the other terminal and run:

```
$ npm i
$ npm start
```

---

## Functionality

- Upload any image and save it to the database. along with name description and tags
- view all images
- click an image to view individually
- delete button removes clicked images
- Filter images by tag through the dropdown
- click all images to return home


### Notes

Front end is loosely typed and low effort as this challenge is for a systems-design position.

Currently the Postgres database is being hosted by ElephantSQL to speed up development, this may yield a slightly slower experience than running postgres on your local machine through docker.

---

## License & Copyright

© Nicholas Sabadicci
