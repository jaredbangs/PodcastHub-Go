var DownloadDirectoryItemChildView = require('../views/downloadDirectoryItemForList.js');

module.exports = Marionette.CollectionView.extend({
	tagName: "div",
	childView: DownloadDirectoryItemChildView

});
