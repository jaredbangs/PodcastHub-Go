var Feed = require('../models/feed.js');
var Marionette = require("backbone.marionette");
var FeedArchiveStrategyCollection = require('../collections/feedArchiveStrategies.js');
var ItemCollection = require('../collections/items.js');
var FeedView = require('./feed.js')
var Template = require('../../templates/downloadDirectoryItemForList.handlebars');

module.exports = Marionette.ItemView.extend({
	
	tagName: "div",
	className: "row feed-info",

	template: Template,

	events: {
                "click .archive": "archive",
        },

        archive: function(domEvent) {

                var self = this;

                self.model.save({IsToBeArchived: true}, {
                        success: function() {
                                console.log("Item updated");
                                self.remove();
                        }
                 });
	},

	id: function() {
		return this.model.cid;
	}
});
