var Channel = require("./channel.js")

module.exports = Feed = Backbone.Model.extend({

	urlRoot: "/feeds",

	parse: function (response, options) {
		
		var channel = new Channel(response.Channel);
		response.Channel = channel;
		return response;
	}

});
