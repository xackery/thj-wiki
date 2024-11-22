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
  - A more complicated version, create a hotkey to target nearest corpse, and also bind it to a hotbutton slot, and in this hotbutton write in:
    - `/loot`
    - `/notify LootWnd LW_LootAllButton leftmouseup`
    - `/notify ConfirmationDialogBox No_Button leftmouseup`