var DownloadDirectoryView = require('./downloadDirectory.js')
var Marionette = require("backbone.marionette");
var ItemCollection = require('../collections/items.js')
var Template = require('../../templates/downloadDirectoryForList.handlebars');

module.exports = Marionette.ItemView.extend({
	
	tagName: "div",
	className: "row feed-info",

	template: Template,

	triggers: {
		"click": "show-directory"
	}
});
