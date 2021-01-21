
const Controller = {
  search: (ev) => {
    ev.preventDefault();
    const form = document.getElementById("form");
    const data = Object.fromEntries(new FormData(form));
    const response = fetch(`/search?q=${data.query}`).then((response) => {
      response.json().then((results) => {
        Controller.updateTable(results);
      });
    });
  },

  updateTable: (results) => {
    const table = document.getElementById("table");
    const rows = [];
    table.innerHTML = ""
    for (let result of results) {
      table.innerHTML += `<p class="table">${result}</p>`
    }
  },

  recommendation: (ev) => {
    let query = ev.target.value
    if(query.length < 3){
      return null
    }
    fetch(`/recommendations?q=${query}`).then((response) => {
      let sugg = document.getElementById("suggestions")
      if(sugg != null) {
        document.getElementById("form").removeChild(sugg)
      }
      response.json().then((results) => {
        let suggestions = document.createElement("div")
        let _class = document.createAttribute("class")
        let _id = document.createAttribute("id")
        _class.value = "suggestion-box"
        _id.value = "suggestions"
        suggestions.setAttributeNode(_class)
        suggestions.setAttributeNode(_id)

        for(let i=0; i <results.length; i++) {
          let p = document.createElement("div")
          p.addEventListener("click", ()=>{
            document.getElementById("query").value = document.activeElement.innerHTML
          })
          p.setAttribute("class", "suggestion")
          p.setAttribute("id", i.toString())
          p.setAttribute("tabindex", i.toString())
          p.innerHTML = results[i]
          suggestions.appendChild(p)
        }
        document.getElementById("form").appendChild(suggestions)
      });
    });
  },
};

const changeVisibility = () => {
  let value = "hidden"
  var ele = document.activeElement
  if(ele.id === "query" || !Number.isNaN(parseInt(ele.id))){
    value = "visible"
  }
  let sugg = document.getElementById("suggestions")
  if(sugg == null) {
    return
  }
  sugg.style.visibility = value
}

const form = document.getElementById("form");
const sbar = document.getElementById("query");
const body = document.getElementById("body")

form.addEventListener("submit", Controller.search);
sbar.addEventListener("input", Controller.recommendation);
window.addEventListener("click", changeVisibility)
