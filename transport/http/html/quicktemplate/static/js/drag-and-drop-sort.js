const el = document.getElementsByClassName('list');

for(let i = 0; i < el.length; i++) {
  var sortable = Sortable.create(el[i], {
    scroll: false,
    animation: 0,
    delayOnTouchOnly: true,
    delay: 1000,
    fallbackTolerance: 3,
    ghostClass: "--hidden",

    /*onStart: function (evt) {
      document.body.classList.add("lock-screen");
    },
    onEnd: function (evt) {
      document.body.classList.remove("lock-screen");
    },*/

    onUpdate: function (e) {
      const el = e.item.id;
      const i = e.newIndex;
      const all = e.to.childNodes;

      let ids = [all.length]
      for(let i = 0; i < all.length; i++) {
        ids[i] = parseInt(all[i].id.split("-")[1]);
      }
      
      const apiURL = "/" + url + "/" + e.to.id.replaceAll("-", "/") + "/reorder";
      updateOrder(apiURL, ids);
    },
  });
}

function updateOrder(url, ids) {
  fetch(url, {
    method: "POST", 
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(ids)
  }).then(response => response.json()).then(result => {
  
    if(result.error_code !== 0) {
      console.error(result);
    }
  }).catch((error) => {
    console.error('Error:', error);
  });
}
