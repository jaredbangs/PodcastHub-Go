var Backbone = require("backbone")

module.exports = Backbone.Model.extend({

	idAttribute: "ItemId",

	initialize: function () {

		if (this.has("PubTime") && this.get("PubTime") !== "0001-01-01T00:00:00Z") {
			
			var pubTime = moment(this.get("PubTime"))

			this.unset("PubTime");
			this.set("PubTime", pubTime._d);
			this.set("PubDisplayDate", pubTime.format("dddd, MMMM Do YYYY"));
		} else if (this.has("PubDate")) {

			this.set("PubDisplayDate", this.get("PubDate"));
		}

		this.set("Archived", this.get("IsArchived") || this.get("IsToBeArchived")); 

		if (this.has("FeedId") && this.get("FeedId") !== "") {
			
			var feedInfo = podcasthub.FeedList.findWhere({ Id: this.get("FeedId") });

			if (feedInfo !== undefined) {
				this.set("FeedName", feedInfo.get("Title"));
				this.set("HasArchiveStrategy", feedInfo.has("ArchiveStrategy") && feedInfo.get("ArchiveStrategy") !== "");
			} else {
				console.log("Feed not found: " + this.get("FeedId"));
			}
		} else {
			console.log("Item has no FeedId");
			console.log(JSON.stringify(this));
		}
	},

	shouldDisplayByDefault: function () {

		var enclosures = this.get("Enclosures");

		return !this.get("Archived") && _.some(enclosures, function (enclosure) { return enclosure.DownloadedFilePath !== "" });
	}
});
