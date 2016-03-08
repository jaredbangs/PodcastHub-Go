var Feed = require('../models/feed.js');
var FeedArchiveStrategyCollection = require('../collections/feedArchiveStrategies.js');
var ItemCollection = require('../collections/items.js');
var FeedView = require('./feed.js')
var Template = require('../../templates/downloadDirectoryItemForList.handlebars');

module.exports = Marionette.ItemView.extend({
	
	tagName: "div",
	className: "row feed-info",

	template: Template,

	events: {
                "click .archive": "archive",
		"click .feed-name": "goToFeed"
        },

        archive: function(domEvent) {

                var self = this;

                self.model.save({IsToBeArchived: true}, {
                        success: function() {
                                console.log("Item updated");
                                self.remove();
                        }
                 });

        },

	goToFeed: function(domEvent) {

	   	podcasthub.feed = new Feed({ id: this.model.get("FeedId")});

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
