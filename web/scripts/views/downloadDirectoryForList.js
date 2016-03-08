var DownloadDirectoryView = require('./downloadDirectory.js')
var ItemCollection = require('../collections/items.js')
var Template = require('../../templates/downloadDirectoryForList.handlebars');

module.exports = Marionette.ItemView.extend({
	
	tagName: "div",
	className: "row feed-info",

	template: Template,

	events: {
		'click': 'goToDownloadDirectory'
	},

	goToDownloadDirectory: function(domEvent) {

		var view = new DownloadDirectoryView({ model: this.model, collection: new ItemCollection(this.model.get("Items").filter(function (item) { return item.shouldDisplayByDefault(); })) });
		this._parent._parent._parent._parent.RegionOne.show(view);
	}
});
