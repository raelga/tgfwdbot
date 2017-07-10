[![Build Status](https://travis-ci.org/raelga/BanMeNotBot.svg?branch=master)](https://travis-ci.org/raelga/BanMeNotBot)
[![Go Report](https://goreportcard.com/badge/github.com/raelga/BanMeNotBot)](https://goreportcard.com/report/github.com/raelga/BanMeNotBot)

# BanMeNotBot

Bot to allow talking in Telegram supergroups when banned.

## Why

Sometimes, due to trolls reporters, users can be banned from any Telegram supergroup, not only the one where reported. 
With this bot, users will be able to participate in specific groups through the bot.

This bot can be added to the group, identified by `TELEGRAM_GROUP_ID`, and will:

- Forward any message from the Group to the private conversation with the banned user
- Forward any message from the private conversation with the banned user to the Group

The admins of the groups where the bot is a member will decide if the allows the banned users to participate or not.

## Disclaimer

**This bot is not intended to be a tool for spammers and malicious users to bypass Telegram bans.**
The purpose is to allow users to interact in the groups with the authorization of the group admins when globally banned from supergroups due to troll reportings.
