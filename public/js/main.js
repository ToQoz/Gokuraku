(function(window, document, $) {
  $(function() {
    main();
  });

  function main() {
    var form = new gokuraku.AddTrackForm(),
        player = new gokuraku.Player(),
        trackList = new gokuraku.TrackList(),
        ws = null,
        webSocketPortFetcher = new gokuraku.API.WebSocketPortFetcher();

    webSocketPortFetcher.onDone = function(port) {
      ws = new gokuraku.WebSocket(port);
      ws.onmessage(function(msg) {
        var data = JSON.parse(msg.data);

        if (data.Type === "play") {
          player.play(data.Track);
        }
      });
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
