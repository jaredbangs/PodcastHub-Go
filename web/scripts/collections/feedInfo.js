var FeedInfo = require('../models/feedInfo');

module.exports = FeedInfoCollection = Backbone.Collection.extend({
	
	model:  FeedInfo,

	url: "/feeds",

	comparator: function(modelA, modelB) {

		if (modelA.get("LastFileDownloaded") > modelB.get("LastFileDownloaded")) return -1; 
		if (modelA.get("LastFileDownloaded") < modelB.get("LastFileDownloaded")) return 1; 
		if (modelA.get("LastUpdated") > modelB.get("LastUpdated")) return -1; 
		if (modelA.get("LastUpdated") < modelB.get("LastUpdated")) return 1; 
		return 0; // equal
	},

	withLoadedFeed: function (id, func) {

		var feedInfo = this.findWhere({ Id: id });

		feedInfo.withLoadedFeed(func);
	}
});
