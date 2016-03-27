var Backbone = require("backbone")
var Strategy = require('../models/feedArchiveStrategy.js');

module.exports = Backbone.Collection.extend({
	
	model:  Strategy,

	url: "/feedArchiveStrategies"

});
