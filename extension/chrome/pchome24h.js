var pchome24h = (function() {
  var re = /https:\/\/24h\.pchome\.com\.tw\/prod\/([^\?]+)/;
  var match = re.exec(location.href);
  var prod = match[1] + '-000';
  console.log(prod);

  function processText(text) {
    jsonp = unescape(text);
    var re = /.*jsonp_prod\((.*?)\);.*/;
    var match = re.exec(jsonp);
    var obj = JSON.parse(match[1]);
    console.log(obj);

    return {
      name: obj[prod].Name == obj[prod].Nick ? obj[prod].Name : obj[prod].Name + " " + obj[prod].Nick,
      price: obj[prod].Price.P,
      url: location.href
    };
  }

  var delegate = {
    fetch: function() {
      return fetch('https://ecapi.pchome.com.tw/ecshop/prodapi/v2/prod/' + prod + '&fields=Name,Nick,Price&_callback=jsonp_prod', {method: 'get'})
        .then(common.processStatus)
        .then(processText)
        .catch(common.processError);
    }
  }

  return delegate;
});
