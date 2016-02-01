(function e(t,n,r){function s(o,u){if(!n[o]){if(!t[o]){var a=typeof require=="function"&&require;if(!u&&a)return a(o,!0);if(i)return i(o,!0);var f=new Error("Cannot find module '"+o+"'");throw f.code="MODULE_NOT_FOUND",f}var l=n[o]={exports:{}};t[o][0].call(l.exports,function(e){var n=t[o][1][e];return s(n?n:e)},l,l.exports,e,t,n,r)}return n[o].exports}var i=typeof require=="function"&&require;for(var o=0;o<r.length;o++)s(r[o]);return s})({1:[function(require,module,exports){
var FeedInfo = require('../models/feedInfo');

module.exports = FeedInfoCollection = Backbone.Collection.extend({
	model:  FeedInfo,
	url: "/feeds"
});

},{"../models/feedInfo":3}],2:[function(require,module,exports){
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

},{"./views/appLayout.js":4}],3:[function(require,module,exports){
module.exports = FeedInfo = Backbone.Model.extend({
});

},{}],4:[function(require,module,exports){
var FeedInfoCollection = require('../collections/feedInfo.js');
var FeedListView = require('./feedList.js');

module.exports = AppLayout = Marionette.LayoutView.extend({
	tagName: "div",

	template: "#layout-template",

	regions: {
		'RegionOne' : '#feed-list'
	},

	onRender: function() {
		
		var feedList = new FeedInfoCollection();
		feedList.fetch();

		var view = new FeedListView({ collection: feedList });

		/* display the collection view in region 1 */
		this.RegionOne.show(view);
	}

});

},{"../collections/feedInfo.js":1,"./feedList.js":6}],5:[function(require,module,exports){
module.exports = FeedInfoForList = Marionette.ItemView.extend({
	tagName: "li",

	template: "#itemView-template"
});

},{}],6:[function(require,module,exports){
var FeedInfoForList = require("./feedInfoForList.js")

module.exports = FeedList = Marionette.CollectionView.extend({
	tagName: "ul",
	childView: FeedInfoForList
});

},{"./feedInfoForList.js":5}]},{},[2]);
