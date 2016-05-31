var Template = require('../../templates/appLayout.handlebars');

module.exports = Marionette.LayoutView.extend({

	template: Template,

	el: "#app-base",

	regions: {
		"Header" : "#header",
		"Main" : "#main",
		"Sidebar" : "#sidebar"
	}
});
