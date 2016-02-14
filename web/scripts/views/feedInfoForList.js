var Feed = require('../models/feed');
var ItemCollection = require('../collections/items');
var FeedView = require('./feed');
var Template = require('../../templates/feedInfoForList.handlebars');

module.exports = FeedInfoForList = Marionette.ItemView.extend({
	
	tagName: "div",
	className: "row feed-info",

	template: Template,

	events: {
		'click': 'goToFeed'
	},

	goToFeed: function(domEvent) {

		podcasthub.feed = new Feed({ id: this.model.get("Id")});

		podcasthub.feed.fetch({
			reset: true,
			success: function (model, response, options) {
		
				var nonArchivedItems = new ItemCollection(podcasthub.feed.get("Channel").get("Items").where({Archived: false}));

				var view = new FeedView({ model: podcasthub.feed, collection: nonArchivedItems });
				view.render();
			},
			error: function (model, response, options) {

			}
		});
	}
});
