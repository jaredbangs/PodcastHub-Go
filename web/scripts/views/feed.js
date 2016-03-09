var ItemView = require("./item.js")
var Template = require('../../templates/feed.handlebars');

module.exports = Marionette.CompositeView.extend({
	
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

		var self = this;

		this.model.save(null, {
			success: function() {

				var feedInfo = podcasthub.FeedList.findWhere({ Id: self.model.get("Id") });
				
				if (feedInfo !== undefined) {

					feedInfo.updateFromModel(self.model);
				}
			}
		});
	},

  	serializeData: function(){
	      var data = this.model.toJSON();
	      data.feedArchiveStrategyNames = this.options.feedArchiveStrategyNames;
	      return data;
	}
});
