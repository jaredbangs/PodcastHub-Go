require("../styles/feed.less")

var $ = require("jquery");
var Backbone = require("backbone");
var Marionette = require("backbone.marionette");


var AppLayoutView = require('./views/appLayout.js');
var AppRouter = require('./routers/app.js');

$(document).ready(function() {

	podcasthub = new Marionette.Application();

	podcasthub.on("before:start", function() {

		podcasthub.layout = new AppLayoutView();
	});

	podcasthub.on("start", function() {
		
		podcasthub.router = new AppRouter({controller: this });

		podcasthub.layout.render();

		Backbone.history.start();
	});

	podcasthub.start();
});
