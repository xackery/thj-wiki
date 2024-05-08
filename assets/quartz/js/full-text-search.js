; (async function() {
  const encoder = (str) => str.toLowerCase().split(/([^a-z]|[^\x00-\x7F])/)
  const contentIndex = new FlexSearch.Document({
    cache: true,
    charset: "latin:extra",
    optimize: true,
    index: [
      {
        field: "content",
        tokenize: "reverse",
        encode: encoder,
      },
      {
        field: "title",
        tokenize: "forward",
        encode: encoder,
      },
    ],
  })

  const { content } = await fetchData
  for (const [key, value] of Object.entries(content)) {
    contentIndex.add({
      id: key,
      title: value.title,
      content: removeMarkdown(value.content),
    })
  }

  const formatForDisplay = (id) => ({
    id,
    url: id,
    title: content[id].title,
    content: content[id].content,
  })

  registerHandlers((e) => {
    const term = e.target.value
    let lunrResp = lunrSearch(term)
    if (lunrResp.length < 3) {
      lunrResp = lunrSearch("*"+term+"*")
    }
    const searchResults = contentIndex.search(term, [
      {
        field: "content",
        limit: 10,
      },
      {
        field: "title",
        limit: 5,
      },
    ])
    const getByField = (field) => {
      const results = searchResults.filter((x) => x.field === field)
      if (results.length === 0) {
        return []
      } else {
        return [...results[0].result]
      }
    }
    const allIds = new Set([...getByField("title"), ...getByField("content")])
    const finalResults = [...allIds].map(formatForDisplay)

    //console.log("finalResults:",finalResults)
    //console.log("lunrResp", lunrResp)

    displayLunrResults(term, lunrResp, true)
  })
})()

// index lunr
var idx;
var searchDatabase = {};

var feedLoaded = function (e) {
  
    searchDatabase = JSON.parse(e.target.response)

    idx = lunr(function() {        
        this.ref("id")
        this.field("href")
        this.field("title", { boost: 10 })
        this.field("content", { boost: 1 })
        for (var i = 0; i < searchDatabase.length; i++) {
            let entry = searchDatabase[i] 
            
            this.add({
                id: i,
                href: entry.href,
                title: entry.title,
                content: entry.content,
            })
        }
    })
    
    //console.log('idx ready')
}


{{ $searchDataFile := printf "%s.search-data.json" .Language.Lang }}
{{ $searchData := resources.Get "search-data.json" | resources.ExecuteAsTemplate $searchDataFile . | resources.Minify | resources.Fingerprint }}
const searchDataURL = '{{ $searchData.RelPermalink }}';  

var xhr = new XMLHttpRequest

xhr.open('get', searchDataURL) //'/js/lunr/search.json')
xhr.addEventListener('load', feedLoaded)
xhr.send()

function lunrSearch(query) {
    let response = []
    let results = idx.search(query)
    for (let i = 0; i < results.length; i++) {
        let result = results[i]
        response[i] = searchDatabase[result.ref]
    }
        
    return response
}


const displayLunrResults = (term, finalResults, extractHighlight = false) => {

  const results = document.getElementById("results-container")
  if (finalResults.length === 0) {
    results.innerHTML = `<button class="result-card">
                    <h3>No results.</h3>
                    <p>Try another search term?</p>
                </button>`
  } else {
    results.innerHTML = finalResults
      .map((result) => {
        return resultToHTML({
          id: result.href,
          url: result.href,
          title: highlight(result.title, term),
          content: highlight(removeMarkdown(result.content), term)
        })
      }
      )
      .join("\n")
    if (LATEX_ENABLED) {
      renderMathInElement(results, {
        delimiters: [
          { left: '$$', right: '$$', display: false },
          { left: '$', right: '$', display: false },
        ],
        throwOnError: false
      })
    }

    const anchors = [...document.getElementsByClassName("result-card")]
    anchors.forEach((anchor) => {
      anchor.onclick = () => redir(anchor.id, term)
    })
  }
}
