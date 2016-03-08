var DownloadDirectoryChildView = require('../views/downloadDirectoryForList.js');

module.exports = Marionette.CollectionView.extend({
	tagName: "div",
	childView: DownloadDirectoryChildView

});
