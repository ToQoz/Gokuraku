(function(window, document, $) {
  if (window.API) {
    return;
  }

  function APIClient() {
    this._maxRetryCount = 10;
    this._retryCount = 0;
    this._defaultRetryInterval = 3000;
    this.onDone = null;
  }

  APIClient.clone = function APIClient_clone() {
    var f = function() {};
    f.prototype = Object.create(APIClient.prototype);
    f.constructor = APIClient;
    return f;
  };

  APIClient.prototype = {
    _retryInterval: APIClient_retryInterval
  };

  function APIClient_retryInterval() {
    return this._defaultRetryInterval * this._retryCount;
  }

  WebSocketPortFetcher = APIClient.clone();
  WebSocketPortFetcher.prototype.fetch = function WebSocketPortFetcher_fetch() {
    return $.ajax({
      type: "GET",
      url: "/websocket_port"
    }).done(function(data) {
      if (this.onDone) {
        this.onDone(data.d);
      }
    }.bind(this)).fail(function(xhr) {
      if (this._retryCount < this._maxRetryCount) {
        Util.alert('error', 'Fail to get /websocket_port. Retry after ' + this._retryInterval() / 1000 + "sec");
        setTimeout(function() {
          Util.clearAlert();
          this._retryCount += 1;
          this.getWebSocketPort();
        }.bind(this), this._retryInterval());
        return;
      }
    }.bind(this));
  };

  AllTrackFetcher = APIClient.clone();
  AllTrackFetcher.prototype.fetch = function AllTrackFetcher_fetch() {
    return $.ajax({
      type: "GET",
      url: "/tracks",
      contentType: 'application/json',
      dataType: 'json'
    }).done(function(data) {
      if (this.onDone) {
        this.onDone(data.d);
      }
    }.bind(this));
  };

  SoundcloudClientIdFetcher = APIClient.clone();
  SoundcloudClientIdFetcher.prototype.fetch = function SoundcloudClientIdFetcher_fetch() {
    return $.ajax({
      type: "GET",
      url: "/soundcloud_client_id"
      // contentType: 'application/json',
      // dataType: 'json'
    }).done(function(data) {
      if (this.onDone) {
        this.onDone(data.d);
      }
    }.bind(this));
  };

  SoundcloudSDKFetcher = APIClient.clone();
  SoundcloudSDKFetcher.prototype.fetch = function SoundcloudSDKFetcher_fetch() {
    return $.getScript('//connect.soundcloud.com/sdk.js').done(function() {
      if (this.onDone) {
        this.onDone();
      }
    }.bind(this));
  };

  TrackPoster = APIClient.clone();
  TrackPoster.prototype.post = function TrackPoster_post() {
    return $.ajax({
      type: "POST",
      url: "/tracks",
      contentType: 'application/json',
      dataType: 'json',
      data: JSON.stringify({
        Url: url
      })
    });
  };


  // exports
  API = {
    AllTrackFetcher: AllTrackFetcher,
    WebSocketPortFetcher: WebSocketPortFetcher,
    TrackPoster: TrackPoster,
    SoundcloudClientIdFetcher: SoundcloudClientIdFetcher,
    SoundcloudSDKFetcher: SoundcloudSDKFetcher
  };

  window.API = API;
})(window, window.document, jQuery);
