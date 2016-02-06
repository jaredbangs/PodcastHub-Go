var Feed = require('../models/feed');
var FeedView = require('./feed');
var Template = require('../../templates/feedInfoForList.handlebars');

module.exports = FeedInfoForList = Marionette.ItemView.extend({
	tagName: "li",

	template: Template,

	events: {
		'click .feed-info': 'goToFeed'
	},

	goToFeed: function(domEvent) {

		var feed = new Feed({ id: this.model.get("Id")});

		feed.fetch({
			reset: true,
			success: function (model, response, options) {
			
				var view = new FeedView({ model: feed });
				view.render();
			},
			error: function (model, response, options) {

			}
		});
	}
});
