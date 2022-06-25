<p align="center">
    <img src="https://i.imgur.com/g9MuGLw.png" alt="Logo" width="160" height="160">

  <h3 align="center">📔 <b>Menhera Van GOgh</b> 📔</h3>

  <p align="center">
    An Application to manipulate images for MenheraBot written in Go. 
    <br />
    <a href="https://github.com/MenheraBot/MenheraBot"><strong>MenheraBot »</strong></a>
    <br />
    <br />
  </p>
</p>

## 👨‍💻 | Contributing

You may contribute to this project by opening an issue or creating a pull request on GitHub. If you want to add a new asset, you need to follow this document, and send the asset to [MenheraBot's Suppport Server](https://discord.com/invite/fZMdQbA).

## 🔥 | Running

To run Menhera Van GOgh, you need to have [Docker](https://www.docker.com/) in your machine. You have two options of installation, follow the one that applies to you.

### 🔮 | Building the Image

> If you want to build the image yourself, you can do it by following these steps:

1. 🧹 Clone the repository

```bash
git clone https://github.com/MenheraBot/MenheraVanGOgh.git
```

2. 💻 Building the Image

```bash
docker build . --tag vangogh
```

3. 🏃‍♂️ Running a Container

```bash
docker run --name VangoghServer -p 2080:2080 -e "TOKEN=" --restart unless-stopped -d -t vangogh
```

> Obs: the `TOKEN` is just for authentication purpuses. The `restart` policy used is because, well, no one wants a server down!

Now we can connect to WS or HTTP to 2080 port!

### 🎉 | Downloading the Image

> If you don't really want all the source code, and just want to execute the bot, you can just donwload the image from the Container Registry.

1. 📥 Download the image

```bash
docker pull ghcr.io/menherabot/vangogh:latest
```

> You need to be [logged in](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry#authenticating-to-the-container-registry)

2. 🏃‍♂️ Running a Container

```bash
docker run --name VangoghServer -p 2080:2080 -e "TOKEN=ReplaceWithToken" --restart unless-stopped -d -t ghcr.io/menherabot/vangogh:latest
```

> Obs: the `TOKEN` is just for authentication purpuses. The `restart` policy used is because, well, no one wants a server down!

Creeper? Awww maan. Van GOgh is on!

## 🔨 | Made With

- [Go](https://go.dev/)
- [Gin](https://github.com/gin-gonic/gin)
- [Go Graphics](https://github.com/fogleman/gg)
- [Gorilla WebSocket](https://github.com/gorilla/websocket)

## 💖 | Special Thanks

I shall thank [RabbitHouseCorp](https://github.com/RabbitHouseCorp) very much for [providing inspiration](https://github.com/RabbitHouseCorp) for this repository. It's the first time I've used Go, totally a long shot, not knowing anything, and thanks to them, I was able to ~~steal most of the code~~ start making this project,  mainly using the Go Graphics base as inspiration. Thank you very much.

## ⚖️ | License

Distributed under the MIT License. See `LICENSE` for more information.

## 📧 | Contact

Discord: **Luxanna#5757**

Twitter: **[@Luxanna_Dev](https://twitter.com/Luxanna_Dev)**

---

MenheraBot was made with ❤️ by Luxanna.
