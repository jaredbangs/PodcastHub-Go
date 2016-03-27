var FeedInfoForList = require('../views/feedInfoForList.js');
var Marionette = require("backbone.marionette");

module.exports = Marionette.CollectionView.extend({
	tagName: "div",
	childView: FeedInfoForList

});
