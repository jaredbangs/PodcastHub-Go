var DownloadDirectoryCollection = require('../collections/downloadDirectories.js');
var DownloadDirectoryListView = require('../views/downloadDirectoryList.js');
var DownloadDirectoryView = require('../views/downloadDirectory.js');
var FeedArchiveStrategyCollection = require('../collections/feedArchiveStrategies.js');
var FeedListView = require('../views/feedList.js');
var FeedView = require('../views/feed.js');
var ItemCollection = require('../collections/items.js');

module.exports = Marionette.AppRouter.extend({

	appRoutes: {
		"showFeed/:id" : "showFeed",
		"showFeedList" : "showFeedList"
	},

	routes : {
		"downloadDirectory/:directory" : "downloadDirectory",
		"feed/:id" : "feed",
		"feedList" : "feedList",
		"itemsByDownloadDirectory" : "itemsByDownloadDirectory"
	},

	downloadDirectory: function(directory) {

		console.log(directory);

		var ddsWithItems = podcasthub.dd.filter( function (item) { return item.get("ItemCount") > 0;});

		var collection = new DownloadDirectoryCollection(ddsWithItems);

		var model = collection.findWhere({ Name: directory });

		model.set("DirectoryNames", collection.pluck("UrlName"));

		if (model === undefined) {
			console.log("Could not find " + directory);
		} else {
		
			var items = new ItemCollection(model.get("Items").filter(function (item) { return item.shouldDisplayByDefault(); }));
			items.comparator = "FeedName";
			var view = new DownloadDirectoryView({ model: model, collection: items });
                	podcasthub.appLayout.RegionOne.show(view);
		}
	},

	feed: function(feedId) {

		var self = this;

		podcasthub.feed = new Feed({ id: feedId });

		podcasthub.feed.fetch({
			reset: true,
			success: function (model, response, options) {

				if (podcasthub.feedArchiveStrategyNames === undefined) {

	                        	new FeedArchiveStrategyCollection().fetch({
		                        	success: function (collection, feedArchiveStrategyNames) {
							podcasthub.feedArchiveStrategyNames = feedArchiveStrategyNames;
							self.showCurrentFeed();
                                        	}
                                	});
				} else {
					self.showCurrentFeed();
				}
                        },
                        error: function (model, response, options) {
                        }
		});
	},

	feedList: function() {

              	var view = new FeedListView({ collection: podcasthub.FeedList });

                podcasthub.appLayout.RegionOne.show(view);
	},

	itemsByDownloadDirectory : function() {

		var self = this;

                podcasthub.dd = new DownloadDirectoryCollection();

                podcasthub.dd.fetch({
			reset: true,
			success: function (model, response, options) {
				self.showItemsByDownloadDirectory();
                       	}
		});
	},

	showCurrentFeed : function() {

		var nonArchivedItems = new ItemCollection(podcasthub.feed.get("Channel").get("Items").filter(function (item) { return item.shouldDisplayByDefault(); }));
                var view = new FeedView({ model: podcasthub.feed, collection: nonArchivedItems, feedArchiveStrategyNames: podcasthub.feedArchiveStrategyNames });
                podcasthub.appLayout.RegionOne.show(view);
	},

	showItemsByDownloadDirectory : function() {

		var ddsWithItems = podcasthub.dd.filter( function (item) { return item.get("ItemCount") > 0;});

		var view = new DownloadDirectoryListView({ collection: new DownloadDirectoryCollection(ddsWithItems) });
                podcasthub.appLayout.RegionOne.show(view);
	}
});
