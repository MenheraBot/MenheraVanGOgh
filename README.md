<p align="center">
    <img src="https://i.imgur.com/g9MuGLw.png" alt="Logo" width="160" height="160">

  <h3 align="center">๐ <b>Menhera Van GOgh</b> ๐</h3>

  <p align="center">
    An Application to manipulate images for MenheraBot written in Go. 
    <br />
    <a href="https://github.com/MenheraBot/MenheraBot"><strong>MenheraBot ยป</strong></a>
    <br />
    <br />
  </p>
</p>

## ๐จโ๐ป | Contributing

You may contribute to this project by opening an issue or creating a pull request on GitHub. If you want to add a new asset, you need to follow this document, and send the asset to [MenheraBot's Suppport Server](https://discord.com/invite/fZMdQbA).

## ๐ฅ | Running

To run Menhera Van GOgh, you need to have [Docker](https://www.docker.com/) in your machine. You have two options of installation, follow the one that applies to you. It is good to have a [Redis](https://redis.io/) instance to improve the performance of the application.

### ๐ฎ | Building the Image

> If you want to build the image yourself, you can do it by following these steps:

1. ๐งน Clone the repository

```bash
git clone https://github.com/MenheraBot/MenheraVanGOgh.git
```

2. ๐ป Building the Image

```bash
docker build . --tag vangogh
```

3. ๐โโ๏ธ Running a Container

```bash
docker run --name VangoghServer -p 2080:2080 -e "TOKEN=" -e "REDIS_URL=" -e "REDIS_DB=" --restart unless-stopped -d -t vangogh
```

> Obs: the `TOKEN` is just for authentication purpuses. `REDIS_URL` and `REDIS_DB` are the variables to connect to your redis instance. The `restart` policy used is because, well, no one wants a server down!

Now we can connect to HTTP to 2080 port!

### ๐ | Downloading the Image

> If you don't really want all the source code, and just want to execute the bot, you can just donwload the image from the Container Registry.

1. ๐ฅ Download the image

```bash
docker pull ghcr.io/menherabot/vangogh:latest
```

> You need to be [logged in](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry#authenticating-to-the-container-registry)

2. ๐โโ๏ธ Running a Container

```bash
docker run --name VangoghServer -p 2080:2080 -e "TOKEN=ReplaceWithToken" -e "REDIS_URL=" -e "REDIS_DB=" --restart unless-stopped -d -t ghcr.io/menherabot/vangogh:latest
```

> Obs: the `TOKEN` is just for authentication purpuses.`REDIS_URL` and `REDIS_DB` are the variables to connect to your redis instance. The `restart` policy used is because, well, no one wants a server down!

Creeper? Awww maan. Van GOgh is on!

## ๐จ | Made With

- [Go](https://go.dev/)
- [Gin](https://github.com/gin-gonic/gin)
- [Go Graphics](https://github.com/fogleman/gg)

## ๐ | Special Thanks

I shall thank [RabbitHouseCorp](https://github.com/RabbitHouseCorp) very much for [providing inspiration](https://github.com/RabbitHouseCorp) for this repository. It's the first time I've used Go, totally a long shot, not knowing anything, and thanks to them, I was able to ~~steal most of the code~~ start making this project, mainly using the Go Graphics base as inspiration. Thank you very much.

## โ๏ธ | License

Distributed under the MIT License. See `LICENSE` for more information.

## ๐ง | Contact

Discord: **Luxanna#5757**

Twitter: **[@Luxanna_Dev](https://twitter.com/Luxanna_Dev)**

---

MenheraBot was made with โค๏ธ by Luxanna.
