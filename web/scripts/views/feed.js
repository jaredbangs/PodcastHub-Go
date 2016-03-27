var ItemView = require("./item.js")
var Marionette = require("backbone.marionette");
var Template = require('../../templates/feed.handlebars');

module.exports = Marionette.CompositeView.extend({
	
	events: {
		"change .archive-path": "changeArchivePath",
		"change .archive-strategy": "changeArchiveStrategy",
		"click .archive-selected": "archiveSelected",
		"click .save": "save",
		"click .select-all": "selectAll"
	},

	template: Template,

	childView: ItemView, 

	childViewContainer: "#feed-items",

	archiveNext: function () {

		var itemId, self;

		self = this;

		if (self.itemIdsToArchive.length > 0) {

			itemId = self.itemIdsToArchive.pop();

			self.collection.each(function (item) {

				if (itemId === item.get("ItemId")) {

					_.defer(function () {
						item.save({IsToBeArchived: true}, {
		        				success: function() {
								console.log("Archived " + item.get("Title"));
								self.archiveNext();
							}
	               				});
					});
				}
			});	
		}
	},

	archiveSelected: function () {

		var itemId, self;

		self = this;

		self.itemIdsToArchive = [];

		this.$("input:checkbox:checked").each(function(index, itemCheckbox) {
			
			itemId = $(itemCheckbox).data("item-id");

			self.itemIdsToArchive.push(itemId);
		});

		self.archiveNext();
	},

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

	selectAll: function () {
		this.$("input").prop("checked", true);
	},

  	serializeData: function(){
	      var data = this.model.toJSON();
	      data.feedArchiveStrategyNames = this.options.feedArchiveStrategyNames;
	      return data;
	}
});
