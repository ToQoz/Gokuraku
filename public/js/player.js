(function(window, document, $) {
  if (window.Player) {
    return;
  }

  function Player() {
    this._isSDKLoaded = false;
    this._isSDKInitialized = false;
    this._sdk = null;
    this._volume = 20;
    this._sound = null;
    this.onPlay = null;
    this.onFinish = null;
    this.onFailure = null;
  }

  Player.prototype = {
    play: play,
    _play: _play,
    _loadSDK: _loadSDK,
    _initializeSDK: _initializeSDK,
    _fetchSoundcloudClientId: _fetchSoundcloudClientId
  };

  function play(track) {
    $.when(this._loadSDK(), this._fetchSoundcloudClientId())
      .done(function() {
          this._initializeSDK();
          this._play(track);
      }.bind(this)).
      fail(function() {
        if (this.onFailure) {
          this.onFailure();
        }
        setTimeout(function() {
          alert("Fail to initialize soundcloud sdk. Retry now.");
          this.play();
        }, 1000);
      }.bind(this));
  }

  function _play(track) {
    var pos = +(new Date()) - +(new Date(track.StartedAt * 1000));

    if (this._sound) {
      this._sound.stop();
    }

    this._sdk.stream("/tracks/" + track.Id, {
      autoPlay: true,
      position: pos,
      volume: this._volume,
      onfinish: function() {
        Util.log("Finish to play: " + track.Title, [track]);
        if (this.onFinish) {
          this.onFinish();
        }
      }.bind(this),
      onplay: function() {
        Util.log("Now Playing(play at pos: " + pos + "): " + track.Title, [track]);
        if (this.onPlay) {
          this.onPlay(track);
        }
        document.title = "▶ " + track.Title + " | 極楽";
      }.bind(this),
      onfailure: function() {
        Util.log("Fail to  play: " + track.Title, [track]);
        if (this.onFailure) {
          this.onFailure();
        }
        setTimeout(function() {
          alert("Fail to play Track. Retry now.");
          this._play();
        }, 1000);
      }.bind(this)
    }, function(sound) {
      this._sound = sound;
    });
  }

  function _fetchSoundcloudClientId() {
    API.getSoundcloudClientId().done(function(data) {
      this._soundcloud_client_id = data.d;
    }.bind(this));
  }

  function _initializeSDK() {
    this._isSDKInitialized = true;
    this._sdk.initialize({
      client_id: this._soundcloud_client_id,
      redirect_uri: ""
    });
  }

  function _loadSDK() {
    return $.getScript('//connect.soundcloud.com/sdk.js').pipe(function() {
      if (typeof SC !== "undefined") {
        this._isSDKLoaded = true;
        this._sdk = SC;
      }

      return this;
    }.bind(this));
  }

  window.Player = Player;
})(window, window.document, jQuery);
