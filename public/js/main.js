(function(window, document, $) {
  $(function() {
    main();
  });

  function main() {
    var form = new AddTrackForm(),
        player = new Player(),
        ws = new WebSocket("ws://0.0.0.0:9099/ws"),
        trackList = new TrackList();

    // websocket
    ws.onmessage = function(msg) {
      var data = JSON.parse(msg.data);

      if (data.Type === "play") {
        player.play(data.Track);
      }
    };

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
      ws.send(JSON.stringify({ Type: "waiting" }));
    };

    player.onFinish = function() {
      ws.send(JSON.stringify({ Type: "waiting" }));
    };
  }
})(window, window.document, jQuery);
