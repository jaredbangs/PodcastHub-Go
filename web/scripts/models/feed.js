var Backbone = require("backbone")
var Channel = require("./channel.js")

module.exports = Backbone.Model.extend({

	urlRoot: "/feeds",

	parse: function (response, options) {
		
		var channel = new Channel(response.Channel);
		response.Channel = channel;
		return response;
	}

});
