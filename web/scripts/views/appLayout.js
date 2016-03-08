var DownloadDirectoryCollection = require('../collections/downloadDirectories.js');
var DownloadDirectoryListView = require('./downloadDirectoryList.js');
var FeedInfoCollection = require('../collections/feedInfo.js');
var FeedListView = require('./feedList.js');
var Template = require('../../templates/appLayout.handlebars');

module.exports = Marionette.LayoutView.extend({
	tagName: "div",

	template: Template,

	regions: {
		'RegionOne' : '#feed-list'
	},

  	events: {
		'click .itemsByDownloadDirectory': 'itemsByDownloadDirectory'
	},

	itemsByDownloadDirectory: function(domEvent) {

		var dd, ddsWithItems, self;

		self = this;

		dd = new DownloadDirectoryCollection();

		dd.fetch({
	               reset: true,
	               success: function (model, response, options) {

			       ddsWithItems = dd.filter( function (item) { return item.get("ItemCount") > 0;});

				var view = new DownloadDirectoryListView({ collection: new DownloadDirectoryCollection(ddsWithItems) });
				self.RegionOne.show(view);
		       }
	        });
        },

	onRender: function() {
		
		podcasthub.FeedList = new FeedInfoCollection();
		podcasthub.FeedList.fetch();

		var view = new FeedListView({ collection: podcasthub.FeedList });

		/* display the collection view in region 1 */
		this.RegionOne.show(view);
	}

});
