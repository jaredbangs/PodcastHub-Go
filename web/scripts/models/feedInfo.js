"use strict";

var moment = require("moment")
var Feed = require('../models/feed.js');

module.exports = Backbone.Model.extend({

	idAttribute: "Id",

	parse: function (response, options) {

		var lastUpdated = moment(response.LastUpdated);
		response.LastUpdated = lastUpdated._d;
		response.LastUpdatedDisplay = lastUpdated.format("MMMM Do YYYY");
		var lastFileDownloaded = moment(response.LastFileDownloaded);
		response.LastFileDownloaded = lastFileDownloaded._d;
		response.LastFileDownloadedDisplay = lastFileDownloaded.format("YYYY-MM-DD");
		return response;
	},

	updateFromModel: function (feedModel) {
		
		this.set("ArchivePath", feedModel.get("ArchivePath"));
                this.set("ArchiveStrategy", feedModel.get("ArchiveStrategy"));
	},

	withLoadedFeed: function (func) {

		var feed;
		
		if (!this.has("FeedModel")) {

			feed = new Feed({ id: this.get("Id")});
			this.set("FeedModel", feed);

			feed.fetch({
				reset: true,
				success: function () { 
					func(feed); 
				}
			});

		} else {
			func(this.get("FeedModel"));
		}
	}
});
