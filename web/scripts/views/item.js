var Template = require('../../templates/item.handlebars');
var Marionette = require("backbone.marionette");

module.exports = Marionette.ItemView.extend({
	
	tagName: "div",
	className: "col-md-4",

	template: Template,

	events: {
		"click .archive": "archive"
	},

	archive: function(domEvent) {

		var self = this;

		self.model.save({IsToBeArchived: true}, { 
			success: function() {
				console.log("Item updated");	
				self.remove();
			}
		});

	}
});
