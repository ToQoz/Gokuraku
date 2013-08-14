(function(window, document, $) {
  if (window.AddTrackForm) {
    return;
  }

  function AddTrackForm() {
    this.$el = $('#post-track');
    this.$el.on('submit', this.onSubmitHandler.bind(this));
    this.afterSubmit = null;
  }

  AddTrackForm.prototype = {
    urlField: urlField,
    submitButton: submitButton,
    onSubmitHandler: onSubmitHandler,
    onSubmitStart: onSubmitStart,
    onSubmitSuccess: onSubmitSuccess,
    onSubmitFail: onSubmitFail,
    onSubmitFinish: onSubmitFinish
  };

  function urlField(e) {
    if (!this._urlField) {
      this._urlField = this.$el.find('input[name="Url"]');
    }

    return this._urlField;
  }

  function submitButton() {
    if (!this._submitButton) {
      this._submitButton = this.$el.find('button[type="submit"]');
    }

    return this._submitButton;
  }

  function onSubmitHandler(e) {
    var track_url = this.urlField().val();

    e.preventDefault();

    this.onSubmitStart();
    var trackPoster = new API.TrackPoster();
    trackPoster.post().
      done(this.onSubmitSuccess.bind(this)).
      fail(this.onSubmitFail.bind(this)).
      always(this.onSubmitFinish.bind(this));
  }

  function onSubmitSuccess(resp) {
    var track = resp.d;

    Util.alert("info", 'Added Track: ' + track.Title);

    this.urlField().val("");

    if (this.onSubmit) {
      this.onSubmit();
    }
  }

  function onSubmitFail(xhr) {
    Util.alert("error", xhr.responseJSON.e.join(" "));
  }

  function onSubmitFinish() {
    this.submitButton().removeClass('disabled').text("Add Track");
  }


  function onSubmitStart() {
    this.submitButton().addClass('disabled').text("Adding...");
  }

  window.AddTrackForm = AddTrackForm;
})(window, window.document, jQuery);
