var Item = require('../models/item');

module.exports = Backbone.Collection.extend({
	model:  Item,

	url: "/items",

	comparator: function(model) {
		return model.get("PubTime");
	}
});
