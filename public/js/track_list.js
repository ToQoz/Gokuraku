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
    var allTrackFetcher = new API.AllTrackFetcher();
    allTrackFetcher.onDone = function(tracks) {
      var output = Mustache.render(this._template, {tracks: tracks});
      this.$el.html(output);
    }.bind(this);
    allTrackFetcher.fetch();
  }

  window.TrackList = TrackList;
})(window, window.document, jQuery);
