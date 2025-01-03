# AREA

The goal of the AREA project is to build an Actions REAction applications

In this project we will have Services (as Google, Github, Spotify...) that will offer Action and Reaction, we will be able to **link actions to reactions**

by exemple: *if I receive an email I launch a Lo-Fi Playlist on Spotify*

This project is composed of three parts:

- Server 
    - Implement Services
    - Implement Database
    - Implement OAUTH2
    - Implement Actions
    - Implement Reactions
- Web
    - Implement OAUTH2
    - Application for user to make Actions Reaction
- Mobile
    - Implement OAUTH2
    - Application for user to make Actions Reaction

We use ***GoLang*** for the Server, ***React Typescript*** for the Web and ***Flutter*** for the Mobile.

### Services implemented:

- Date Time
- Meteo
- Discord
- Google
- Github
- Gitlab
- Spotify
- DropBox
- TicketMaster
- Asana
- Twilio

## Usage

To use the Application you only need to
```sh
    docker-compose up -d
```

And it will launch the application

## Documentation:

[Wiki](https://shorturl.at/1yxb8) Containing everything you need to know to be able to develop on this project

### Swagger Link in Local Host:

- Api-Gateway
    - http://127.0.0.1:8080/swagger/index.html
- Date Time Api
    - http://127.0.0.1:8082/swagger/index.html
- Discord Api
    - http://127.0.0.1:8083/swagger/index.html
- Asana Api
    - http://127.0.0.1:8092/swagger/index.html
- Dropbox Api
    - http://127.0.0.1:8096/swagger/index.html
- Github Api
    - http://127.0.0.1:8086/swagger/index.html
- Gitlab Api
    - http://127.0.0.1:8087/swagger/index.html
- Google Api
    - http://127.0.0.1:8088/swagger/index.html
- Meteo Api
    - http://127.0.0.1:8089/swagger/index.html
- Spotify Api
    - http://127.0.0.1:8091/swagger/index.html
- User Services
    - http://127.0.0.1:8085/swagger/index.html

## Co-Contributors

<table>
  <tr>
    <td align="center">
      <a href="https://github.com/paulbardeur">
        <img src="https://avatars.githubusercontent.com/u/114899301?v=4" width=85><br>
        <sub>Paul Bardeur</sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/edward-lalande">
        <img src="https://avatars.githubusercontent.com/u/114470214?v=4" width=85><br>
        <sub>Edward Lalande</sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/timotheeplisson">
        <img src="https://avatars.githubusercontent.com/u/91876984?v=4" width=85><br>
        <sub>Timothee Plisson</sub>
      </a>
    </td>
  </tr>
</table>
