let isBazaarAllowed = true;

class World {
  constructor() {
    this.adjacencyList = {};
    this.fullNames = {};
    this.notes = {};
  }


  addZone(shortName, fullName) {
    if (!this.adjacencyList[shortName]) {
      this.adjacencyList[shortName] = [];
      this.fullNames[shortName] = fullName;
    }
  }

  addZoneLine(source, destination, weight, note = '') {
    this.adjacencyList[source].push({ node: destination, weight, note });
    this.notes[source + destination] = note;
  }

  addBiZoneLine(source, destination, weight, note = '') {
    this.adjacencyList[source].push({ node: destination, weight, note });
    this.adjacencyList[destination].push({ node: source, weight, note });
    this.notes[source + destination] = note;
    this.notes[destination + source] = note;
  }

  findShortestPath(start, end) {
    const nodes = new PriorityQueue();
    const distances = {};
    const previous = {};
    let shortestPath = []; // Stores the shortest path
    let smallest;

    // Initial state setup
    for (let node in this.adjacencyList) {
      if (node === start) {
        distances[node] = 0;
        nodes.enqueue(node, 0);
      } else {
        distances[node] = Infinity;
        nodes.enqueue(node, Infinity);
      }
      previous[node] = null;
    }

    // Main algorithm loop
    while (!nodes.isEmpty()) {
      smallest = nodes.dequeue().val;

      if (smallest === end) {
        // We have reached the destination
        // Construct the path by backtracking from the end
        while (previous[smallest]) {
          shortestPath.push(smallest);
          smallest = previous[smallest];
        }
        break;
      }

      if (smallest || distances[smallest] !== Infinity) {
        for (let neighbor in this.adjacencyList[smallest]) {
          let nextNode = this.adjacencyList[smallest][neighbor];

          let note = this.notes[smallest + nextNode.node];
          if (note == 'use Bazaar Portal AA to' && !isBazaarAllowed) {
            continue;
          }


          // Calculate new distance to neighboring node
          let candidate = distances[smallest] + nextNode.weight;
          let nextNeighbor = nextNode.node;
          if (candidate < distances[nextNeighbor]) {
            // Updating new smallest distance to neighbor
            distances[nextNeighbor] = candidate;
            // Updating previous - How we got to next neighbor
            previous[nextNeighbor] = smallest;
            // Enqueue in priority queue with new priority
            nodes.enqueue(nextNeighbor, candidate);
          }
        }
      }
    }

    return shortestPath.concat(smallest).reverse();
  }
}

class PriorityQueue {
  constructor() {
    this.values = [];
  }

  enqueue(val, priority) {
    this.values.push({ val, priority });
    this.sort();
  }

  dequeue() {
    return this.values.shift();
  }

  isEmpty() {
    return this.values.length === 0;
  }

  sort() {
    this.values.sort((a, b) => a.priority - b.priority);
  }
}

const w = new World();
w.addZone("qeynos", "South Qeynos"); // Classic
w.addZone("qeynos2", "North Qeynos"); // Classic
w.addZone("qrg", "Surefall Glade"); // Classic
w.addZone("qeytoqrg", "Qeynos Hills"); // Classic
w.addZone("highkeep", "HighKeep"); // Classic
w.addZone("freportn", "North Freeport"); // Classic
w.addZone("freportw", "West Freeport"); // Classic
w.addZone("freporte", "East Freeport"); // Classic
w.addZone("runnyeye", "Clan RunnyEye"); // Classic
w.addZone("qey2hh1", "West Karana"); // Classic
w.addZone("northkarana", "North Karana"); // Classic
w.addZone("southkarana", "South Karana"); // Classic
w.addZone("eastkarana", "East Karana"); // Classic
w.addZone("beholder", "Gorge of King Xorbb"); // Classic
w.addZone("blackburrow", "BlackBurrow"); // Classic
w.addZone("paw", "Infected Paw"); // Classic
w.addZone("rivervale", "Rivervale"); // Classic
w.addZone("kithicor", "Kithicor Forest (A)"); // Classic
w.addZone("commons", "West Commonlands"); // Classic
w.addZone("ecommons", "East Commonlands"); // Classic
w.addZone("erudnint", "Erudin Palace"); // Classic
w.addZone("erudnext", "Erudin"); // Classic
w.addZone("nektulos", "Nektulos Forest"); // Classic
w.addZone("cshome", "Sunset Home"); // Classic
w.addZone("lavastorm", "Lavastorm Mountains"); // Classic
w.addZone("nektropos", "Nektropos"); // Classic
w.addZone("halas", "Halas"); // Classic
w.addZone("everfrost", "Everfrost Peaks"); // Classic
w.addZone("soldunga", "Solusek's Eye"); // Classic
w.addZone("soldungb", "Nagafen's Lair"); // Classic
w.addZone("misty", "Misty Thicket"); // Classic
w.addZone("nro", "North Ro"); // Classic
w.addZone("sro",  "South Ro"); // Classic
//w.addZone("southro", "South Ro (A)"); // Classic v2
w.addZone("befallen", "Befallen"); // Classic
w.addZone("oasis", "Oasis of Marr"); // Classic
w.addZone("tox", "Toxxulia Forest"); // Classic
w.addZone("hole", "The Ruins of Old Paineel"); // Classic
w.addZone("neriaka", "Neriak Foreign Quarter"); // Classic
w.addZone("neriakb", "Neriak Commons"); // Classic
w.addZone("neriakc", "Neriak Third Gate"); // Classic
w.addZone("neriakd", "Neriak Palace"); // Classic
w.addZone("najena", "Najena"); // Classic
w.addZone("qcat", "Qeynos Catacombs"); // Classic
w.addZone("innothule", "Innothule Swamp"); // Classic
w.addZone("feerrott", "The Feerrott"); // Classic
w.addZone("cazicthule", "Cazic-Thule"); // Classic
w.addZone("oggok", "Oggok"); // Classic
w.addZone("rathemtn", "Mountains of Rathe"); // Classic
w.addZone("lakerathe", "Lake Rathetear"); // Classic
w.addZone("grobb", "Grobb"); // Classic
w.addZone("aviak", "Aviak Village"); // Classic
w.addZone("gfaydark", "The Greater Faydark"); // Classic
w.addZone("akanon", "Ak'Anon"); // Classic
w.addZone("steamfont", "Steamfont Mountains"); // Classic
w.addZone("lfaydark", "The Lesser Faydark"); // Classic
w.addZone("crushbone", "Clan Crushbone"); // Classic
w.addZone("mistmoore", "Castle Mistmoore"); // Classic
w.addZone("kaladima", "North Kaladim"); // Classic
w.addZone("felwithea", "Felwithe"); // Classic
w.addZone("felwitheb", "FelwitheB"); // Classic
w.addZone("unrest", "Estate of Unrest"); // Classic
w.addZone("kedge", "Kedge Keep"); // Classic
w.addZone("guktop", "Upper Guk"); // Classic
w.addZone("gukbottom", "Lower Guk"); // Classic
w.addZone("kaladimb", "South Kaladim"); // Classic
w.addZone("butcher", "Butcherblock Mountains"); // Classic
w.addZone("oot", "Ocean of Tears"); // Classic
w.addZone("cauldron", "Dagnor's Cauldron"); // Classic
w.addZone("airplane", "Plane of Sky"); // Classic
w.addZone("fearplane", "Plane of Fear"); // Classic
w.addZone("permafrost", "Permafrost Keep"); // Classic
w.addZone("kerraridge", "Kerra Isle"); // Classic
w.addZone("paineel", "Paineel"); // Classic
//w.addZone("hateplane", "The Plane of Hate"); // Classic
w.addZone("arena", "The Arena"); // Classic
w.addZone("soltemple", "Temple of Solusek Ro"); // Classic
w.addZone("erudsxing", "Erud's Crossing"); // Classic
w.addZone("stonebrunt", "Stonebrunt Mountains"); // Classic
w.addZone("warrens", "The Warrens"); // Classic
w.addZone("erudsxing2", "Marauder's Mire"); // Classic
w.addZone("bazaar", "The Bazaar"); // Classic
//w.addZone("bazaar2", "The Bazaar (2)"); // Classic
//w.addZone("arena2", "The Arena"); // Classic
w.addZone("jaggedpine", "The Jaggedpine Forest"); // Classic
w.addZone("nedaria", "Nedaria's Landing"); // Classic
//w.addZone("tutorial", "Tutorial Zone"); // Classic
//w.addZone("load", "Loading (A)"); // Classic
//w.addZone("load2", "Loading (B)"); // Classic
w.addZone("hateplaneb", "The Plane of Hate"); // Classic
w.addZone("shadowrest", "Shadowrest"); // Classic
//w.addZone("tutoriala", "The Mines of Gloomingdeep (A)"); // Classic
//w.addZone("tutorialb", "The Mines of Gloomingdeep (B)"); // Classic
//w.addZone("clz", "Loading (C)"); // Classic
//w.addZone("poknowledge", "Plane of Knowledge"); // Classic
w.addZone("soldungc", "The Caverns of Exile"); // Classic
//w.addZone("guildlobby", "The Guild Lobby"); // Classic
//w.addZone("barter", "The Barter Hall"); // Classic
//w.addZone("takishruins", "Ruins of Takish-Hiz"); // Classic
//w.addZone("freeporteast", "East Freeport"); // Classic v2
//w.addZone("freeportwest", "West Freeport"); // Classic v2
//w.addZone("freeportsewers", "Freeport Sewers"); // Classic v2
//w.addZone("northro", "North Ro (B)"); // Classic v2
//w.addZone("southro", "South Ro (B)"); // Classic v2
w.addZone("highpasshold", "Highpass Hold"); // Classic v2
//w.addZone("commonlands", "Commonlands"); // Classic v2
//w.addZone("oceanoftears", "Ocean Of Tears"); // Classic v2
//w.addZone("kithforest", "Kithicor Forest (B)"); // Classic v2
//w.addZone("befallenb", "Befallen (B)"); // Classic v2
//w.addZone("highpasskeep", "Highpass Keep"); // Classic v2
//w.addZone("innothuleb", "Innothule Swamp (B)"); // Classic v2
//w.addZone("toxxulia", "Toxxulia Forest"); // Classic v2
//w.addZone("mistythicket", "Misty Thicket (B)"); // Classic v2
w.addZone("steamfontmts", "Steamfont Mountains"); // Classic
w.addZone("dragonscalea", "Tinmizer's Wunderwerks"); // Classic
w.addZone("crafthalls", "Ngreth's Den"); // Classic
w.addZone("weddingchapel", "Wedding Chapel"); // Classic
w.addZone("weddingchapeldark", "Wedding Chapel"); // Classic
w.addZone("dragoncrypt", "Lair of the Fallen"); // Classic
w.addZone("arttest", "Art Testing Domain"); // Classic
w.addZone("fhalls", "The Forgotten Halls"); // Classic
w.addZone("apprentice", "Designer Apprentice"); // Classic

w.addZone("fieldofbone", "The Field of Bone"); // Kunark
w.addZone("warslikswood", "Warsliks Wood"); // Kunark
w.addZone("droga", "Temple of Droga"); // Kunark
w.addZone("cabwest", "West Cabilis"); // Kunark
w.addZone("swampofnohope", "Swamp of No Hope"); // Kunark
w.addZone("firiona", "Firiona Vie"); // Kunark
w.addZone("lakeofillomen", "Lake of Ill Omen"); // Kunark
w.addZone("dreadlands", "Dreadlands"); // Kunark
w.addZone("burningwood", "Burning Woods"); // Kunark
w.addZone("kaesora", "Kaesora"); // Kunark
w.addZone("sebilis", "Old Sebilis"); // Kunark
w.addZone("citymist", "City of Mist"); // Kunark
w.addZone("skyfire", "Skyfire Mountains"); // Kunark
w.addZone("frontiermtns", "Frontier Mountains"); // Kunark
w.addZone("overthere", "The Overthere"); // Kunark
w.addZone("emeraldjungle", "The Emerald Jungle"); // Kunark
w.addZone("trakanon", "Trakanon's Teeth"); // Kunark
w.addZone("timorous", "Timorous Deep"); // Kunark
w.addZone("kurn", "Kurn's Tower"); // Kunark
w.addZone("karnor", "Karnor's Castle"); // Kunark
w.addZone("chardok", "Chardok"); // Kunark
w.addZone("dalnir", "Dalnir"); // Kunark
w.addZone("charasis", "Howling Stones"); // Kunark
w.addZone("cabeast", "East Cabilis"); // Kunark
w.addZone("nurga", "Mines of Nurga"); // Kunark
w.addZone("veeshan", "Veeshan's Peak"); // Kunark
w.addZone("veksar", "Veksar"); // Kunark
w.addZone("chardokb", "The Halls of Betrayal"); // Kunark
w.addZone("gunthak", "Gulf of Gunthak"); // LoY
w.addZone("dulak", "Dulak's Harbor"); // LoY
w.addZone("torgiran", "Torgiran Mines"); // LoY
w.addZone("nadox", "Crypt of Nadox"); // LoY
w.addZone("hatesfury", "Hate's Fury, The Scorned Maiden"); // LoY
w.addZone("iceclad", "Iceclad Ocean"); // Scars of Velious
w.addZone("frozenshadow", "Tower of Frozen Shadow"); // Scars of Velious
w.addZone("velketor", "Velketor's Labyrinth"); // Scars of Velious
w.addZone("kael", "Kael Drakkal"); // Scars of Velious
w.addZone("skyshrine", "Skyshrine"); // Scars of Velious
w.addZone("thurgadina", "Thurgadin"); // Scars of Velious
w.addZone("eastwastes", "Eastern Wastes"); // Scars of Velious
w.addZone("cobaltscar", "Cobalt Scar"); // Scars of Velious
w.addZone("greatdivide", "Great Divide"); // Scars of Velious
w.addZone("wakening", "The Wakening Land"); // Scars of Velious
w.addZone("westwastes", "Western Wastes"); // Scars of Velious
w.addZone("crystal", "Crystal Caverns"); // Scars of Velious
w.addZone("necropolis", "Dragon Necropolis"); // Scars of Velious
w.addZone("templeveeshan", "Temple of Veeshan"); // Scars of Velious
w.addZone("sirens", "Siren's Grotto"); // Scars of Velious
w.addZone("mischiefplane", "Plane of Mischief"); // Scars of Velious
w.addZone("growthplane", "Plane of Growth"); // Scars of Velious
w.addZone("sleeper", "Sleeper's Tomb"); // Scars of Velious
w.addZone("thurgadinb", "Icewell Keep"); // Scars of Velious
// w.addZone("shadowhaven", "Shadow Haven"); // Luclin
// w.addZone("nexus", "The Nexus"); // Luclin
// w.addZone("echo", "Echo Caverns"); // Luclin
// w.addZone("acrylia", "Acrylia Caverns"); // Luclin
// w.addZone("sharvahl", "Shar Vahl"); // Luclin
// w.addZone("paludal", "Paludal Caverns"); // Luclin
// w.addZone("fungusgrove", "Fungus Grove"); // Luclin
// w.addZone("vexthal", "Vex Thal"); // Luclin
// w.addZone("sseru", "Sanctus Seru"); // Luclin
// w.addZone("katta", "Katta Castellum"); // Luclin
// w.addZone("netherbian", "Netherbian Lair"); // Luclin
// w.addZone("ssratemple", "Ssraeshza Temple"); // Luclin
// w.addZone("griegsend", "Grieg's End"); // Luclin
// w.addZone("thedeep", "The Deep"); // Luclin
// w.addZone("shadeweaver", "Shadeweaver's Thicket"); // Luclin
// w.addZone("hollowshade", "Hollowshade Moor"); // Luclin
// w.addZone("grimling", "Grimling Forest"); // Luclin
// w.addZone("mseru", "Marus Seru"); // Luclin
// w.addZone("letalis", "Mons Letalis"); // Luclin
// w.addZone("twilight", "The Twilight Sea"); // Luclin
// w.addZone("thegrey", "The Grey"); // Luclin
// w.addZone("tenebrous", "The Tenebrous Mountains"); // Luclin
// w.addZone("maiden", "The Maiden's Eye"); // Luclin
// w.addZone("dawnshroud", "Dawnshroud Peaks"); // Luclin
// w.addZone("scarlet", "The Scarlet Desert"); // Luclin
// w.addZone("umbral", "The Umbral Plains"); // Luclin
// w.addZone("akheva", "Akheva Ruins"); // Luclin
// w.addZone("codecay", "Ruins of Lxanvom"); // Planes of Power
// w.addZone("pojustice", "Plane of Justice"); // Planes of Power
// w.addZone("potranquility", "Plane of Tranquility"); // Planes of Power
// w.addZone("ponightmare", "Plane of Nightmare"); // Planes of Power
// w.addZone("podisease", "Plane of Disease"); // Planes of Power
// w.addZone("poinnovation", "Plane of Innovation"); // Planes of Power
// w.addZone("potorment", "Plane of Torment"); // Planes of Power
// w.addZone("povalor", "Plane of Valor"); // Planes of Power
// w.addZone("bothunder", "Torden, The Bastion of Thunder"); // Planes of Power
// w.addZone("postorms", "Plane of Storms"); // Planes of Power
// w.addZone("hohonora", "Halls of Honor"); // Planes of Power
// w.addZone("solrotower", "Solusek Ro's Tower"); // Planes of Power
// w.addZone("powar", "Plane of War"); // Planes of Power
// w.addZone("potactics", "Drunder, Fortress of Zek"); // Planes of Power
// w.addZone("poair", "Eryslai, the Kingdom of Wind"); // Planes of Power
// w.addZone("powater", "Reef of Coirnav"); // Planes of Power
// w.addZone("pofire", "Doomfire, The Burning Lands"); // Planes of Power
// w.addZone("poeartha", "Vegarlson, The Earthen Badlands"); // Planes of Power
// w.addZone("potimea", "Plane of Time (A)"); // Planes of Power
// w.addZone("hohonorb", "Temple of Marr (A)"); // Planes of Power
// w.addZone("nightmareb", "Lair of Terris Thule"); // Planes of Power
// w.addZone("poearthb", "Stronghold of the Twelve"); // Planes of Power
// w.addZone("potimeb", "Plane of Time (B)"); // Planes of Power
// w.addZone("gunthak", "Gulf of Gunthak"); // LoY
// w.addZone("dulak", "Dulak's Harbor"); // LoY
// w.addZone("torgiran", "Torgiran Mines"); // LoY
// w.addZone("nadox", "Crypt of Nadox"); // LoY
// w.addZone("hatesfury", "Hate's Fury, The Scorned Maiden"); // LoY
// w.addZone("guka", "The Cauldron of Lost Souls"); // LDoN
// w.addZone("ruja", "The Bloodied Quarries"); // LDoN
// w.addZone("taka", "The Sunken Library"); // LDoN
// w.addZone("mira", "The Silent Gallery"); // LDoN
// w.addZone("mmca", "The Forlorn Caverns"); // LDoN
// w.addZone("gukb", "The Drowning Crypt"); // LDoN
// w.addZone("rujb", "The Halls of War"); // LDoN
// w.addZone("takb", "The Shifting Tower"); // LDoN
// w.addZone("mirb", "The Maw of the Menagerie"); // LDoN
// w.addZone("mmcb", "The Dreary Grotto"); // LDoN
// w.addZone("gukc", "The Ancient Aqueducts"); // LDoN
// w.addZone("rujc", "The Wind Bridges"); // LDoN
// w.addZone("takc", "The Fading Temple"); // LDoN
// w.addZone("mirc", "The Spider Den"); // LDoN
// w.addZone("mmcc", "The Asylum of Invoked Stone"); // LDoN
// w.addZone("gukd", "The Mushroom Grove"); // LDoN
// w.addZone("rujd", "The Gladiator Pits"); // LDoN
// w.addZone("takd", "The Royal Observatory"); // LDoN
// w.addZone("mird", "The Hushed Banquet"); // LDoN
// w.addZone("mmcd", "The Chambers of Eternal Affliction"); // LDoN
// w.addZone("guke", "The Foreboding Prison"); // LDoN
// w.addZone("ruje", "The Drudge Hollows"); // LDoN
// w.addZone("take", "The River of Recollection"); // LDoN
// w.addZone("mire", "The Frosted Halls"); // LDoN
// w.addZone("mmce", "The Sepulcher of the Damned"); // LDoN
// w.addZone("gukf", "The Chapel of the Witnesses"); // LDoN
// w.addZone("rujf", "The Fortified Lair of the Taskmasters"); // LDoN
// w.addZone("takf", "The Sandfall Corridors"); // LDoN
// w.addZone("mirf", "The Forgotten Wastes"); // LDoN
// w.addZone("mmcf", "The Ritualistic Summoning Grounds"); // LDoN
// w.addZone("gukg", "The Root Garden"); // LDoN
// w.addZone("rujg", "The Hidden Vale"); // LDoN
// w.addZone("takg", "The Balancing Chamber"); // LDoN
// w.addZone("mirg", "The Heart of the Menagerie"); // LDoN
// w.addZone("mmcg", "The Cesspits of Putrescence"); // LDoN
// w.addZone("gukh", "The Accursed Sanctuary"); // LDoN
// w.addZone("rujh", "The Blazing Forge"); // LDoN
// w.addZone("takh", "The Sweeping Tides"); // LDoN
// w.addZone("mirh", "The Morbid Laboratory"); // LDoN
// w.addZone("mmch", "The Aisles of Blood"); // LDoN
// w.addZone("ruji", "The Arena of Chance"); // LDoN
// w.addZone("taki", "The Antiquated Palace"); // LDoN
// w.addZone("miri", "The Theater of Imprisoned Horrors"); // LDoN
// w.addZone("mmci", "The Halls of Sanguinary Rites"); // LDoN
// w.addZone("rujj", "The Barracks of War"); // LDoN
// w.addZone("takj", "The Prismatic Corridors"); // LDoN
// w.addZone("mirj", "The Grand Library"); // LDoN
// w.addZone("mmcj", "The Infernal Sanctuary"); // LDoN
// w.addZone("abysmal", "Abysmal Sea"); // GoD
// w.addZone("natimbi", "Natimbi, The Broken Shores"); // GoD
// w.addZone("qinimi", "Qinimi, Court of Nihilia"); // GoD
// w.addZone("riwwi", "Riwwi, Coliseum of Games"); // GoD
// w.addZone("barindu", "Barindu, Hanging Gardens"); // GoD
// w.addZone("ferubi", "Ferubi, Forgotten Temple of Taelosia"); // GoD
// w.addZone("snpool", "Sewers of Nihilia, Pool of Sludge"); // GoD
// w.addZone("snlair", "Sewers of Nihilia, Lair of Trapped Ones"); // GoD
// w.addZone("snplant", "Sewers of Nihilia, Purifying Plant"); // GoD
// w.addZone("sncrematory", "Sewers of Nihilia, the Crematory"); // GoD
// w.addZone("tipt", "Tipt, Treacherous Crags"); // GoD
// w.addZone("vxed", "Vxed, The Crumbling Caverns"); // GoD
// w.addZone("yxtta", "Yxtta, Pulpit of Exiles"); // GoD
// w.addZone("uqua", "Uqua, The Ocean God Chantry"); // GoD
// w.addZone("kodtaz", "Kod'Taz, Broken Trial Grounds"); // GoD
// w.addZone("ikkinz", "Ikkinz, Chambers of Destruction"); // GoD
// w.addZone("qvic", "Qvic, Prayer Grounds of Calling"); // GoD
// w.addZone("inktuta", "Inktu`Ta, The Unmasked Chapel"); // GoD
// w.addZone("txevu", "Txevu, Lair of the Elite"); // GoD
// w.addZone("tacvi", "Tacvi, Seat of the Slaver"); // GoD
// w.addZone("qvibc", "Qvic, the Hidden Vault"); // GoD
// w.addZone("wallofslaughter", "Wall of Slaughter"); // OoW
// w.addZone("bloodfields", "The Bloodfields"); // OoW
// w.addZone("draniksscar", "Dranik's Scar"); // OoW
// w.addZone("causeway", "Nobles' Causeway"); // OoW
// w.addZone("chambersa", "Muramite Proving Grounds (A)"); // OoW
// w.addZone("chambersb", "Muramite Proving Grounds (B)"); // OoW
// w.addZone("chambersc", "Muramite Proving Grounds (C)"); // OoW
// w.addZone("chambersd", "Muramite Proving Grounds (D)"); // OoW
// w.addZone("chamberse", "Muramite Proving Grounds (E)"); // OoW
// w.addZone("chambersf", "Muramite Proving Grounds (F)"); // OoW
// w.addZone("provinggrounds", "Muramite Proving Grounds"); // OoW
// w.addZone("anguish", "Asylum of Anguish"); // OoW
// w.addZone("dranikhollowsa", "Dranik's Hollows (A)"); // OoW
// w.addZone("dranikhollowsb", "Dranik's Hollows (B)"); // OoW
// w.addZone("dranikhollowsc", "Dranik's Hollows (C)"); // OoW
// w.addZone("dranikhollowsd", "Dranik's Hollows (D)"); // OoW
// w.addZone("dranikhollowse", "Dranik's Hollows (E)"); // OoW
// w.addZone("dranikhollowsf", "Dranik's Hollows (F)"); // OoW
// w.addZone("dranikhollowsg", "Dranik's Hollows (G)"); // OoW
// w.addZone("dranikhollowsh", "Dranik's Hollows (H)"); // OoW
// w.addZone("dranikhollowsi", "Dranik's Hollows (I)"); // OoW
// w.addZone("dranikhollowsj", "Dranik's Hollows (J)"); // OoW
// w.addZone("dranikcatacombsa", "Catacombs of Dranik (A)"); // OoW
// w.addZone("dranikcatacombsb", "Catacombs of Dranik (B)"); // OoW
// w.addZone("dranikcatacombsc", "Catacombs of Dranik (C)"); // OoW
// w.addZone("draniksewersa", "Sewers of Dranik (A)"); // OoW
// w.addZone("draniksewersb", "Sewers of Dranik (B)"); // OoW
// w.addZone("draniksewersc", "Sewers of Dranik (C)"); // OoW
// w.addZone("riftseekers", "Riftseekers' Sanctum"); // OoW
// w.addZone("harbingers", "Harbingers' Spire"); // OoW
// w.addZone("dranik", "The Ruined City of Dranik"); // OoW
// w.addZone("broodlands", "The Broodlands"); // DoN
// w.addZone("stillmoona", "Stillmoon Temple"); // DoN
// w.addZone("stillmoonb", "The Ascent"); // DoN
// w.addZone("thundercrest", "Thundercrest Isles"); // DoN
// w.addZone("delvea", "Lavaspinner's Lair"); // DoN
// w.addZone("delveb", "Tirranun's Delve"); // DoN
// w.addZone("thenest", "The Accursed Nest"); // DoN
// w.addZone("guildhall", "Guild Hall"); // DoN
// w.addZone("illsalin", "Ruins of Illsalin"); // DoDH
// w.addZone("illsalina", "Imperial Bazaar"); // DoDH
// w.addZone("illsalinb", "Temple of the Korlach"); // DoDH
// w.addZone("illsalinc", "The Nargilor Pits"); // DoDH
// w.addZone("dreadspire", "Dreadspire Keep"); // DoDH
// w.addZone("drachnidhive", "The Hive"); // DoDH
// w.addZone("drachnidhivea", "Living Larder"); // DoDH
// w.addZone("drachnidhiveb", "Coven of the Skinwalkers"); // DoDH
// w.addZone("drachnidhivec", "Queen Sendaii's Lair"); // DoDH
// w.addZone("westkorlach", "Stoneroot Falls"); // DoDH
// w.addZone("westkorlacha", "Chambers of Xill"); // DoDH
// w.addZone("westkorlachb", "Caverns of the Lost"); // DoDH
// w.addZone("westkorlachc", "Lair of the Korlach"); // DoDH
// w.addZone("eastkorlach", "Undershore"); // DoDH
// w.addZone("eastkorlacha", "Snarlstone Dens"); // DoDH
// w.addZone("shadowspine", "Shadowspine"); // DoDH
// w.addZone("corathus", "Corathus Creep"); // DoDH
// w.addZone("corathusa", "Sporali Caverns"); // DoDH
// w.addZone("corathusb", "Corathus Lair"); // DoDH
// w.addZone("nektulosa", "Shadowed Grove"); // DoDH
// w.addZone("arcstone", "Arcstone"); // PoR
// w.addZone("relic", "Relic"); // PoR
// w.addZone("skylance", "Skylance"); // PoR
// w.addZone("devastation", "The Devastation"); // PoR
// w.addZone("devastationa", "The Seething Wall"); // PoR
// w.addZone("rage", "Sverag, Stronghold of Rage"); // PoR
// w.addZone("ragea", "Razorthorn, Tower of Sullon Zek"); // PoR
// w.addZone("takishruinsa", "The Root of Ro"); // PoR
// w.addZone("elddar", "The Elddar Forest"); // PoR
// w.addZone("elddara", "Tunare's Shrine"); // PoR
// w.addZone("theater", "Theater of Blood"); // PoR
// w.addZone("theatera", "Deathknell, Tower of Dissonance"); // PoR
// w.addZone("freeportacademy", "Academy of Arcane Sciences"); // PoR
// w.addZone("freeporttemple", "Temple of Marr (B)"); // PoR
// w.addZone("freeportmilitia", "Freeport Militia House"); // PoR
// w.addZone("freeportarena", "Arena"); // PoR
// w.addZone("freeportcityhall", "City Hall"); // PoR
// w.addZone("freeporttheater", "Theater"); // PoR
// w.addZone("freeporthall", "Hall of Truth"); // PoR
// w.addZone("crescent", "Crescent Reach"); // TSS
// w.addZone("moors", "Blightfire Moors"); // TSS
// w.addZone("stonehive", "Stone Hive"); // TSS
// w.addZone("mesa", "Goru`kar Mesa"); // TSS
// w.addZone("roost", "Blackfeather Roost"); // TSS
// w.addZone("steppes", "The Steppes"); // TSS
// w.addZone("icefall", "Icefall Glacier"); // TSS
// w.addZone("valdeholm", "Valdeholm"); // TSS
// w.addZone("frostcrypt", "Frostcrypt, Throne of the Shade King"); // TSS
// w.addZone("sunderock", "Sunderock Springs"); // TSS
// w.addZone("vergalid", "Vergalid Mines"); // TSS
// w.addZone("direwind", "Direwind Cliffs"); // TSS
// w.addZone("ashengate", "Ashengate, Reliquary of the Scale"); // TSS
// w.addZone("kattacastrum", "Katta Castrum"); // TBS
// w.addZone("thalassius", "Thalassius, the Coral Keep"); // TBS
// w.addZone("atiiki", "Jewel of Atiiki"); // TBS
// w.addZone("zhisza", "Zhisza, the Shissar Sanctuary"); // TBS
// w.addZone("silyssar", "Silyssar, New Chelsith"); // TBS
// w.addZone("solteris", "Solteris, the Throne of Ro"); // TBS
// w.addZone("barren", "Barren Coast"); // TBS
// w.addZone("buriedsea", "The Buried Sea"); // TBS
// w.addZone("jardelshook", "Jardel's Hook"); // TBS
// w.addZone("monkeyrock", "Monkey Rock"); // TBS
// w.addZone("suncrest", "Suncrest Isle"); // TBS
// w.addZone("deadbone", "Deadbone Reef"); // TBS
// w.addZone("blacksail", "Blacksail Folly"); // TBS
// w.addZone("maidensgrave", "Maiden's Grave"); // TBS
// w.addZone("redfeather", "Redfeather Isle"); // TBS
// w.addZone("shipmvp", "The Open Sea (A)"); // TBS
// w.addZone("shipmvu", "The Open Sea (B)"); // TBS
// w.addZone("shippvu", "The Open Sea (C)"); // TBS
// w.addZone("shipuvu", "The Open Sea (D)"); // TBS
// w.addZone("shipmvm", "The Open Sea (E)"); // TBS
// w.addZone("mechanotus", "Fortress Mechanotus"); // SoF
// w.addZone("mansion", "Meldrath's Majestic Mansion"); // SoF
// w.addZone("steamfactory", "The Steam Factory"); // SoF
// w.addZone("shipworkshop", "S.H.I.P. Workshop"); // SoF
// w.addZone("gyrospireb", "Gyrospire Beza"); // SoF
// w.addZone("gyrospirez", "Gyrospire Zeka"); // SoF
// w.addZone("dragonscale", "Dragonscale Hills"); // SoF
// w.addZone("lopingplains", "Loping Plains"); // SoF
// w.addZone("hillsofshade", "Hills of Shade"); // SoF
// w.addZone("bloodmoon", "Bloodmoon Keep"); // SoF
// w.addZone("crystallos", "Crystallos, Lair of the Awakened"); // SoF
// w.addZone("guardian", "The Mechamatic Guardian"); // SoF
// w.addZone("cryptofshade", "Crypt of Shade"); // SoF
// w.addZone("dragonscaleb", "Deepscar's Den"); // SoF
// w.addZone("oldfieldofbone", "Old Field of Scale"); // SoD
// w.addZone("oldkaesoraa", "Kaesora Library"); // SoD
// w.addZone("oldkaesorab", "Hatchery Wing"); // SoD
// w.addZone("oldkurn", "Old Kurn's Tower"); // SoD
// w.addZone("oldkithicor", "Bloody Kithicor"); // SoD
// w.addZone("oldcommons", "Old Commonlands"); // SoD
// w.addZone("oldhighpass", "Old Highpass Hold"); // SoD
// w.addZone("thevoida", "The Void (A)"); // SoD
// w.addZone("thevoidb", "The Void (B)"); // SoD
// w.addZone("thevoidc", "The Void (C)"); // SoD
// w.addZone("thevoidd", "The Void (D)"); // SoD
// w.addZone("thevoide", "The Void (E)"); // SoD
// w.addZone("thevoidf", "The Void (F)"); // SoD
// w.addZone("thevoidg", "The Void (G)"); // SoD
// w.addZone("oceangreenhills", "Oceangreen Hills"); // SoD
// w.addZone("oceangreenvillage", "Oceangreen Village"); // SoD
// w.addZone("oldblackburrow", "Old Blackburrow"); // SoD
// w.addZone("bertoxtemple", "Temple of Bertoxxulous"); // SoD
// w.addZone("discord", "Korafax, Home of the Riders"); // SoD
// w.addZone("discordtower", "Citadel of the Worldslayer"); // SoD
// w.addZone("oldbloodfield", "Old Bloodfields"); // SoD
// w.addZone("precipiceofwar", "The Precipice of War"); // SoD
// w.addZone("olddranik", "City of Dranik"); // SoD
// w.addZone("toskirakk", "Toskirakk"); // SoD
// w.addZone("korascian", "Korascian Warrens"); // SoD
// w.addZone("rathechamber", "Rathe Council Chambers"); // SoD
// w.addZone("oldfieldofboneb", "Field of Scale"); // SoD
// w.addZone("brellsrest", "Brell's Rest"); // UF
// w.addZone("fungalforest", "Fungal Forest"); // UF
// w.addZone("underquarry", "The Underquarry"); // UF
// w.addZone("coolingchamber", "The Cooling Chamber"); // UF
// w.addZone("shiningcity", "Kernagir, The Shining City"); // UF
// w.addZone("arthicrex", "Arthicrex"); // UF
// w.addZone("foundation", "The Foundation"); // UF
// w.addZone("lichencreep", "Lichen Creep"); // UF
// w.addZone("pellucid", "Pellucid Grotto"); // UF
// w.addZone("stonesnake", "Volska's Husk"); // UF
// w.addZone("brellstemple", "Brell's Temple"); // UF
// w.addZone("convorteum", "The Convorteum"); // UF
// w.addZone("brellsarena", "Brell's Arena"); // UF
// w.addZone("crafthalls", "Ngreth's Den"); // UF
// w.addZone("weddingchapel", "Wedding Chapel"); // UF
// w.addZone("dragoncrypt", "Lair of the Fallen"); // UF
// w.addZone("feerrott2", "The Feerrott (B)"); // HoT
// w.addZone("thulehouse1", "House of Thule"); // HoT
// w.addZone("thulehouse2", "House of Thule, Upper Floors"); // HoT
// w.addZone("housegarden", "The Grounds"); // HoT
// w.addZone("thulelibrary", "The Library"); // HoT
// w.addZone("well", "The Well"); // HoT
// w.addZone("fallen", "Erudin Burning"); // HoT
// w.addZone("morellcastle", "Morell's Castle"); // HoT
// w.addZone("morelltower",  "Morell's Tower"); // HoT
// w.addZone("alkabormare", "Al`Kabor's Nightmare"); // HoT
// w.addZone("miragulmare", "Miragul's Nightmare"); // HoT
// w.addZone("thuledream", "Fear Itself"); // HoT
// w.addZone("somnium", "Sanctum Somnium"); // HoT
// w.addZone("neighborhood", "Sunrise Hills"); // HoT
// w.addZone("phylactery", "Miragul's Phylactery"); // HoT
// w.addZone("argath", "Argath"); // HoT
// w.addZone("arelis", "Valley of Lunanyn"); // HoT
// w.addZone("beastdomain", "Beast's Domain"); // HoT
// w.addZone("cityofbronze", "City of Bronze"); // HoT
// w.addZone("eastsepulcher", "East Sepulcher"); // HoT
// w.addZone("sarithcity", "Sarith City"); // HoT
// w.addZone("rubak", "Rubak Oseka"); // HoT
// w.addZone("resplendent", "Resplendent Temple"); // HoT
// w.addZone("pillarsalra", "Pillars of Alra"); // HoT
// w.addZone("windsong", "Windsong"); // HoT
// w.addZone("guildhalllrg", "Palatial Guidhall"); // HoT
// w.addZone("sepulcher", "Sepulcher of Order"); // HoT
// w.addZone("westsepulcher", "West Sepulcher"); // HoT
// w.addZone("resplendent", "Resplendent Temple"); // HoT
// w.addZone("shadowedmount", "Shadowed Mount"); // HoT
// w.addZone("guildhalllrg", "Grand Guild Hall"); // HoT
// w.addZone("guildhallsml", "Greater Guild Hall"); // HoT
// w.addZone("plhogrinteriors1a1", "One Bedroom House Interior"); // HoT
// w.addZone("plhogrinteriors1a2", "One Bedroom House Interior"); // HoT
// w.addZone("plhogrinteriors3a1", "Three Bedroom House Interior"); // HoT
// w.addZone("plhogrinteriors3a2", "Three Bedroom House Interior"); // HoT
// w.addZone("plhogrinteriors3b1", "Three Bedroom House Interior"); // HoT
// w.addZone("plhogrinteriors3b2", "Three Bedroom House Interior"); // HoT
// w.addZone("plhdkeinteriors1a1", "One Bedroom House Interior"); // HoT
// w.addZone("plhdkeinteriors1a2", "One Bedroom House Interior"); // HoT
// w.addZone("plhdkeinteriors1a3", "One Bedroom House Interior"); // HoT
// w.addZone("plhdkeinteriors3a1", "Three Bedroom House Interior"); // HoT
// w.addZone("plhdkeinteriors3a2", "Three Bedroom House Interior"); // HoT
// w.addZone("plhdkeinteriors3a3", "Three Bedroom House Interior"); // HoT
// w.addZone("guildhall3", "Modest Guild Hall"); // HoT
// w.addZone("kaelshard", "Kael Drakkel: The King's Madness"); // RoF
// w.addZone("eastwastesshard", "East Wastes: Zeixshi-Kar's Awakening"); // RoF
// w.addZone("crystalshard", "The Crystal Caverns: Fragment of Fear"); // RoF
// w.addZone("shardslanding", "Shard's Landing"); // RoF
// w.addZone("xorbb", "Valley of King Xorbb"); // RoF
// w.addZone("breedinggrounds", "The Breeding Grounds"); // RoF
// w.addZone("eviltree", "Evantil, the Vile Oak"); // RoF
// w.addZone("grelleth", "Grelleth's Palace, the Chateau of Filth"); // RoF
// w.addZone("chapterhouse", "Chapterhouse of the Fallen"); // RoF
// w.addZone("phinteriortree", "Evantil's Abode"); // RoF
// w.addZone("chelsithreborn", "Chelsith Reborn"); // RoF
// w.addZone("poshadow", "Plane of Shadow"); // RoF
// w.addZone("pomischief", "The Plane of Mischief"); // RoF
// w.addZone("The Burned Woods", "burnedwoods"); // RoF
// w.addZone("heartoffear", "Heart of Fear: The Threshold"); // RoF
// w.addZone("heartoffearb", "Heart of Fear: The Rebirth"); // RoF
// w.addZone("heartoffearc", "Heart of Fear: The Epicenter"); // RoF
// w.addZone("thevoidh", "The Void (H)"); // CoTF
// w.addZone("ethernere", "Ethernere Tainted West Karana"); // CoTF
// w.addZone("neriakd", "Neriak - Fourth Gate"); // CoTF
// w.addZone("deadhills", "The Dead Hills"); // CoTF
// w.addZone("bixiewarfront", "Bix Warfront"); // CoTF
// w.addZone("towerofrot", "Tower of Rot"); // CoTF
// w.addZone("arginhiz", "Argin-Hiz"); // CoTF
// w.addZone("arxmentis", "Arx Mentis"); // TDS
// w.addZone("brotherisland", "Brother Island"); // TDS
// w.addZone("endlesscaverns", "Caverns of Endless Song"); // TDS
// w.addZone("dredge", "Combine Dredge"); // TDS
// w.addZone("degmar", "Degmar, the Lost Castle"); // TDS
// w.addZone("kattacastrumb", "Katta Castrum, The Deluge"); // TDS
// w.addZone("tempesttemple", "Tempest Temple"); // TDS
// w.addZone("thuliasaur", "Thuliasaur Island"); // TDS
// w.addZone("exalted", "Sul Vius: Demiplane of Life"); // TBM
// w.addZone("exaltedb", "Sul Vius: Demiplane of Decay"); // TBM
// w.addZone("cosul", "Crypt of Sul"); // TBM
// w.addZone("codecayb", "Ruins of Lxanvom"); // TBM
// w.addZone("pohealth", "The Plane of Health"); // TBM
// w.addZone("chardoktwo", "Chardok"); // EoK
// w.addZone("frontiermtnsb", "Frontier Mountains"); // EoK
// w.addZone("korshaext", "Gates of Kor-Sha"); // EoK
// w.addZone("korshaint", "Kor-Sha Laboratory"); // EoK
// w.addZone("lceanium", "Lceanium"); // EoK
// w.addZone("scorchedwoods", "Scorched Woods"); // EoK
// w.addZone("drogab", "Temple of Droga"); // EoK
// w.addZone("charasisb", "Sathir's Tomb"); // RoS
// w.addZone("gorowyn", "Gorowyn"); // RoS
// w.addZone("charasistwo", "Howling Stones"); // RoS
// w.addZone("skyfiretwo", "Skyfire Mountains"); // RoS
// w.addZone("overtheretwo", "The Overthere"); // RoS
// w.addZone("veeshantwo", "Veeshan's Peak"); // RoS
// w.addZone("esianti", "Esianti: Palace of the Winds"); // TBL
// w.addZone("trialsofsmoke", "Plane of Smoke"); // TBL
// w.addZone("stratos", "Stratos: Zephyr's Flight"); // TBL
// w.addZone("empyr", "Empyr: Realms of Ash"); // TBL
// w.addZone("aalishai", "AAlishai: Palace of Embers"); // TBL
// w.addZone("mearatas", "Mearatas: The Stone Demesne"); // TBL
// w.addZone("chamberoftears", "The Chamber of Tears"); // TBL
// w.addZone("gnomemtn", "Gnome Memorial Mountain"); // TBL
// w.addZone("eastwastestwo", "The Eastern Wastes"); // ToV
// w.addZone("frozenshadowtwo", "The Tower of Frozen Shadow"); // ToV
// w.addZone("crystaltwoa", "The Ry`Gorr Mines"); // ToV
// w.addZone("greatdividetwo", "The Great Divide"); // ToV
// w.addZone("velketortwo", "Velketor's Labyrinth"); // ToV
// w.addZone("kaeltwo", "Kael Drakkel"); // ToV
// w.addZone("crystaltwob", "Crystal Caverns"); // ToV
// w.addZone("sleepertwo", "The Sleeper's Tomb"); // CoV
// w.addZone("necropolistwo", "Dragon Necropolis"); // CoV
// w.addZone("cobaltscartwo", "Cobalt Scar"); // CoV
// w.addZone("westwastestwo", "The Western Wastes"); // CoV
// w.addZone("skyshrinetwo", "Skyshrine"); // CoV
// w.addZone("templeveeshantwo", "The Temple of Veeshan"); // CoV
// w.addZone("maidentwo", "Maiden's Eye"); // ToL
// w.addZone("umbraltwo", "Umbral Plains"); // ToL
// w.addZone("akhevatwo", "Ka Vethan"); // ToL
// w.addZone("vexthaltwo", "Vex Thal"); // ToL
// w.addZone("shadowvalley", "Shadow Valley"); // ToL
// w.addZone("basilica", "Basilica of Adumbration"); // ToL
// w.addZone("bloodfalls", "Bloodfalls"); // ToL
// w.addZone("maidenhouseint", "Coterie Chambers"); // ToL
// w.addZone("shadowhaventwo", "Ruins of Shadow Haven"); // NoS
// w.addZone("sharvahltwo", "Shar Vahl, Divided"); // NoS
// w.addZone("shadeweavertwo", "Shadeweaver's Tangle"); // NoS
// w.addZone("paludaltwo", "Paludal Caverns"); // NoS
// w.addZone("deepshade", "Deepshade"); // NoS
// w.addZone("firefallpass", "Firefall Pass"); // NoS
// w.addZone("hollowshadetwo", "Hollowshade Moor"); // NoS
// w.addZone("darklightcaverns", "Darklight Caverns"); // NoS
// w.addZone("laurioninn", "Laurion Inn"); // LS
// w.addZone("timorousfalls", "Timorous Falls"); // LS
// w.addZone("ankexfen", "Ankexfen Keep"); // LS
// w.addZone("moorsofnokk", "Moors of Nokk"); // LS
// w.addZone("unkemptwoods", "Unkempt Woods"); // LS
// w.addZone("herosforge", "The Hero's Forge"); // LS
// w.addZone("pallomen", "Pal'Lomen"); // LS

w.addBiZoneLine('qeytoqrg', 'qeynos2', 2);
w.addBiZoneLine('qeynos2', 'qeynos', 2);
w.addBiZoneLine('qeynos', 'erudsxing', 2);
w.addBiZoneLine('qeytoqrg', 'blackburrow', 2);
w.addBiZoneLine('blackburrow', 'everfrost', 2);
w.addBiZoneLine('blackburrow', 'jaggedpine', 2);
w.addBiZoneLine('jaggedpine', 'nedaria', 2);
w.addBiZoneLine('everfrost', 'halas', 2);
w.addBiZoneLine('everfrost', 'permafrost', 2);
w.addBiZoneLine('qeytoqrg', 'qey2hh1', 2);
w.addBiZoneLine('qey2hh1', 'northkarana', 2);
w.addBiZoneLine('northkarana', 'eastkarana', 2);
w.addBiZoneLine('northkarana', 'southkarana', 2);
w.addBiZoneLine('southkarana', 'paw', 2);
w.addBiZoneLine('southkarana', 'lakerathe', 2);
w.addBiZoneLine('eastkarana', 'highpasshold', 2);
w.addBiZoneLine('highpasshold', 'highkeep', 2);
w.addBiZoneLine('highpasshold', 'kithicor', 2);
w.addBiZoneLine('kithicor', 'rivervale', 2);
w.addBiZoneLine('kithicor', 'commons', 2);
w.addBiZoneLine('rivervale', 'misty', 2);
w.addBiZoneLine('misty', 'runnyeye', 2);
w.addBiZoneLine('runnyeye', 'beholder', 2);
w.addBiZoneLine('eastkarana', 'beholder', 2);
w.addBiZoneLine('lakerathe', 'arena', 2);
w.addBiZoneLine('lakerathe', 'rathemtn', 2);
w.addBiZoneLine('rathemtn', 'feerrott', 2);
w.addBiZoneLine('feerrott', 'fearplane', 2);
w.addBiZoneLine('feerrott', 'cazicthule', 2);
w.addBiZoneLine('qrg', 'qeytoqrg', 2);
w.addBiZoneLine('qeynos', 'qcat', 2);
w.addBiZoneLine('qeynos2', 'qcat', 2);
w.addBiZoneLine('erudsxing', 'erudnext', 2);
w.addBiZoneLine('erudnext', 'tox', 2);
w.addBiZoneLine('tox', 'kerraridge', 2);
w.addBiZoneLine('erudnext', 'erudnint', 2);
w.addBiZoneLine('feerrott', 'innothule', 2);
w.addBiZoneLine('innothule', 'grobb', 2);
w.addBiZoneLine('innothule', 'sro', 2);
w.addBiZoneLine('innothule', 'guktop', 2);
w.addBiZoneLine('guktop', 'gukbottom', 2);
w.addBiZoneLine('sro', 'oasis', 2);
w.addBiZoneLine('oasis', 'nro', 2);
w.addBiZoneLine('nro', 'freporte', 2);
w.addBiZoneLine('freporte', 'freportw', 2);
w.addBiZoneLine('freportw', 'freportn', 2);
w.addBiZoneLine('freportw', 'ecommons', 2);
w.addBiZoneLine('ecommons', 'commons', 2);
w.addBiZoneLine('ecommons', 'nektulos', 2);
w.addBiZoneLine('ecommons', 'nro', 2);
w.addBiZoneLine('nektulos', 'lavastorm', 2);
w.addBiZoneLine('lavastorm', 'soldungb', 2);
w.addBiZoneLine('lavastorm', 'soldunga', 2);
w.addBiZoneLine('lavastorm', 'soltemple', 2);
w.addBiZoneLine('lavastorm', 'najena', 2);
w.addBiZoneLine('freporte', 'oot', 2);
w.addBiZoneLine('oot', 'butcher', 2);
w.addBiZoneLine('butcher', 'kaladimb', 2);
w.addBiZoneLine('kaladimb', 'kaladima', 2);
w.addBiZoneLine('butcher', 'cauldron', 2);
w.addBiZoneLine('butcher', 'gfaydark', 2);
w.addBiZoneLine('gfaydark', 'crushbone', 2);
w.addBiZoneLine('cauldron', 'unrest', 2);
w.addBiZoneLine('gfaydark', 'felwithea', 2);
w.addBiZoneLine('felwithea', 'felwitheb', 2);
w.addBiZoneLine('oasis', 'timorous');

w.addBiZoneLine('timorous', 'overthere');
w.addBiZoneLine('overthere', 'skyfire');
w.addBiZoneLine('skyfire', 'veeshan');
w.addBiZoneLine('skyfire', 'burningwood');
w.addBiZoneLine('burningwood', 'chardok');
w.addBiZoneLine('chardok', 'chardokb');
w.addBiZoneLine('burningwood', 'dreadlands');

w.addBiZoneLine('oasis', 'iceclad');
w.addBiZoneLine('iceclad', 'frozenshadow');
w.addBiZoneLine('iceclad', 'eastwastes');
w.addBiZoneLine('eastwastes', 'greatdivide');
w.addBiZoneLine('greatdivide', 'thurgadina');
w.addBiZoneLine('thurgadina', 'thurgadinb');
w.addBiZoneLine('greatdivide', 'velketor');
w.addBiZoneLine('greatdivide', 'mischiefplane');
w.addBiZoneLine('eastwastes', 'kael');
w.addBiZoneLine('kael', 'wakening');
w.addBiZoneLine('wakening', 'growthplane');
w.addBiZoneLine('wakening', 'skyshrine');
w.addBiZoneLine('skyshrine', 'cobaltscar');
w.addBiZoneLine('cobaltscar', 'sirens');
w.addBiZoneLine('sirens', 'westwastes');
w.addBiZoneLine('westwastes', 'necropolis');
w.addBiZoneLine('westwastes', 'templeveeshan');
w.addBiZoneLine('eastwastes', 'crystal');


w.addZoneLine('airplane', 'ecommons', 2, 'jump off island to');

const adjacencyListKeys = Object.keys(w.adjacencyList);
for (let zone of adjacencyListKeys) {
  if (zone ==  'bazaar') continue;
  w.addZoneLine(zone, 'bazaar', 2, 'use Bazaar Portal AA to');
}

['akanon', 'blackburrow', 'burningwood', 'cabeast', 'cauldron', 'cauldron', 'cazicthule', 'citymist', 'commons', 'dreadlands', 'erudnext', 'erudsxing', 'felwithea', 'fieldofbone', 'firiona', 'freportw', 'gfaydark', 'grobb', 'halas', 'hateplaneb', 'highpasshold', 'hole', 'kaladima', 'lakeofillomen', 'lavastorm', 'mistmoore', 'neriakb', 'northkarana', 'oasis', 'oggok', 'oot', 'overthere', 'paineel', 'permafrost', 'qeynos2', 'qrg', 'rathemtn', 'rivervale', 'soldungb', 'southkarana', 'trakanon', ].forEach(function(zone) {
  if (zone ==  'bazaar') return;
  w.addZoneLine('bazaar', zone, 2, 'use Tearel to teleport to');
});


let searchForm = document.getElementById("searchForm");

searchForm.addEventListener("submit", (e) => {
    e.preventDefault();
    let from = document.getElementById("from");
    let to = document.getElementById("to");
    let isBazaarChk = document.getElementById("isBazaarPortalAllowed");
    isBazaarAllowed = isBazaarChk.checked;

    if (from.value === "" || to.value === "" || from.value === undefined || to.value === undefined) {
      document.getElementById("results").innerHTML = "Please select a starting and ending zone.";
      return;
    }


    if (from.value === to.value) {
      document.getElementById("results").innerHTML = "You are already in <b>"+w.fullNames[from.value]+"</b>!";
      return;
    }


    let nav = w.findShortestPath(from.value, to.value);
    let out = ""
    out += "<ol>";
    if (nav.length === 0) {
      document.getElementById("results").innerHTML = "No route to from <b>"+w.fullNames[from.value]+"</b> to <b>"+w.fullNames[to.value]+"</b> found.";
      return;
    }

    out = "To get from <b>"+w.fullNames[from.value]+"</b> to <b>"+w.fullNames[to.value]+"</b>:<br>";
    out += "<ol>"

    let src = ""
    let adj = ""
    for (let i = 0; i < nav.length; i++) {
      if (src == "") {
        src = nav[i];
        continue;
      }
      adj = w.notes[src+nav[i]];
      if (adj == "") {
        adj = "zone to";
      }
      out += "<li>From <b>"+w.fullNames[src]+" ("+src+")</b> "+adj+" <b>"+w.fullNames[nav[i]]+" ("+nav[i]+")</b></li>";
      src = nav[i];
      adj = w.notes[""+src+nav[i]];

    }
    out += "</ol>"
    out += "You are now in <b>"+w.fullNames[to.value]+"</b>!<br>";
    document.getElementById("results").innerHTML = out;
});

let zonesList = document.getElementById("zones");

for (let zone of adjacencyListKeys) {
  let option = document.createElement("option");
  option.value = zone;
  option.text = w.fullNames[zone];
  zonesList.appendChild(option);
}

// const shortestPath = w.findShortestPath('qeynos', 'qeynos2');
// console.log(shortestPath);
// for (let shortName of shortestPath) {
//   console.log(w.fullNames[shortName], shortName);
// }

