---
title: "Map"
weight: 3
---
![Map](images/map.png)

Insert interactive map here

<!--more-->
{{<rawhtml>}}<pre class="mermaid">
            graph LR
            sfg[Surefall Glade] <--> qeynos[Qeynos Hills]
            qeynos <--> blackburrow[Blackburrow]
            qeynos <--> nqeynos[North Qeynos]
            nqeynos <--> sqeynos[South Qeynos]
            sqeynos <-. boat .-> erudsxing[Erud's Crossing]
            blackburrow <--> everfrost[Everfrost Peaks]
            everfrost <--> halas[Halas]
            everfrost <--> permafrost[Permafrost Keep]
            qeynos <--> wkarana[West Karana]
            wkarana <--> nkarana[North Karana]
            nkarana <--> ekarana[East Karana]
            nkarana <--> skarana[South Karana]
            nkarana <-.spire.-> nexus[Nexus]
</pre>{{</rawhtml>}}    

    {{<rawhtml>}}<script type="module">
      import mermaid from 'https://cdn.jsdelivr.net/npm/mermaid@10/dist/mermaid.esm.min.mjs';
      mermaid.initialize({ startOnLoad: true });
    </script>{{</rawhtml>}}  
