chrome.runtime.onMessage.addListener(
  function(request, sender) {
    var options = {
      type: "basic",
      title: "點我去 honestbee 買更便宜",
      message: request.name,
      priority: 2,
      iconUrl: chrome.runtime.getURL("cartoon_honey_bee_flying_around_0071-0905-2616-0020_SMU.jpg")
    };

    chrome.notifications.create(request.name, options, function(id) {
      timer = setTimeout(function() { chrome.notifications.clear(id); }, 60000);
    });

    chrome.notifications.onClicked.addListener(function(id) {
      chrome.tabs.create({ url: request.url });
      chrome.notifications.clear(id);
    });
});

// http://www.acclaimclipart.com/free_clipart_images/cartoon_honey_bee_flying_around_0071-0905-2616-0020.html
