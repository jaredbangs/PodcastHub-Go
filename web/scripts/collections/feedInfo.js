var FeedInfo = require('../models/feedInfo');

module.exports = FeedInfoCollection = Backbone.Collection.extend({
	model:  FeedInfo,
	url: "/feeds"
});
