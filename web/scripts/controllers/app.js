"use strict";

var DownloadDirectoryListView = require('../views/downloadDirectoryList.js');
var DownloadDirectoryCollection = require('../collections/downloadDirectories.js');
var DownloadDirectoryView = require('../views/downloadDirectory.js')
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

	showDownloadDirectories: function () {

		var self, view;

		self = this;

		this._withLoadedDownloadDirectories(function () {
			
			view = new DownloadDirectoryListView({ collection: self.downloadDirectories.getNonArchivedDirectories() });
			//self.listenTo(view, "all", function (eventName) { console.log(eventName) });
			self.listenTo(view, "childview:show-directory", self._onShowDirectory);

			self._showMainView(view);
			self._updateUrl("/showDownloadDirectories");
		});
	},

	showDirectory: function (name) {

		var self = this;

		this._withLoadedFeedList(function () {
			self._withLoadedDownloadDirectories(function () {
				self._showDirectory(self.downloadDirectories.findWhere({ Name: name }));
			});
		});
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

	_onShowDirectory: function (view, childViewTriggerArguments) {
		this._showDirectory(childViewTriggerArguments.model);		
	},

	_onShowFeed: function (view, childViewTriggerArguments) {
		this.showFeed(childViewTriggerArguments.model.get("Id"));		
	},

	_showDirectory: function (model) {

		var self, view;
		self = this;
	
		var view = new DownloadDirectoryView({ model: model, collection: new ItemCollection(model.get("Items").filter(function (item) { return item.shouldDisplayByDefault(); })) });
		self.listenTo(view, "refresh:download:directories", function () {
			self.downloadDirectories = undefined;
			self.showDownloadDirectories();
		});
		self._showMainView(view);
		self._updateUrl("/showDirectory/" + model.get("UrlName"));
	},

	_showMainView: function (view) {
		this.application.layout.showChildView("Main", view);
	},

	_updateUrl: function (url) {
		this.application.router.navigate(url);
	},

	_withLoadedDownloadDirectories: function (func) {

		if (this.downloadDirectories === undefined) {

			this.downloadDirectories = new DownloadDirectoryCollection();
			this.downloadDirectories.fetch({
				reset: true,
				success: func
			});
			this.application.DownloadDirectories = this.downloadDirectories;

		} else {
			func();
		}
	},

	_withLoadedFeed: function (id, func) {
		
		var self = this;

		this._withLoadedFeedList(function () {
			self.feedList.withLoadedFeed(id, function (feed) { func(feed); });
		});
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
