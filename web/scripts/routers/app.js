var DownloadDirectoryCollection = require('../collections/downloadDirectories.js');
var DownloadDirectoryListView = require('../views/downloadDirectoryList.js');
var FeedArchiveStrategyCollection = require('../collections/feedArchiveStrategies.js');
var FeedListView = require('../views/feedList.js');
var FeedView = require('../views/feed.js');
var ItemCollection = require('../collections/items.js');

module.exports = Marionette.AppRouter.extend({

	routes : {
		"feed/:id" : "feed",
		"feedList" : "feedList",
		"itemsByDownloadDirectory" : "itemsByDownloadDirectory"
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

//		if (podcasthub.dd === undefined) {

                	podcasthub.dd = new DownloadDirectoryCollection();

                	podcasthub.dd.fetch({
				reset: true,
				success: function (model, response, options) {
					self.showItemsByDownloadDirectory();
                        	}
			});
//		} else {
//			self.showItemsByDownloadDirectory();
//		}
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
