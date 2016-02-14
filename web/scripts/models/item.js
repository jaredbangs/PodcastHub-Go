module.exports = Item = Backbone.Model.extend({

	idAttribute: "ItemId",

	initialize: function () {

		if (this.has("PubTime")) {
			
			var pubTime = moment(this.get("PubTime"))

			this.unset("PubTime");
			this.set("PubTime", pubTime._d);
			this.set("PubDisplayDate", pubTime.format("dddd, MMMM Do YYYY"));
		}

		this.set("Archived", this.get("IsArchived") || this.get("IsToBeArchived")); 
	}
});
