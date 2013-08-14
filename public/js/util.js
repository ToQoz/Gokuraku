(function(window, document, $) {
  if (window.Util) {
    return;
  }

  var Util = {
    log: log,
    info: info,
    error: error,
    alert: alert,
    clearAlert: clearAlert
  };

  function log(message, relatedObjects) {
    _log("log", message, relatedObjects);
  }

  function info(message, relatedObjects) {
    _log("info", message, relatedObjects);
  }

  function error(message, relatedObjects) {
    _log("error", message, relatedObjects);
  }

  function _log(type, message, relatedObjects) {
    var now = new Date(),
        y = now.getFullYear(),
        m = _fillWithZero(2, now.getMonth() + 1),
        d = _fillWithZero(2, now.getDate()),
        h = _fillWithZero(2, now.getHours()),
        min = _fillWithZero(2, now.getMinutes()),
        s = _fillWithZero(2, now.getSeconds()),
        formatedTime = y + "/" + m + "/" + d + " " + h + ":" + min + ":" + s;

    relatedObjects = relatedObjects ? relatedObjects : [];

    relatedObjects.unshift("[" + formatedTime + "] " + message);
    console[type].apply(console, relatedObjects);
  }

  function clearAlert() {
    $("#alert-container").html("");
  }

  function alert(type, message) {
    var  template = $("#alert-error-templates").html(),
         container = $("#alert-container"),
         output = Mustache.render(template, {
           type: type,
           message: message
         });

    container.html(output);
  }

  function _fillWithZero(width, val) {
    var str_val = "" + val;

    if (str_val.length > width) {
      var err = "val(" + val + ") length is over " + width;
      throw err;
    }

    if (str_val.length == width) {
      return str_val;
    }

    return _fillWithZero(width, "0" + str_val);
  }

  window.Util = Util;
})(window, window.document, jQuery);
