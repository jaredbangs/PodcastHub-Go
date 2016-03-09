var Feed = require('../models/feed.js');
var FeedArchiveStrategyCollection = require('../collections/feedArchiveStrategies.js');
var ItemCollection = require('../collections/items.js');
var FeedView = require('./feed.js')
var Template = require('../../templates/feedInfoForList.handlebars');

module.exports = FeedInfoForList = Marionette.ItemView.extend({
	
	tagName: "div",
	className: "row feed-info",

	template: Template,

	events: {
	},

	goToFeed: function(domEvent) {

		podcasthub.feed = new Feed({ id: this.model.get("Id")});

		podcasthub.feed.fetch({
			reset: true,
			success: function (model, response, options) {

				new FeedArchiveStrategyCollection().fetch({
					success: function (collection, feedArchiveStrategyNames) {
					
						var nonArchivedItems = new ItemCollection(podcasthub.feed.get("Channel").get("Items").filter(function (item) { return item.shouldDisplayByDefault(); }));

						var view = new FeedView({ model: podcasthub.feed, collection: nonArchivedItems, feedArchiveStrategyNames: feedArchiveStrategyNames });
						view.render();
					}
				});

			},
			error: function (model, response, options) {

			}
		});
	}
});
