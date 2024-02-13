# gosmoclash

> ATTENTION: There is nothing here yet. I would sure like to make this one day, but now I have other stuff to do.

This is a rewrite of my 2D spaceship management game [cosmoclash](github.com/nexovec/cosmoclash) in go. The original was never finished, because the network code is basically unmanagable in lua, even though the architecture isn't super horrible.

Cosmoclash is my exploratory programming project, so there aren't any firm requirements. You control a character manning a spaceship, you start on a space station where you undock your spaceship to warp into a battle-royale like map where you compete for resources, avoid asteroid dust storms, do PvP or team up with others. The big dream of this game is to have a single persistent world for all players with an expansive economy system to truly simulate difficult political relationship in a space-born civilization, so code needs to be carefully designed for this right from the get-go.

A very interesting part about this the unique possibility to compare this code to the original lua code I've spent a couple months working on. Because this could be a long-running project, I will carefully document the process too.

## Getting started

### Requirements

Have a Ubuntu or debian based system.
Launch `./development-setup.sh`.

### Launch

`go run main.go`

### Configuration

No configuration yet

## Used technologies

- [Go](https://go.dev/)
- [Ebiten](https://ebitengine.org/)
- [Chipmunk](https://github.com/jakecoffman/cp)

## Backlog

- [ ] Make a walkable ship
  - [ ] Walking character + platform
  - [ ] Save and load a ship file
    - NOTE: use introspection
- [ ] Multiplayer
  - [ ] Parse configuration on what to run
  - [ ] Launch game or server or both
  - [ ] Synchronize player state
    - NOTE: use enet
  - [ ] Server-side physics
  - [ ] Local physics with prediction
  - [ ] Account management
  - [ ] Session management
- [ ] Assets
  - [ ] Load images
  - [ ] Load audio
  - [ ] Togglable Hot-reload
- [ ] Rendering
  - [ ] Render sprite animations
  - [ ] 2D light
  - [ ] blur
  - [ ] screenshake
  - [ ] flashes
  - [ ] antialiasing
  - [ ] Stars
- [ ] Integrations
  - [ ] Tiled
  - [ ] Aseprite
- [ ] Gameplay
  - [ ] Inventory system
    - [ ] Item interaction
    - [ ] Teleports
    - [ ] Player inventory
    - [ ] Dropping items
  - [ ] Ship navigation
    - NOTE: requires world switching to be done
    - [ ] Zoomed out camera view
    - [ ] Visualize movement path
    - [ ] Move ship
    - [ ] Ship bounding box + collisions
    - [ ] Ship spawn/despawn
      - NOTE: ensure uniqueness of a ship
  - [ ] Player interaction
    - [ ] Chat
    - [ ] Friends list
    - [ ] Corporation(guild system)
    - [ ] Player party(intra-ship party)
    - [ ] Alliance(inter-ship party)
  - [ ] PvP
    - NOTE: requires the particle system to be done
    - [ ] Player combat
      - [ ] Equip weapon
      - [ ] Shoot the particle
      - [ ] Visual effect on collision
      - [ ] Player HP
        - [ ] Health bar
        - [ ] Store on server
        - [ ] Player dies
    - [ ] Ship combat
      - [ ] Health bar
      - [ ] Lasers
      - [ ] Shields
      - [ ] Rockets
      - [ ] Mines
      - [ ] Bullets - ?? do we want this?
  - [ ] Mobs
    - NOTE: requires the AI toolkit to be done
    - [ ] Pirate ships
    - [ ] Home station NPCs
  - [ ] Game loop
    - NOTE: I suggest 15-60 minute raid-based game loop
    - REFERENCES: see cycle frontier, escape from tarkov, sievert
- [ ] Tooling
  - [ ] UI library
  - [ ] Audio queuing
  - [ ] World switching(for both ships and players)
  - [ ] Particle system
    - [ ] Space debris
- [ ] Platform support
  - [ ] Graphics settings
  - [ ] Compilation targets
    - [ ] Windows
    - [ ] Android
    - [ ] WASM
  - [ ] Mobile controls
