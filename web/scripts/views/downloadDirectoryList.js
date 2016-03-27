var DownloadDirectoryChildView = require('../views/downloadDirectoryForList.js');
var Marionette = require("backbone.marionette");

module.exports = Marionette.CollectionView.extend({
	tagName: "div",
	childView: DownloadDirectoryChildView

});
