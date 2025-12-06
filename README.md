# Project エンジョックー

## Background

This is a hobby project that was born from some questions,

1. How to make a game that can be played with other people but I don't 
   have a confidence yet to make a multiplayer game that needs lot of costs
   and complexity?
2. Would it be cool to play a game that somewhat multiplayer but actually single player?

And here I am trying to make it true by making a game that its game world would
be defined by serialized data in a text format (e.g. JSON files). This game
world will have its own git repository and the game client will manage its
update, merge conflict, pull request, as a part of gameplay.

Because its nature of the game world as a serialized data that are shared using
git repository, this game will be a sandbox game. This project might act as
akin of game engine, or game SDK that can access the game data.

## How it works

1. Game client will create unique ID that represents current player
2. Game client then have some option:
   1. Create a new world
   2. Clone a game world from git repository
3. Play at your own pace, make a change in the game world
4. Sync the game world to git repository

## Milestone

1. Game World Data MVP
   1. Game World Loader
   2. Game World Saver
   3. Game World Game World Picker
   4. Game World Terrain
   5. Game World Entity Interface
   6. Game World Structure
2. C# Implementation
3. Typescript Implementation
