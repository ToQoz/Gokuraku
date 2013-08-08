(function(window, document, $) {
  if (window.TrackList) {
    return;
  }

  function TrackList() {
    this.$el = $("#track-list-container");
    this._template = $("#track-list-templates").html();
    this.update();
  }

  TrackList.prototype = {
    update: update
  };

  function update() {
    API.getAllTrack().done(function(tracks) {
      var output = Mustache.render(this._template, {tracks: tracks.d});
      this.$el.html(output);
    }.bind(this));
  }

  window.TrackList = TrackList;
})(window, window.document, jQuery);
