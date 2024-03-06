---
title: "Zone Guide"
weight: 3
---

Discover how to get from zone A to B.

<!--more-->

{{<rawhtml>}}
<div class="container">
  <form action="" id="searchForm">
    <h1>Search For Zone Links</h1>
    From:
    <input list="zones" name="from" id="from">
    To:
    <input list="zones" name="to" id="to">
  <datalist id="zones">
  </datalist>
    <button type="submit">Submit</button>
  </form>
    <div id="results"></div>
</div>


<script>


isAARoutingEnabled = false;

const zones = new Map([
["qeynos","South Qeynos"], // Classic
["qeynos2","North Qeynos"], // Classic
["qrg","Surefall Glade"], // Classic
["qeytoqrg","Qeynos Hills"], // Classic
["highkeep","HighKeep"], // Classic
["freportn","North Freeport"], // Classic
["freportw","West Freeport"], // Classic
["freporte","East Freeport"], // Classic
["runnyeye","Clan RunnyEye"], // Classic
["qey2hh1","West Karana"], // Classic
["northkarana","North Karana"], // Classic
["southkarana","South Karana"], // Classic
["eastkarana","East Karana"], // Classic
["beholder","Gorge of King Xorbb"], // Classic
["blackburrow","BlackBurrow"], // Classic
["paw","Infected Paw"], // Classic
["rivervale","Rivervale"], // Classic
["kithicor","Kithicor Forest (A)"], // Classic
["commons","West Commonlands"], // Classic
["ecommons","East Commonlands"], // Classic
["erudnint","Erudin Palace"], // Classic
["erudnext","Erudin"], // Classic
["nektulos","Nektulos Forest"], // Classic
["cshome","Sunset Home"], // Classic
["lavastorm","Lavastorm Mountains"], // Classic
["nektropos","Nektropos"], // Classic
["halas","Halas"], // Classic
["everfrost","Everfrost Peaks"], // Classic
["soldunga","Solusek's Eye"], // Classic
["soldungb","Nagafen's Lair"], // Classic
["misty","Misty Thicket (A)"], // Classic
["nro","North Ro"], // Classic
["sro", "South Ro"], // Classic
//["southro","South Ro (A)"], // Classic v2
["befallen","Befallen (A)"], // Classic
["oasis","Oasis of Marr"], // Classic
["tox","Toxxulia Forest"], // Classic
["hole","The Ruins of Old Paineel"], // Classic
["neriaka","Neriak Foreign Quarter"], // Classic
["neriakb","Neriak Commons"], // Classic
["neriakc","Neriak Third Gate"], // Classic
["neriakd","Neriak Palace"], // Classic
["najena","Najena"], // Classic
["qcat","Qeynos Catacombs"], // Classic
["innothule","Innothule Swamp (A)"], // Classic
["feerrott","The Feerrott(A)"], // Classic
["cazicthule","Cazic-Thule"], // Classic
["oggok","Oggok"], // Classic
["rathemtn","Mountains of Rathe"], // Classic
["lakerathe","Lake Rathetear"], // Classic
["grobb","Grobb"], // Classic
["aviak","Aviak Village"], // Classic
["gfaydark","The Greater Faydark"], // Classic
["akanon","Ak'Anon"], // Classic
["steamfont","Steamfont Mountains"], // Classic
["lfaydark","The Lesser Faydark"], // Classic
["crushbone","Clan Crushbone"], // Classic
["mistmoore","Castle Mistmoore"], // Classic
["kaladima","Kaladim (A)"], // Classic
["felwithea","Felwithe (A)"], // Classic
["felwitheb","Felwithe (B)"], // Classic
["unrest","Estate of Unrest"], // Classic
["kedge","Kedge Keep"], // Classic
["guktop","Upper Guk"], // Classic
["gukbottom","Lower Guk"], // Classic
["kaladimb","Kaladim (B)"], // Classic
["butcher","Butcherblock Mountains"], // Classic
["oot","Ocean of Tears"], // Classic
["cauldron","Dagnor's Cauldron"], // Classic
["airplane","Plane of Sky"], // Classic
["fearplane","Plane of Fear"], // Classic
["permafrost","Permafrost Keep"], // Classic
["kerraridge","Kerra Isle"], // Classic
["paineel","Paineel"], // Classic
["hateplane","The Plane of Hate"], // Classic
["arena","The Arena (A)"], // Classic
["soltemple","Temple of Solusek Ro"], // Classic
["erudsxing","Erud's Crossing"], // Classic
["stonebrunt","Stonebrunt Mountains"], // Classic
["warrens","The Warrens"], // Classic
["erudsxing2","Marauder's Mire"], // Classic
["bazaar","The Bazaar"], // Classic
["bazaar2","The Bazaar (2)"], // Classic
["arena2","The Arena (B)"], // Classic
["jaggedpine","The Jaggedpine Forest"], // Classic
["nedaria","Nedaria's Landing"], // Classic
["tutorial","Tutorial Zone"], // Classic
["load","Loading (A)"], // Classic
["load2","Loading (B)"], // Classic
["hateplaneb","The Plane of Hate"], // Classic
["shadowrest","Shadowrest"], // Classic
//["tutoriala","The Mines of Gloomingdeep (A)"], // Classic
//["tutorialb","The Mines of Gloomingdeep (B)"], // Classic
//["clz","Loading (C)"], // Classic
//["poknowledge","Plane of Knowledge"], // Classic
["soldungc","The Caverns of Exile"], // Classic
//["guildlobby","The Guild Lobby"], // Classic
//["barter","The Barter Hall"], // Classic
//["takishruins","Ruins of Takish-Hiz"], // Classic
//["freeporteast","East Freeport"], // Classic v2
//["freeportwest","West Freeport"], // Classic v2
//["freeportsewers","Freeport Sewers"], // Classic v2
//["northro","North Ro (B)"], // Classic v2
//["southro","South Ro (B)"], // Classic v2
//["highpasshold","Highpass Hold"], // Classic v2
//["commonlands","Commonlands"], // Classic v2
//["oceanoftears","Ocean Of Tears"], // Classic v2
//["kithforest","Kithicor Forest (B)"], // Classic v2
//["befallenb","Befallen (B)"], // Classic v2
//["highpasskeep","Highpass Keep"], // Classic v2
//["innothuleb","Innothule Swamp (B)"], // Classic v2
//["toxxulia","Toxxulia Forest"], // Classic v2
//["mistythicket","Misty Thicket (B)"], // Classic v2
["steamfontmts","Steamfont Mountains"], // Classic
["dragonscalea","Tinmizer's Wunderwerks"], // Classic
["crafthalls","Ngreth's Den"], // Classic
["weddingchapel","Wedding Chapel"], // Classic
["weddingchapeldark","Wedding Chapel"], // Classic
["dragoncrypt","Lair of the Fallen"], // Classic
["arttest","Art Testing Domain"], // Classic
["fhalls","The Forgotten Halls"], // Classic
["apprentice","Designer Apprentice"], // Classic
// ["fieldofbone","The Field of Bone"], // Kunark
// ["warslikswood","Warsliks Wood"], // Kunark
// ["droga","Temple of Droga"], // Kunark
// ["cabwest","West Cabilis"], // Kunark
// ["swampofnohope","Swamp of No Hope"], // Kunark
// ["firiona","Firiona Vie"], // Kunark
// ["lakeofillomen","Lake of Ill Omen"], // Kunark
// ["dreadlands","Dreadlands"], // Kunark
// ["burningwood","Burning Woods"], // Kunark
// ["kaesora","Kaesora"], // Kunark
// ["sebilis","Old Sebilis"], // Kunark
// ["citymist","City of Mist"], // Kunark
// ["skyfire","Skyfire Mountains"], // Kunark
// ["frontiermtns","Frontier Mountains"], // Kunark
// ["overthere","The Overthere"], // Kunark
// ["emeraldjungle","The Emerald Jungle"], // Kunark
// ["trakanon","Trakanon's Teeth"], // Kunark
// ["timorous","Timorous Deep"], // Kunark
// ["kurn","Kurn's Tower"], // Kunark
// ["karnor","Karnor's Castle"], // Kunark
// ["chardok","Chardok"], // Kunark
// ["dalnir","Dalnir"], // Kunark
// ["charasis","Howling Stones"], // Kunark
["cabeast","East Cabilis"], // Kunark
// ["nurga","Mines of Nurga"], // Kunark
// ["veeshan","Veeshan's Peak"], // Kunark
// ["veksar","Veksar"], // Kunark
// ["chardokb","The Halls of Betrayal"], // Kunark
// ["iceclad","Iceclad Ocean"], // Scars of Velious
// ["frozenshadow","Tower of Frozen Shadow"], // Scars of Velious
// ["velketor","Velketor's Labyrinth"], // Scars of Velious
// ["kael","Kael Drakkal"], // Scars of Velious
// ["skyshrine","Skyshrine"], // Scars of Velious
// ["thurgadina","Thurgadin"], // Scars of Velious
// ["eastwastes","Eastern Wastes"], // Scars of Velious
// ["cobaltscar","Cobalt Scar"], // Scars of Velious
// ["greatdivide","Great Divide"], // Scars of Velious
// ["wakening","The Wakening Land"], // Scars of Velious
// ["westwastes","Western Wastes"], // Scars of Velious
// ["crystal","Crystal Caverns"], // Scars of Velious
// ["necropolis","Dragon Necropolis"], // Scars of Velious
// ["templeveeshan","Temple of Veeshan"], // Scars of Velious
// ["sirens","Siren's Grotto"], // Scars of Velious
// ["mischiefplane","Plane of Mischief"], // Scars of Velious
// ["growthplane","Plane of Growth"], // Scars of Velious
// ["sleeper","Sleeper's Tomb"], // Scars of Velious
// ["thurgadinb","Icewell Keep"], // Scars of Velious
// ["shadowhaven","Shadow Haven"], // Luclin
// ["nexus","The Nexus"], // Luclin
// ["echo","Echo Caverns"], // Luclin
// ["acrylia","Acrylia Caverns"], // Luclin
// ["sharvahl","Shar Vahl"], // Luclin
// ["paludal","Paludal Caverns"], // Luclin
// ["fungusgrove","Fungus Grove"], // Luclin
// ["vexthal","Vex Thal"], // Luclin
// ["sseru","Sanctus Seru"], // Luclin
// ["katta","Katta Castellum"], // Luclin
// ["netherbian","Netherbian Lair"], // Luclin
// ["ssratemple","Ssraeshza Temple"], // Luclin
// ["griegsend","Grieg's End"], // Luclin
// ["thedeep","The Deep"], // Luclin
// ["shadeweaver","Shadeweaver's Thicket"], // Luclin
// ["hollowshade","Hollowshade Moor"], // Luclin
// ["grimling","Grimling Forest"], // Luclin
// ["mseru","Marus Seru"], // Luclin
// ["letalis","Mons Letalis"], // Luclin
// ["twilight","The Twilight Sea"], // Luclin
// ["thegrey","The Grey"], // Luclin
// ["tenebrous","The Tenebrous Mountains"], // Luclin
// ["maiden","The Maiden's Eye"], // Luclin
// ["dawnshroud","Dawnshroud Peaks"], // Luclin
// ["scarlet","The Scarlet Desert"], // Luclin
// ["umbral","The Umbral Plains"], // Luclin
// ["akheva","Akheva Ruins"], // Luclin
// ["codecay","Ruins of Lxanvom"], // Planes of Power
// ["pojustice","Plane of Justice"], // Planes of Power
// ["potranquility","Plane of Tranquility"], // Planes of Power
// ["ponightmare","Plane of Nightmare"], // Planes of Power
// ["podisease","Plane of Disease"], // Planes of Power
// ["poinnovation","Plane of Innovation"], // Planes of Power
// ["potorment","Plane of Torment"], // Planes of Power
// ["povalor","Plane of Valor"], // Planes of Power
// ["bothunder","Torden, The Bastion of Thunder"], // Planes of Power
// ["postorms","Plane of Storms"], // Planes of Power
// ["hohonora","Halls of Honor"], // Planes of Power
// ["solrotower","Solusek Ro's Tower"], // Planes of Power
// ["powar","Plane of War"], // Planes of Power
// ["potactics","Drunder, Fortress of Zek"], // Planes of Power
// ["poair","Eryslai, the Kingdom of Wind"], // Planes of Power
// ["powater","Reef of Coirnav"], // Planes of Power
// ["pofire","Doomfire, The Burning Lands"], // Planes of Power
// ["poeartha","Vegarlson, The Earthen Badlands"], // Planes of Power
// ["potimea","Plane of Time (A)"], // Planes of Power
// ["hohonorb","Temple of Marr (A)"], // Planes of Power
// ["nightmareb","Lair of Terris Thule"], // Planes of Power
// ["poearthb","Stronghold of the Twelve"], // Planes of Power
// ["potimeb","Plane of Time (B)"], // Planes of Power
// ["gunthak","Gulf of Gunthak"], // LoY
// ["dulak","Dulak's Harbor"], // LoY
// ["torgiran","Torgiran Mines"], // LoY
// ["nadox","Crypt of Nadox"], // LoY
// ["hatesfury","Hate's Fury, The Scorned Maiden"], // LoY
// ["guka","The Cauldron of Lost Souls"], // LDoN
// ["ruja","The Bloodied Quarries"], // LDoN
// ["taka","The Sunken Library"], // LDoN
// ["mira","The Silent Gallery"], // LDoN
// ["mmca","The Forlorn Caverns"], // LDoN
// ["gukb","The Drowning Crypt"], // LDoN
// ["rujb","The Halls of War"], // LDoN
// ["takb","The Shifting Tower"], // LDoN
// ["mirb","The Maw of the Menagerie"], // LDoN
// ["mmcb","The Dreary Grotto"], // LDoN
// ["gukc","The Ancient Aqueducts"], // LDoN
// ["rujc","The Wind Bridges"], // LDoN
// ["takc","The Fading Temple"], // LDoN
// ["mirc","The Spider Den"], // LDoN
// ["mmcc","The Asylum of Invoked Stone"], // LDoN
// ["gukd","The Mushroom Grove"], // LDoN
// ["rujd","The Gladiator Pits"], // LDoN
// ["takd","The Royal Observatory"], // LDoN
// ["mird","The Hushed Banquet"], // LDoN
// ["mmcd","The Chambers of Eternal Affliction"], // LDoN
// ["guke","The Foreboding Prison"], // LDoN
// ["ruje","The Drudge Hollows"], // LDoN
// ["take","The River of Recollection"], // LDoN
// ["mire","The Frosted Halls"], // LDoN
// ["mmce","The Sepulcher of the Damned"], // LDoN
// ["gukf","The Chapel of the Witnesses"], // LDoN
// ["rujf","The Fortified Lair of the Taskmasters"], // LDoN
// ["takf","The Sandfall Corridors"], // LDoN
// ["mirf","The Forgotten Wastes"], // LDoN
// ["mmcf","The Ritualistic Summoning Grounds"], // LDoN
// ["gukg","The Root Garden"], // LDoN
// ["rujg","The Hidden Vale"], // LDoN
// ["takg","The Balancing Chamber"], // LDoN
// ["mirg","The Heart of the Menagerie"], // LDoN
// ["mmcg","The Cesspits of Putrescence"], // LDoN
// ["gukh","The Accursed Sanctuary"], // LDoN
// ["rujh","The Blazing Forge"], // LDoN
// ["takh","The Sweeping Tides"], // LDoN
// ["mirh","The Morbid Laboratory"], // LDoN
// ["mmch","The Aisles of Blood"], // LDoN
// ["ruji","The Arena of Chance"], // LDoN
// ["taki","The Antiquated Palace"], // LDoN
// ["miri","The Theater of Imprisoned Horrors"], // LDoN
// ["mmci","The Halls of Sanguinary Rites"], // LDoN
// ["rujj","The Barracks of War"], // LDoN
// ["takj","The Prismatic Corridors"], // LDoN
// ["mirj","The Grand Library"], // LDoN
// ["mmcj","The Infernal Sanctuary"], // LDoN
// ["abysmal","Abysmal Sea"], // GoD
// ["natimbi","Natimbi, The Broken Shores"], // GoD
// ["qinimi","Qinimi, Court of Nihilia"], // GoD
// ["riwwi","Riwwi, Coliseum of Games"], // GoD
// ["barindu","Barindu, Hanging Gardens"], // GoD
// ["ferubi","Ferubi, Forgotten Temple of Taelosia"], // GoD
// ["snpool","Sewers of Nihilia, Pool of Sludge"], // GoD
// ["snlair","Sewers of Nihilia, Lair of Trapped Ones"], // GoD
// ["snplant","Sewers of Nihilia, Purifying Plant"], // GoD
// ["sncrematory","Sewers of Nihilia, the Crematory"], // GoD
// ["tipt","Tipt, Treacherous Crags"], // GoD
// ["vxed","Vxed, The Crumbling Caverns"], // GoD
// ["yxtta","Yxtta, Pulpit of Exiles"], // GoD
// ["uqua","Uqua, The Ocean God Chantry"], // GoD
// ["kodtaz","Kod'Taz, Broken Trial Grounds"], // GoD
// ["ikkinz","Ikkinz, Chambers of Destruction"], // GoD
// ["qvic","Qvic, Prayer Grounds of Calling"], // GoD
// ["inktuta","Inktu`Ta, The Unmasked Chapel"], // GoD
// ["txevu","Txevu, Lair of the Elite"], // GoD
// ["tacvi","Tacvi, Seat of the Slaver"], // GoD
// ["qvibc","Qvic, the Hidden Vault"], // GoD
// ["wallofslaughter","Wall of Slaughter"], // OoW
// ["bloodfields","The Bloodfields"], // OoW
// ["draniksscar","Dranik's Scar"], // OoW
// ["causeway","Nobles' Causeway"], // OoW
// ["chambersa","Muramite Proving Grounds (A)"], // OoW
// ["chambersb","Muramite Proving Grounds (B)"], // OoW
// ["chambersc","Muramite Proving Grounds (C)"], // OoW
// ["chambersd","Muramite Proving Grounds (D)"], // OoW
// ["chamberse","Muramite Proving Grounds (E)"], // OoW
// ["chambersf","Muramite Proving Grounds (F)"], // OoW
// ["provinggrounds","Muramite Proving Grounds"], // OoW
// ["anguish","Asylum of Anguish"], // OoW
// ["dranikhollowsa","Dranik's Hollows (A)"], // OoW
// ["dranikhollowsb","Dranik's Hollows (B)"], // OoW
// ["dranikhollowsc","Dranik's Hollows (C)"], // OoW
// ["dranikhollowsd","Dranik's Hollows (D)"], // OoW
// ["dranikhollowse","Dranik's Hollows (E)"], // OoW
// ["dranikhollowsf","Dranik's Hollows (F)"], // OoW
// ["dranikhollowsg","Dranik's Hollows (G)"], // OoW
// ["dranikhollowsh","Dranik's Hollows (H)"], // OoW
// ["dranikhollowsi","Dranik's Hollows (I)"], // OoW
// ["dranikhollowsj","Dranik's Hollows (J)"], // OoW
// ["dranikcatacombsa","Catacombs of Dranik (A)"], // OoW
// ["dranikcatacombsb","Catacombs of Dranik (B)"], // OoW
// ["dranikcatacombsc","Catacombs of Dranik (C)"], // OoW
// ["draniksewersa","Sewers of Dranik (A)"], // OoW
// ["draniksewersb","Sewers of Dranik (B)"], // OoW
// ["draniksewersc","Sewers of Dranik (C)"], // OoW
// ["riftseekers","Riftseekers' Sanctum"], // OoW
// ["harbingers","Harbingers' Spire"], // OoW
// ["dranik","The Ruined City of Dranik"], // OoW
// ["broodlands","The Broodlands"], // DoN
// ["stillmoona","Stillmoon Temple"], // DoN
// ["stillmoonb","The Ascent"], // DoN
// ["thundercrest","Thundercrest Isles"], // DoN
// ["delvea","Lavaspinner's Lair"], // DoN
// ["delveb","Tirranun's Delve"], // DoN
// ["thenest","The Accursed Nest"], // DoN
// ["guildhall","Guild Hall"], // DoN
// ["illsalin","Ruins of Illsalin"], // DoDH
// ["illsalina","Imperial Bazaar"], // DoDH
// ["illsalinb","Temple of the Korlach"], // DoDH
// ["illsalinc","The Nargilor Pits"], // DoDH
// ["dreadspire","Dreadspire Keep"], // DoDH
// ["drachnidhive","The Hive"], // DoDH
// ["drachnidhivea","Living Larder"], // DoDH
// ["drachnidhiveb","Coven of the Skinwalkers"], // DoDH
// ["drachnidhivec","Queen Sendaii's Lair"], // DoDH
// ["westkorlach","Stoneroot Falls"], // DoDH
// ["westkorlacha","Chambers of Xill"], // DoDH
// ["westkorlachb","Caverns of the Lost"], // DoDH
// ["westkorlachc","Lair of the Korlach"], // DoDH
// ["eastkorlach","Undershore"], // DoDH
// ["eastkorlacha","Snarlstone Dens"], // DoDH
// ["shadowspine","Shadowspine"], // DoDH
// ["corathus","Corathus Creep"], // DoDH
// ["corathusa","Sporali Caverns"], // DoDH
// ["corathusb","Corathus Lair"], // DoDH
// ["nektulosa","Shadowed Grove"], // DoDH
// ["arcstone","Arcstone"], // PoR
// ["relic","Relic"], // PoR
// ["skylance","Skylance"], // PoR
// ["devastation","The Devastation"], // PoR
// ["devastationa","The Seething Wall"], // PoR
// ["rage","Sverag, Stronghold of Rage"], // PoR
// ["ragea","Razorthorn, Tower of Sullon Zek"], // PoR
// ["takishruinsa","The Root of Ro"], // PoR
// ["elddar","The Elddar Forest"], // PoR
// ["elddara","Tunare's Shrine"], // PoR
// ["theater","Theater of Blood"], // PoR
// ["theatera","Deathknell, Tower of Dissonance"], // PoR
// ["freeportacademy","Academy of Arcane Sciences"], // PoR
// ["freeporttemple","Temple of Marr (B)"], // PoR
// ["freeportmilitia","Freeport Militia House"], // PoR
// ["freeportarena","Arena"], // PoR
// ["freeportcityhall","City Hall"], // PoR
// ["freeporttheater","Theater"], // PoR
// ["freeporthall","Hall of Truth"], // PoR
// ["crescent","Crescent Reach"], // TSS
// ["moors","Blightfire Moors"], // TSS
// ["stonehive","Stone Hive"], // TSS
// ["mesa","Goru`kar Mesa"], // TSS
// ["roost","Blackfeather Roost"], // TSS
// ["steppes","The Steppes"], // TSS
// ["icefall","Icefall Glacier"], // TSS
// ["valdeholm","Valdeholm"], // TSS
// ["frostcrypt","Frostcrypt, Throne of the Shade King"], // TSS
// ["sunderock","Sunderock Springs"], // TSS
// ["vergalid","Vergalid Mines"], // TSS
// ["direwind","Direwind Cliffs"], // TSS
// ["ashengate","Ashengate, Reliquary of the Scale"], // TSS
// ["kattacastrum","Katta Castrum"], // TBS
// ["thalassius","Thalassius, the Coral Keep"], // TBS
// ["atiiki","Jewel of Atiiki"], // TBS
// ["zhisza","Zhisza, the Shissar Sanctuary"], // TBS
// ["silyssar","Silyssar, New Chelsith"], // TBS
// ["solteris","Solteris, the Throne of Ro"], // TBS
// ["barren","Barren Coast"], // TBS
// ["buriedsea","The Buried Sea"], // TBS
// ["jardelshook","Jardel's Hook"], // TBS
// ["monkeyrock","Monkey Rock"], // TBS
// ["suncrest","Suncrest Isle"], // TBS
// ["deadbone","Deadbone Reef"], // TBS
// ["blacksail","Blacksail Folly"], // TBS
// ["maidensgrave","Maiden's Grave"], // TBS
// ["redfeather","Redfeather Isle"], // TBS
// ["shipmvp","The Open Sea (A)"], // TBS
// ["shipmvu","The Open Sea (B)"], // TBS
// ["shippvu","The Open Sea (C)"], // TBS
// ["shipuvu","The Open Sea (D)"], // TBS
// ["shipmvm","The Open Sea (E)"], // TBS
// ["mechanotus","Fortress Mechanotus"], // SoF
// ["mansion","Meldrath's Majestic Mansion"], // SoF
// ["steamfactory","The Steam Factory"], // SoF
// ["shipworkshop","S.H.I.P. Workshop"], // SoF
// ["gyrospireb","Gyrospire Beza"], // SoF
// ["gyrospirez","Gyrospire Zeka"], // SoF
// ["dragonscale","Dragonscale Hills"], // SoF
// ["lopingplains","Loping Plains"], // SoF
// ["hillsofshade","Hills of Shade"], // SoF
// ["bloodmoon","Bloodmoon Keep"], // SoF
// ["crystallos","Crystallos, Lair of the Awakened"], // SoF
// ["guardian","The Mechamatic Guardian"], // SoF
// ["cryptofshade","Crypt of Shade"], // SoF
// ["dragonscaleb","Deepscar's Den"], // SoF
// ["oldfieldofbone","Old Field of Scale"], // SoD
// ["oldkaesoraa","Kaesora Library"], // SoD
// ["oldkaesorab","Hatchery Wing"], // SoD
// ["oldkurn","Old Kurn's Tower"], // SoD
// ["oldkithicor","Bloody Kithicor"], // SoD
// ["oldcommons","Old Commonlands"], // SoD
// ["oldhighpass","Old Highpass Hold"], // SoD
// ["thevoida","The Void (A)"], // SoD
// ["thevoidb","The Void (B)"], // SoD
// ["thevoidc","The Void (C)"], // SoD
// ["thevoidd","The Void (D)"], // SoD
// ["thevoide","The Void (E)"], // SoD
// ["thevoidf","The Void (F)"], // SoD
// ["thevoidg","The Void (G)"], // SoD
// ["oceangreenhills","Oceangreen Hills"], // SoD
// ["oceangreenvillage","Oceangreen Village"], // SoD
// ["oldblackburrow","Old Blackburrow"], // SoD
// ["bertoxtemple","Temple of Bertoxxulous"], // SoD
// ["discord","Korafax, Home of the Riders"], // SoD
// ["discordtower","Citadel of the Worldslayer"], // SoD
// ["oldbloodfield","Old Bloodfields"], // SoD
// ["precipiceofwar","The Precipice of War"], // SoD
// ["olddranik","City of Dranik"], // SoD
// ["toskirakk","Toskirakk"], // SoD
// ["korascian","Korascian Warrens"], // SoD
// ["rathechamber","Rathe Council Chambers"], // SoD
// ["oldfieldofboneb","Field of Scale"], // SoD
// ["brellsrest","Brell's Rest"], // UF
// ["fungalforest","Fungal Forest"], // UF
// ["underquarry","The Underquarry"], // UF
// ["coolingchamber","The Cooling Chamber"], // UF
// ["shiningcity","Kernagir, The Shining City"], // UF
// ["arthicrex","Arthicrex"], // UF
// ["foundation","The Foundation"], // UF
// ["lichencreep","Lichen Creep"], // UF
// ["pellucid","Pellucid Grotto"], // UF
// ["stonesnake","Volska's Husk"], // UF
// ["brellstemple","Brell's Temple"], // UF
// ["convorteum","The Convorteum"], // UF
// ["brellsarena","Brell's Arena"], // UF
// ["crafthalls","Ngreth's Den"], // UF
// ["weddingchapel","Wedding Chapel"], // UF
// ["dragoncrypt","Lair of the Fallen"], // UF
// ["feerrott2","The Feerrott (B)"], // HoT
// ["thulehouse1","House of Thule"], // HoT
// ["thulehouse2","House of Thule, Upper Floors"], // HoT
// ["housegarden","The Grounds"], // HoT
// ["thulelibrary","The Library"], // HoT
// ["well","The Well"], // HoT
// ["fallen","Erudin Burning"], // HoT
// ["morellcastle","Morell's Castle"], // HoT
// ["morelltower", "Morell's Tower"], // HoT
// ["alkabormare","Al`Kabor's Nightmare"], // HoT
// ["miragulmare","Miragul's Nightmare"], // HoT
// ["thuledream","Fear Itself"], // HoT
// ["somnium","Sanctum Somnium"], // HoT
// ["neighborhood","Sunrise Hills"], // HoT
// ["phylactery","Miragul's Phylactery"], // HoT
// ["argath","Argath"], // HoT
// ["arelis","Valley of Lunanyn"], // HoT
// ["beastdomain","Beast's Domain"], // HoT
// ["cityofbronze","City of Bronze"], // HoT
// ["eastsepulcher","East Sepulcher"], // HoT
// ["sarithcity","Sarith City"], // HoT
// ["rubak","Rubak Oseka"], // HoT
// ["resplendent","Resplendent Temple"], // HoT
// ["pillarsalra","Pillars of Alra"], // HoT
// ["windsong","Windsong"], // HoT
// ["guildhalllrg","Palatial Guidhall"], // HoT
// ["sepulcher","Sepulcher of Order"], // HoT
// ["westsepulcher","West Sepulcher"], // HoT
// ["resplendent","Resplendent Temple"], // HoT
// ["shadowedmount","Shadowed Mount"], // HoT
// ["guildhalllrg","Grand Guild Hall"], // HoT
// ["guildhallsml","Greater Guild Hall"], // HoT
// ["plhogrinteriors1a1","One Bedroom House Interior"], // HoT
// ["plhogrinteriors1a2","One Bedroom House Interior"], // HoT
// ["plhogrinteriors3a1","Three Bedroom House Interior"], // HoT
// ["plhogrinteriors3a2","Three Bedroom House Interior"], // HoT
// ["plhogrinteriors3b1","Three Bedroom House Interior"], // HoT
// ["plhogrinteriors3b2","Three Bedroom House Interior"], // HoT
// ["plhdkeinteriors1a1","One Bedroom House Interior"], // HoT
// ["plhdkeinteriors1a2","One Bedroom House Interior"], // HoT
// ["plhdkeinteriors1a3","One Bedroom House Interior"], // HoT
// ["plhdkeinteriors3a1","Three Bedroom House Interior"], // HoT
// ["plhdkeinteriors3a2","Three Bedroom House Interior"], // HoT
// ["plhdkeinteriors3a3","Three Bedroom House Interior"], // HoT
// ["guildhall3","Modest Guild Hall"], // HoT
// ["kaelshard","Kael Drakkel: The King's Madness"], // RoF
// ["eastwastesshard","East Wastes: Zeixshi-Kar's Awakening"], // RoF
// ["crystalshard","The Crystal Caverns: Fragment of Fear"], // RoF
// ["shardslanding","Shard's Landing"], // RoF
// ["xorbb","Valley of King Xorbb"], // RoF
// ["breedinggrounds","The Breeding Grounds"], // RoF
// ["eviltree","Evantil, the Vile Oak"], // RoF
// ["grelleth","Grelleth's Palace, the Chateau of Filth"], // RoF
// ["chapterhouse","Chapterhouse of the Fallen"], // RoF
// ["phinteriortree","Evantil's Abode"], // RoF
// ["chelsithreborn","Chelsith Reborn"], // RoF
// ["poshadow","Plane of Shadow"], // RoF
// ["pomischief","The Plane of Mischief"], // RoF
// ["The Burned Woods","burnedwoods"], // RoF
// ["heartoffear","Heart of Fear: The Threshold"], // RoF
// ["heartoffearb","Heart of Fear: The Rebirth"], // RoF
// ["heartoffearc","Heart of Fear: The Epicenter"], // RoF
// ["thevoidh","The Void (H)"], // CoTF
// ["ethernere","Ethernere Tainted West Karana"], // CoTF
// ["neriakd","Neriak - Fourth Gate"], // CoTF
// ["deadhills","The Dead Hills"], // CoTF
// ["bixiewarfront","Bix Warfront"], // CoTF
// ["towerofrot","Tower of Rot"], // CoTF
// ["arginhiz","Argin-Hiz"], // CoTF
// ["arxmentis","Arx Mentis"], // TDS
// ["brotherisland","Brother Island"], // TDS
// ["endlesscaverns","Caverns of Endless Song"], // TDS
// ["dredge","Combine Dredge"], // TDS
// ["degmar","Degmar, the Lost Castle"], // TDS
// ["kattacastrumb","Katta Castrum, The Deluge"], // TDS
// ["tempesttemple","Tempest Temple"], // TDS
// ["thuliasaur","Thuliasaur Island"], // TDS
// ["exalted","Sul Vius: Demiplane of Life"], // TBM
// ["exaltedb","Sul Vius: Demiplane of Decay"], // TBM
// ["cosul","Crypt of Sul"], // TBM
// ["codecayb","Ruins of Lxanvom"], // TBM
// ["pohealth","The Plane of Health"], // TBM
// ["chardoktwo","Chardok"], // EoK
// ["frontiermtnsb","Frontier Mountains"], // EoK
// ["korshaext","Gates of Kor-Sha"], // EoK
// ["korshaint","Kor-Sha Laboratory"], // EoK
// ["lceanium","Lceanium"], // EoK
// ["scorchedwoods","Scorched Woods"], // EoK
// ["drogab","Temple of Droga"], // EoK
// ["charasisb","Sathir's Tomb"], // RoS
// ["gorowyn","Gorowyn"], // RoS
// ["charasistwo","Howling Stones"], // RoS
// ["skyfiretwo","Skyfire Mountains"], // RoS
// ["overtheretwo","The Overthere"], // RoS
// ["veeshantwo","Veeshan's Peak"], // RoS
// ["esianti","Esianti: Palace of the Winds"], // TBL
// ["trialsofsmoke","Plane of Smoke"], // TBL
// ["stratos","Stratos: Zephyr's Flight"], // TBL
// ["empyr","Empyr: Realms of Ash"], // TBL
// ["aalishai","AAlishai: Palace of Embers"], // TBL
// ["mearatas","Mearatas: The Stone Demesne"], // TBL
// ["chamberoftears","The Chamber of Tears"], // TBL
// ["gnomemtn","Gnome Memorial Mountain"], // TBL
// ["eastwastestwo","The Eastern Wastes"], // ToV
// ["frozenshadowtwo","The Tower of Frozen Shadow"], // ToV
// ["crystaltwoa","The Ry`Gorr Mines"], // ToV
// ["greatdividetwo","The Great Divide"], // ToV
// ["velketortwo","Velketor's Labyrinth"], // ToV
// ["kaeltwo","Kael Drakkel"], // ToV
// ["crystaltwob","Crystal Caverns"], // ToV
// ["sleepertwo","The Sleeper's Tomb"], // CoV
// ["necropolistwo","Dragon Necropolis"], // CoV
// ["cobaltscartwo","Cobalt Scar"], // CoV
// ["westwastestwo","The Western Wastes"], // CoV
// ["skyshrinetwo","Skyshrine"], // CoV
// ["templeveeshantwo","The Temple of Veeshan"], // CoV
// ["maidentwo","Maiden's Eye"], // ToL
// ["umbraltwo","Umbral Plains"], // ToL
// ["akhevatwo","Ka Vethan"], // ToL
// ["vexthaltwo","Vex Thal"], // ToL
// ["shadowvalley","Shadow Valley"], // ToL
// ["basilica","Basilica of Adumbration"], // ToL
// ["bloodfalls","Bloodfalls"], // ToL
// ["maidenhouseint","Coterie Chambers"], // ToL
// ["shadowhaventwo","Ruins of Shadow Haven"], // NoS
// ["sharvahltwo","Shar Vahl, Divided"], // NoS
// ["shadeweavertwo","Shadeweaver's Tangle"], // NoS
// ["paludaltwo","Paludal Caverns"], // NoS
// ["deepshade","Deepshade"], // NoS
// ["firefallpass","Firefall Pass"], // NoS
// ["hollowshadetwo","Hollowshade Moor"], // NoS
// ["darklightcaverns","Darklight Caverns"], // NoS
// ["laurioninn","Laurion Inn"], // LS
// ["timorousfalls","Timorous Falls"], // LS
// ["ankexfen","Ankexfen Keep"], // LS
// ["moorsofnokk","Moors of Nokk"], // LS
// ["unkemptwoods","Unkempt Woods"], // LS
// ["herosforge","The Hero's Forge"], // LS
// ["pallomen","Pal'Lomen"], // LS
]);

// bidirectional routes
const biRoutes = [
    ['qeytoqrg', 'qeynos2'],
    ['qeynos2', 'qeynos'],
    ['qeynos', 'erudsxing'],
    ['qeytoqrg', 'blackburrow'],
    ['blackburrow', 'everfrost'],
    ['blackburrow', 'jaggedpine'],
    ['jaggedpine', 'nedaria'],
    ['everfrost', 'halas'],
    ['everfrost', 'permafrost'],
    ['qeytoqrg', 'qey2hh1'],
    ['qey2hh1', 'northkarana'],
    ['northkarana', 'eastkarana'],
    ['eastkarana', 'southkarana'],
    ['southkarana', 'lakerathe'],
    ['lakerathe', 'arena'],
    ['lakerathe', 'rathemtn'],
    ['rathemtn', 'feerrott'],
    ['feerrott', 'fearplane'],
    ['feerrott', 'cazicthule'],
    ['qrg', 'qeytoqrg'],
    ['qeynos', 'qcat'],
    ['qeynos2', 'qcat'],
    ['erudsxing', 'erudnext'],
    ['erudnext', 'tox'],
    ['tox', 'kerraridge'],
    ['erudnext', 'erudnint'],
    ['feerrott', 'innothule'],
    ['innothule', 'grobb'],
    ['innothule', 'sro'],
    ['innothule', 'guktop'],
    ['guktop', 'gukbottom'],
    ['sro', 'oasis'],
    ['oasis', 'nro'],
    ['nro', 'freporte'],
    ['freporte', 'freportw'],
    ['freportw', 'freportn'],
    ['freportw', 'ecommons'],
    ['ecommons', 'commons'],
    ['commons', 'nektulos'],
];

// unidirectional (one way) routes
const uniRoutes = [
    'airplane', 'ecommons',
];


// special notes for connections
notes = new Map([
    ['qeynoserudsxing', 'use boat to'],
    ['erudsxingsqeynos', 'use boat to'],
]);

// The graph
const adjacencyList = new Map();

// Add node
function addNode(zone) {
    adjacencyList.set(zone, []);
}

// Add edge, undirected
function addBiEdge(origin, destination) {
    src = adjacencyList.get(origin)
    if (!src) {
        console.log("addBiEdge "+origin+", "+destination+": no src for "+origin);
        return
    }
    src.push(destination);
    dst = adjacencyList.get(destination)
    if (!dst) {
        console.log("addBiEdge "+origin+", "+destination+": no dst for "+destination);
        return
    }
    dst.push(origin);
}

function addEdge(origin, destination) {
    src = adjacencyList.get(origin)
    if (!src) {
        console.log("addEdge "+origin+", "+destination+": no src for "+origin);
        return
    }
    src.push(destination);
}

// Create the Graph
zones.forEach(function(value, key) {
    addNode(key);
});
biRoutes.forEach(route => addBiEdge(...route))
uniRoutes.forEach(route => addEdge(...route))

// special universal transport things

// Bazaar AA
for (const [key, value] of zones) {
     addEdge(key, "bazaar"); // every zone can get to bazaar
     notes.set(key+"bazaar", "use bazaar portal AA to");
}

// Tearel in bazaar
tearelZones = [
'akanon',
'cabeast', // kunark
'erudnext',
'felwithea',
'freportw',
'gfaydark',
'grobb',
'halas',
'kaladima',
'neriakb',
'oggok',
'paineel',
'qeynos2',
'rivervale',
'airplane',
'hateplane',
].forEach(function(value) {
    addEdge("bazaar", value);
    notes.set("bazaar"+value, "use Tearel to teleport to");
});


function bfs(start, end) {

    const visited = new Set();

    const queue = [start];
    var last = start;

    while (queue.length > 0) {

        const zone = queue.shift(); // mutates the queue

        let destinations = adjacencyList.get(zone);

        for (const destination of destinations) {
            if (destination === end)  {
                while (val = queue.pop()) {
                    adj = notes.get(last+val);
                    console.log(last+" "+(adj ? adj : "zone to")+" "+val)

                    last = val;
                }
                adj = notes.get(last+destination);
                console.log(last+" "+(adj ? adj : "zone to")+" "+end)
                return;

            }

            adj = notes.get(zone+destination);
            if (!isAA(adj) && !visited.has(destination)) {
                visited.add(destination);
                queue.push(destination);
            }
        }
    }

    console.log("no route to "+end);
    return;
}

function dfs(start, end, nav, visited = new Set()) {
    visited.add(start);

    const destinations = adjacencyList.get(start);
    if (!destinations) {
        console.log("no destinations for "+start);
        return false;
    }

    for (const destination of destinations) {
        if (destination === end)  {
            nav.push(destination);
            return true;
        }

        if (!visited.has(destination)) {
            if (dfs(destination, end, nav, visited)) {
                nav.push(destination);
                return true;
            }
        }
    }
    return false;
}

nav = [];
start = "qeytoqrg";
end = "erudsxing";
if (dfs(start, end, nav)) {
    src = start
    while (dst = nav.pop()) {
        adj = notes.get(start+dst);
        srcFullName = zones.get(src);
        dstFullName = zones.get(dst);

        console.log("from "+srcFullName+" ("+src+") "+(adj ? adj : "zone to")+" "+dstFullName+" ("+dst+")");
        src = dst;
    }
} else {
    console.log("no route to "+end);
}

let searchForm = document.getElementById("searchForm");

searchForm.addEventListener("submit", (e) => {
    e.preventDefault();
    let from = document.getElementById("from");
    let to = document.getElementById("to");

    nav = [];
    if (dfs(from.value, to.value, nav)) {
        src = from.value;
        out = "To get from "+zones.get(from.value)+" to "+zones.get(to.value)+":<br>";

        while (dst = nav.pop()) {
            adj = notes.get(src+dst);
            srcFullName = zones.get(src);
            dstFullName = zones.get(dst);

            out += "From "+srcFullName+" ("+src+") "+(adj ? adj : "zone to")+" "+dstFullName+" ("+dst+")<br>";
            src = dst;
        }
        out += "You are now in "+zones.get(to.value)+"!<br>";
        document.getElementById("results").innerHTML = out;
        return;
    }
    document.getElementById("results").innerHTML = "No route to "+to.value+" found.";
});

let zonesList = document.getElementById("zones");
zones.forEach(function(value, key) {
    let option = document.createElement("option");
    option.value = key;
    option.text = value;
    zonesList.appendChild(option);
});

</script>



{{</rawhtml>}}

