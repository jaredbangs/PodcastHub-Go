var Backbone = require("backbone")
var DownloadDirectory = require("../models/downloadDirectory.js");
var ItemCollection = require("../collections/items.js");

module.exports = Backbone.Collection.extend({
	model:  DownloadDirectory,

	url: "/itemsByDownloadDirectory",

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
