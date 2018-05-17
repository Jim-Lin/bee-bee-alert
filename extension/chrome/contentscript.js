var adapter;

if (location.href.indexOf('pchome.com.tw') > -1) {
  adapter = pchome24h();
}

adapter.fetch()
.then(function(value) {
  return fetch("http://localhost/comparison", {
    method: "POST",
    headers: {
      "Accept": "application/json",
      "Content-Type": "application/json"
    },
    body: JSON.stringify(value)
  });
})
.then(common.processStatus)
.then(processText)
.catch(common.processError);

function processText(text) {
  var prod = JSON.parse(text);
  console.log(prod);

  if (prod.name != "") {
    chrome.runtime.sendMessage(prod);
  }
}
