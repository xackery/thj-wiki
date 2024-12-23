---
title: Macros
weight: 0
---

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
- Donate to Tribute Merchant:
    - `/hotbutton Donate /notify TributeMasterWnd TMW_DonateButton leftmouseup`
- Set trader prices
    - To use this optimally, hotkey the "Set" hotkey to 1. hotkey the "Save" to 2. Then, click an item in your trader bag, press 1, type price, enter, press 2 to save. Click next item
    - `/hotbutton Set /notify BazaarWnd BZW_Clear_Button leftmouseup`
    - then edit button, second line:
    - `/notify BazaarWnd BZW_Money0 leftmouseup`
    - after pressing Set, you can hotkey saving the price with this:
    - `/hotbutton Save /notify BazaarWnd BZW_SetPrice_Button leftmouseup`
