var Backbone = require("backbone")
var DownloadDirectory = require("../models/downloadDirectory.js");
var ItemCollection = require("../collections/items.js");

DownloadDirectories = Backbone.Collection.extend({
	model:  DownloadDirectory,

	url: "/itemsByDownloadDirectory",

	getNonArchivedDirectories: function () {
		return new DownloadDirectories(this.filter(function (dir) { return dir.get("ItemCount") > 0; }));
	},

	parse: function(response) {

		return _.map(response, function (value, key) {

			var dd, items
			       
			dd = new DownloadDirectory();

			dd.set("Name", key);
			dd.set("UrlName", encodeURIComponent(key));

			dd.set("Items", new ItemCollection(response[key]));

			items = dd.get("Items").filter(function (item) { return item.shouldDisplayByDefault(); });

			dd.set("ItemCount", items.length);

			return dd;	
            	});
      	}

});

module.exports = DownloadDirectories;
