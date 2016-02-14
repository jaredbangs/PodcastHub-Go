var Template = require('../../templates/feed.handlebars');
var ItemView = require("./item.js")

module.exports = FeedView = Marionette.CompositeView.extend({
	
	el: "#app-base",

	template: Template,

	childView: ItemView, 

	childViewContainer: "#feed-items"
});
