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

## Macros

These macros can be copy pasted in game by setting the Paste From Clipboard hotkey in game.
    - Press ALT+O, Keys, UI, scroll down to Paste From Clipboard, and set it to CTRL+V

- Sell highlighted item to merchant: `/hotbutton Sell /notify MerchantWND MW_Sell_Button leftmouseup`
    - To add any quantity, add next line: `/notify QuantityWnd QTYW_Accept_Button leftmouseup`
- Press autobank when bank window is open: `/hotbutton Autobank /notify BigBankWnd AutoButton leftmouseup`
- Press combine on Tradeskill Window:
    - `/hotbutton Combine /notify TradeskillWnd CombineButton leftmouseup`
    - set the hotbutton down to a hot bar, then edit it to add into the button manually (can't paste into the hot button editor):
    - `/pause 1`
    - `/autoinv`
- Destroy item on cursor: `/hotbutton Destroy /notify IW_InvPage IW_Destroy leftmouseup`
    - To bypass confirmation dialog, can add: `/notify ConfirmationDialogBox CD_Yes_Button leftmouseup`
- Loot active corpse: `/hotbutton Loot /notify LootWnd LW_LootAllButton leftmouseup`
    - Note: if you shift right click a corpse, you will auto loot it without needing any of these macros
    - A more complicated version, create a hotkey to target nearest corpse, and also bind it to a hotbutton slot, and in this hotbutton write in:
        - `/loot`
        - `/notify LootWnd LW_LootAllButton leftmouseup`
        - `/notify ConfirmationDialogBox No_Button leftmouseup`
- Mass turn in: Idea is to put the stack of items you're turning in to bag 1 slot 1, hold CTRL, and click this hotbutton:
    - `/hotbutton TurnIn /itemnotify in pack1 1 leftmouseup`
    - You can repeat the above /itemnotify line multiple times in the hotkey if multiple handins are needed of same item
    - Next, go into ALT+O, Keys, Commands, and rebind Use centerscreen from 'u' to 'CTRL+Z'
    - Now, you can automate the give button by: `/hotbutton Give /notify GiveWnd GVW_Give_Button leftmouseup`
    - So in summary, hold CTRL, left click TurnIn hotbutton, press Z (with ctrl still hand), repeat until give window is full, then press Give hotbutton
- Accept items from parcel:
    - `/hotbutton Parcel /notify MerchantWnd MW_MerchantSubwindows tabselect 2`
    - Then, you'll write manually in next line
    - `/notify MerchantWnd MW_ItemlistMail listselect 1`
    - `/notify MerchantWnd MW_Retrieve_Button leftmouseup`

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