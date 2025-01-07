---
title: Macros
weight: 0
---

These macros can be copy pasted in game by setting the Paste From Clipboard hotkey in game.
    - Press ALT+O, Keys, UI, scroll down to Paste From Clipboard, and set it to CTRL+V

## Sell Item To Merchant

This is handy when you are going to be selling a bunch to a merchant.

- `/hotbutton Sell /notify MerchantWND MW_Sell_Button leftmouseup` Press sell button
- `/notify QuantityWnd QTYW_Accept_Button leftmouseup` Accept the max (default) quantity on stacks

## Auto Bank

This presses the auto bank button. Hold an item on your cursor and press this key while a bank is open, to deposit an item

- `/hotbutton Autobank /notify BigBankWnd AutoButton leftmouseup`

## Tradeskill Combine

When working with tradeskills, it can get annoying to press combine, then auto inventory after. This simplifies the flow to one button press spam

- `/notify TradeskillWnd CombineButton leftmouseup`
- `/pause 1, /autoinv` You can repeat this line and tweak delays for big combines


## Destroy Item On Cursor

This requires your inventory window to be open, if it isn't, the hotbutton is ignored

- `/notify IW_InvPage IW_Destroy leftmouseup`
- `/notify ConfirmationDialogBox CD_Yes_Button leftmouseup` bypress YES confirmation

## Loot Corpse

Go into your options, Keys, Target, and bind the target nearest corpse key to a key of your choice. Then, create a hotbutton with the following:

- `/loot`
- `/notify LootWnd LW_LootAllButton leftmouseup`
- `/notify ConfirmationDialogBox No_Button leftmouseup`

## Mass Give Items to NPC

Idea is to put the stack of items you're turning in to bag 1 slot 1, hold CTRL, and click this hotbutton:

- Go to Options, Keys, Commands, and rebind Use centerscreen from 'u' to 'CTRL+Z'
- Create a new hotbutton with the following:
- `/itemnotify in pack1 1 leftmouseup`
- `/itemnotify in pack1 1 leftmouseup`
- `/itemnotify in pack1 1 leftmouseup`
- `/itemnotify in pack1 1 leftmouseup`
- Hold CTRL down. Click the hotkey above, then Z, repeat until give window is full, then press:
- `/notify GiveWnd GVW_Give_Button leftmouseup` to turn in items

## Accept Items From Parcel

Got a lot of items in your parcel window to grab? Here's to speed it up

- `/notify MerchantWnd MW_MerchantSubwindows tabselect 2`
- `/notify MerchantWnd MW_ItemlistMail listselect 1`
- `/notify MerchantWnd MW_Retrieve_Button leftmouseup`

## Donate to Tribute Merchant

Speed up giving items for tribute:

- `/notify TributeMasterWnd TMW_DonateButton leftmouseup`


## Set Trader Prices

To use this optimally, hotkey the "Set" hotkey to 1. hotkey the "Save" to 2. Then, click an item in your trader bag, press 1, type price, enter, press 2 to save. Click next item

- `/hotbutton Set /notify BazaarWnd BZW_Clear_Button leftmouseup` "Set" hotkey
- `/notify BazaarWnd BZW_Money0 leftmouseup` "Set" hotkey line 2

- `/notify BazaarWnd BZW_SetPrice_Button leftmouseup` "Save" hotkey
