var AppLayoutView = require('./views/appLayout.js');

$( document ).ready(function() {
	
	var app = new Backbone.Marionette.Application();

	app.addRegions({
		/* reference to container element in the HTML file */
		appRegion: '#app-base'
	});

	app.start();

	/* create a new instance of the layout from the module */
	var layout = new AppLayoutView();

	/* display the layout in the region defined at the top of this file */
	app.appRegion.show(layout);

});
