(function(window, document, $) {
  if (gokuraku.Player) {
    return;
  }

  function Player() {
    this._soundcloud_client_id = null;
    this._track = null;
    this._sdk = null;
    // SDK
    this._isSDKLoaded = false;
    this._isSDKInitialized = false;
    // Sound
    this._volume = 20;
    this._sound = null;
    // Retry
    this._initSDKRetryCount = 0;
    this._playRetryCount = 0;
    this._defaultRetryInterval = 3 * 1000;
    this._maxRetryInterval = 10 * 1000;
    // Callback
    this.onPlay = null;
    this.onFinish = null;
    this.onFailure = null;
  }

  Player.prototype = {
    play: play,
    _play: _play,
    _playRetryInterval: _playRetryInterval,
    _onFailToPlay: _onFailToPlay,
    _loadSDK: _loadSDK,
    _initializeSDK: _initializeSDK,
    _initSDKRetryInterval: _initSDKRetryInterval,
    _fetchSoundcloudClientId: _fetchSoundcloudClientId
  };

  function play(track) {
    if (this._isSDKLoaded && this._isSDKInitialized) {
      this._play(track);
      return;
    }

    $.when(this._loadSDK(), this._fetchSoundcloudClientId())
      .done(function() {
          this._initSDKRetryCount = 0;
          this._initializeSDK();
          this._play(track);
      }.bind(this)).
      fail(function(xhr) {
        if (this.onFailure) {
          this.onFailure();
        }

        // Retry init SDK
        gokuraku.Util.alert("error", "Fail to initialize soundcloud sdk. Retry after " + this._initSDKRetryInterval() / 1000 + "sec.");
        setTimeout(function() {
          gokuraku.Util.clearAlert();
          this._initSDKRetryCount += 1;
          this.play(track);
        }.bind(this), this._initSDKRetryInterval());
        return;
      }.bind(this));
  }

  function _play(track) {
    var pos = +(new Date()) - +(new Date(track.StartedAt * 1000));

    this._track = track;

    if (this._sound) {
      this._sound.stop();
      this._sound.destruct();
      this._sound = null;
    }

    this._sdk.stream("/tracks/" + this._track.Id, {
      autoPlay: false,
      position: pos,
      volume: this._volume,
      onfinish: function Player_onfinish() {
        gokuraku.Util.info("Finish to play: " + this._track.Title, [this._track, this._sound]);
        if (this.onFinish) {
          this.onFinish();
        }
      }.bind(this),
      onplay: function Player_onplay() {
        if (this.onPlay) {
          this.onPlay(this._track);
        }
      }.bind(this),
      ontimeout: function() {
        gokuraku.Util.error("error: Sound#ontimeout", [this._track, this._sound]);
        this._onFailToPlay();
      }.bind(this),
      ondataerror: function() {
        gokuraku.Util.error("error: Sound#ondataerror", [this._track, this._sound]);
        this._onFailToPlay();
      }.bind(this),
      onsuspend: function() {
        gokuraku.Util.error("error: Sound#onsuspend", [this._track, this._sound]);
        this._onFailToPlay();
      }.bind(this),
      onfailure: function() {
        gokuraku.Util.error("error: Sound#onfailure", [this._track, this._sound]);
        this._onFailToPlay();
      }.bind(this)
    }, function(sound) {
      this._sound = sound;
      this._sound.play();

      // if fail to get get sound resource in Sound#play(), any event don't occur...
      // (e.g ontimeout, ondataerror, onsuspend, onfailure)
      // So monior Sound#readyState
      var i = setInterval(function() {
        if (+this._sound.readyState === 2) {
          clearInterval(i);

          gokuraku.Util.error("Sound#readyState is 2", [this._track, this._sound]);
          this._onFailToPlay();
        } else if (+this._sound.readyState === 3) {
          clearInterval(i);
          gokuraku.Util.clearAlert();

          gokuraku.Util.info("Now Playing(play at pos: " + pos + "): " + this._track.Title, [this._track, this._sound]);
          document.title = "▶ " + this._track.Title + " | 極楽";
          this._playRetryCount = 0;
        }
      }.bind(this), 1000);
    }.bind(this));
  }

  function _onFailToPlay() {
    gokuraku.Util.error("Fail to  play: " + this._track.Title, [this._track, this._sound]);
    if (this.onFailure) {
      this.onFailure();
    }

    // Retry play sound
    gokuraku.Util.alert("error", "Fail to play Track. Retry after " + this._playRetryInterval() / 1000 + "sec.");
    setTimeout(function() {
      this._playRetryCount += 1;
      this._play(this._track);
    }.bind(this), this._playRetryInterval());
  }

  // Retry interval calculators
  function _initSDKRetryInterval() {
    var i = this._initSDKRetryCount * this._defaultRetryInterval;

    if (i <= this._maxRetryInterval) {
      return i;
    } else {
      return this._maxRetryInterval;
    }
  }

  function _playRetryInterval() {
    var i = this._playRetryCount * this._defaultRetryInterval;

    if (i <= this._maxRetryInterval) {
      return i;
    } else {
      return this._maxRetryInterval;
    }
  }

  // Client key
  function _fetchSoundcloudClientId() {
    var soundcloudClientIdFetcher = new gokuraku.API.SoundcloudClientIdFetcher();
    soundcloudClientIdFetcher.onDone = function(soundcloud_client_id) {
      this._soundcloud_client_id = soundcloud_client_id;
    }.bind(this);

    return soundcloudClientIdFetcher.fetch();
  }

  // SDK
  function _initializeSDK() {
    this._isSDKInitialized = true;
    this._sdk.initialize({
      client_id: this._soundcloud_client_id,
      redirect_uri: ""
    });
  }

  function _loadSDK() {
    soundcloudSDKFetcher = new gokuraku.API.SoundcloudSDKFetcher();
    soundcloudSDKFetcher.onDone = function() {
      if (typeof SC !== "undefined") {
        this._isSDKLoaded = true;
        this._sdk = SC;
      }
    }.bind(this);
    return soundcloudSDKFetcher.fetch();
  }

  gokuraku.Player = Player;
})(window, window.document, jQuery);
