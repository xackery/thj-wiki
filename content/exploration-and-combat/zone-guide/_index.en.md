---
title: "Zone Guide"
images: [zone-guide.png]
weight: 3
description: "Discover how to get from zone A to B."
---

![Zone Guide](/images/zone-guide.webp)

{{<rawhtml>}}
<div class="container">
  <form action="" id="searchForm">
    <h1>Search For Zone Links</h1>
    <input type="checkbox" id="isBazaarPortalAllowed" name="isBazaarPortalAllowed" checked>
    <label for="isBazaarPortalAllowed">Allow Bazaar Portal</label><br>

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
<script src="zone-guide.js"></script>
{{</rawhtml>}}


