var Marionette = require("backbone.marionette");
var DownloadDirectoryItemChildView = require('../views/downloadDirectoryItemForList.js');
var Template = require('../../templates/downloadDirectory.handlebars');

module.exports = Marionette.CollectionView.extend({
	Template: Template,
	childView: DownloadDirectoryItemChildView,
	childViewContainer: "#directory-items"
});
