---
title: Game Mechanics
weight: 0
---

Game Mechanics

## Spell Damage

- gear with int on it (regular, enchanted, and legendary) tends to have spell damage on legendary gear.
- spell damage has a cap of 750
- spell damage increases single target DoT spells with a 2 to 1 ratio (2 spell damage = 1 damage increase)
- spell damage does not influence bard song AE dots
- spell damage is found in your inventory window, on the stats page.
- spell damage increases proc DD spell effect damage amount with a 1 to 1 ratio
- spell damage increases sympethic strike DD damage amount with a 1 to 1 ratio
- spell damage increases AE DD spell damage amount with a 1 to 1 ratio

## Heal Amount

- gear with wis on it (regular, enchanted, and legendary) tends to have heal amount on legendary gear.
- heal amount has a cap of 750
- heal amount increases single target heal spells heal amount with a 1 to 1 ratio
- heal amount is found in your inventory window, on the stats page.
- heal amount affects runes, their absorption amount is increased with a 1 to 1 ratio
- heal amount affects sympethic healing value amount with a 1 to 1 ratio

## Flurry

- Warriors and Berserkers get an AA called flurry
- You must have the triple attack skill and succeed a roll to trigger flurry
- You must succeed a roll on: double attack, triple attack.
    - Two rolls are then executed for flurry, if you succeed each roll, you can attack up to 5 times in a single round


## Upgrading Items

- if an item can be placed in your power slot, you will find that mobs that give you experience will also give item experience
- if you use the Consume Item AA while holding the same item as the one in your power slot of any quality, it will be destroyed to increase the level of the item in your power slot
- if you do level up an item via the power slot, it immediately becomes NO TRADE
- to upgrade a regular item to enchanted, it requires 3 regular items to be merged with an existing one (aka 4 total)
- to upgrade a enchanted item to legendary, it requires 3 enchanted items to be merged with an existing one (aka 4 total)
-

## Upgrading Bags and Augs

- augs and Bags are intentionally flagged to not be able to go into the power slot
- items not in the power slot are not able to be leveled or use the consume item AA to level them up
- this was a design decision, to make legendary bags and legendary augs to be more valuable

## Tribute

- in the bazaar at the east bank near a Mischievous Halfling, you'll find Aeyln D`Sai (Tribute Master)
- many items can be given to this NPC to gain a point system called favor
- this NPC also let's you set various focus effects and stat bonuses that favor will cost to maintain
- press ALT+U to open your tribute window
    - Press the activate button to turn on tribute. Press deactive to disable it
- tribute will consume the Active Cost amount every 10 minutes until turned off. Current Favor shows your remaining tribute points you have to consume


## UI Tips and Tricks

- You can pick up a bag and ctrl+left click another bag to transfer all items in your held bag into the targetted clicked bag
- Hold shift and right click a corpse to auto loot all contents
- Hold shift while pressing sell to sell regardless of quantity of the highlighted item
  - This also applies to buying from a merchant
- Hold control while pressing sell to sell one quantity of the highlighted item
  - This also applies to buying from a merchant


## Procs

- Weapon and aug procs do not have a PPM limit timer
- 2 Handed weapons and ranged weapons only tap into the "double proc" mechanic if the weapon has a proc, and at least one aug has a proc
- Spell Buff have a PPM timer, and are prioritized by buff slot ID (lower ID's are prioritized)
- Priority of buff procs are in this order for proc attempt priority:
    - perma buffs -> no timer (things like AAs that give you a permanent proc, since no timer can roll to next in line)
    - spell buff procs -> timer (things like Vampiric Embrace)
    - ranged attack procs -> timer (procs that only work with archery/throwing)
    - aa melee procs -> timer (things like rabid bear or dance of the blade)
    - one shot poison/apply poison junk (just removes no timer) (things that use apply poison and old school poison system)
- When you hit a buff proc on a single attack and succeed to proc, it increments a counter by 1. Once you hit 4 total successful spell procs, it breaks out of the spell proc loop
- How procs are prioritized for weapons is:
    - attempt to proc built in weapon
    - if weapon proc fails roll, try to proc slot 1 aug
    - if slot 1 aug proc fails roll, try to proc slot 2 aug

## Macros

- [Macros were moved to it's own page](/getting-started/macros)

## Pet Mechanics

> [!tip]- Pet bonds are very strong in this world
>
> In The Heroes Journey universe, pets are much more strongly tied to their owners. So much so that NPC's will sometimes be able to detect this connection. For example, if you command your pet to attack, your enemies may very well ignore the pet and come for its owner, YOU!  You may need to instruct your pets to taunt these enemies. Beware, if a pet dies, the severed connection will most likely disrupt the stasis state known as Feign Death.

> [!tip]- Pet specific spells affect all summoned pets
>
> You may attune yourself to up to three pets so long as your class selections support this (one pet per class). Because of the strong ties to your pets, many of the pet-specific magicks will affect all of them. While worldy charmed pets will benefit from many spells that specialize on pets, they do not share the same Pet Affinity that summoned pets share in order to receive the benefits of general group spells.

> [!tip]- You can issue commands to all pets simultaneously or to a single pet
>
> By default, your pets will operate in harmony: attacking as a group, holding as a group, even releasing the bonds formed with you as a group. If you wish for a particular pet to execute a command on its own, you must direct your full attention to it by targeting the pet before issuing its command. Note that if you issue a stance command such as taunt or guard to a single pet, then issue the same command to all pets, they will harmoniously remain out of sync.

> [!tip] - Check the status of those pets and the bling they're carrying!
>
> You can check the status of your pets, including the stats,  weapons, and armor  equipped etc. by using the command:
/say #mystats

## Maps

> [!tip]- Making use of your map
>
> On your map, you can see a location of mobs moving around the map.  Some MQ2 functionality was imported into THJ, one of these being the /mapfilter command.  You can do /mapfilter named and it will filter out other mobs, you can also do /mapfilter custom to enter the name of a specific named mob "(/mapfilter custom Slate)"  Now, when I'm in east commons, I can open my map and track Seargent Slate in real time.  In addition, if the mob is within range to you, you can right click it on map to target it as well making it a very handy tool for pacify and the sort!

Note, when you make a new character, go to your skills tab and open the "track" window one time, and you will be able to see the names of the mobs on your map within range of your track skill as normal, with increased range for higher skill ranks.


## AA's

> [!tip]- AA's that stack
>
>We always want to look for synergy with our AA's, but we also can't ignore the little, yet arguably more important.  Everyone gets Run 5 as an AA, and that gets you up to Spirit of the Wolf speed.  Bard gets to additional ranks of run with Fleet of Foot.  Monk gets 3 additional ranks via Weightless Steps.  You get a combined Run 10, which is on par with Selo's using Naggy drums.  Keep an eye out for class AA's that stack in that way.

## Resists

- 2 resists equates to roughly 1% chance to resist
- player resists go up to 550 base. If you're above this number, subtract your heroic resists (the gold +#) from the total, and if that's less than 550, you get benefit from normal resist buffss and gear still. Otherwise, heroic is the only way to increase them.
- spells can have min resists and max resists fields set to modify final resist checks
- a dice is rolled of 0 to 200 to determine if a spell lands, and returns partial landings as well. The enemy's natural resists, and factors noted below, among some other edge cases are used

### Outgoing damage spell landing

When you cast a spell on an NPC:

- an enemy with a level greater than yours adds a +1 to all resists with a cap of +9
- an enemy with a level lower than yours adds a -1 to all resists until the cap of -9
- an enemy with a level lower than 20 levels of yours adds -1000 to all resist checks
- each point of heroic charisma adds -1 to resist check

### Incoming damage spell landing

When an NPC casts on you:

- an enemy with a level greater than yours adds -1 to all resists with a cap of -15

## Dispelling Mobs

This is just to bring to attention a trend I've seen about bards being able to dispel Innoruuk DS vs others, and a deep dive into how Cancel Magic works as a whole.

- Dispel doesn't use resists in any manner.
- TryDispel is a function that takes a player's caster level, the buff level, and a level_modifier
- When you try to dispel a mob, the server iterates all buffs on the target mob, and rolls a d100 vs every buff

The formula for dispel is as follows:
- start with a dispel chance of 32%
- take your caster level, and subtract the enemy's level to get a level_diff
    - there are two ways to modify your caster level, one is is with Jamfest AA, which gives +6 caster level for this check, and the other is with the spell's level mod, I listed below the values of commonly used spells
- if the level_diff is > 0, level_diff is multiplied by 7 for dispel chance. (Basically every level_diff below a target is a 7% penalty chance)
- if the level_diff is < 0, level_diff is multiplied by 2. (Basically every level_diff above a target is a 2% bonus chance)
- if dispel_chance is greater than or equal to 100, return success
- if dispel_chance is less than 10, set to 10
- roll a d100 vs the final value, return true if value is greater or equal to dice


A quick reference of dispel spells and their level mods:

- Cancel Magic +1 level, 1 try
- Syvelian's Anti-Magic Aria +4 level, 1 try
- Nullify Magic +4 level, 2 tries
- Annul Magic +9 level, 2 tries
- Abolish Magic +9 level, 3 tries
- Purge Magic +9 level, 2 tries
- Recant Magic +9 level, 4 tries
- Pillage Magic +9 level, 4 tries

Scenario 1: Non-bard Cancel Magic
```
To show this in action, if I'm level 60, non-bard and fight Innoruuk (who is level 70)

dispel_chance = 32 // baseline 32% chance
my_level = 60 // my level is 60
my_level += 1 // add cancel magic level mod 1
level_diff = my_level-70 = -9 // i'm 9 levels below Innoruuk with cancel magic bonus 1
dispel_chance += level_diff * 2 = 18 // since it's lower than 0, we lower chance by x2 of diff

so any time I use cancel magic, I have a 18% chance to land vs innoruuk
```

Scenario 2: Bard with Jamfest uses Cancel Magic
```
If I'm a level 60, bard with Jamfest max, and fight Innoruuk (who is level 70)

dispel_chance = 32 // baseline 32% chance
my_level = 60 // my level is 60
my_level += 6 // add jamfest level mod 6
my_level += 1 // add cancel magic level mod 1
level_diff = 67-70 =  -3 // i'm 4 levels below Innoruuk with the jamfest +6
dispel_chance += level_diff * 2 = 26 // since it's lower than 0, we lower chance by x2 of diff

so any time a bard uses cancel magic, they have a 26% chance to land vs innoruuk
```

Scenario 3: Non-bard Annul Magic
```
dispel_chance = 32 // baseline 32% chance
my_level = 60 // my level is 60
my_level += 9 // add annul magic level mod 9
level_diff = my_level-70 = -1 // i'm 1 level below Innoruuk with annul magic bonus 9
dispel_chance += level_diff * 2 = 30 // since it's lower than 0, we lower chance by x2 of diff

so any time I use annul magic, I have a 30% chance to land vs innoruuk, but two rolls, so roughly ~ 51% chance to land (2x30%)
```



- If an ally puts a debuff on a mob, it has a ~32% chance to be dispelled, since they're equal level to you
- Using higher levels of cancel magic give you multiple rolls per cast, look at the SPA data and how many slots have cancel magic, and that's how many times you are rolling vs each buff/debuff on a target