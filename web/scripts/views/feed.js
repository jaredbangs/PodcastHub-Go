var Template = require('../../templates/feed.handlebars');

module.exports = FeedView = Marionette.LayoutView.extend({
	el: "#app-base",

	template: Template
});
