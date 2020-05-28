# Smegbot

A discord bot that enables you to easily replace all of a given user's roles
with a predefined one. The user's current roles will be saved in case they
need to be restored at a later date.

## Installation

Smegbot is written in Go, so to install you can run the following:
```
go install github.com/dshoreman/smegbot
```

## Usage

To add your Smegbot to a server, you'll first need to create an Application with a bot user at https://discord.com/developers/applications. Make a note of the Client ID as you'll need it soon.

Run the bot with `$GOBIN/smegbot -t BOT_TOKEN` where `BOT_TOKEN` is the token you were given for the Application's bot user.

### Adding Smegbot to Discord

Once you've created an application, replace `CLIENT_ID` in the URL below with your  Client ID and paste it into a browser.
```
https://discord.com/api/oauth2/authorize?permissions=268823616&scope=bot&client_id=CLIENT_ID
```

## Commands

### `ping`
Responds with "Pong!" for testing the bot.

### `.roles <user_mention>`
Lists all the roles that the mentioned user has in the server.
