(function(window, document, $) {
  $(function() {
    main();
  });

  function main() {
    if (!window.WebSocket) {
      Util.alert("error", "This browser is not support WebSocket. Please access from browser that support WebSocket");
      return;
    }

    var form = new AddTrackForm(),
        player = new Player(),
        ws,
        trackList = new TrackList();

    var retry_count = 0;


    connectToWebSocket = function(port) {
      ws = new WebSocket("ws://" + location.hostname + ":" + port + "/ws");

      ws.onmessage = function(msg) {
        var data = JSON.parse(msg.data);

        if (data.Type === "play") {
          player.play(data.Track);
        }
      };

      ws.onclose = function() {
        Util.alert("error", "Closed connection to WebSocket. Reconnecting after 3 sec.");
        setTimeout(function () {
          Util.clearAlert();
          connectToWebSocket(port);
        }, 3000);
      };
    };

    var webSocketPortFetcher = new API.WebSocketPortFetcher();

    webSocketPortFetcher.onDone = function(port) {
      connectToWebSocket(port);
    };

    webSocketPortFetcher.fetch();

    form.onSubmit = function() {
      trackList.update();
    };

    player.onPlay = function(track) {
      var template = $('#track-templates').html(),
          track_for_template = $.extend(track);

      track_for_template.br = function () {
        return function (text, render) {
          return render(text).replace(/(\r?\n)/g, "<br>");
        };
      };

      $('#track-container').html(Mustache.render(template, track_for_template));
      ws.send(JSON.stringify({ Type: "playing" }));
    };

    player.onFailure = function() {
      ws.send(JSON.stringify({ Type: "invalid" }));
    };

    player.onFinish = function() {
      ws.send(JSON.stringify({ Type: "waiting" }));
    };
  }
})(window, window.document, jQuery);
