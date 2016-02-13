var Item = require('../models/item');

module.exports = ItemCollection = Backbone.Collection.extend({
	model:  Item,

	url: "/items",

	comparator: function(model) {
		return model.get("PubTime");
	}
});
