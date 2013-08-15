(function(window, document, $) {
  function WebSocket(port) {
    if (!window.WebSocket) {
      gokuraku.Util.alert("error", "This browser is not support WebSocket. Please access from browser that support WebSocket");
      return;
    }

    this._port = port;
    this._ws = null;
    this._connect();
  }

  WebSocket.prototype = {
    _connect: WebSocket_connect,
    onmessage: WebSocket_onmessage,
    send: WebSocket_send
  };

  function WebSocket_send(msg) {
    if (this._ws) {
      this._ws.send(msg);
    }
  }

  function WebSocket_onmessage(fn) {
    if (this._ws) {
      this._ws.onmessage = fn;
    }
  }

  function WebSocket_connect() {
    this._ws = new window.WebSocket("ws://" + location.hostname + ":" + this._port + "/ws");
    this.onclose = function() {
      gokuraku.Util.alert("error", "Closed connection to WebSocket. Reconnecting after 3 sec.");
      setTimeout(function () {
        gokuraku.Util.clearAlert();
        this._connect();
      }.bind(this), 3000);
    }.bind(this);
  }

  gokuraku.WebSocket = WebSocket;
})(window, window.document, jQuery);
