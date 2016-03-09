require('../styles/feed.less')

var AppLayoutView = require('./views/appLayout.js');
var AppRouter = require('./routers/app.js');

$( document ).ready(function() {

	podcasthub.app = new Backbone.Marionette.Application();
	
	podcasthub.app.addRegions({
		/* reference to container element in the HTML file */
		appRegion: '#app-base'
	});

	podcasthub.app.start();

	/* create a new instance of the layout from the module */
	podcasthub.appLayout = new AppLayoutView();

	/* display the layout in the region defined at the top of this file */
	podcasthub.app.appRegion.show(podcasthub.appLayout);

	var router = new AppRouter({controller: this });
	Backbone.history.start();
});
