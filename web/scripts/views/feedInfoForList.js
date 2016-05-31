var Feed = require('../models/feed.js');
var FeedArchiveStrategyCollection = require('../collections/feedArchiveStrategies.js');
var ItemCollection = require('../collections/items.js');
var FeedView = require('./feed.js')
var Template = require('../../templates/feedInfoForList.handlebars');

module.exports = Marionette.ItemView.extend({
	
	tagName: "tr",
	className: "feed-info",

	template: Template,

	triggers: {
		"click" : "show-feed"
	}
});
