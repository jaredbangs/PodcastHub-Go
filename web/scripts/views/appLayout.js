var FeedInfoCollection = require('../collections/feedInfo.js');
var FeedListView = require('./feedList.js');
var Template = require('../../templates/appLayout.handlebars');

module.exports = AppLayout = Marionette.LayoutView.extend({
	tagName: "div",

	template: Template,

	regions: {
		'RegionOne' : '#feed-list'
	},

	onRender: function() {
		
		var feedList = new FeedInfoCollection();
		feedList.fetch();

		var view = new FeedListView({ collection: feedList });

		/* display the collection view in region 1 */
		this.RegionOne.show(view);
	}

});
