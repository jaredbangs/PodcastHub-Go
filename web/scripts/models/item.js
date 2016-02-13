module.exports = Item = Backbone.Model.extend({

	initialize: function () {

		if (this.has("PubTime")) {
			
			var pubTime = moment(this.get("PubTime"))

			this.unset("PubTime");
			this.set("PubTime", pubTime._d);
			this.set("PubDisplayDate", pubTime.format("dddd, MMMM Do YYYY"));
		}
	}
});
