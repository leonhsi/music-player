# music-player

## Build and Run

Clone this repo:
```
git clone git@github.com:leonhsi/music-player.git <dir_name>
```

### Crawler

The web crawler will download the music mp3 and image covers.

Run:
```
cd crawler
python3 crawler.py
```

This will download and mp3 to `frontend/src/audios` and cover images to `frontend/src/covers`

### Frontend

Install the node modules:
```
cd frontend
npm install
```

Run the react project on default port (3000):
```
npm start
```

Then the music player would show on `http://localhost:3000`:

![image](https://hackmd.io/_uploads/Bkw-Ht136.png)


### Backend

Useful app: 
1. [docker desktop](https://www.docker.com/products/docker-desktop/) to manage docker containers, 
2. [TablePlus](https://tableplus.com) to view databases, 
3. [postman](https://www.postman.com/downloads/) to send simple HTTP request.


Download postgres image on [docker hub](https://hub.docker.com/_/postgres):
```bash
docker pull postgres
docker images
```

Create postgres container:
```bash
cd backend
make postgres
```

Create database for music-player in postgres container:
```bash
make createdb

# if need to delete database
make dropdb
```

Create schema for music-player in database:
```bash
make migrateup

# if need to delete schema
make migratedown
```
The schema:

![image](https://hackmd.io/_uploads/SysxuYyn6.png)

Generate golang code to manipulate database using sqlc:
```
make sqlc
```

Run the service on port 8080:
```
make run
```

![image](https://hackmd.io/_uploads/S1o7sKkha.png)


#### TablePlus
We can use TablePlus to view our database.

open TablePlus, connect to the postgres container
(password is 123 according to the `make postgres` command)

![image](https://hackmd.io/_uploads/By-NFY12p.png)


after running `make createdb` command, the music-player database is created.
click the upper left icon, choose music-player database

![image](https://hackmd.io/_uploads/HkPsFFkhp.png)

after running `make migrateup` command, the tables are created, it should be empty:
(We could clean the tables by `make migratedown`)
![image](https://hackmd.io/_uploads/Syov5t1na.png)

after running `make run` command, the initial data is written to the tables:

![image](https://hackmd.io/_uploads/HkIncYk3p.png)

#### Postman

after the serve is running, we could use postman to test the HTTP APIs.

current APIs:
* `POST` method
    * `POST /songs`: create songs
    * `POST /artists`: create artists
* `GET` method
    * `GET /songs/name/:name`: get song metadata by song name
    * `GET /songs/name/:id`: get song metadata by song id
    * `GET /songs`: list all songs' metadata
    * `GET /artists/:name`: get artist by artist name
    * `GET /artists`: list all artists

We could test these APIs with postman. 

Let's say I want to get the metadata of song `10 Minute`:

select `GET` method, and type `http://localhost:8080/songs/name/10%20Minute`:
(`%20` denotes a space)

![image](https://hackmd.io/_uploads/SJFYpYJ26.png)

Send the request, and we should get the correct response from server:

![image](https://hackmd.io/_uploads/HJKzAtkhp.png)

The server would show that it receive a HTTP request, successfully handled it and return status which 200 means OK:

![image](https://hackmd.io/_uploads/S1jo0Y1np.png)

If we want to list songs, add params in the request (for example, list 3 songs):

![image](https://hackmd.io/_uploads/Sk-Z1cyh6.png)
