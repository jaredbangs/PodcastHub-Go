var FeedInfoForList = require('../views/feedInfoForList.js');

module.exports = Marionette.CollectionView.extend({
	tagName: "div",
	childView: FeedInfoForList

});
