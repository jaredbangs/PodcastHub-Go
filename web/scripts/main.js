require("../styles/PodcastHub.less")

var $ = require("jquery");

var App = require('./app.js');

$(document).ready(function() {
	podcasthub = new App();
	podcasthub.start();
});
