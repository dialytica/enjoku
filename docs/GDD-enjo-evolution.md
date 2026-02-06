# Game Design Documentation

## Basic Info

* Game Name: Enjo Evolution
* Genre: Clicker, Creature Collections, Casuals

## Game Elements

* Cultivating creatures, grow and evolve based on its diet and biome
* Expand areas and biomes
* Collect in game currencies that are recharges from daily check in and
  activities
* Collect various creature, based on its diet and biome

## Player

Single player game, but open to be scalable into multi player interaction

## Technical Specs

### Technical Form

2D Graphics

### View

Top Down View, 3rd person perspective. Player act as omniscience entity.

### Platform

Android, Steam

### Language

Either C# or GDScript (in Godot)

### Device

Mobile, PC, Web(?)

## Game Play

Grow an Enjo, mysterious creatures that devour Enjorodients and evolve into more
complex creatures. Collects the various form of Enjo by evolving to the next
evolution or adopt new Enjo! Expand the plot of land to accomodate more Enjo!
Fend off your Enjo from Stray Enjos that might be more advanced and try to
devour your Enjos! Experience various random events to enrich Enjo Experiences.

### Winning Condition

* Evolve Enjo into the most advanced being
* Collects all of Enjos Variants

### Losing Condition

* Your Enjo are being devoured by strays

### Game Play Outline

* Opening Game Application
* Start the Game Menu
* Game Options Menu
* Enjo Evolution
* Enjo Events
* Stray Enjo Encounters
* Enjo Index (Collections index, and integrated with achievement)
* Story Synopsis
* Plot of Lands
* Lists of Enjos
* List of Plot of Lands
* Enjo Store
* End Game: Collect all of possible creatures, biomes, and events

### Key Features

* Growing and nurturing unique Enjo creature
* Digital Collectibles and Achievements
* Defend your Enjo from Strays
* Casual simple gameplay

## Design Document

### Design Guidelines

* Ensure the freemium part of the game is designed tightly with game play
* Be flexible with premium costs. Adjust the pricing to promote player purchase
* This is casual game. The control must be as a simple few point and clicks
* As a reference, please look at Lonely Guy Game in Android

### Game Design Definition

You start a game with a single mysterious creature called Enjo, 200 Life Gems,
100 Evo Points and a single plot of land. You interact with Enjo by clicking on
them to select them, and then you can click various Enjorodients that spawn and
will be consumed by Enjo.

Enjorodients will yields Evo Points, and EnjoXP. Each Enjorodients have its own
EnjoXP and by achieving specific number of various EnjoXP, Enjo can evolve into
various form. The Enjorodients spawn is dependent to its biome. Such as in a
dark forest, there will be mushroom, bugs, and fruits. While in a plain biome we
might find small animals such as rabbits, small birds, and saplings.

The Enjorodients will spawn for every 30s, and have a maximum capacity of
spawned Enjorodients objects. Default capacity is 10. Player can force spawn to
maximum capacity by using 150 Life Gems. Player also can increase the maximum
spawn capacity on each plot of land by using 400 Life Gems or 1000 Evo Points.
The capacity price will increase at 2x rate for every increased capacity.
E.g. upgrade to capacity to 11 will costs 1000 Evo Points, then upgrade
capacity to 12 will costs 2000 Evo Points and so on. However, upgrading using
Life Gem have a flat price at 400 Life Gems.

Player can buy more plot of land too starting at 10000 Evo Points and will
increase its price at 2x rate every owned plot of land. However player can
buy with a flat price at 3000 Life Gems.

Life Gems is a premium in game currency. Current conversion rate is 1$ for
100 Life Gem. There will be a discount for more Enjorodients.

There is also random event that can be triggered by consuming 200 Evo Points
or 150 Life Gems. The random events has its own EnjoXP and might impact Enjo
Evolution. The random events also has its own rarity and the % of the rarity
can be boosted by using 300 Life Gems.

Enjo can also consume lesser evolved other Enjo creature. The consumed Enjo
creature will be converted into Enjorodients EnjoXP. Some special Enjo creature
will have its own EnjoXP when consumed.

There also Stray Adavanced Enjo that will try to consume your Enjo. Fight the
strays and get extra EnjoXPs.

### Game Flowchart

<!-- TODO: create flowchart and embed the image into document -->
1. Enjorodients Devour Flow
2. Enjorodients Spawn Flow
3. Evolution Flow
4. Upgrade Plot of Land Capacity Flow
5. Buy Plot of Land Capacity Flow
6. Booster Spawn Flow
7. Stray Enjos Encounter Flow

### Player Definition

Player is yourself, detached from the game itself. However you will have
access to the following section of the games.

* Game Scene with plot of land and Enjo and Enjorodients inside them
* List of Enjos, and view Enjo Status
* List of plot of lands and view its status
* Enjo Index, contains various discovery when playing game
* Game Options / Settings

Evo Points and Life Gems will be bound to Player too.

### Enjo Definition

Enjo is the main focus of the gameplay and main character. We collect and
guide them to consume Enjorodients. After obtaining specific number of EnjoXP
from Enjorodients, we can evolve Enjo into next step of evolution.

Every player will start with a single Enjo at Prime Enjo stage. This is the
most basic Enjo and Enjo that we can adopt from the shop.

#### Enjo Properties

Every Enjo have this basic properties that can be viewed in Enjo status

* Name, yes you can name your Enjo, or will be randomly generated
* Evolution Stage, e.g. Prime Enjo
* Stamina, Enjo's stamina that will be consumed after triggering random event
* EnjoXPs, an exhaustive list of consumed Enjorodients, Events, and other Enjo
  evolution
* Evolution History, showing its timeline of evolution

#### Enjo Evolution

Enjo will evovle after consuming and obtaining specific amount of EnjoXPs.
<!-- TODO: details more on Enjo Evolution -->
#### Enjo Adoption

Adopting new Enjo costs 15000 Evo Points for the first new adoption. The
adoption cost using Evo Points is using progressive pricing. The price increase
depends on the number of Enjo you owned. If you already had 2 Enjos, and
want to adopt new one, it will costs 30000 Evo Points, and then after having
this 3 Enjo you want adopt a new one it will costs 60000 Evo Points. We can
refer to the following formula:

  `cost(n)=initial_cost*2^n`

* n: number of owned Enjo
* cost(n): cost of new adoption after owning n Enjo
* initial_cost: The starting cost currently at 15000 Evo Points

Adoption using Life Gem will have a flat fixed cost at 4500 Life Gems.

<!-- TODO: Please finish the Game Design Definitions -->
