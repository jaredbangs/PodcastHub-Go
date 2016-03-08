var ItemView = require("./item.js")
var Template = require('../../templates/feed.handlebars');

module.exports = FeedView = Marionette.CompositeView.extend({
	
	el: "#app-base",

	events: {
		"change .archive-path": "changeArchivePath",
		"change .archive-strategy": "changeArchiveStrategy",
		"click .save": "save"
	},

	template: Template,

	childView: ItemView, 

	childViewContainer: "#feed-items",

	changeArchivePath: function(e) {
		this.model.set("ArchivePath", $(e.currentTarget).val());
	},

	changeArchiveStrategy: function(e) {
		this.model.set("ArchiveStrategy", $(e.currentTarget).val());
	},

	save: function() {
		this.model.save({
			success: function() {
				 podcasthub.FeedList = new FeedInfoCollection();
		                 podcasthub.FeedList.fetch();
			}
		});
	},

  	serializeData: function(){
	      var data = this.model.toJSON();
	      data.feedArchiveStrategyNames = this.options.feedArchiveStrategyNames;
	      return data;
	}
});
