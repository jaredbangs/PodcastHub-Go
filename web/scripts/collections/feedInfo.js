var Backbone = require("backbone")
var FeedInfo = require('../models/feedInfo');

module.exports = FeedInfoCollection = Backbone.Collection.extend({
	
	model:  FeedInfo,

	url: "/feeds",

	comparator: function(modelA, modelB) {

		if (modelA.get("LastUpdated") > modelB.get("LastUpdated")) return -1; 
		if (modelA.get("LastUpdated") < modelB.get("LastUpdated")) return 1; 
		return 0; // equal
	}
});
