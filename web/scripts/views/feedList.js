"use strict";

var FeedInfoForList = require('../views/feedInfoForList.js');
var Template = require("../../templates/feedList.handlebars");

module.exports = Marionette.CompositeView.extend({
	childViewContainer: ".feeds",
	childView: FeedInfoForList,
	tagName: "div",
	template: Template,

	onDomRefresh: function () {
	    $("#FeedList").DataTable({
		"aoColumnDefs": [
			{ 'bSortable': false, 'aTargets': [ 0 ] }
	         ]
	    });
	}
});
