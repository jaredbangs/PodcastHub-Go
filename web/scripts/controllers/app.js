"use strict";

var Feed = require('../models/feed.js');
var FeedArchiveStrategyCollection = require('../collections/feedArchiveStrategies.js');
var FeedInfoCollection = require('../collections/feedInfo.js');
var FeedListView = require('../views/feedList.js');
var FeedView = require('../views/feed.js')
var ItemCollection = require('../collections/items.js');

module.exports = Marionette.Object.extend({

	initialize: function (options) {

		var self = this;

		this.application = options.application;

		if (this.feedArchiveStrategyNames === undefined) {
			new FeedArchiveStrategyCollection().fetch({
				success: function (collection, feedArchiveStrategyNames) {
					self.feedArchiveStrategyNames = feedArchiveStrategyNames;
				}
			});
		}
	},

	showFeed: function (id) {

		var nonArchivedItems, self, view;
		
		self = this;

		this._withLoadedFeed(id, function (feed) {

			nonArchivedItems = new ItemCollection(feed.get("Channel").get("Items").filter(function (item) { return item.shouldDisplayByDefault(); }));
			
			view = new FeedView({ model: feed, collection: nonArchivedItems, feedArchiveStrategyNames: self.feedArchiveStrategyNames });
			self._showMainView(view);
			self._updateUrl("/showFeed/" + feed.id);
		});
	},

	showFeedList: function () {

		var self, view;

		self = this;

		this._withLoadedFeedList(function () {
			
			view = new FeedListView({ collection: self.feedList });
			//this.listenTo(view, "all", function (eventName) { console.log(eventName) });
			self.listenTo(view, "childview:show-feed", self._onShowFeed);

			self._showMainView(view);
			self._updateUrl("/showFeedList");
		});
	},

	_onShowFeed: function (view, childViewTriggerArguments) {
		this.showFeed(childViewTriggerArguments.model.get("Id"));		
	},

	_showMainView: function (view) {
		this.application.layout.showChildView("Main", view);
	},

	_updateUrl: function (url) {
		this.application.router.navigate(url);
	},

	_withLoadedFeed: function (id, func) {
		this.feedList.withLoadedFeed(id, function (feed) { func(feed); });
	},

	_withLoadedFeedList: function (func) {
		if (this.feedList === undefined) {

			this.feedList = new FeedInfoCollection();
			this.feedList.fetch({
				reset: true,
				success: func
			});
			this.application.FeedList = this.feedList;

		} else {
			func();
		}
	}
});
