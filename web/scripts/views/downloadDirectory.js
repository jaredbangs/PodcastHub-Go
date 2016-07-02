var Marionette = require("backbone.marionette");
var DownloadDirectoryItemChildView = require('../views/downloadDirectoryItemForList.js');
var Template = require('../../templates/downloadDirectory.handlebars');

module.exports = Marionette.CompositeView.extend({
	childView: DownloadDirectoryItemChildView,
	childViewContainer: "#directory-items",
	template: Template,

	events: {
		"click .archive-all": "archiveAll"
	},

	triggers: {
		"click .page-back": "page:back",
		"click .page-forward": "page:forward"
	},

        archiveAll: function(domEvent) {

                var saveRequests, self;
		
		self = this;

		saveRequests = [];

		self.collection.each(function (model) {
		
			saveRequests.push(model.save({IsToBeArchived: true}, {
				success: function() {

					self.$("#" + model.cid).remove();
				}
			}));
		});

		$.when.apply($, saveRequests).done(function () {
			self.triggerMethod("refresh:download:directories");
		});
	}
});
