var FeedInfoForList = require("./feedInfoForList.js")

module.exports = FeedList = Marionette.CollectionView.extend({
	tagName: "ul",
	childView: FeedInfoForList
});
