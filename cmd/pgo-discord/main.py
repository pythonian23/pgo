#!/usr/bin/env python3
import discord
import shellwords
import logging
import sys


parser = shellwords.ShellWords()
intents = discord.Intents.default()


class Client(discord.Client):
    async def on_ready(self):
        logging.info(f"{self.user} Online")
    async def on_message(self, msg: discord.Message):
        if msg.author.id == self.user.id:
            return
        command = parser.parse(msg.content)
        if command[0].lower() == "pgo":
            await msg.reply("Command output goes here")

client = Client(intents=intents)
client.run(sys.argv[1])
