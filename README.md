# gosmoclash

This is a rewrite of my 2D spaceship management game [cosmoclash](github.com/nexovec/cosmoclash) in go. The original was never finished, because the network code is basically unmanagable in lua, even though the architecture isn't super horrible.

Cosmoclash is my exploratory programming project, so there aren't any firm requirements. You control a character manning a spaceship, you start on a space station where you undock your spaceship to warp into a battle-royale like map where you compete for resources, avoid asteroid dust storms, do PvP or team up with others. The big dream of this game is to have a single persistent world for all players with an expansive economy system to truly simulate difficult political relationship in a space-born civilization, so code needs to be carefully designed for this right from the get-go.

A very interesting part about this the unique possibility to compare this code to the original lua code I've spent a couple months working on. Because this could be a long-running project, I will carefully document the process too.