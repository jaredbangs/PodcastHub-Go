var Backbone = require("backbone")
var ItemCollection = require("../collections/items.js")

module.exports = Channel = Backbone.Model.extend({

	initialize: function () {

		var itemList = this.get("ItemList");

		if (_.isArray(itemList)) {
			this.unset("ItemList");
			this.set("Items", new ItemCollection(itemList));
		}
	}

});
