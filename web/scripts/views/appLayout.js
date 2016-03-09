var FeedInfoCollection = require('../collections/feedInfo.js');
var FeedListView = require('./feedList.js');
var Template = require('../../templates/appLayout.handlebars');

module.exports = Marionette.LayoutView.extend({
	tagName: "div",

	template: Template,

	regions: {
		'RegionOne' : '#feed-list'
	},

	onRender: function() {
		
		podcasthub.FeedList = new FeedInfoCollection();
		podcasthub.FeedList.fetch();

		var view = new FeedListView({ collection: podcasthub.FeedList });

		/* display the collection view in region 1 */
		this.RegionOne.show(view);
	}

});
