(function(window, document, $) {
  if (window.API) {
    return;
  }

  API = {
    getAllTrack: getAllTrack,
    addTrack: addTrack,
    getSoundcloudClientId: getSoundcloudClientId
  };

  function getAllTrack() {
    return $.ajax({
      type: "GET",
      url: "/tracks",
      contentType: 'application/json',
      dataType: 'json'
    });
  }

  function getSoundcloudClientId() {
    return $.ajax({
      type: "GET",
      url: "/soundcloud_client_id"
      // contentType: 'application/json',
      // dataType: 'json'
    });
  }

  function addTrack(url) {
    return $.ajax({
      type: "POST",
      url: "/tracks",
      contentType: 'application/json',
      dataType: 'json',
      data: JSON.stringify({
        Url: url
      })
    });
  }

  window.API = API;
})(window, window.document, jQuery);
