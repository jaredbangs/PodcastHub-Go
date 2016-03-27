
var Marionette = require("backbone.marionette");
var FeedInfoCollection = require('../collections/feedInfo.js');
var FeedListView = require('./feedList.js');
var Template = require('../../templates/appLayout.handlebars');

module.exports = Marionette.LayoutView.extend({

	template: Template,

	el: "#app-base",

	regions: {
		"Header" : "#header",
		"Main" : "#main",
		"Sidebar" : "#sidebar"
	},

	onRender: function() {
		
		podcasthub.FeedList = new FeedInfoCollection();
		podcasthub.FeedList.fetch();

		var view = new FeedListView({ collection: podcasthub.FeedList });

		/* display the collection view in region 1 */
		this.regions.Main.show(view);
	}

});
