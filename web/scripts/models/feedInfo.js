module.exports = Backbone.Model.extend({

	idAttribute: "Id",

	parse: function (response, options) {

		var lastUpdated = moment(response.LastUpdated);
		response.LastUpdated = lastUpdated._d;
		response.LastUpdatedDisplay = lastUpdated.format("MMMM Do YYYY");
		return response;
	}
});
