var common = {
  processStatus: function(response) {
    if (response.status === 200 || response.status === 0) {
        return response.text();
    } else {
        return Promise.reject(new Error(response.statusText));
    }
  },

  processError: function(error) {
    console.log(error);
  }
}
